package service

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/godcong/ipfs-node-service/config"
	"github.com/godcong/ipfs-node-service/openssl"
	"github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

// StatusUploaded 已上传
const StatusUploaded = "uploaded"

// StatusDownloading 正在下载
const StatusDownloading = "downloading"

// StatusDownloaded 已下载
const StatusDownloaded = "downloaded"

// StatusQueuing 队列中
const StatusQueuing = "queuing"

// StatusTransferring 转换中
const StatusTransferring = "transferring"

// StatusFileWrong 文件错误
const StatusFileWrong = "wrong file"

//StatusFailed 失败
const StatusFailed = "failed"

// StatusFinished 完成
const StatusFinished = "finished"

// HandleFunc ...
type HandleFunc func(name, key string) error

// QueueServer ...
type QueueServer struct {
	*redis.Client
	config    *config.Configure
	Processes int
	cancel    context.CancelFunc
}

var globalQueue *QueueServer

// Push ...
func Push(v *Streamer) {
	globalQueue.Push(v)
}

// Push ...
func (s *QueueServer) Push(v *Streamer) {
	s.RPush("node_queue", v.JSON())
}

// Pop ...
func Pop() *Streamer {
	return globalQueue.Pop()
}

// Pop ...
func (s *QueueServer) Pop() *Streamer {
	pop := s.LPop("node_queue").Val()
	return ParseStreamer(pop)
}

func transfer(ch chan<- string, info *Streamer) {
	var e error
	chanRes := info.FileName()
	defer func() {
		if e != nil {
			log.Error(e)
			globalQueue.Set(info.ID, StatusFailed, 0)
			chanRes = info.FileName() + ":[" + e.Error() + "]"
		}
		ch <- chanRes
	}()

	globalQueue.Set(info.ID, StatusDownloading, 0)
	e = downloadFromOSS(info)
	if e != nil {
		return
	}

	globalQueue.Set(info.ID, StatusTransferring, 0)
	log.Infof("%+v", info)
	if info.Encrypt {
		if info.Key == "" {
			b, e := openssl.Base64Key()
			if e != nil {
				log.Error(e)
				return
			}
			info.Key = string(b)
		}
		name := info.KeyFile()
		e = toM3U8WithKey(info.ID, info.SourceFile(), info.Transfer, name)
	} else {
		e = toM3U8(info.ID, info.SourceFile(), info.Transfer)
	}

	if e != nil {
		return
	}

	detail, e := commitToIPNS(info.ID, info.DestPath())
	if e != nil {
		return
	}

	var qr QueueResult

	e = mapstructure.Decode(detail, &qr)
	if e != nil {
		log.Error(e)
		return
	}
	qr.Key = info.Key
	log.Info(qr)
	e = NewBack().Callback(&qr)

	if e != nil {
		log.Error(e)
		return
	}

	cfg := config.Config()
	if cfg.Node.Clear {
		e = clearDownload(info.SourceFile())
		if e != nil {
			log.Error(e)
			return
		}
		//e = clearTransfer(info.Transfer, info.ID)
		//if e != nil {
		//	log.Error(e)
		//	return
		//}
	}

}

func clearTransfer(file string, id string) error {
	filename := filepath.Join(file, id)
	info, e := os.Stat(filename)
	log.Printf("%s,%+v,%+v\n", filename, info, e) //has nil
	if e == nil {
		e = os.RemoveAll(filename)
		if e == nil {
			return nil
		}
	}
	return e
}

func clearDownload(filename string) error {
	info, e := os.Stat(filename)
	log.Printf("%s,%+v,%+v\n", filename, info, e) //has nil
	if e == nil {
		e = os.Remove(filename)
		if e == nil {
			return nil
		}
	}
	return e
}

// QueueResult ...
type QueueResult struct {
	ID     string `json:"id" mapstructure:"id"`
	FSInfo struct {
		Hash string `json:"hash" mapstructure:"hash"`
		Name string `json:"name" mapstructure:"name"`
		Size string `json:"size" mapstructure:"size"`
	} `json:"fs_info" mapstructure:"fs_info"`
	NSInfo struct {
		Name  string `json:"name" mapstructure:"name"`
		Value string `json:"value" mapstructure:"value"`
	} `json:"ns_info" mapstructure:"ns_info"`
	Key string `json:"key" mapstructure:"key"`
}

// JSON ...
func (r *QueueResult) JSON() string {
	s, _ := jsoniter.MarshalToString(r)
	return s
}

func transferNothing(threads chan<- string) {
	time.Sleep(3 * time.Second)
	threads <- ""
}

// NewQueueServer ...
func NewQueueServer(cfg *config.Configure) *QueueServer {
	client := redis.NewClient(&redis.Options{
		Addr:     config.DefaultString(cfg.Queue.HostPort, ":6379"),
		Password: config.DefaultString(cfg.Queue.Password, ""), // no password set
		DB:       cfg.Queue.DB,                                 // use default DB
	})
	return &QueueServer{
		config: cfg,
		Client: client,
	}
}

// Start ...
func (s *QueueServer) Start() {
	pong, err := s.Ping().Result()
	if err != nil {
		panic(err)
	}
	globalQueue = s
	log.Println(pong)

	var c context.Context
	c, s.cancel = context.WithCancel(context.Background())
	//run with a new go channel
	go func() {
		threads := make(chan string, s.Processes)

		for i := 0; i < s.Processes; i++ {
			log.Println("start", i)
			go transferNothing(threads)
		}

		for {
			select {
			case v := <-threads:
				if v != "" {
					log.Println("success: ", v)
				}

				if s := Pop(); s != nil {
					go transfer(threads, s)
				} else {
					go transferNothing(threads)
				}
			case <-c.Done():
				return
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}()
}

// Stop ...
func (s *QueueServer) Stop() {
	if s.cancel == nil {
		return
	}
	s.cancel()
}
