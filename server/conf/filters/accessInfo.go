package filters

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func accessLog(next web.FilterFunc) web.FilterFunc{
	return func(ctx *context.Context) {
		//接到请求的时间
		startTime := time.Now().UnixNano()
		//ip
		var ip = ctx.Input.IP()
		//请求地址
		uri := ctx.Request.RequestURI
		//request ID处理
		reqID := strconv.Itoa(int(startTime))
		ctx.Request.Header.Add("X-Request-Id",reqID)
		ctx.ResponseWriter.Header().Add("X-Request-Id",reqID)

		log.WithFields(log.Fields{
			"reqID":reqID,
			"url": uri,
			"ip":ip,
		}).Info("start request")

		next(ctx)

		//处理完后相应的时间
		endTime := time.Now().UnixNano()

		log.WithFields(log.Fields{
			"reqID":reqID,
			"responseMS":(endTime-startTime)/1e6,
		}).Info("end request")
	}
}