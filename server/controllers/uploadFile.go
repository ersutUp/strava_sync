package controllers

import (
	"fit_sync_server/utils"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/sirupsen/logrus"
)

type UploadFileController struct {
	beego.Controller
}

func (uf *UploadFileController) URLMapping()  {
	//一定要写

}

// @router / [post]
func (uf *UploadFileController) Post() {
	f, _, _ := uf.GetFile("data")//获取上传的文件
	if f == nil {
		uf.Ctx.ResponseWriter.Status = 500
		uf.Ctx.WriteString("没有接收到文件")
		return
	}
	defer f.Close()//关闭上传的文件，不然的话会出现临时文件不能清除的情况

	id := uf.GetString("id")

	dir := "./fit/"
	//如果目录不存在则创建目录
	createFolderErr := utils.CreateFolder(dir)
	if createFolderErr != nil {
		logrus.Info("文件夹创建失败",createFolderErr)
		uf.Ctx.ResponseWriter.Status = 500
		uf.Ctx.WriteString("文件夹创建失败")
		return
	}

	//保存文件
	err := uf.SaveToFile("data", dir+id+".fit")
	if err != nil {
		uf.Ctx.ResponseWriter.Status = 500
		uf.Ctx.WriteString("文件保存失败")
		return
	}

	//todo 更新数据库

	uf.Ctx.WriteString( "上传成功" )
}