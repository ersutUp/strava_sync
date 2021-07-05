package controllers

import (
	"fit_sync_server/conf/db"
	"fit_sync_server/models"
	"fit_sync_server/utils"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/sirupsen/logrus"
)


var fileFolder = utils.GetAppConfig("fileFolder")

type FileController struct {
	beego.Controller
}

func (this *FileController) URLMapping()  {
	//一定要写
	this.Mapping("GetFit",this.GetFit)
}

// @router / [post]
func (this *FileController) Post() {
	f, _, _ := this.GetFile("data")//获取上传的文件
	if f == nil {
		this.Ctx.ResponseWriter.WriteHeader(500)
		this.Ctx.WriteString("没有接收到文件")
		return
	}
	defer f.Close()//关闭上传的文件，不然的话会出现临时文件不能清除的情况

	//strava的id
	id := this.GetString("id")

	dir := "fit/"
	//如果目录不存在则创建目录
	createFolderErr := utils.CreateFolder(dir)
	if createFolderErr != nil {
		logrus.Info("文件夹创建失败",createFolderErr)
		this.Ctx.ResponseWriter.WriteHeader(500)
		this.Ctx.WriteString("文件夹创建失败")
		return
	}

	filePath := dir+id+".fit"

	//保存文件
	err := this.SaveToFile("data", fileFolder+filePath)
	if err != nil {
		this.Ctx.ResponseWriter.WriteHeader(500)
		this.Ctx.WriteString("文件保存失败")
		return
	}

	//更新数据库
	var trainingDB models.Training
	db.Mydb.Model(&models.Training{}).Where("strava_id = ?",id).Find(&trainingDB)
	if trainingDB.ID > 0 {
		trainingDB.FitPath = filePath
		db.Mydb.Select("fit_path").Updates(trainingDB)
	} else {
		logrus.Warn("strava id is ["+id+"] not find db row")
	}

	this.Ctx.WriteString( "上传成功" )
}


// @router /fit [get]
func (this *FileController) GetFit() {

	id := this.Ctx.Input.Query("id")

	//根据id查询数据库
	var trainingDB models.Training
	db.Mydb.Model(&models.Training{}).Where("id = ?",id).Find(&trainingDB)

	if trainingDB.FitPath == "" {
		this.Ctx.ResponseWriter.WriteHeader(404)
		this.Ctx.WriteString("fit文件未入库")
		return
	} else {
		this.Ctx.Output.Download(fileFolder+trainingDB.FitPath)
	}
}