package services

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
	"github.com/udistrital/utils_oas/requestresponse"
)

func GetAllPersonas_parametros() (APIResponseDTO requestresponse.APIResponse) {
	var personas map[string]interface{}

	fmt.Println("http://" + beego.AppConfig.String("UrlCrudContactos") + fmt.Sprintf("/contacto?query=Activo:true&limit=0"))

	errPersona := request.GetJson("http://"+beego.AppConfig.String("UrlCrudContactos")+fmt.Sprintf("/contacto?query=Activo:true&limit=0"), &personas)

	if errPersona == nil {
		APIResponseDTO = requestresponse.APIResponseDTO(true, 200, nil, personas)
		return APIResponseDTO
	}
	APIResponseDTO = requestresponse.APIResponseDTO(true, 200, nil, nil)
	return APIResponseDTO
}
