package internal

import (
	"context"
	"github.com/celt237/gin-enhance.demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiHandlerImpl struct{}

func (a *ApiHandlerImpl) WrapContext(c *gin.Context) context.Context {
	return c.Request.Context()
}

func (a *ApiHandlerImpl) Success(c *gin.Context, contentType string, data interface{}) {
	result := model.ResultGeneric[interface{}]{
		Code:    200,
		Data:    data,
		Message: "Success",
		Time:    0,                      // You might want to implement a timing mechanism
		TraceId: c.GetString("TraceId"), // Assuming you set TraceId in the context
	}
	c.JSON(http.StatusOK, result)
}

func (a *ApiHandlerImpl) CodeError(c *gin.Context, produceType string, data interface{}, code int, err error) {
	result := model.ResultGeneric[interface{}]{
		Code:    code,
		Data:    data,
		Message: err.Error(),
		Time:    0,                      // You might want to implement a timing mechanism
		TraceId: c.GetString("TraceId"), // Assuming you set TraceId in the context
	}
	c.JSON(http.StatusInternalServerError, result)
}

func (a *ApiHandlerImpl) Error(c *gin.Context, contentType string, data interface{}, err error) {
	result := model.ResultGeneric[interface{}]{
		Code:    500,
		Data:    data,
		Message: err.Error(),
		Time:    0,                      // You might want to implement a timing mechanism
		TraceId: c.GetString("TraceId"), // Assuming you set TraceId in the context
	}
	c.JSON(http.StatusInternalServerError, result)
}

func (a *ApiHandlerImpl) HandleCustomerAnnotation(c *gin.Context, name string, params ...string) error {
	// Implement custom annotation handling if needed
	return nil
}
