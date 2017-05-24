package controllers

import (
	"github.com/revel/revel"
	"net/http"
	"github.com/elirenato/golangseed/app/commons"
	"github.com/elirenato/golangseed/app/models"
	"log"
	"strconv"
	"strings"
	"fmt"
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

func (c BaseController) parsePageableRequest(defaultSort string) models.Pageable {
	var pageable models.Pageable
	pageable.Page, _ = strconv.ParseInt(c.Params.Get("page"), 10, 64)
	if pageable.Page < 0 {
		pageable.Page = 0
	}
	str := c.Params.Get("size")
	if !commons.IsBlank(str) {
		pageable.Size, _ = strconv.ParseInt(str, 10, 64)
	}
	if pageable.Size <= 0 || pageable.Size > models.MaxPageSize {
		pageable.Size = 20
	}
	sortRequest := c.Params.Get("sort")
	if commons.IsBlank(sortRequest) {
		sortRequest = defaultSort
	}
	sortFields := strings.Split(sortRequest, ",")
	sortSize := len(sortFields)
	if sortSize > 0 {
		for _, jsonSortField := range sortFields {
			sort := models.SortOption{}
			direction := jsonSortField[0:1]
			if  direction == "-" {
				sort.Direction = models.DirectionDesc
			} else {
				sort.Direction = models.DirectionAsc
			}
			if  direction == "-" || direction == "+" {
				sort.Property = jsonSortField[1:]
			} else {
				sort.Property = jsonSortField
			}
			pageable.Sort = append(pageable.Sort, sort)
		}
	}
	return pageable
}

//Pagination uses the same principles as the <a href="https://developer.github.com/v3/#pagination">Github API</a>,
//and follow <a href="http://tools.ietf.org/html/rfc5988">RFC 5988 (Link header)</a>.
func (c BaseController) generatePaginationHttpHeaders(page *models.Page) {
	headers := c.Response.Out.Header()
	headers.Set("X-Total-Count", fmt.Sprintf("%d", page.TotalElements))
	link := ""
	baseUrl := c.Request.URL.Path
	if (page.PageNumber + 1 < page.TotalPages) {
		link = "<" + c.generateUri(baseUrl, page.PageNumber + 1, page.PageSize) + ">; rel=\"next\","
	}
	// prev link
	if (page.PageNumber > 0) {
		link += "<" + c.generateUri(baseUrl, page.PageNumber - 1, page.PageSize) + ">; rel=\"prev\","
	}
	// last and first link
	lastPage := int64(0);
	if (page.TotalPages > 0) {
		lastPage = page.TotalPages - 1
	}
	link += "<" + c.generateUri(baseUrl, lastPage, page.PageSize) + ">; rel=\"last\",";
	link += "<" + c.generateUri(baseUrl, 0, page.PageSize) + ">; rel=\"first\"";
	headers.Set("Link", link)
}

func (c BaseController)  generateUri(baseUrl string, page int64, size int64) string {
	params := map[string]string{}
	params["page"] = fmt.Sprintf("%d", page)
	params["size"] = fmt.Sprintf("%d", size)
	return commons.GenerateUri(baseUrl, &params)
}