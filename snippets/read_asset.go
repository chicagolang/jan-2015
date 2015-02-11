// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
  cannonicalName := strings.Replace(name, "\\", "/", -1)
  if f, ok := _bindata[cannonicalName]; ok {
    a, err := f()
    if err != nil {
      return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
    }
    return a.bytes, nil
  }
  return nil, fmt.Errorf("Asset %s not found", name)
}