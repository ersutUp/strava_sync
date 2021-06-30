package filters

import (
	"github.com/beego/beego/v2/server/web"
)

func init()  {
	web.InsertFilterChain("/*", accessLog)
}
