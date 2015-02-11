// Returns MIME content type based on filename
func assetContentType(name string) string {
  ext := filepath.Ext(name)
  return mime.TypeByExtension(ext)
}

// Responds with asset data for a given path, or with error if asset is missing
func serveAsset(path string, c *gin.Context) {
  buff, err := Asset(path)

  if err != nil {
    c.String(400, err.Error())
    return
  }

  c.Data(200, assetContentType(path), buff)
}