package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order/common/util"
	"order/service"
	"path"
)

func UploadHtml(ctx *gin.Context) {
	files := service.File()
	ctx.HTML(http.StatusOK, "upload.html", gin.H{"data": files})
}

func Upload(ctx *gin.Context) {
	f, err := ctx.FormFile("f1") //和从请求中获取携带的参数一样
	status := ctx.PostForm("status")
	folder := ctx.DefaultPostForm("folder", "")
	if err != nil {
		ctx.JSON(http.StatusOK, util.Result().SetError(http.StatusBadRequest, err.Error(), nil))
		return
	}
	uploadPath := "./upload"
	if folder != "" {
		uploadPath += "/" + folder
	}

	//将读取到的文件保存到本地(服务端)
	dst := path.Join(uploadPath, f.Filename)
	_ = ctx.SaveUploadedFile(f, dst)
	service.UploadFile(f, dst, status)
	ctx.Redirect(http.StatusFound, "/api/v1/html/index")

}

func UploadDelete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id != "" {
		if err := service.DeleteFile(id); err != nil {
			ctx.Redirect(http.StatusFound, "/api/v1/html/index")
			return
		}
	}
	ctx.Redirect(http.StatusFound, "/api/v1/html/index")
}
