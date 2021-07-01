package models

import "time"

type Training struct {

	ID int64 `json:"-"`

	Type string
	OriginalData string
	Name string
	StartDate string `json:"start_date"`
	//里程
	Distance string
	StravaId int64 `json:"id"`
	FitPath string `json:"-"`
	IsUploadXingzhe   bool `json:"-"`
	IsUploadBlackbird bool `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

}
