package rest

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/adriendomoison/go-boot-api/user/service/model"
	"github.com/adriendomoison/go-boot-api/user/rest/jsonmodel"
	"github.com/adriendomoison/go-boot-api/apicore/helpers/apihelper"
)

// Make sure the interface is implemented correctly
var _ jsonmodel.Interface = (*rest)(nil)

// Implement interface
type rest struct {
	service model.Interface
}

// New return a new rest instance
func New(service model.Interface) *rest {
	return &rest{service}
}

// Post allows to access the service to create a user
func (r *rest) Post(c *gin.Context) {
	var reqDTO jsonmodel.RequestDTO
	if err := c.BindJSON(&reqDTO); err != nil {
		c.JSON(apihelper.BuildRequestError(err))
	} else {
		if resDTO, err := r.service.Add(reqDTO); err != nil {
			c.JSON(apihelper.BuildResponseError(err))
		} else {
			c.JSON(http.StatusCreated, resDTO)
		}
	}
}

// Get allows to access the service to retrieve a user when sending its email
func (r *rest) Get(c *gin.Context) {
	if resDTO, err := r.service.Retrieve(c.Param("email")); err != nil {
		c.JSON(apihelper.BuildResponseError(err))
	} else {
		c.JSON(http.StatusOK, resDTO)
	}
}

// PutEmail allows to access the service to update the email of a user
func (r *rest) PutEmail(c *gin.Context) {
	var reqDTO jsonmodel.RequestDTOPutEmail
	if err := c.BindJSON(&reqDTO); err != nil {
		c.JSON(apihelper.BuildRequestError(err))
	} else {
		if resDTO, err := r.service.EditEmail(reqDTO); err != nil {
			c.JSON(apihelper.BuildResponseError(err))
		} else {
			c.JSON(http.StatusOK, resDTO)
		}
	}
}

// PutPassword allows to access the service to update the password of a user
func (r *rest) PutPassword(c *gin.Context) {
	var reqDTO jsonmodel.RequestDTOPutPassword
	if err := c.BindJSON(&reqDTO); err != nil {
		c.JSON(apihelper.BuildRequestError(err))
	} else {
		if reqDTO, err := r.service.EditPassword(reqDTO); err != nil {
			c.JSON(apihelper.BuildResponseError(err))
		} else {
			c.JSON(http.StatusOK, reqDTO)
		}
	}
}

// Delete allows to access the service to remove a user from the records
func (r *rest) Delete(c *gin.Context) {
	if err := r.service.Remove(c.Param("email")); err != nil {
		c.JSON(apihelper.BuildResponseError(err))
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "user has been deleted successfully"})
	}
}
