package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/AgendaBeego/parametros_mid_v1/controllers:Personas_parametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/AgendaBeego/parametros_mid_v1/controllers:Personas_parametrosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/AgendaBeego/parametros_mid_v1/controllers:Personas_parametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/AgendaBeego/parametros_mid_v1/controllers:Personas_parametrosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/AgendaBeego/parametros_mid_v1/controllers:Personas_parametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/AgendaBeego/parametros_mid_v1/controllers:Personas_parametrosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/AgendaBeego/parametros_mid_v1/controllers:Personas_parametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/AgendaBeego/parametros_mid_v1/controllers:Personas_parametrosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/AgendaBeego/parametros_mid_v1/controllers:Personas_parametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/AgendaBeego/parametros_mid_v1/controllers:Personas_parametrosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
