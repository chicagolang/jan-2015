func static_index_html() (*asset, error) {
  path := "/Users/sosedoff/example_app/static/index.html"
  name := "static/index.html"
  bytes, err := bindata_read(path, name)
  if err != nil {
    return nil, err
  }
  fi, err := os.Stat(path)
  if err != nil {
    err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
  }
  a := &asset{bytes: bytes, info: fi}
  return a, err
}