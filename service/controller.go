package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/go-ffmpeg/openssl"
	"log"
	"net/http"
)

/**
 * @apiDefine Success
 * @apiSuccess {string} msg 返回具体消息
 * @apiSuccess {int} code 返回状态码：【正常：0】，【失败，-1】
 * @apiSuccess {json} [detail]  如正常则返回detail
 */
/**
 * @apiDefine Failed
 * @apiErrorExample {json} Error-Response:
 *     {
 *       "code":-1,
 *       "msg":"error message",
 *     }
 */
const _ = "apiDefine"

//UploadPost 文件上传接口
/**
* @api {post} /v1/upload 文件上传接口
* @apiName upload
* @apiGroup Upload
* @apiVersion  0.0.1
*
* @apiUse Success
* @apiParam  {Binary} binary 媒体文件二进制文件
* @apiParamExample  {Binary} Request-Example:
*    upload a binary file from local
* @apiSuccessExample {json} Success-Response:
*     {
*       "code":0,
*       "msg":"ok",
*       "detail":{
*			"id":"9FCp2x2AeEWNobvzKA3vRgqzZNqFWEJTMpLAz2hLhQGEd3URD5VTwDdTwrjTu2qm"
*		 }
*     }
* @apiSuccess (detail) {string} id 文件名ID
* @apiUse Failed
* @apiSampleRequest /v1/upload
 */
func UploadPost(vertion string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer ctx.Request.Body.Close()
		src := "./upload/"
		fileName, err := writeTo(src, ctx.Request.Body)
		if err != nil {
			ctx.JSON(http.StatusOK, JSON(-1, err.Error()))
			return
		}

		client.Set(fileName, string(b), 0)
		log.Println(fileName)
		ctx.JSON(http.StatusOK, JSON(0, "ok", gin.H{"id": fileName}))
		return
	}
}

// TransferPost 视频转换处理
/**
* @api {post} /v1/transfer 视频转换处理
* @apiName transfer
* @apiGroup Transfer
* @apiVersion  0.0.1
*
* @apiUse Success
* @apiParam  {string} id 文件名ID
* @apiParam  {string} url KeyFile存放URL
* @apiParamExample  {string} Request-Example:
* {
*     "id":"9FCp2x2AeEWNobvzKA3vRgqzZNqFWEJTMpLAz2hLhQGEd3URD5VTwDdTwrjTu2qm"
* }
*
* @apiSuccessExample {json} Success-Response:
*     {
*       "code":0,
*       "msg":"ok",
*       "detail":{
*			"id":"9FCp2x2AeEWNobvzKA3vRgqzZNqFWEJTMpLAz2hLhQGEd3URD5VTwDdTwrjTu2qm"
*		 }
*     }
* @apiSuccess (detail) {string} id 文件名ID
* @apiUse Failed
 */
func TransferPost(version string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		b, err := openssl.HexKey()
		if err != nil {
			ctx.JSON(http.StatusOK, JSON(-1, err.Error()))
			return
		}
		id := ctx.PostForm("id")
		stream := NewStreamer(string(b), id)
		stream.SetURI("http://localhost:8080/infos" + "/" + id + "/key")
		stream.SetDst("./transfer/")
		stream.SetSrc(src)
		queue.Push(stream)
		ResultOK(ctx)
	}
}

// StatusGet 获取视频转换状态
/**
*
* @api {get} /v1/info/:id 获取视频转换状态
* @apiName info
* @apiGroup Info
* @apiVersion  0.0.1
*
* @apiParam  {String} id 文件名ID
*
* @apiUse Success
* @apiSuccess  {string} code 返回状态码：【正常：0】，【处理中：1】，【ID不存在：2】
*
* @apiSampleRequest /v1/status/:id
* @apiParamExample  {string} Request-Example:
* 	http://localhost:8080/v1/status/9FCp2x2AeEWNobvzKA3vRgqzZNqFWEJTMpLAz2hLhQGEd3URD5VTwDdTwrjTu2qm
*
* @apiSuccessExample {json} Success-Response OK:
* {
*       "code":0,
*       "msg":"ok",
* }
* @apiSuccessExample {json} Success-Response Processing:
* {
*       "code":1,
*       "msg":"processing",
* }
* @apiUse Failed
 */
func InfoGet(version string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		key, err := client.Get(id).Result()
		if err != nil {
			ctx.JSON(http.StatusOK, JSON(2, "data not found"))
			return
		}

		if key != "" {
			ctx.JSON(http.StatusOK, JSON(1, "processing"))
			return
		}

		ResultOK(ctx, gin.H{})
		return
	}
}

// StatusGet 获取视频转换状态
/**
*
* @api {get} /v1/list 获取所有视频列表
* @apiName list
* @apiGroup List
* @apiVersion  0.0.1
*
* @apiUse Success
* @apiSuccess  {string} code 返回状态码：【正常：0】，【处理中：1】
*
* @apiSampleRequest /v1/list
* @apiParamExample  {string} Request-Example:
* 	http://localhost:8080/v1/list
*
* @apiSuccessExample {json} Success-Response OK:
* {
*       "code":0,
*       "msg":"ok",
* }
* @apiUse Failed
 */
func ListGet(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ResultOK(ctx)
	}
}

func ResultOK(ctx *gin.Context, h ...gin.H) {
	if h != nil {
		ctx.JSON(http.StatusOK, JSON(0, "ok", h...))
		return
	}
	ctx.JSON(http.StatusOK, JSON(0, "ok"))
}

func ResultFail(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, JSON(-1, "fail"))
}
