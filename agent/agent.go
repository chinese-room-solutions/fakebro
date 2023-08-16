package agent

import (
	"context"
	"net"
	"net/http"
	"time"

	tls "github.com/refraction-networking/utls"
)

var BaseHeaders = map[string]map[string]string{
	"firefox": {
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
		"Accept-Language":           "en-US,en;q=0.5",
		"DNT":                       "1",
		"Upgrade-Insecure-Requests": "1",
		"Connection":                "keep-alive",
		"Sec-Fetch-Dest":            "document",
		"Sec-Fetch-Mode":            "navigate",
		"Sec-Fetch-Site":            "none",
		"Sec-Fetch-User":            "?1",
	},
}

var BaseTLSConfigs = map[string]*tls.ClientHelloSpec{}

// Predefined list of user agents.
// Borrowed from: https://github.com/refraction-networking/utls/blob/8199306255caf0d870f69cb36f6b440b33dbf7c5/u_parrots.go
var Agents = []*Agent{
	{
		Name:    "firefox",
		Version: 99,
		Headers: map[string]string{
			"Host":                      "indeed.com",
			"User-Agent":                "Mozilla/5.0 (X11; Linux x86_64; rv:99.0) Gecko/20100101 Firefox/99.0",
			"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
			"Accept-Language":           "en-US,en;q=0.5",
			"DNT":                       "1",
			"Upgrade-Insecure-Requests": "1",
			"Connection":                "keep-alive",
			"Sec-Fetch-Dest":            "document",
			"Sec-Fetch-Mode":            "navigate",
			"Sec-Fetch-Site":            "none",
			"Sec-Fetch-User":            "?1",
		},
		TLSConfig: &tls.ClientHelloSpec{
			TLSVersMin: tls.VersionTLS10,
			TLSVersMax: tls.VersionTLS13,
			CipherSuites: []uint16{
				tls.TLS_AES_128_GCM_SHA256,
				tls.TLS_CHACHA20_POLY1305_SHA256,
				tls.TLS_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
			},
			CompressionMethods: []byte{
				0x0,
			},
			Extensions: []tls.TLSExtension{
				&tls.SNIExtension{},                  //server_name
				&tls.ExtendedMasterSecretExtension{}, //extended_master_secret
				&tls.RenegotiationInfoExtension{Renegotiation: tls.RenegotiateOnceAsClient}, //extensionRenegotiationInfo
				&tls.SupportedCurvesExtension{
					Curves: []tls.CurveID{ //supported_groups
						tls.X25519,
						tls.CurveP256,
						tls.CurveP384,
						tls.CurveP521,
						tls.CurveID(tls.FakeFFDHE2048),
						tls.CurveID(tls.FakeFFDHE3072),
					},
				},
				&tls.SupportedPointsExtension{
					SupportedPoints: []byte{ //ec_point_formats
						0x0,
					},
				},
				&tls.SessionTicketExtension{},
				&tls.ALPNExtension{AlpnProtocols: []string{"h2", "http/1.1"}}, //application_layer_protocol_negotiation
				&tls.StatusRequestExtension{},
				&tls.FakeDelegatedCredentialsExtension{
					SupportedSignatureAlgorithms: []tls.SignatureScheme{ //signature_algorithms
						tls.ECDSAWithP256AndSHA256,
						tls.ECDSAWithP384AndSHA384,
						tls.ECDSAWithP521AndSHA512,
						tls.ECDSAWithSHA1,
					},
				},
				&tls.KeyShareExtension{
					KeyShares: []tls.KeyShare{
						{Group: tls.X25519},
						{Group: tls.CurveP256}, //key_share
					},
				},
				&tls.SupportedVersionsExtension{
					Versions: []uint16{
						tls.VersionTLS13, //supported_versions
						tls.VersionTLS12,
						tls.VersionTLS11,
						tls.VersionTLS10,
					},
				},
				&tls.SignatureAlgorithmsExtension{
					SupportedSignatureAlgorithms: []tls.SignatureScheme{ //signature_algorithms
						tls.ECDSAWithP256AndSHA256,
						tls.ECDSAWithP384AndSHA384,
						tls.ECDSAWithP521AndSHA512,
						tls.PSSWithSHA256,
						tls.PSSWithSHA384,
						tls.PSSWithSHA512,
						tls.PKCS1WithSHA256,
						tls.PKCS1WithSHA384,
						tls.PKCS1WithSHA512,
						tls.ECDSAWithSHA1,
						tls.PKCS1WithSHA1,
					},
				},
				&tls.PSKKeyExchangeModesExtension{
					Modes: []uint8{ //psk_key_exchange_modes
						tls.PskModeDHE,
					},
				},
				&tls.FakeRecordSizeLimitExtension{Limit: 0x4001},                 //record_size_limit
				&tls.UtlsPaddingExtension{GetPaddingLen: tls.BoringPaddingStyle}, //padding
			},
		},
	},
	{
		Name:    "firefox",
		Version: 102,
		Headers: map[string]string{
			"Host":                      "indeed.com",
			"User-Agent":                "Mozilla/5.0 (X11; Linux x86_64; rv:102.0) Gecko/20100101 Firefox/102.0",
			"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
			"Accept-Language":           "en-US,en;q=0.5",
			"DNT":                       "1",
			"Upgrade-Insecure-Requests": "1",
			"Connection":                "keep-alive",
			"Sec-Fetch-Dest":            "document",
			"Sec-Fetch-Mode":            "navigate",
			"Sec-Fetch-Site":            "none",
			"Sec-Fetch-User":            "?1",
		},
		TLSConfig: &tls.ClientHelloSpec{
			TLSVersMin: tls.VersionTLS10,
			TLSVersMax: tls.VersionTLS13,
			CipherSuites: []uint16{
				tls.TLS_AES_128_GCM_SHA256,
				tls.TLS_CHACHA20_POLY1305_SHA256,
				tls.TLS_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
			CompressionMethods: []byte{
				0,
			},
			Extensions: []tls.TLSExtension{
				&tls.SNIExtension{},                  //server_name
				&tls.ExtendedMasterSecretExtension{}, //extended_master_secret
				&tls.RenegotiationInfoExtension{Renegotiation: tls.RenegotiateOnceAsClient}, //extensionRenegotiationInfo
				&tls.SupportedCurvesExtension{
					Curves: []tls.CurveID{ //supported_groups
						tls.X25519,
						tls.CurveP256,
						tls.CurveP384,
						tls.CurveP521,
						tls.CurveID(tls.FakeFFDHE2048),
						tls.CurveID(tls.FakeFFDHE3072),
					},
				},
				&tls.SupportedPointsExtension{
					SupportedPoints: []byte{ //ec_point_formats
						0,
					},
				},
				&tls.SessionTicketExtension{},
				&tls.ALPNExtension{AlpnProtocols: []string{"h2"}}, //application_layer_protocol_negotiation
				&tls.StatusRequestExtension{},
				&tls.FakeDelegatedCredentialsExtension{
					SupportedSignatureAlgorithms: []tls.SignatureScheme{ //signature_algorithms
						tls.ECDSAWithP256AndSHA256,
						tls.ECDSAWithP384AndSHA384,
						tls.ECDSAWithP521AndSHA512,
						tls.ECDSAWithSHA1,
					},
				},
				&tls.KeyShareExtension{
					KeyShares: []tls.KeyShare{
						{Group: tls.X25519},
						{Group: tls.CurveP256}, //key_share
					},
				},
				&tls.SupportedVersionsExtension{
					Versions: []uint16{
						tls.VersionTLS13, //supported_versions
						tls.VersionTLS12,
					},
				},
				&tls.SignatureAlgorithmsExtension{
					SupportedSignatureAlgorithms: []tls.SignatureScheme{ //signature_algorithms
						tls.ECDSAWithP256AndSHA256,
						tls.ECDSAWithP384AndSHA384,
						tls.ECDSAWithP521AndSHA512,
						tls.PSSWithSHA256,
						tls.PSSWithSHA384,
						tls.PSSWithSHA512,
						tls.PKCS1WithSHA256,
						tls.PKCS1WithSHA384,
						tls.PKCS1WithSHA512,
						tls.ECDSAWithSHA1,
						tls.PKCS1WithSHA1,
					},
				},
				&tls.PSKKeyExchangeModesExtension{
					Modes: []uint8{ //psk_key_exchange_modes
						tls.PskModeDHE,
					},
				},
				&tls.FakeRecordSizeLimitExtension{Limit: 0x4001},                 //record_size_limit
				&tls.UtlsPaddingExtension{GetPaddingLen: tls.BoringPaddingStyle}, //padding
			},
		},
	},
	{
		Name:    "firefox",
		Version: 105,
		Headers: map[string]string{
			"Host":                      "indeed.com",
			"User-Agent":                "Mozilla/5.0 (X11; Linux x86_64; rv:105.0) Gecko/20100101 Firefox/105.0",
			"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
			"Accept-Language":           "en-US,en;q=0.5",
			"DNT":                       "1",
			"Upgrade-Insecure-Requests": "1",
			"Connection":                "keep-alive",
			"Sec-Fetch-Dest":            "document",
			"Sec-Fetch-Mode":            "navigate",
			"Sec-Fetch-Site":            "none",
			"Sec-Fetch-User":            "?1",
		},
		TLSConfig: &tls.ClientHelloSpec{
			TLSVersMin: tls.VersionTLS12,
			TLSVersMax: tls.VersionTLS13,
			CipherSuites: []uint16{
				tls.TLS_AES_128_GCM_SHA256,
				tls.TLS_CHACHA20_POLY1305_SHA256,
				tls.TLS_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
			CompressionMethods: []uint8{
				0x0, // no compression
			},
			Extensions: []tls.TLSExtension{
				&tls.SNIExtension{},
				&tls.ExtendedMasterSecretExtension{},
				&tls.RenegotiationInfoExtension{
					Renegotiation: tls.RenegotiateOnceAsClient,
				},
				&tls.SupportedCurvesExtension{
					Curves: []tls.CurveID{
						tls.X25519,
						tls.CurveP256,
						tls.CurveP384,
						tls.CurveP521,
						256,
						257,
					},
				},
				&tls.SupportedPointsExtension{
					SupportedPoints: []uint8{
						0x0, // uncompressed
					},
				},
				&tls.SessionTicketExtension{},
				&tls.ALPNExtension{
					AlpnProtocols: []string{
						"h2",
						"http/1.1",
					},
				},
				&tls.StatusRequestExtension{},
				&tls.FakeDelegatedCredentialsExtension{
					SupportedSignatureAlgorithms: []tls.SignatureScheme{
						tls.ECDSAWithP256AndSHA256,
						tls.ECDSAWithP384AndSHA384,
						tls.ECDSAWithP521AndSHA512,
						tls.ECDSAWithSHA1,
					},
				},
				&tls.KeyShareExtension{
					KeyShares: []tls.KeyShare{
						{
							Group: tls.X25519,
						},
						{
							Group: tls.CurveP256,
						},
					},
				},
				&tls.SupportedVersionsExtension{
					Versions: []uint16{
						tls.VersionTLS13,
						tls.VersionTLS12,
					},
				},
				&tls.SignatureAlgorithmsExtension{
					SupportedSignatureAlgorithms: []tls.SignatureScheme{
						tls.ECDSAWithP256AndSHA256,
						tls.ECDSAWithP384AndSHA384,
						tls.ECDSAWithP521AndSHA512,
						tls.PSSWithSHA256,
						tls.PSSWithSHA384,
						tls.PSSWithSHA512,
						tls.PKCS1WithSHA256,
						tls.PKCS1WithSHA384,
						tls.PKCS1WithSHA512,
						tls.ECDSAWithSHA1,
						tls.PKCS1WithSHA1,
					},
				},
				&tls.PSKKeyExchangeModesExtension{
					Modes: []uint8{
						tls.PskModeDHE,
					},
				},
				&tls.FakeRecordSizeLimitExtension{
					Limit: 0x4001,
				},
				&tls.UtlsPaddingExtension{
					GetPaddingLen: tls.BoringPaddingStyle,
				},
			},
		},
	},
}

type Agent struct {
	Name        string
	Version     uint
	Headers     map[string]string
	TLSConfig   *tls.ClientHelloSpec
	DialTimeout time.Duration
}

type ActiveAgent struct {
	T       http.Transport
	Headers map[string]string
}

// NewAgent creates a new Agent.
// name is the name of the agent.
// version is the version of the agent.
// headers is the headers to send with the request.
// tlsConf is the tls configuration to use.
// diealTimeout is the timeout for the dial.
func NewAgent(
	name string,
	version uint,
	headers map[string]string,
	tlsConf *tls.ClientHelloSpec,
	diealTimeout time.Duration,
) *Agent {
	return &Agent{name, version, headers, tlsConf, diealTimeout}
}

// NewActiveAgent creates a new ActiveAgent
// agent is the agent to use.
func NewActiveAgent(agent *Agent) *ActiveAgent {
	return &ActiveAgent{
		T: http.Transport{
			DialTLSContext: agent.DialTLS,
		},
		Headers: agent.Headers,
	}
}

// DialTLSContext is the dial function for creating TLS connections.
// ctx is a context provided for cancellation.
// network is the network on which to open the connection ("tcp", "tcp4" or "tcp6").
// addr is the address of the server.
func (a *Agent) DialTLS(ctx context.Context, network string, addr string) (net.Conn, error) {
	config := tls.Config{ServerName: addr, InsecureSkipVerify: true}
	dialConn, err := net.DialTimeout(network, addr, a.DialTimeout)
	if err != nil {
		return nil, err
	}
	conn := tls.UClient(dialConn, &config, tls.HelloCustom)
	err = conn.ApplyPreset(a.TLSConfig)
	if err != nil {
		return nil, err
	}
	err = conn.Handshake()
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// Stop closes all idle connections.
func (aa *ActiveAgent) Stop() {
	aa.T.CloseIdleConnections()
}

// SelectOne selects the first agent that matches the given condition.
func SelectOne(condition func(a *Agent) bool) *Agent {
	for _, a := range Agents {
		if condition(a) {
			return a
		}
	}

	return nil
}

// SelectMany selects all agents that match the given condition.
func SelectMany(condition func(a *Agent) bool) []*Agent {
	var agents []*Agent

	for _, a := range Agents {
		if condition(a) {
			agents = append(agents, a)
		}
	}

	return agents
}
