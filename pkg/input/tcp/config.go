package tcp

import "github.com/ph/doc-experiment/pkg/typ"

// Defines a tcp input configuration
type Config struct {
	
	//   description: |
	//     The TCP port to listen on for event streams.
	Port int `yaml:"port"`

	//   description: |
	//     The host to listen on for event streams.
	Host string `yaml:"host"`

	//   description: |
	//     Configuration options for SSL parameters like the certificate, key, and the certificate authorities to use.
	TLS  typ.TLSConfig `yaml:"ssl"`
}
