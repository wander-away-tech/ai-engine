package utils

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// This function allows to read a request body into a variable.
// Pass in the pointer to the variable as the argument `v`.
// If an error occurs, it responds to the request with a BAD REQUEST BODY error,
// and returns the error from the function.
func GetJSONBody(ctx *gin.Context, v any) error {
	bodyBytes, err := io.ReadAll(ctx.Request.Body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request, Unable to read Request Body.",
			"error":   "ERR_BAD_REQUEST_BODY",
		})
		return err
	}

	err = json.Unmarshal(bodyBytes, v)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request, Unable to read Request Body.",
			"error":   "ERR_BAD_REQUEST_BODY",
		})
		return err
	}

	return nil
}
