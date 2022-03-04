package typ

import (
	"crypto/tls"
	"crypto/x509"
	"time"
)

// TLSConfig is the interface used to configure a tcp client or server from a `Config`
type TLSConfig struct {

	// List of allowed SSL/TLS protocol versions. Connections might be dropped
	// after handshake succeeded, if TLS version in use is not listed.
	Versions []TLSVersion

	// Configure SSL/TLS verification mode used during handshake. By default
	// VerifyFull will be used.
	Verification TLSVerificationMode

	// List of certificate chains to present to the other side of the
	// connection.
	Certificates []tls.Certificate

	// Set of root certificate authorities use to verify server certificates.
	// If RootCAs is nil, TLS might use the system its root CA set (not supported
	// on MS Windows).
	RootCAs *x509.CertPool

	// Set of root certificate authorities use to verify client certificates.
	// If ClientCAs is nil, TLS might use the system its root CA set (not supported
	// on MS Windows).
	ClientCAs *x509.CertPool

	// List of supported cipher suites. If nil, a default list provided by the
	// implementation will be used.
	CipherSuites []CipherSuite

	// Types of elliptic curves that will be used in an ECDHE handshake. If empty,
	// the implementation will choose a default.
	CurvePreferences []tls.CurveID

	// Renegotiation controls what types of renegotiation are supported.
	// The default, never, is correct for the vast majority of applications.
	Renegotiation tls.RenegotiationSupport

	// ClientAuth controls how we want to verify certificate from a client, `none`, `optional` and
	// `required`, default to required. Do not affect TCP client.
	ClientAuth tls.ClientAuthType

	// CASha256 is the CA certificate pin, this is used to validate the CA that will be used to trust
	// the server certificate.
	CASha256 []string

	// CATrustedFingerprint is the HEX encoded fingerprint of a CA certificate. If present in the chain
	// this certificate will be added to the list of trusted CAs (RootCAs) during the handshake.
	CATrustedFingerprint string

	// time returns the current time as the number of seconds since the epoch.
	// If time is nil, TLS uses time.Now.
	time func() time.Time
}

type TLSVersion uint16

type TLSVerificationMode uint8

// Constants of the supported verification mode.
const (
	VerifyFull TLSVerificationMode = iota
	VerifyNone
	VerifyCertificate
	VerifyStrict
)

type CipherSuite uint16
