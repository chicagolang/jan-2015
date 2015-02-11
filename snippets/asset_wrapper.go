type _bintree_t struct {
  Func func() (*asset, error)
  Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
  "static": &_bintree_t{nil, map[string]*_bintree_t{
    "index.html": &_bintree_t{static_index_html, map[string]*_bintree_t{
    }},
  }},
}}