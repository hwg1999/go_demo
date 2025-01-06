package webframework_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_single_file_trans(t *testing.T) {
	e := gin.Default()
	e.POST("/upload", uploadFile)
	log.Fatalln(e.Run(":8080"))
}

func uploadFile(ctx *gin.Context) {
	// 获取文件
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, "%+v", err)
		return
	}
	// 保存在本地
	err = ctx.SaveUploadedFile(file, "./"+file.Filename)
	if err != nil {
		ctx.String(http.StatusBadRequest, "%+v", err)
		return
	}
	// 返回结果
	ctx.String(http.StatusOK, "upload %s size:%d byte successfully!", file.Filename, file.Size)
}

func Test_multi_file_trans(t *testing.T) {
	e := gin.Default()
	e.POST("/upload", uploadFile)
	e.POST("/uploadFiles", uploadFiles)
	log.Fatalln(e.Run(":8080"))
}

func uploadFiles(ctx *gin.Context) {
	// 获取gin解析好的multipart表单
	form, _ := ctx.MultipartForm()
	// 根据键值取得对应的文件列表
	files := form.File["files"]
	// 遍历文件列表，保存到本地
	for _, file := range files {
		err := ctx.SaveUploadedFile(file, "./"+file.Filename)
		if err != nil {
			ctx.String(http.StatusBadRequest, "upload failed")
			return
		}
	}
	// 返回结果
	ctx.String(http.StatusOK, "upload %d files successfully!", len(files))
}

func Test_file_download(t *testing.T) {
	e := gin.Default()
	e.POST("/upload", uploadFile)
	e.POST("/uploadFiles", uploadFiles)
	e.GET("/download/:filename", download)
	log.Fatalln(e.Run(":8080"))
}

func download(ctx *gin.Context) {
	// 获取文件名
	filename := ctx.Param("filename")
	// 返回对应文件
	ctx.FileAttachment(filename, filename)
}
