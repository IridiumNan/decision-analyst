package web

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type testResp struct {
	Data any
}

func NewTestResp(data any) *testResp {
	return &testResp{
		Data: data,
	}
}

// handlerCreate handle request for creating a new event
func handlerCreate(c *gin.Context) {
	c.JSON(http.StatusOK, NewTestResp("hello from create endpoint"))
}

// handlerFetch query specific event by id then return all details
// call [models.Database.Fetch](id)
func handlerFetch(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(strings.Trim(idStr, "/ "))
	if err != nil {
		slog.Error("error while parsing id from client", "err", err, "idStr", idStr)
		return
	}

	c.JSON(http.StatusOK, NewTestResp(fmt.Sprintf("receive request for id = %d", id)))
}

// handlerList return a list events without [models.Event.Source]
func handlerList(c *gin.Context) {
	c.JSON(http.StatusOK, NewTestResp("hello from list endpoint"))
}

func handlerDelete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(strings.Trim(idStr, "/ "))
	if err != nil {
		slog.Error("error while parsing id from client", "err", err, "idStr", id)
		return
	}
	c.JSON(http.StatusOK, NewTestResp(fmt.Sprintf("deleting an event whose id = %d", id)))
}

// handlerUpdate handle request for updating event field
func handlerUpdate(c *gin.Context) {
	idStr := c.Param("id")

	c.JSON(http.StatusOK, NewTestResp(fmt.Sprintf("idStr: %s", idStr)))
}
