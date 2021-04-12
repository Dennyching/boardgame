// handlers.article.go

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Myform struct {
	Tags []string `form:"tags[]"`
}

func showgame(c *gin.Context) {
	// Obtain the POSTed title and content values
	Select := c.Request.URL.Query().Get("Tag")

	log.Println(Select)
	if a, err := getSelect(Select); err == nil {
		// If the article is created successfully, show success message
		render(c, gin.H{
			"payload": a}, "select-successful.html")
	} else {
		// if there was an error while creating the article, abort with an error
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
