package models

type UserInfo struct {
	ID int64

	StravaEmail string
	StravaPass string
	//获取多少天内的数据
	BeforeDay int16
	StravaSyncSecond int32

	//server酱
	SendKey string
}
