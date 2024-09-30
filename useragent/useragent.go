package useragent

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	_ Header = iota

	SecCHUAPlatformHeader
	SecCHUAPlatformVersionHeader
	SecCHUAArchHeader
	SecCHUABitnessHeader
	UserAgentHeader
)

type Header int

func (h Header) String() string {
	switch h {
	case SecCHUAPlatformHeader:
		return "sec-ch-ua-platform"
	case SecCHUAPlatformVersionHeader:
		return "sec-ch-ua-platform-version"
	case SecCHUAArchHeader:
		return "sec-ch-ua-arch"
	case SecCHUABitnessHeader:
		return "sec-ch-ua-bitness"
	case UserAgentHeader:
		return "user-agent"
	default:
		return ""
	}
}

const (
	_ TokenType = iota

	// Platform token types
	StartPlatform
	PlatformLinux
	PlatformMacOS
	PlatformWindows
	EndPlatform

	// Platform version token types
	StartLinuxPlatformVersion
	LinuxPlatformVersion_5_18_11
	LinuxPlatformVersion_5_19_15
	LinuxPlatformVersion_6_7_11
	LinuxPlatformVersion_6_8_12
	LinuxPlatformVersion_6_9_10
	LinuxPlatformVersion_6_10_11
	EndLinuxPlatformVersion

	StartMacOSPlatformVersion
	MacOSPlatformVersion_13_6_6
	MacOSPlatformVersion_13_7
	MacOSPlatformVersion_14_4_1
	MacOSPlatformVersion_14_6_1
	MacOSPlatformVersion_14_7
	MacOSPlatformVersion_15_0
	EndMacOSPlatformVersion

	StartWindowsPlatformVersion
	WindowsPlatformVersion_10_0_0
	WindowsPlatformVersion_14_0_0
	EndWindowsPlatformVersion

	// Architecture token types
	StartArch
	ArchX86
	ArchX64
	ArchARM
	EndArch

	// Bitness token types
	StartBitness
	Bit64
	EndBitness

	// Browser identifier token types
	Mozilla5BrowserIdentifier

	// Window system token types
	X11WindowSystem

	// Device type token types
	MacintoshDevice

	// OS token types
	Linux

	StartMacOS
	MacOS_13_6_6
	MacOS_13_7
	MacOS_14_4_1
	MacOS_14_6_1
	MacOS_14_7
	MacOS_15_0
	EndMacOS

	StartWindows
	WindowsNT_10_0
	EndWindows

	// OS bitness token types
	Win64Arch

	// Processor architecture token types
	StartProcArch
	X64ProcArch
	X86_64ProcArch
	EndProcArch

	// Rendering engine token types
	StartAppleWebKit
	AppleWebKit_537_36
	EndAppleWebKit

	StartSafariWebKit
	SafariWebKit_537_36
	EndSafariWebKit

	// Additional info token types
	KHTMLAdditionalInfo

	// Browser name token types
	StartChrome
	Chrome_120_0
	Chrome_121_0
	Chrome_122_0
	Chrome_123_0
	Chrome_124_0
	Chrome_125_0
	Chrome_126_0
	Chrome_127_0
	Chrome_128_0
	Chrome_129_0
	EndChrome
)

type TokenType int

func (t TokenType) String() string {
	switch t {
	case PlatformLinux:
		return "Linux"
	case PlatformMacOS:
		return "macOS"
	case PlatformWindows:
		return "Windows"
	case LinuxPlatformVersion_5_18_11:
		return "5.18.11"
	case LinuxPlatformVersion_5_19_15:
		return "5.19.15"
	case LinuxPlatformVersion_6_7_11:
		return "6.7.11"
	case LinuxPlatformVersion_6_8_12:
		return "6.8.12"
	case LinuxPlatformVersion_6_9_10:
		return "6.9.10"
	case LinuxPlatformVersion_6_10_11:
		return "6.10.11"
	case MacOSPlatformVersion_13_6_6:
		return "13.6.6"
	case MacOSPlatformVersion_13_7:
		return "13.7"
	case MacOSPlatformVersion_14_4_1:
		return "14.4.1"
	case MacOSPlatformVersion_14_6_1:
		return "14.6.1"
	case MacOSPlatformVersion_14_7:
		return "14.7"
	case MacOSPlatformVersion_15_0:
		return "15.0"
	case WindowsPlatformVersion_10_0_0:
		return "10.0.0"
	case WindowsPlatformVersion_14_0_0:
		return "14.0.0"
	case ArchX86:
		return "x86"
	case ArchX64:
		return "x64"
	case ArchARM:
		return "arm"
	case Bit64:
		return "64"
	case Mozilla5BrowserIdentifier:
		return "Mozilla/5.0"
	case X11WindowSystem:
		return "(X11;"
	case MacintoshDevice:
		return "(Macintosh;"
	case Linux:
		return "Linux"
	case MacOS_13_6_6:
		return "Intel Mac OS X 13_6_6)"
	case MacOS_13_7:
		return "Intel Mac OS X 13_7)"
	case MacOS_14_4_1:
		return "Intel Mac OS X 14_4_1)"
	case MacOS_14_6_1:
		return "Intel Mac OS X 14_6_1)"
	case MacOS_14_7:
		return "Intel Mac OS X 14_7)"
	case MacOS_15_0:
		return "Intel Mac OS X 15_0)"
	case WindowsNT_10_0:
		return "(Windows NT 10.0;"
	case Win64Arch:
		return "Win64;"
	case X64ProcArch:
		return "x64)"
	case X86_64ProcArch:
		return "x86_64)"
	case AppleWebKit_537_36:
		return "AppleWebKit/537.36"
	case KHTMLAdditionalInfo:
		return "(KHTML, like Gecko)"
	case SafariWebKit_537_36:
		return "Safari/537.36"
	case Chrome_120_0:
		return "Chrome/120.0.0.0"
	case Chrome_121_0:
		return "Chrome/121.0.0.0"
	case Chrome_122_0:
		return "Chrome/122.0.0.0"
	case Chrome_123_0:
		return "Chrome/123.0.0.0"
	case Chrome_124_0:
		return "Chrome/124.0.0.0"
	case Chrome_125_0:
		return "Chrome/125.0.0.0"
	case Chrome_126_0:
		return "Chrome/126.0.0.0"
	case Chrome_127_0:
		return "Chrome/127.0.0.0"
	case Chrome_128_0:
		return "Chrome/128.0.0.0"
	case Chrome_129_0:
		return "Chrome/129.0.0.0"
	default:
		return ""
	}
}

type Token struct {
	Possibilities []TokenType
	rand          *rand.Rand
}

type UserAgent struct {
	Headers map[string]string
	tokens  []*Token
}

func NewToken(seed int64, allowedTokens ...TokenType) *Token {
	possibilities := make([]TokenType, 0, EndChrome)
	if len(allowedTokens) > 0 {
		possibilities = make([]TokenType, len(allowedTokens))
		copy(possibilities, allowedTokens)
	} else {
		for i := TokenType(0); i < EndChrome; i++ {
			possibilities = append(possibilities, TokenType(i))
		}
	}

	return &Token{
		Possibilities: possibilities,
		rand:          rand.New(rand.NewSource(seed)),
	}
}

// NewUserAgent generates a new user agent headers with the given length and seed.
// The allowedTokens parameter is used to limit the possible token types.
func NewUserAgent(length int, seed int64, allowedTokens ...TokenType) *UserAgent {
	tokens := make([]*Token, length)
	for i := range tokens {
		tokens[i] = NewToken(seed, allowedTokens...)
	}
	tokens[0].Possibilities = filterTokens(tokens[0].Possibilities, StartPlatform, EndPlatform)

	ua := &UserAgent{
		Headers: map[string]string{
			SecCHUAPlatformHeader.String():        "",
			SecCHUAPlatformVersionHeader.String(): "",
			SecCHUAArchHeader.String():            "",
			SecCHUABitnessHeader.String():         "",
			UserAgentHeader.String():              "",
		},
		tokens: tokens,
	}

	ua.generate()
	ua.updateHeaders()

	return ua
}

func (ua *UserAgent) generate() {
	for i, token := range ua.tokens {
		tt := token.Collapse()
		if tt == 0 {
			break
		}

		for j := i + 1; j < len(ua.tokens); j++ {
			ua.tokens[j].Observe(token, ua.tokens[j-1])
		}
	}
}

func (ua *UserAgent) updateHeaders() {
	for i, token := range ua.tokens {
		if i < 4 {
			ua.Headers[Header(i+1).String()] = token.Possibilities[0].String()
		} else if i < len(ua.tokens) && len(token.Possibilities) > 0 {
			ua.Headers[UserAgentHeader.String()] += fmt.Sprintf("%s ", token.Possibilities[0].String())
		}
	}
	ua.Headers[UserAgentHeader.String()] = strings.TrimSpace(ua.Headers[UserAgentHeader.String()])
}

func (t *Token) Collapse() TokenType {
	if len(t.Possibilities) == 0 {
		return 0
	}
	t.Possibilities = []TokenType{
		t.Possibilities[t.rand.Intn(len(t.Possibilities))],
	}
	return t.Possibilities[0]
}

func (t *Token) Observe(collapsed, prev *Token) {
	reduced := []TokenType{}
	for _, currentType := range t.Possibilities {
		t.rand.Shuffle(len(prev.Possibilities), func(i, j int) {
			prev.Possibilities[i], prev.Possibilities[j] = prev.Possibilities[j], prev.Possibilities[i]
		})
		for _, prevType := range prev.Possibilities {
			if isCompatible(collapsed.Possibilities[0], prevType, currentType) {
				reduced = append(reduced, currentType)
				break
			}
		}
	}
	t.Possibilities = reduced
}

func isCompatible(collapsed, prev, current TokenType) bool {
	isPlatform := func(token TokenType) bool {
		return in(token, StartPlatform, EndPlatform)
	}

	isPlatformVersion := func(token TokenType) bool {
		return in(token, StartLinuxPlatformVersion, EndLinuxPlatformVersion) ||
			in(token, StartMacOSPlatformVersion, EndMacOSPlatformVersion) ||
			in(token, StartWindowsPlatformVersion, EndWindowsPlatformVersion)
	}

	isArch := func(token TokenType) bool {
		return in(token, StartArch, EndArch)
	}

	isBitness := func(token TokenType) bool {
		return in(token, StartBitness, EndBitness)
	}

	isChromeVersion := func(token TokenType) bool {
		return in(token, StartChrome, EndChrome)
	}

	// First token must be a platform
	if prev == 0 {
		return isPlatform(current)
	}

	// Second token must be a platform version corresponding to the platform
	if isPlatform(prev) {
		switch prev {
		case PlatformLinux:
			return in(current, StartLinuxPlatformVersion, EndLinuxPlatformVersion)
		case PlatformMacOS:
			return in(current, StartMacOSPlatformVersion, EndMacOSPlatformVersion)
		case PlatformWindows:
			return in(current, StartWindowsPlatformVersion, EndWindowsPlatformVersion)
		}
		return false
	}

	// Third token must be an architecture
	if isPlatformVersion(prev) {
		switch collapsed {
		case PlatformLinux:
			return current == ArchX86
		case PlatformWindows:
			return current == ArchX64
		case PlatformMacOS:
			return current == ArchARM
		}
		return true
	}

	// Fourth token must be a bitness
	if isArch(prev) {
		return isBitness(current)
	}

	// Fifth token must be Mozilla5BrowserIdentifier
	if isBitness(prev) {
		return current == Mozilla5BrowserIdentifier
	}

	// After Mozilla5BrowserIdentifier, we expect the window system or device type or OS
	if prev == Mozilla5BrowserIdentifier {
		switch collapsed {
		case PlatformLinux:
			return current == X11WindowSystem
		case PlatformMacOS:
			return current == MacintoshDevice
		case PlatformWindows:
			return current == WindowsNT_10_0
		}
		return true
	}

	// After window system or device type, we expect the OS
	if prev == X11WindowSystem {
		return current == Linux
	}

	if prev == MacintoshDevice {
		if in(collapsed, StartMacOSPlatformVersion, EndMacOSPlatformVersion) {
			platformVersionOffset := int(collapsed) - int(StartMacOSPlatformVersion)
			return int(current) == int(StartMacOS)+platformVersionOffset
		}
		return true
	}

	if prev == WindowsNT_10_0 {
		return current == Win64Arch
	}

	// After OS, we expect the processor architecture
	if prev == Linux {
		return current == X86_64ProcArch
	}

	if prev == Win64Arch {
		return current == X64ProcArch
	}

	// After processor architecture, we expect AppleWebKit
	if in(prev, StartProcArch, EndProcArch) || in(prev, StartMacOS, EndMacOS) {
		return current == AppleWebKit_537_36
	}

	// After AppleWebKit, we expect KHTML additional info
	if prev == AppleWebKit_537_36 {
		return current == KHTMLAdditionalInfo
	}

	// After KHTML, we expect Chrome version
	if prev == KHTMLAdditionalInfo {
		return isChromeVersion(current)
	}

	// After Chrome version, we expect SafariWebKit
	if isChromeVersion(prev) {
		if collapsed == AppleWebKit_537_36 {
			return current == SafariWebKit_537_36
		}
		return true
	}

	// No more tokens expected after SafariWebKit
	return false
}

func in(token, start, end TokenType) bool {
	return token > start && token < end
}

func filterTokens(tokens []TokenType, start, end TokenType) []TokenType {
	filtered := make([]TokenType, 0, len(tokens))
	for _, token := range tokens {
		if in(token, start, end) {
			filtered = append(filtered, token)
		}
	}

	return filtered
}
