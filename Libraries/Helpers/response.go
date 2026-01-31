package Helpers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	TimeLocation    string = "Asia/Jakarta"
	ImageJpeg       string = "image/jpeg"
	ContentTypeJSON string = "application/json"
)

type setResponse struct {
	Status     string      `json:"status"`
	Data       interface{} `json:"data"`
	Code       int         `json:"code"`
	AccessTime string      `json:"accessTime"`
}

func HttpResponseSuccess(g *gin.Context, data interface{}) {
	location, _ := time.LoadLocation(TimeLocation)
	responseData := setResponse{
		Code:       http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Data:       data,
		AccessTime: time.Now().In(location).Format("02-01-2006 15:04:05")}

	g.Header("Content-Type", ContentTypeJSON)
	g.JSON(http.StatusOK, responseData)
}

func HttpResponseError(g *gin.Context, data interface{}, code int) {
	location, _ := time.LoadLocation(TimeLocation)
	responseData := setResponse{
		Code:       code,
		Status:     http.StatusText(code),
		AccessTime: time.Now().In(location).Format("02-01-2006 15:04:05"),
		Data:       data,
	}
	g.Header("Content-Type", ContentTypeJSON)
	g.JSON(code, responseData)
}
func HttpResponsePDF(g *gin.Context, dataPdfByte []byte) {
	g.Header("Content-Type", "application/pdf")
	g.Header("Content-Length", strconv.Itoa(len(dataPdfByte)))
	g.Data(http.StatusOK, "application/pdf", dataPdfByte)
}

func HttpResponsedDownloadPDF(g *gin.Context, dataPdfByte []byte, fileName string) {
	g.Header("Content-Type", "application/pdf")
	g.Header("Content-Disposition", "attachment; filename="+fileName)
	g.Header("Content-Length", strconv.Itoa(len(dataPdfByte)))
	g.Data(http.StatusOK, "application/pdf", dataPdfByte)
}

func HttpResponsedDownloadFile(g *gin.Context, dataPdfByte []byte, fileName string) {
	g.Header("Content-Disposition", "attachment; filename="+fileName)
	g.Header("Content-Length", strconv.Itoa(len(dataPdfByte)))
	g.Data(http.StatusOK, "application/octet-stream", dataPdfByte)
}

func HttpPreviewImages(g *gin.Context, dataImg []byte) {
	g.Header("Content-Type", ImageJpeg)
	g.Header("Content-Length", strconv.Itoa(len(dataImg)))
	g.Data(http.StatusOK, ImageJpeg, dataImg)
}
