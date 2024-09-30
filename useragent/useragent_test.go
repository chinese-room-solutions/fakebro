package useragent

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTokenCollapse(t *testing.T) {
	token := NewToken(42, func(tt TokenType) bool {
		return tt == PlatformLinux || tt == PlatformMacOS
	})
	collapsed := token.Collapse()
	require.Contains(t, []TokenType{PlatformLinux, PlatformMacOS}, collapsed)
	require.Equal(t, 1, len(token.Possibilities))
}

func TestTokenObserve(t *testing.T) {
	collapsed := NewToken(42, func(tt TokenType) bool { return tt == PlatformLinux })
	collapsed.Collapse()
	prev := NewToken(42, func(tt TokenType) bool { return tt == LinuxPlatformVersion_5_18_11 })
	prev.Collapse()
	current := NewToken(42, func(tt TokenType) bool {
		return tt == ArchX86 || tt == ArchX64 || tt == ArchARM
	})

	current.Observe(collapsed, prev)

	require.Equal(t, 1, len(current.Possibilities))
	require.Equal(t, ArchX86, current.Possibilities[0])
}

func TestNewUserAgent(t *testing.T) {
	ua := NewUserAgent(20, 42, func(tt TokenType) bool { return true })

	require.NotEmpty(t, ua.Headers[SecCHUAPlatformHeader.String()])
	require.NotEmpty(t, ua.Headers[SecCHUAPlatformVersionHeader.String()])
	require.NotEmpty(t, ua.Headers[SecCHUAArchHeader.String()])
	require.NotEmpty(t, ua.Headers[SecCHUABitnessHeader.String()])
	require.NotEmpty(t, ua.Headers[UserAgentHeader.String()])
}

func TestNewUserAgentWithCondition(t *testing.T) {
	condition := func(tt TokenType) bool {
		allowed := []TokenType{
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
		return slices.Contains(allowed, tt)
	}

	ua := NewUserAgent(20, 42, condition)

	expectedHeaders := map[string]string{
		SecCHUAPlatformHeader.String():        "Linux",
		SecCHUAPlatformVersionHeader.String(): "5.18.11",
		SecCHUAArchHeader.String():            "x86",
		SecCHUABitnessHeader.String():         "64",
		UserAgentHeader.String():              "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
	}

	for header, expectedValue := range expectedHeaders {
		require.Equal(t, expectedValue, ua.Headers[header], "Mismatch in header %s", header)
	}
}

func TestIsCompatible(t *testing.T) {
	testCases := []struct {
		collapsed TokenType
		prev      TokenType
		current   TokenType
		expected  bool
	}{
		{PlatformLinux, 0, PlatformLinux, true},
		{PlatformLinux, PlatformLinux, LinuxPlatformVersion_5_18_11, true},
		{PlatformMacOS, MacOSPlatformVersion_13_6_6, ArchARM, true},
		{PlatformWindows, WindowsPlatformVersion_10_0_0, ArchX64, true},
		{ArchX86, ArchX86, Bit64, true},
		{Bit64, Bit64, Mozilla5BrowserIdentifier, true},
		{PlatformLinux, Mozilla5BrowserIdentifier, X11WindowSystem, true},
		{X11WindowSystem, X11WindowSystem, Linux, true},
		{PlatformMacOS, MacintoshDevice, MacOS_13_6_6, true},
		{WindowsNT_10_0, WindowsNT_10_0, Win64Arch, true},
		{Win64Arch, Win64Arch, X64ProcArch, true},
		{X64ProcArch, X64ProcArch, AppleWebKit_537_36, true},
		{AppleWebKit_537_36, AppleWebKit_537_36, KHTMLAdditionalInfo, true},
		{KHTMLAdditionalInfo, KHTMLAdditionalInfo, Chrome_120_0, true},
		{Chrome_120_0, Chrome_120_0, SafariWebKit_537_36, true},
		{PlatformLinux, PlatformLinux, MacOSPlatformVersion_13_6_6, false},
		{PlatformMacOS, MacOSPlatformVersion_13_6_6, ArchX86, false},
		{PlatformWindows, WindowsPlatformVersion_10_0_0, ArchARM, false},
	}

	for _, tc := range testCases {
		result := isCompatible(tc.collapsed, tc.prev, tc.current)
		require.Equal(t, tc.expected, result)
	}
}

func TestIn(t *testing.T) {
	testCases := []struct {
		token    TokenType
		start    TokenType
		end      TokenType
		expected bool
	}{
		{PlatformLinux, StartPlatform, EndPlatform, true},
		{StartPlatform, StartPlatform, EndPlatform, false},
		{EndPlatform, StartPlatform, EndPlatform, false},
		{LinuxPlatformVersion_5_18_11, StartLinuxPlatformVersion, EndLinuxPlatformVersion, true},
		{Chrome_120_0, StartChrome, EndChrome, true},
	}

	for _, tc := range testCases {
		result := in(tc.token, tc.start, tc.end)
		require.Equal(t, tc.expected, result)
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

	expected := []TokenType{
		PlatformLinux,
	}

	filtered := filterTokens(tokens, StartPlatform, EndPlatform)
	require.Equal(t, expected, filtered)
}
