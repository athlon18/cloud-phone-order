package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"order/common/util"
	"order/service"
	"os"
	"time"
)

func File(ctx *gin.Context) {
	file := service.GetFile()
	ctx.JSON(http.StatusOK, util.Result().SetSuccess(file))
}

func FileHtml(ctx *gin.Context) {
	files, list := service.FileList()
	ctx.HTML(http.StatusOK, "file.html", gin.H{
		"data":     files,
		"time":     time.Now().Unix(),
		"filelist": list,
	})
}

func FolderCreate(ctx *gin.Context) {
	_dir := ctx.DefaultPostForm("folder", "")
	_dir = "./upload/" + _dir
	exist, err := PathExists(_dir)
	if err != nil {
		fmt.Printf("get1 dir error![%v]\n", err)
		ctx.Redirect(http.StatusFound, "/api/v1/file/index")
		return
	}

	if !exist {
		// 创建文件夹
		if err := os.Mkdir(_dir, os.ModePerm); err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		}
	}
	ctx.Redirect(http.StatusFound, "/api/v1/file/index")
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func FileDelete(ctx *gin.Context) {
	_dir := ctx.DefaultPostForm("file", "")
	_dir = "./upload/" + _dir
	exist, err := PathExists(_dir)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		ctx.Redirect(http.StatusFound, "/api/v1/file/index")
		return
	}
	if exist {
		if err := os.RemoveAll(_dir); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	ctx.Redirect(http.StatusFound, "/api/v1/file/index")
}
