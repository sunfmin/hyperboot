package react

import (
	"bytes"
)

func MainLess(data ReactData) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n@import \"node_modules/bootstrap/less/bootstrap.less\";\n\n@font-family-sans-serif: \"Open Sans\",Arial,\"Hiragino Sans GB\",\"Microsoft YaHei\",\"微软雅黑\",\"STHeiti\",\"WenQuanYi Micro Hei\", SimSun, sans-serif;\n\n@line-height-base: 1.8;\n@font-size-base: 16px;\n@blockquote-font-size:        (@font-size-base * 1.05);\n@headings-line-height:    1.2;\n\n\n@font-size-h1:            floor((@font-size-base * 1.6));\n@font-size-h2:            floor((@font-size-base * 1.3));\n@font-size-h3:            ceil((@font-size-base * 1.2));\n@font-size-h4:            ceil((@font-size-base * 1.1));\n@font-size-h5:            @font-size-base;\n@font-size-h6:            ceil((@font-size-base * 0.85));\n\n@headings-font-weight:    600;\n// Navbar links\n@navbar-default-link-color:                #6f5499;\n@navbar-default-link-hover-color:          #6f5499;\n@navbar-default-link-hover-bg:             transparent;\n@navbar-default-link-active-color:         #6f5499;\n@navbar-default-link-active-bg:            darken(@navbar-default-bg, 6.5%);\n@navbar-default-link-disabled-color:       #ccc;\n@navbar-default-link-disabled-bg:          transparent;\n\n// Navbar brand label\n@navbar-default-brand-color:               @navbar-default-link-color;\n@navbar-default-brand-hover-color:         darken(@navbar-default-brand-color, 10%);\n@navbar-default-brand-hover-bg:            transparent;")

	return _buffer.String()
}
