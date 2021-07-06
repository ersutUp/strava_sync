package controllers

import (
	"encoding/json"
	"fit_sync_server/conf/db"
	"fit_sync_server/models"
	beego "github.com/beego/beego/v2/server/web"
)

//近期的多少条记录需要同步，防止一直同步个不停
var limit = 20

type TrainingController struct {
	beego.Controller
}


func (this *TrainingController) URLMapping()  {
	//一定要写
	this.Mapping("GetXZ",this.GetXZ)
}


// @router / [post]
func (this *TrainingController) Post() {
	var trainingOriginalDatas []string
	json.Unmarshal(this.Ctx.Input.RequestBody, &trainingOriginalDatas)

	//需要上传fit文件的id(strava_id)
	trainingIDs := []int64{}

	//查询近n条记录中没有fit的记录
	trainingDBNotFits := []int64{}
	//db.Mydb.Model(&models.Training{}).Select("strava_id").Where("fit_path = ''").Order("id desc").Limit(10).Find(&trainingDBNotFits)
	db.Mydb.Table("(?) as u",db.Mydb.Model(&models.Training{}).Order("id desc").Limit(limit)).Select("strava_id").Where("fit_path = ''").Find(&trainingDBNotFits)

	//没有fit文件也需要上传fit文件
	trainingIDs = append(trainingIDs, trainingDBNotFits...)

	//遍历数据
	for _, data := range trainingOriginalDatas {
		var training models.Training
		//json字符串转 结构体
		json.Unmarshal([]byte(data), &training)
		training.OriginalData = data
		training.StravaId = training.ID
		training.ID = 0

		//数据入库
		var trainingDB models.Training
		db.Mydb.Model(&models.Training{}).Where("strava_id = ?",training.StravaId).Find(&trainingDB)
		//数据库是否有数据
		if trainingDB.ID > 0 {
			training.ID = trainingDB.ID
			db.Mydb.Omit("fit_path").Updates(training)
		} else {
			db.Mydb.Create(&training)
			//新数据一定需要上传fit
			trainingIDs = append(trainingIDs, training.StravaId)
		}
	}

	this.Data["json"] = trainingIDs
	this.ServeJSON()

}

//获取近n条记录中没有上传行者的记录
// @router /xz [get]
func (this *TrainingController) GetXZ() {

	//需要上传fit文件的id(strava_id)
	ids := []int64{}

	//查询近n条记录中没有上传行者的记录
	trainingDBNotUploadXZ := []int64{}
	db.Mydb.Table("(?) as u",db.Mydb.Model(&models.Training{}).Order("id desc").Limit(limit)).Select("id").Where("is_upload_xingzhe = 0 and fit_path != ''").Find(&trainingDBNotUploadXZ)

	ids = append(ids, trainingDBNotUploadXZ...)

	this.Data["json"] = ids
	this.ServeJSON()

}

//通知已上传行者
// @router /xz [put]
func (this *TrainingController) UpdateXZ() {

	training := models.Training{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &training)

	if training.IsUploadXingzhe == 0 {
		this.Ctx.WriteString("not XZ id")
		return
	}

	if training.ID == 0 {
		this.Ctx.WriteString("not id")
		return
	}

	db.Mydb.Select("is_upload_xingzhe").Updates(training)

	this.Ctx.WriteString("ok")

}

//获取近n条记录中没有上传黑鸟的记录
// @router /blackbird [get]
func (this *TrainingController) GetBlackbird() {

	//需要上传fit文件的id(strava_id)
	ids := []int64{}

	//查询近n条记录中没有上传黑鸟的记录(过滤掉虚拟骑行的数据)
	trainingDBNotUploadBlackbird := []int64{}
	db.Mydb.Table("(?) as u",db.Mydb.Model(&models.Training{}).Order("id desc").Limit(limit)).Select("id").Where("is_upload_blackbird = 0 and fit_path != '' and type != 'VirtualRide'").Find(&trainingDBNotUploadBlackbird)

	ids = append(ids, trainingDBNotUploadBlackbird...)

	this.Data["json"] = ids
	this.ServeJSON()

}

//通知已上传黑鸟
// @router /blackbird [put]
func (this *TrainingController) UpdateBlackbird() {

	training := models.Training{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &training)

	if training.IsUploadBlackbird == 0 {
		this.Ctx.WriteString("not blackbird id")
		return
	}

	if training.ID == 0 {
		this.Ctx.WriteString("not id")
		return
	}

	db.Mydb.Select("is_upload_blackbird").Updates(training)

	this.Ctx.WriteString("ok")

}

