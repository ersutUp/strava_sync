package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["fit_sync_server/controllers:FileController"] = append(beego.GlobalControllerRouter["fit_sync_server/controllers:FileController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fit_sync_server/controllers:FileController"] = append(beego.GlobalControllerRouter["fit_sync_server/controllers:FileController"],
        beego.ControllerComments{
            Method: "GetFit",
            Router: "/fit",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fit_sync_server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["fit_sync_server/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fit_sync_server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["fit_sync_server/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fit_sync_server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["fit_sync_server/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fit_sync_server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["fit_sync_server/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fit_sync_server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["fit_sync_server/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fit_sync_server/controllers:TrainingController"] = append(beego.GlobalControllerRouter["fit_sync_server/controllers:TrainingController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fit_sync_server/controllers:TrainingController"] = append(beego.GlobalControllerRouter["fit_sync_server/controllers:TrainingController"],
        beego.ControllerComments{
            Method: "GetBlackbird",
            Router: "/blackbird",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fit_sync_server/controllers:TrainingController"] = append(beego.GlobalControllerRouter["fit_sync_server/controllers:TrainingController"],
        beego.ControllerComments{
            Method: "UpdateBlackbird",
            Router: "/blackbird",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fit_sync_server/controllers:TrainingController"] = append(beego.GlobalControllerRouter["fit_sync_server/controllers:TrainingController"],
        beego.ControllerComments{
            Method: "GetXZ",
            Router: "/xz",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fit_sync_server/controllers:TrainingController"] = append(beego.GlobalControllerRouter["fit_sync_server/controllers:TrainingController"],
        beego.ControllerComments{
            Method: "UpdateXZ",
            Router: "/xz",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fit_sync_server/controllers:UserInfoController"] = append(beego.GlobalControllerRouter["fit_sync_server/controllers:UserInfoController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
