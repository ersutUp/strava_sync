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

	var trainings []models.Training

	for _, data := range trainingOriginalDatas {
		var training models.Training
		json.Unmarshal([]byte(data), &training)
		training.OriginalData = data


		var trainingDB models.Training
		db.Mydb.Model(&models.Training{}).Where("strava_id = ?",training.StravaId).Find(&trainingDB)
		if trainingDB.ID > 0 {
			training.ID = trainingDB.ID
			db.Mydb.Omit("fit_path").Updates(training)
		} else {
			db.Mydb.Create(&training)
		}
		trainings = append(trainings, training)
	}


	this.Ctx.WriteString("success")

}