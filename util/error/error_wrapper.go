package err

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type WrapperHandle func(c *gin.Context) (interface{}, error)

func ErrorWrapper(handle WrapperHandle) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := handle(c)
		if err != nil {
			apiError := err.(ApiError)
			c.JSON(apiError.Status, apiError)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
