package service

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/godcong/go-ffmpeg/util"
	"github.com/rakyll/statik/fs"
)

// Router ...
func Router(engine *gin.Engine) error {
	//api document
	//engine.Static("/doc", "./doc")
	st, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	engine.StaticFS("/doc", st)

	ver := "/v1"
	group := engine.Group(ver)
	group.Use()
	group.Static("/transfer", "./transfer")

	group.Any("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "%s", "hello world")
	})

	//上传文件
	group.POST("/upload", UploadPost(ver))

	//视频转换
	group.POST("/transfer", TransferPost(ver))

	//下载并转换
	group.POST("/downloadTransform", func(ctx *gin.Context) {

	})

	//服务器视频列表
	group.GET("/list/:id", ListGet(ver))

	//查看状态
	group.GET("/info/:id", InfoGet(ver))

	return nil
}

func writeTo(path string, reader io.Reader) (string, error) {
	fileName := util.GenerateRandomString(64)
	file, err := os.OpenFile(path+fileName, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return "", err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	i, err := io.Copy(writer, reader)
	if err != nil {
		return "", err
	}
	err = writer.Flush()
	if err != nil {
		return "", err
	}

	if i == 0 {
		return "", fmt.Errorf("upload with %d", i)
	}
	log.Println("success:", i)
	return fileName, nil
}
