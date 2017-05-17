package controllers

import (
	"github.com/revel/revel"
	"github.com/elirenato/golangseed/app/repositories"
	"github.com/elirenato/golangseed/app/models"
	"github.com/elirenato/golangseed/app/commons"
	"time"
	"log"
	"fmt"
)

type GroupController struct {
	BaseController
}

var groupRepository = repositories.NewGroupRepository(&Dbm)

func (c GroupController) Create(dto models.Group) revel.Result {
	//validate the basic fields
	//TODO how to handle the fields of the struct (pointer or nullstring)
	// dto.Validate(c.Validation)
	// if c.Validation.HasErrors() {
	// 	log.Println(c.Validation.Errors)
	// 	return c.RenderBadRequest("group.creation")
	// }	
	if dto.Id != nil {
		fmt.Println("Should be an update instead of create to an existent resource")
		return c.RenderInternalServerError()		
	}
	// jsonByte, _ := json.Marshal(dto)
	// jsonString := string(jsonByte)
	// fmt.Printf(" jsonString: %s ", jsonString)
	model := &models.Group{
		Name: commons.InitializeString(c.Params.Get("Name")),
		CreatedDate: commons.InitializeTime(time.Now()),
	}
	err := groupRepository.Persist(Dbm, model)
	if err != nil {
		log.Println(err)
		return c.RenderInternalServerError()
	}
	return c.RenderOK(model)
}

func (c GroupController) Read(id int64) revel.Result {
	model, err := groupRepository.FindOne(id)
	if err != nil {
		return c.RenderInternalServerError()
	}
	if (model == nil) {
		return c.RenderNotFound()
	}
	return c.RenderOK(model)
}

func (c GroupController) Update(dto models.Group) revel.Result {
	//validate the basic fields
	dto.Validate(c.Validation)
	if dto.Id == nil {
		//should be a create operation instead of updated
		return c.RenderInternalServerError()		
	}
	model, err := groupRepository.FindOne(*dto.Id)
	if err != nil {
		return c.RenderInternalServerError()
	}
	if (model == nil) {
		return c.RenderNotFound()
	}
	model.Name = dto.Name
	model.ImageUrl = dto.ImageUrl
	model.LastModifiedDate = commons.InitializeTime(time.Now())
	err = groupRepository.Update(&model)
	if err != nil {
		log.Println(err)
		return c.RenderInternalServerError()
	}
	return c.RenderOK(model)
}

func (c GroupController) Delete(id int64) revel.Result {
	model, err := groupRepository.FindOne(id)
	if err != nil {
		return c.RenderInternalServerError()
	}
	if (model == nil) {
		return c.RenderNotFound()
	}
	err = groupRepository.Delete(&model)
	if err != nil {
		log.Println(err)
		return c.RenderInternalServerError()
	}
	return c.RenderOK(model)
}