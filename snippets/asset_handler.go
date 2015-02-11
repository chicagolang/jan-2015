// Serves just the home page
func HomeHandler(c *gin.Context) {
	serveAsset("static/index.html", c)
}

// Serves any file under ./static directory
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