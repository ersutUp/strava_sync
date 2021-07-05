package models

import "time"

type Training struct {

	ID int64

	Type string
	OriginalData string
	Name string
	StartDate string `json:"start_date"`
	//里程
	Distance string
	StravaId int64
	FitPath string `json:"-"`
	IsUploadXingzhe   int64
	IsUploadBlackbird bool
	CreatedAt time.Time
	UpdatedAt time.Time

}
