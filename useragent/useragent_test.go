package useragent

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTokenCollapse(t *testing.T) {
	token := NewToken(42, WithAllowedTokens(PlatformLinux, PlatformMacOS))
	collapsed := token.Collapse()
	require.Contains(t, []TokenType{PlatformLinux, PlatformMacOS}, collapsed)
	require.Equal(t, 1, len(token.Possibilities))
}

func TestTokenObserve(t *testing.T) {
	collapsed := NewToken(42, WithAllowedTokens(PlatformLinux))
	collapsed.Collapse()
	prev := NewToken(42, WithAllowedTokens(LinuxPlatformVersion_5_18_11))
	prev.Collapse()
	current := NewToken(42, WithAllowedTokens(ArchX86, ArchX64, ArchARM))

	current.Observe(collapsed, prev)

	require.Equal(t, 1, len(current.Possibilities))
	require.Equal(t, ArchX86, current.Possibilities[0])
}

func TestNewUserAgent(t *testing.T) {
	ua := NewUserAgent(20, 42)

	require.NotEmpty(t, ua.Headers[SecCHUAPlatformHeader.String()])
	require.NotEmpty(t, ua.Headers[SecCHUAPlatformVersionHeader.String()])
	require.NotEmpty(t, ua.Headers[SecCHUAArchHeader.String()])
	require.NotEmpty(t, ua.Headers[SecCHUABitnessHeader.String()])
	require.NotEmpty(t, ua.Headers[UserAgentHeader.String()])

	// Additional checks for header format
	require.Regexp(t, `^(Linux|macOS|Windows)$`, ua.Headers[SecCHUAPlatformHeader.String()])
	require.Regexp(t, `^\d+\.\d+(\.\d+)?$`, ua.Headers[SecCHUAPlatformVersionHeader.String()])
	require.Regexp(t, `^(x86|x64|arm)$`, ua.Headers[SecCHUAArchHeader.String()])
	require.Equal(t, "64", ua.Headers[SecCHUABitnessHeader.String()])
	require.Regexp(t, `^Mozilla/5\.0 .+ AppleWebKit/537\.36 .+ Chrome/\d+\.\d+\.\d+\.\d+ Safari/537\.36$`, ua.Headers[UserAgentHeader.String()])
}

func TestNewUserAgentWithAllowedTokens(t *testing.T) {
	allowedTokens := []TokenType{
		PlatformLinux,
		LinuxPlatformVersion_5_18_11,
		ArchX86,
		Bit64,
		Mozilla5BrowserIdentifier,
		X11WindowSystem,
		Linux,
		X86_64ProcArch,
		AppleWebKit_537_36,
		KHTMLAdditionalInfo,
		Chrome_120_0,
		SafariWebKit_537_36,
	}

	ua := NewUserAgent(20, 42, WithAllowedTokens(allowedTokens...))

	expectedHeaders := map[string]string{
		SecCHUAPlatformHeader.String():        "Linux",
		SecCHUAPlatformVersionHeader.String(): "5.18.11",
		SecCHUAArchHeader.String():            "x86",
		SecCHUABitnessHeader.String():         "64",
		UserAgentHeader.String():              "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
	}

	for header, expectedValue := range expectedHeaders {
		require.Equal(t, expectedValue, ua.Headers[header], "Mismatch In header %s", header)
	}
}

func TestIsCompatible(t *testing.T) {
	testCases := []struct {
		name      string
		collapsed TokenType
		prev      TokenType
		current   TokenType
		expected  bool
	}{
		{"Platform", PlatformLinux, 0, PlatformLinux, true},
		{"Linux Platform Version", PlatformLinux, PlatformLinux, LinuxPlatformVersion_5_18_11, true},
		{"MacOS Architecture", PlatformMacOS, MacOSPlatformVersion_13_6_6, ArchARM, true},
		{"Windows Architecture", PlatformWindows, WindowsPlatformVersion_10_0_0, ArchX64, true},
		{"Bitness", ArchX86, ArchX86, Bit64, true},
		{"Mozilla Identifier", Bit64, Bit64, Mozilla5BrowserIdentifier, true},
		{"Linux Window System", PlatformLinux, Mozilla5BrowserIdentifier, X11WindowSystem, true},
		{"Linux OS", X11WindowSystem, X11WindowSystem, Linux, true},
		{"MacOS Device", PlatformMacOS, MacintoshDevice, MacOS_13_6_6, true},
		{"Windows NT", WindowsNT_10_0, WindowsNT_10_0, Win64Arch, true},
		{"Windows Architecture", Win64Arch, Win64Arch, X64ProcArch, true},
		{"AppleWebKit", X64ProcArch, X64ProcArch, AppleWebKit_537_36, true},
		{"KHTML Info", AppleWebKit_537_36, AppleWebKit_537_36, KHTMLAdditionalInfo, true},
		{"Chrome Version", KHTMLAdditionalInfo, KHTMLAdditionalInfo, Chrome_120_0, true},
		{"Safari WebKit", Chrome_120_0, Chrome_120_0, SafariWebKit_537_36, true},
		{"Incompatible Linux Version", PlatformLinux, PlatformLinux, MacOSPlatformVersion_13_6_6, false},
		{"Incompatible MacOS Architecture", PlatformMacOS, MacOSPlatformVersion_13_6_6, ArchX86, false},
		{"Incompatible Windows Architecture", PlatformWindows, WindowsPlatformVersion_10_0_0, ArchARM, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isCompatible(tc.collapsed, tc.prev, tc.current)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestIn(t *testing.T) {
	testCases := []struct {
		name     string
		token    TokenType
		start    TokenType
		end      TokenType
		expected bool
	}{
		{"Platform Linux", PlatformLinux, StartPlatform, EndPlatform, true},
		{"Start Platform", StartPlatform, StartPlatform, EndPlatform, false},
		{"End Platform", EndPlatform, StartPlatform, EndPlatform, false},
		{"Linux Platform Version", LinuxPlatformVersion_5_18_11, StartLinuxPlatformVersion, EndLinuxPlatformVersion, true},
		{"Chrome Version", Chrome_120_0, StartChrome, EndChrome, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := In(tc.token, tc.start, tc.end)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestFilterTokens(t *testing.T) {
	tokens := []TokenType{
		PlatformLinux,
		LinuxPlatformVersion_5_18_11,
		ArchX86,
		Bit64,
		Mozilla5BrowserIdentifier,
		X11WindowSystem,
		Linux,
		X86_64ProcArch,
		AppleWebKit_537_36,
		KHTMLAdditionalInfo,
		Chrome_120_0,
		SafariWebKit_537_36,
	}

	testCases := []struct {
		name     string
		start    TokenType
		end      TokenType
		expected []TokenType
	}{
		{"Filter Platforms", StartPlatform, EndPlatform, []TokenType{PlatformLinux}},
		{"Filter Linux Platform Versions", StartLinuxPlatformVersion, EndLinuxPlatformVersion, []TokenType{LinuxPlatformVersion_5_18_11}},
		{"Filter Architectures", StartArch, EndArch, []TokenType{ArchX86}},
		{"Filter Chrome Versions", StartChrome, EndChrome, []TokenType{Chrome_120_0}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			filtered := filterTokens(tokens, tc.start, tc.end)
			require.Equal(t, tc.expected, filtered)
		})
	}
}

func BenchmarkNewUserAgent(b *testing.B) {
	benchCases := []struct {
		name          string
		length        int
		seed          int64
		allowedTokens []TokenType
	}{
		{
			name:   "default",
			length: 20,
			seed:   42,
		},
		{
			name:   "short",
			length: 10,
			seed:   42,
		},
		{
			name:   "medium",
			length: 15,
			seed:   42,
		},
		{
			name:   "long",
			length: 30,
			seed:   42,
		},
		{
			name:   "with allowed tokens",
			length: 20,
			seed:   42,
			allowedTokens: []TokenType{
				PlatformLinux,
				LinuxPlatformVersion_5_18_11,
				ArchX86,
				Bit64,
				Mozilla5BrowserIdentifier,
				X11WindowSystem,
				Linux,
				X86_64ProcArch,
				AppleWebKit_537_36,
				KHTMLAdditionalInfo,
				Chrome_120_0,
				SafariWebKit_537_36,
			},
		},
	}

	for _, bc := range benchCases {
		b.Run(bc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if len(bc.allowedTokens) > 0 {
					NewUserAgent(bc.length, bc.seed, WithAllowedTokens(bc.allowedTokens...))
				} else {
					NewUserAgent(bc.length, bc.seed)
				}
			}
		})
	}
}
