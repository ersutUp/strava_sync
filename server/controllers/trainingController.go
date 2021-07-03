package controllers

import (
	"encoding/json"
	"fit_sync_server/conf/db"
	"fit_sync_server/models"
	beego "github.com/beego/beego/v2/server/web"
)

type TrainingController struct {
	beego.Controller
}



// @router / [post]
func (this *TrainingController) Post()  {
	var trainingOriginalDatas []string
	json.Unmarshal(this.Ctx.Input.RequestBody, &trainingOriginalDatas)

	//需要上传fit文件的id(strava_id)
	trainingIDs := []int64{}

	//遍历数据
	for _, data := range trainingOriginalDatas {
		var training models.Training
		//json字符串转 结构体
		json.Unmarshal([]byte(data), &training)
		training.OriginalData = data

		//数据入库
		var trainingDB models.Training
		db.Mydb.Model(&models.Training{}).Where("strava_id = ?",training.StravaId).Find(&trainingDB)
		//数据库是否有数据
		if trainingDB.ID > 0 {
			training.ID = trainingDB.ID
			db.Mydb.Omit("fit_path").Updates(training)
			//没有fit文件也需要上传fit文件
			if trainingDB.FitPath == "" {
				trainingIDs = append(trainingIDs, training.StravaId)
			}
		} else {
			db.Mydb.Create(&training)
			//新数据一定需要上传fit
			trainingIDs = append(trainingIDs, training.StravaId)
		}
	}

	this.Data["json"] = trainingIDs
	this.ServeJSON()

}