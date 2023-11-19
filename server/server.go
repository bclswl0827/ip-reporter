package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func startServer(listen string, options *serverOptions) error {
	r := gin.New()
	r.Use(
		gzip.Gzip(options.Gzip, gzip.WithExcludedPaths([]string{options.APIPrefix})),
		gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			trimmedErr := strings.TrimRight(param.ErrorMessage, "\n")
			return fmt.Sprintf("%s [server] %s %d %s %s %s\n",
				param.TimeStamp.Format("2006/01/02 15:04:05"),
				param.Method, param.StatusCode,
				param.ClientIP, param.Path, trimmedErr,
			)
		}),
	)
	if options.CORS {
		r.Use(AllowCORS([]httpHeader{
			{
				Header: "Access-Control-Allow-Origin",
				Value:  "*",
			}, {
				Header: "Access-Control-Allow-Methods",
				Value:  "POST, OPTIONS, GET",
			}, {
				Header: "Access-Control-Allow-Headers",
				Value:  "Content-Type",
			}, {
				Header: "Access-Control-Expose-Headers",
				Value:  "Content-Length",
			},
		}))
	}
	r.NoRoute(func(c *gin.Context) {
		setErrorResponse(c, http.StatusNotFound)
	})

	r.GET("/", func(c *gin.Context) {
		setHTMLResponse(c, http.StatusOK, INDEX_HTML)
	})
	r.GET("/devices", func(c *gin.Context) {
		for i := 0; i < len(options.DeviceList); i++ {
			timeDiff := (getTimestamp(true) - options.DeviceList[i]["updated_at"].(int64)) / 1000
			if timeDiff > int64(options.DeviceLife) {
				options.DeviceList = append(options.DeviceList[:i], options.DeviceList[i+1:]...)
				i--
			}
		}
		setJSONResponse(c, "设备列表获取成功", options.DeviceList)
	})
	r.POST("/collect", func(c *gin.Context) {
		var req clientRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			setErrorResponse(c, http.StatusBadRequest)
			return
		}
		if req.DeviceTag == "" {
			setErrorResponse(c, http.StatusBadRequest)
			return
		}
		for _, device := range options.DeviceList {
			if device["device_tag"].(string) == req.DeviceTag {
				device["updated_at"] = getTimestamp(true)
				device["ip_address"] = req.IPAddress
				setJSONResponse(c, "IP 地址列表更新成功", nil)
				return
			}
		}
		options.DeviceList = append(options.DeviceList, map[string]any{
			"updated_at": getTimestamp(true),
			"device_tag": req.DeviceTag,
			"ip_address": req.IPAddress,
		})
		setJSONResponse(c, "IP 地址列表上报成功", nil)
	})

	return r.Run(listen)
}
