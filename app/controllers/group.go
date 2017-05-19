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

var groupRepository = repositories.NewGroupRepository(&Dbm)

func (c GroupController) Create() revel.Result {
	dto := models.Group{}
	err := commons.DecodeJsonBody(c.Request.Body, &dto)
	if err != nil {
		log.Println(err)
		return c.RenderBadRequest("error.generic.invalidBody")
	}
	//validate the basic fields
	dto.Validate(c.Validation)
	if c.Validation.HasErrors() {
		log.Println(c.Validation.Errors)
		return c.RenderBadRequest("error.generic.validation")
	}
	if dto.Id.Valid {
		log.Println("Should be an update instead of create to an existent resource")
		return c.RenderInternalServerError()		
	}
	model := &models.Group{
		Name: dto.Name,
		CreatedDate: null.NewTime(time.Now(), true),
	}
	err = groupRepository.Persist(Dbm, model)
	if err != nil {
		// TODO improve the error handler for user already exists
		// pqError := err.(*pq.Error)
		// pqError.Constraint is empty, why?
		if strings.Contains(err.Error(), models.GroupUniqueName) {
			log.Println(fmt.Sprintf("A record already exists with the name %s", model.Name.String))
			return c.RenderBadRequest("error.group.alreadyExists")
		}
		return c.RenderInternalServerError()
	}
	return c.RenderOK(model)
}

func (c GroupController) Read(id int64) revel.Result {
	fmt.Println("### ");
	fmt.Println(id);
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
	// if dto.Id == nil {
	// 	//should be a create operation instead of updated
	// 	return c.RenderInternalServerError()		
	// }
	// model, err := groupRepository.FindOne(*dto.Id)
	// if err != nil {
	// 	return c.RenderInternalServerError()
	// }
	// if (model == nil) {
	// 	return c.RenderNotFound()
	// }
	// model.Name = dto.Name
	// model.ImageUrl = dto.ImageUrl
	// model.LastModifiedDate = commons.InitializeTime(time.Now())
	// err = groupRepository.Update(&model)
	// if err != nil {
	// 	log.Println(err)
	// 	return c.RenderInternalServerError()
	// }
	return c.RenderOK(dto)
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