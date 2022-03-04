package tcp

import "github.com/ph/doc-experiment/pkg/typ"

type Config struct {
	Port int
	Host string
	TLS  typ.TLSConfig
}
