package useragent

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTokenCollapse(t *testing.T) {
	token := NewToken(42, WithAllowedTokens(PLATFORM_LINUX, PLATFORM_MACOS))
	collapsed := token.Collapse()
	require.Contains(t, []TokenType{PLATFORM_LINUX, PLATFORM_MACOS}, collapsed)
	require.Equal(t, 1, len(token.Possibilities))
}

func TestTokenObserve(t *testing.T) {
	collapsed := NewToken(42, WithAllowedTokens(PLATFORM_LINUX))
	collapsed.Collapse()
	prev := NewToken(42, WithAllowedTokens(LINUX_PLATFORM_VERSION_5_18_11))
	prev.Collapse()
	current := NewToken(42, WithAllowedTokens(ARCH_X86, ARCH_X64, ARCH_ARM))

	current.Observe(collapsed, prev)

	require.Equal(t, 1, len(current.Possibilities))
	require.Equal(t, ARCH_X86, current.Possibilities[0])
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
		PLATFORM_LINUX,
		LINUX_PLATFORM_VERSION_5_18_11,
		ARCH_X86,
		BIT_64,
		MOZILLA_5_BROWSER_IDENTIFIER,
		X11_WINDOW_SYSTEM,
		LINUX,
		X86_64_PROC_ARCH,
		APPLE_WEBKIT_537_36,
		KHTML_ADDITIONAL_INFO,
		CHROME_120_0,
		SAFARI_WEBKIT_537_36,
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
		{"Platform", PLATFORM_LINUX, 0, PLATFORM_LINUX, true},
		{"Linux Platform Version", PLATFORM_LINUX, PLATFORM_LINUX, LINUX_PLATFORM_VERSION_5_18_11, true},
		{"MacOS Architecture", PLATFORM_MACOS, MACOS_PLATFORM_VERSION_13_6_6, ARCH_ARM, true},
		{"Windows Architecture", PLATFORM_WINDOWS, WINDOWS_PLATFORM_VERSION_10_0_0, ARCH_X64, true},
		{"Bitness", ARCH_X86, ARCH_X86, BIT_64, true},
		{"Mozilla Identifier", BIT_64, BIT_64, MOZILLA_5_BROWSER_IDENTIFIER, true},
		{"Linux Window System", PLATFORM_LINUX, MOZILLA_5_BROWSER_IDENTIFIER, X11_WINDOW_SYSTEM, true},
		{"Linux OS", X11_WINDOW_SYSTEM, X11_WINDOW_SYSTEM, LINUX, true},
		{"MacOS Device", PLATFORM_MACOS, MACINTOSH_DEVICE, MACOS_13_6_6, true},
		{"Windows NT", WINDOWS_NT_10_0, WINDOWS_NT_10_0, WIN64_ARCH, true},
		{"Windows Architecture", WIN64_ARCH, WIN64_ARCH, X64_PROC_ARCH, true},
		{"AppleWebKit", X64_PROC_ARCH, X64_PROC_ARCH, APPLE_WEBKIT_537_36, true},
		{"KHTML Info", APPLE_WEBKIT_537_36, APPLE_WEBKIT_537_36, KHTML_ADDITIONAL_INFO, true},
		{"Chrome Version", KHTML_ADDITIONAL_INFO, KHTML_ADDITIONAL_INFO, CHROME_120_0, true},
		{"Safari WebKit", CHROME_120_0, CHROME_120_0, SAFARI_WEBKIT_537_36, true},
		{"Incompatible Linux Version", PLATFORM_LINUX, PLATFORM_LINUX, MACOS_PLATFORM_VERSION_13_6_6, false},
		{"Incompatible MacOS Architecture", PLATFORM_MACOS, MACOS_PLATFORM_VERSION_13_6_6, ARCH_X86, false},
		{"Incompatible Windows Architecture", PLATFORM_WINDOWS, WINDOWS_PLATFORM_VERSION_10_0_0, ARCH_ARM, false},
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
		{"Platform Linux", PLATFORM_LINUX, START_PLATFORM, END_PLATFORM, true},
		{"Start Platform", START_PLATFORM, START_PLATFORM, END_PLATFORM, false},
		{"End Platform", END_PLATFORM, START_PLATFORM, END_PLATFORM, false},
		{"Linux Platform Version", LINUX_PLATFORM_VERSION_5_18_11, START_LINUX_PLATFORM_VERSION, END_LINUX_PLATFORM_VERSION, true},
		{"Chrome Version", CHROME_120_0, START_CHROME, END_CHROME, true},
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
		PLATFORM_LINUX,
		LINUX_PLATFORM_VERSION_5_18_11,
		ARCH_X86,
		BIT_64,
		MOZILLA_5_BROWSER_IDENTIFIER,
		X11_WINDOW_SYSTEM,
		LINUX,
		X86_64_PROC_ARCH,
		APPLE_WEBKIT_537_36,
		KHTML_ADDITIONAL_INFO,
		CHROME_120_0,
		SAFARI_WEBKIT_537_36,
	}

	testCases := []struct {
		name     string
		start    TokenType
		end      TokenType
		expected []TokenType
	}{
		{"Filter Platforms", START_PLATFORM, END_PLATFORM, []TokenType{PLATFORM_LINUX}},
		{"Filter Linux Platform Versions", START_LINUX_PLATFORM_VERSION, END_LINUX_PLATFORM_VERSION, []TokenType{LINUX_PLATFORM_VERSION_5_18_11}},
		{"Filter Architectures", START_ARCH, END_ARCH, []TokenType{ARCH_X86}},
		{"Filter Chrome Versions", START_CHROME, END_CHROME, []TokenType{CHROME_120_0}},
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
				PLATFORM_LINUX,
				LINUX_PLATFORM_VERSION_5_18_11,
				ARCH_X86,
				BIT_64,
				MOZILLA_5_BROWSER_IDENTIFIER,
				X11_WINDOW_SYSTEM,
				LINUX,
				X86_64_PROC_ARCH,
				APPLE_WEBKIT_537_36,
				KHTML_ADDITIONAL_INFO,
				CHROME_120_0,
				SAFARI_WEBKIT_537_36,
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
