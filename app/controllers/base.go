package controllers

import (
	"github.com/revel/revel"
    "net/http"
)

type BaseController struct {
    GorpController
}

func (c BaseController) RenderOK(result interface{}) (revel.Result) {
    c.Response.Status = http.StatusOK
    c.Response.ContentType = "application/json; charset=utf-8"
    return c.RenderJSON(result)
}

func (c BaseController) RenderBadRequest(messageKey string, args ...interface{}) (revel.Result) {
    c.Response.Status = http.StatusBadRequest
    c.Response.ContentType = "application/json; charset=utf-8"
    result := make(map[string]interface{})
    result["error"] = messageKey
    result["message"] = c.Message(messageKey)
    return c.RenderJSON(result)
}

func (c BaseController) RenderInternalServerError() (revel.Result) {
    c.Response.Status = http.StatusInternalServerError
    return c.Render();
}