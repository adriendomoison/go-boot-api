/*
	apihelper automate construction of http response
*/
package apihelper

import (
	"net/http"
	"github.com/adriendomoison/go-boot-api/tool"
	"github.com/adriendomoison/go-boot-api/apicore/helpers/servicehelper"
	"gopkg.in/go-playground/validator.v8"
)

// Interface for all API error messages
type apiError interface {
}

// Error is the default error message structure the API returns
type Error struct {
	apiError apiError
	Param    string `json:"param"`
	Detail   string `json:"detail"`
	Message  string `json:"message"`
}

// ApiErrors carry the list of errors returned by the API from a request
type ApiErrors struct {
	Errors []apiError
}

// BuildRequestError build a usable JSON error object from an error string generated by the structure validator
func BuildRequestError(err error) (int, ApiErrors) {
	var apiErrors ApiErrors
	switch err.(type) {
	case validator.ValidationErrors:
		for _, v := range err.(validator.ValidationErrors) {
			var validationError Error
			validationError.Param = tool.ToSnakeCase(v.Field)
			validationError.Detail = "Field validation for " + tool.ToSnakeCase(v.Field) + " failed on the " + v.Tag + " tag."
			if v.Tag == "required" {
				validationError.Message = "This field is required"
			}
			if v.Tag == "email" {
				validationError.Message = "Invalid email address. Valid e-mail can contain only latin letters, numbers, '@' and '.'"
			}
			if v.Tag == "url" {
				validationError.Message = "Invalid URL address. Valid URL start with http:// or https://"
			}
			apiErrors.Errors = append(apiErrors.Errors, validationError)
		}
		return http.StatusBadRequest, apiErrors
	default:
		apiErrors.Errors = append(apiErrors.Errors, Error{
			Detail: err.Error(),
		})
		return http.StatusBadRequest, apiErrors
	}
}

// BuildResponseError apply the right status to the http response and build the error JSON object
func BuildResponseError(err *servicehelper.Error) (int, ApiErrors) {
	var apiErrors ApiErrors
	apiErrors.Errors = append(apiErrors.Errors, Error{
		Detail:  err.Detail.Error(),
		Message: err.Message,
		Param:   err.Param,
	})
	if err.Code == servicehelper.BadRequest {
		return http.StatusBadRequest, apiErrors
	} else if err.Code == servicehelper.AlreadyExist {
		return http.StatusConflict, apiErrors
	} else if err.Code == servicehelper.NotFound {
		return http.StatusNotFound, apiErrors
	}
	return http.StatusInternalServerError, apiErrors
}
