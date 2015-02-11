package main

import (
	"mime"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func assetContentType(name string) string {
	ext := filepath.Ext(name)
	return mime.TypeByExtension(ext)
}

func serveAsset(path string, c *gin.Context) {
	buff, err := Asset(path)

	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.Data(200, assetContentType(path), buff)
}

func HomeHandler(c *gin.Context) {
	serveAsset("static/index.html", c)
}

func AssetHandler(c *gin.Context) {
	path := "static" + c.Params.ByName("path")
	serveAsset(path, c)
}

func main() {
	router := gin.Default()
	router.GET("/", HomeHandler)
	router.GET("/static/*path", AssetHandler)
	router.Run(":5000")
}
