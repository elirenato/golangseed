package controllers

import (
	"github.com/revel/revel"
    "net/http"
    "github.com/elirenato/golangseed/app/commons"
)

type BaseController struct {
    *revel.Controller
}

func (c BaseController) RenderOK(o interface{}) (revel.Result) {
    c.Response.Status = http.StatusOK
    c.Response.ContentType = commons.ApplicationJsonContentType
    return c.RenderJSON(o)
}

func (c BaseController) RenderNotFound() (revel.Result) {
    c.Response.Status = http.StatusNotFound
    c.Response.ContentType = commons.ApplicationJsonContentType
    return c.RenderJSON(nil)
}

func (c BaseController) RenderBadRequest(messageKey string, args ...interface{}) (revel.Result) {
    c.Response.Status = http.StatusBadRequest
    c.Response.ContentType = commons.ApplicationJsonContentType
    result := make(map[string]interface{})
    result["error"] = messageKey
    result["message"] = c.Message(messageKey)
    return c.RenderJSON(result)
}

func (c BaseController) RenderInternalServerError() (revel.Result) {
    c.Response.Status = http.StatusInternalServerError
    c.Response.ContentType = commons.ApplicationJsonContentType
    return c.RenderJSON(nil)
}