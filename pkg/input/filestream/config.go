package filestream

import "github.com/ph/doc-experiment/pkg/typ"

type Config struct {
	Id         string
	DataStream typ.DataStream
	Paths      []string
	UseOutput  string
}
