package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/AgendaBeego/parametros_mid_v1/services"
	"github.com/udistrital/utils_oas/errorhandler"
)

// Personas_parametrosController operations for Personas_parametros
type Personas_parametrosController struct {
	beego.Controller
}

// URLMapping ...
func (c *Personas_parametrosController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Personas_parametros
// @Param	body		body 	models.Personas_parametros	true		"body for Personas_parametros content"
// @Success 201 {object} models.Personas_parametros
// @Failure 403 body is empty
// @router / [post]
func (c *Personas_parametrosController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Personas_parametros by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Personas_parametros
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Personas_parametrosController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Personas_parametros
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Personas_parametros
// @Failure 403
// @router / [get]
func (c *Personas_parametrosController) GetAll() {
	defer errorhandler.HandlePanic(&c.Controller)
	//fmt.Println("UWU")
	respuesta := services.GetAllPersonas_parametros()
	c.Ctx.Output.SetStatus(respuesta.Status)
	c.Data["json"] = respuesta
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Personas_parametros
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Personas_parametros	true		"body for Personas_parametros content"
// @Success 200 {object} models.Personas_parametros
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Personas_parametrosController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Personas_parametros
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Personas_parametrosController) Delete() {

}
