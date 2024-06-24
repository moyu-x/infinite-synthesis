package generate

import (
	_ "embed"
)

//go:embed template.go.tpl
var tpl string

type service struct {
	Name     string
	FullName string
	FilePath string

	Methods   []*method
	MethodSet map[string]*method
}

type method struct {
	Name     string
	Num      int
	Request  string
	Response string

	Path         string
	Method       string
	Body         string
	ResponseBody string
}
