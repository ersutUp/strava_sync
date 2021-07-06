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

	//行者账号
	XzAccount string
	//行者密码
	XzPass string

	//黑鸟账号
	BlackbirdUsername string
	//黑鸟密码
	BlackbirdPass string
}
