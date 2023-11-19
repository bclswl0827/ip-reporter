package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

func setHTMLResponse(c *gin.Context, code int, html string) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(code, html)
}

func setJSONResponse(c *gin.Context, message string, data any) error {
	currentTime := time.Now().UTC().Format(time.RFC3339)
	currentPath := c.Request.URL.Path

	response := httpResponse{
		Error:   false,
		Path:    currentPath,
		Time:    currentTime,
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	}

	jsonBytes, err := jsoniter.Marshal(response)
	if err != nil {
		setErrorResponse(c, http.StatusInternalServerError)
		return err
	}

	c.Data(http.StatusOK, "application/json", jsonBytes)
	return nil
}

func setErrorResponse(c *gin.Context, code int) {
	currentTime := time.Now().UTC().Format(time.RFC3339)
	currentPath := c.Request.URL.Path

	switch code {
	case 400:
		c.JSON(http.StatusBadRequest, httpResponse{
			Error:   true,
			Path:    currentPath,
			Time:    currentTime,
			Status:  http.StatusBadRequest,
			Message: "Unable to process this request",
		})
	case 401:
		c.JSON(http.StatusUnauthorized, httpResponse{
			Error:   true,
			Path:    currentPath,
			Time:    currentTime,
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized access",
		})
	case 403:
		c.JSON(http.StatusForbidden, httpResponse{
			Error:   true,
			Path:    currentPath,
			Time:    currentTime,
			Status:  http.StatusForbidden,
			Message: "Request is forbidden",
		})
	case 404:
		c.JSON(http.StatusNotFound, httpResponse{
			Error:   true,
			Path:    currentPath,
			Time:    currentTime,
			Status:  http.StatusNotFound,
			Message: "Could not find this resource",
		})
	case 405:
		c.JSON(http.StatusMethodNotAllowed, httpResponse{
			Error:   true,
			Path:    currentPath,
			Time:    currentTime,
			Status:  http.StatusMethodNotAllowed,
			Message: "Request method is not allowed",
		})
	case 410:
		c.JSON(http.StatusGone, httpResponse{
			Error:   true,
			Path:    currentPath,
			Time:    currentTime,
			Status:  http.StatusGone,
			Message: "The requested resource is not available",
		})
	case 413:
		c.JSON(http.StatusRequestEntityTooLarge, httpResponse{
			Error:   true,
			Path:    currentPath,
			Time:    currentTime,
			Status:  http.StatusMethodNotAllowed,
			Message: "Request exceeds data limit",
		})
	case 429:
		c.JSON(http.StatusTooManyRequests, httpResponse{
			Error:   true,
			Path:    currentPath,
			Time:    currentTime,
			Status:  http.StatusTooManyRequests,
			Message: "Too many requests",
		})
	case 500:
		c.JSON(http.StatusInternalServerError, httpResponse{
			Error:   true,
			Path:    currentPath,
			Time:    currentTime,
			Status:  http.StatusInternalServerError,
			Message: "Server internal error",
		})
	case 502:
		c.JSON(http.StatusBadGateway, httpResponse{
			Error:   true,
			Path:    currentPath,
			Time:    currentTime,
			Status:  http.StatusBadGateway,
			Message: "Server gateway error",
		})
	case 503:
		c.JSON(http.StatusServiceUnavailable, httpResponse{
			Error:   true,
			Path:    currentPath,
			Time:    currentTime,
			Status:  http.StatusServiceUnavailable,
			Message: "Service is currently unavailable",
		})
	case 504:
		c.JSON(http.StatusGatewayTimeout, httpResponse{
			Error:   true,
			Path:    currentPath,
			Time:    currentTime,
			Status:  http.StatusGatewayTimeout,
			Message: "Server gateway timeout",
		})
	default:
		c.JSON(http.StatusSeeOther, httpResponse{
			Error:   true,
			Path:    currentPath,
			Time:    currentTime,
			Status:  http.StatusSeeOther,
			Message: "Unknown error occurred",
		})
	}

	c.Abort()
}
