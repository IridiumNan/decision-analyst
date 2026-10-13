package web

import "github.com/gin-gonic/gin"

func NewRouter() (router *gin.Engine) {
	r := gin.Default()

	// endpoint for creating a new event
	r.POST("/event/core/create", handlerCreate)

	// for fech event detail information
	r.GET("/event/core/fetch/:id", handlerFetch)

	r.GET("/event/core/list", handlerList)

	r.DELETE("/event/core/delete/:id", handlerDelete)

	// for evnet modified actions
	// including soft-delelte, update event info
	r.PATCH("/event/core/update/:id", handlerUpdate)

	// TODO:
	// for plugin resigter endpoint
	// r.GET("/event/plugin/:id/*action", handlers ...gin.handlerFunc)

	return r
}
