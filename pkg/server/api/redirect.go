package api

import (
	"log"
	"net/http"
	"short-url/pkg/server/service"

	"github.com/gin-gonic/gin"
)

func RedirectOriginUrl(ctx *gin.Context) {
	tinyUrl := ctx.Param("short")

	realUrl, err := service.NewTinyUrlService().GerOriginUrlByTinyUrl(ctx, tinyUrl)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.Redirect(http.StatusFound, realUrl)
}

func GetShortUrl(ctx *gin.Context) {
	url, ok := ctx.GetQuery("url")
	if !ok {
		log.Println("err url")
		return
	}
	realUrl, err := service.NewTinyUrlService().GerTinyUrl(ctx, url)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.String(http.StatusOK, realUrl)
}
