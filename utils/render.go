package utils

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, files []string, data gin.H) {

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = tmpl.ExecuteTemplate(c.Writer, "base", data)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
}