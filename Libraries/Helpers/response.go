package Helpers

import (
	"fmt"
	"math"
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

type setResponseWithMeta struct {
	Status     string      `json:"status"`
	Data       interface{} `json:"data"`
	Code       int         `json:"code"`
	AccessTime string      `json:"accessTime"`
	Meta       *Meta       `json:"meta,omitempty"`
	Links      *Links      `json:"links,omitempty"`
}

type Meta struct {
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
	Total   int `json:"total"`
}

type Links struct {
	First string `json:"first,omitempty"`
	Last  string `json:"last"`
	Next  string `json:"next"`
}

func HttpResponseSuccessWithPagination(g *gin.Context, data interface{}, total int) {
	location, _ := time.LoadLocation(TimeLocation)

	page, perPage := getPaginationFromContext(g)
	lastPage := int(math.Ceil(float64(total) / float64(perPage)))

	meta := &Meta{
		Pagination: Pagination{
			Page:    page,
			PerPage: perPage,
			Total:   total,
		},
	}

	baseURL := g.Request.URL.Path

	links := &Links{
		Last: fmt.Sprintf("%s?page=%d&per_page=%d", baseURL, lastPage, perPage),
	}

	if page < lastPage {
		links.Next = fmt.Sprintf("%s?page=%d&per_page=%d", baseURL, page+1, perPage)
	}

	response := setResponseWithMeta{
		Status:     http.StatusText(http.StatusOK),
		Code:       http.StatusOK,
		Data:       data,
		Meta:       meta,
		Links:      links,
		AccessTime: time.Now().In(location).Format("02-01-2006 15:04:05"),
	}

	g.Header("Content-Type", ContentTypeJSON)
	g.JSON(http.StatusOK, response)
}

func getPaginationFromContext(g *gin.Context) (page, perPage int) {
	page, _ = strconv.Atoi(g.DefaultQuery("page", "1"))
	perPage, _ = strconv.Atoi(g.DefaultQuery("per_page", "10"))

	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 10
	}
	return
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
