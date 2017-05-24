package controllers

import (
	"github.com/revel/revel"
	"github.com/elirenato/golangseed/app/repositories"
	"github.com/elirenato/golangseed/app/models"
	"time"
	"log"
	"github.com/elirenato/null"
	"github.com/elirenato/golangseed/app/commons"
	"strings"
	"fmt"
)

type GroupController struct {
	BaseController
}

var groupRepository = repositories.NewGroupRepository(&commons.Dbm)

func (c GroupController) treatAndReturnErrorOnSave(err error, model *models.Group) revel.Result {
	if strings.Contains(err.Error(), models.GroupUniqueName) {
		log.Println(fmt.Sprintf("A record already exists with the name %s", model.Name.String))
		return c.RenderBadRequest("error.group.alreadyExists")
	}
	log.Println(err)
	return c.RenderInternalServerError("error.internalServerError")
}

func (c GroupController) Create() revel.Result {
	dto := models.Group{}
	result := c.decodeAndValidateRequest(&dto)
	if result != nil {
		return result
	}
	if dto.Id.Valid {
		return c.RenderBadRequest("error.generic.invalidCreateId")
	}
	model := &models.Group{
		Name: dto.Name,
		CreatedDate: null.NewTime(time.Now(), true),
	}
	err := groupRepository.Persist(model)
	if err != nil {
		return c.treatAndReturnErrorOnSave(err, model)
	}
	return c.RenderOK(model)
}

func (c GroupController) Update() revel.Result {
	dto := models.Group{}
	result := c.decodeAndValidateRequest(&dto)
	if result != nil {
		return result
	}
	if !dto.Id.Valid {
		return c.RenderBadRequest("error.generic.invalidUpdateId")
	}
	model, err := groupRepository.FindOne(dto.Id.Int64)
	if err != nil {
		log.Println(err)
		return c.RenderInternalServerError("error.internalServerError")
	}
	if model == nil {
		return c.RenderNotFound("error.generic.notFound")
	}
	model.Name = dto.Name
	model.LastModifiedDate = null.NewTime(time.Now(), true)
	err = groupRepository.Update(model)
	if err != nil {
		return c.treatAndReturnErrorOnSave(err, model)
	}
	return c.RenderOK(model)	
}

func (c GroupController) Read(id int64) revel.Result {
	model, err := groupRepository.FindOne(id)
	if err != nil {
		log.Println(err)
		return c.RenderInternalServerError("error.internalServerError")
	}
	if (model == nil) {
		return c.RenderNotFound("error.generic.notFound")
	}
	return c.RenderOK(model)
}

func (c GroupController) List() revel.Result {
	pageable := c.parsePageableRequest("Id")
	page, err := groupRepository.ListAll(pageable)
	if err != nil {
		log.Println(err)
		return c.RenderInternalServerError("error.internalServerError")
	}
	c.generatePaginationHttpHeaders(page)
	return c.RenderOK(page.Items)
}

func (c GroupController) Delete(id int64) revel.Result {
	model, err := groupRepository.FindOne(id)
	if err != nil {
		log.Println(err)
		return c.RenderInternalServerError("error.internalServerError")
	}
	if (model == nil) {
		return c.RenderNotFound("error.generic.notFound")
	}
	err = groupRepository.Delete(model)
	if err != nil {
		log.Println(err)
		return c.RenderInternalServerError("error.internalServerError")
	}
	return c.RenderOK(model)
}