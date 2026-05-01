package utils

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, files []string, data gin.H) {

	baseFiles := []string{
		"templates/layout/base.html",
		"templates/layout/sidebar.html",
		"templates/components/loading.html",
	}

	// gabungkan base + page
	allFiles := append(baseFiles, files...)

	tmpl, err := template.ParseFiles(allFiles...)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = tmpl.ExecuteTemplate(c.Writer, "base", data)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
}
