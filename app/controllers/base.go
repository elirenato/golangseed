package controllers

import (
	"github.com/revel/revel"
    "net/http"
    "github.com/elirenato/golangseed/app/commons"
    "log"
)

type BaseController struct {
    *revel.Controller
}

func (c BaseController) decodeAndValidateRequest(dtoPtr interface{}) revel.Result {
    err := commons.DecodeJsonBody(c.Request.Body, dtoPtr)
	if err != nil {
		log.Println(err)
		return c.RenderBadRequest("error.generic.invalidBody")
	}
	//validate the basic fields
	dtoPtr.(commons.Validate).Validate(c.Validation)
	if c.Validation.HasErrors() {
		log.Println(c.Validation.Errors)
		return c.RenderBadRequest("error.generic.validation")
	}
    return nil
}

func (c BaseController) RenderOK(o interface{}) (revel.Result) {
    c.Response.Status = http.StatusOK
    c.Response.ContentType = commons.ApplicationJsonContentType
    return c.RenderJSON(o)
}

func (c BaseController) RenderNotFound(messageKey string, args ...interface{}) (revel.Result) {
    message := c.Message(messageKey, args)
    log.Println(message)
    c.Response.Status = http.StatusNotFound
    c.Response.ContentType = commons.ApplicationJsonContentType
    result := make(map[string]interface{})
    result["error"] = messageKey
    result["message"] = message
    return c.RenderJSON(result)
}

func (c BaseController) RenderBadRequest(messageKey string, args ...interface{}) (revel.Result) {
    message := c.Message(messageKey, args)
    log.Println(message)
    c.Response.Status = http.StatusBadRequest
    c.Response.ContentType = commons.ApplicationJsonContentType
    result := make(map[string]interface{})
    result["error"] = messageKey
    result["message"] = message
    return c.RenderJSON(result)
}

func (c BaseController) RenderInternalServerError(messageKey string, args ...interface{}) (revel.Result) {
    message := c.Message(messageKey, args)
    log.Println(message)
    c.Response.Status = http.StatusInternalServerError
    c.Response.ContentType = commons.ApplicationJsonContentType
    result := make(map[string]interface{})
    result["error"] = messageKey
    result["message"] = message
    return c.RenderJSON(result)
}