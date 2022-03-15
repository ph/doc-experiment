package typ

import (
	"crypto/tls"
	"crypto/x509"
	"time"
)

// TLSConfig is the interface used to configure a tcp client or server from a `Config`
//
//  examples:
//     - value: tlsConfigExample
type TLSConfig struct {

	//   description: |
	//     List of allowed SSL/TLS versions.
	//     If the SSL/TLS server supports none of the specified versions, the connection will be dropped during or after the handshake.
	//   values:
	//     - SSLv3
	//     - TLSv1
	//     - TLSv1.0
	//     - TLSv1.1
	//     - TLSv1.2
	//     - TLSv1.3
	//   default: [TLSv1.1, TLSv1.2, TLSv1.3]
	Versions []TLSVersion `yaml:"supported_protocols"`

	//   description: |
	//     Configures the SSL/TLS verification mode used during handshake.
	//   values:
	//     - full
	//     - strict
	//     - certificate
	//     - none
	//   default: full
	Verification TLSVerificationMode `yaml:"verification_mode"`

	//   description: |
	//     List of certificate chains to present to the other side of the connection.
	Certificates []tls.Certificate `yaml:"certificates"`

	//   description: |
	//     List of root certificate authorities used to verify server certificates.
	//     If this setting is empty or unset, TLS may use the system's root CA (not supported on MS Windows).
	RootCAs *x509.CertPool `yaml:"certificate_authorities"`

	//   description: |
	//     List of root certificate authorities to use to verify client certificates.
	//     If this setting is empty or unset, TLS may use the system's root CA set (not supported on MS Windows).
	ClientCAs *x509.CertPool `yaml:"client_certificate_authorities"` //I made this name up

	//   description: |
	//     List of supported cipher suites.
	//     If nil, a default list provided by the implementation is used.
	//
	//     For a list of supported cipher suites, see https://path/to/docs/topic.html 
	CipherSuites []CipherSuite `yaml:"cipher_suites"` // Note sure about valid values here.

	//   description: |
	//     List of elliptic curve types to use in an ECDHE handshake.
	//     If empty, the implementation chooses a default.
	//   values:
	//     - P-256
	//     - P-384
	//     - P-521
	//     - X25519
	CurvePreferences []tls.CurveID `yaml:"curve_types"`

	//   description: |
	//     The type of renegotiation that's supported.
	//     The default is correct for the vast majority of applications.
	//   values:
	//     - never
	//     - once
	//     - freely
	//   default: never
	Renegotiation tls.RenegotiationSupport `yaml:"renegotiation"`

	//   description: |
	//     Controls how certificates from the client are verified.
	//     Does not affect the TCP client.
	//   values:
	//     - none
	//     - optional
	//     - required
	//   default: required
	ClientAuth tls.ClientAuthType `yaml:"verification_mode"`

	//   description: |
	//     The CA certificate pin used to validate the CA used to trust the server certificate.
	//     The pin is a base64 encoded string of the SHA-256 of the certificate.
	//
	//     > Note: Do not use this setting with `verification_mode` set to `none`, or the check will always fail because it will not receive any verified chains.
	CASha256 []string `yaml:"ca_sha256"`

	//   description: |
	//     The HEX-encoded fingerprint of a CA certificate.
	//     If present in the chain, this certificate is added to the list of trusted CAs (root CAs) during the handshake.
	CATrustedFingerprint string `yaml:"ca_trusted _fingerprint"`

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
