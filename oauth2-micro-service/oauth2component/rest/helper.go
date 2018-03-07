package rest

import (
	"errors"
	"net/http"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/RangelReale/osin"
	"github.com/adriendomoison/gobootapi/errorhandling/apihelper"
)

func HandleLoginPage(r *rest, ar *osin.AuthorizeRequest, c *gin.Context) (uint, bool) {
	var errorStatus = false
	var errorList []apihelper.ApiError
	if c.Request.Method == "POST" {
		c.Request.ParseForm()
		if userInfo, err := r.service.AskUserServiceToCheckCredentials(c.Request.Form.Get("username"), c.Request.Form.Get("password"), "password"); err == nil {
			return userInfo.UserId, true
		} else {
			errorStatus = true
			errorList = err.Errors
		}
	}
	c.HTML(http.StatusOK, "authentication.tmpl", gin.H{
		"client_id":     ar.Client.GetId(),
		"authorize_url": c.Request.URL,
		"error_status": errorStatus,
		"errors": errorList,
	})
	return 0, false
}

func DownloadAccessToken(url string, auth *osin.BasicAuth, output map[string]interface{}) error {
	// download access token
	preq, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	if auth != nil {
		preq.SetBasicAuth(auth.Username, auth.Password)
	}

	pclient := &http.Client{}
	presp, err := pclient.Do(preq)
	if err != nil {
		return err
	}
	if presp.StatusCode != 200 {
		return errors.New("invalid status code")
	}

	jdec := json.NewDecoder(presp.Body)
	err = jdec.Decode(&output)
	return err
}