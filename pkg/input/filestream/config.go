package filestream

import "github.com/ph/doc-experiment/pkg/typ"

// Defines a filestream input configuration.
type Config struct { `yaml:"???"`
	
	//   description: |
	//     The ID of blah blah blah.
	Id         string `yaml:"id"`

	//   description: |
	//     The data stream to send events to.
	DataStream typ.DataStream `yaml:"data_stream"`
	
	//   description: |
	//     A list of glob-based paths that will be crawled and fetched.
	Paths      []string `yaml:"paths"`

	//   description: |
	//     The output to use for this filestream output.
	UseOutput  string `yaml:"use_output"`
}
