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
	START_PLATFORM
	PLATFORM_LINUX
	PLATFORM_MACOS
	PLATFORM_WINDOWS
	END_PLATFORM

	// Platform version token types
	START_LINUX_PLATFORM_VERSION
	LINUX_PLATFORM_VERSION_5_18_11
	LINUX_PLATFORM_VERSION_5_19_15
	LINUX_PLATFORM_VERSION_6_7_11
	LINUX_PLATFORM_VERSION_6_8_12
	LINUX_PLATFORM_VERSION_6_9_10
	LINUX_PLATFORM_VERSION_6_10_11
	END_LINUX_PLATFORM_VERSION

	START_MACOS_PLATFORM_VERSION
	MACOS_PLATFORM_VERSION_13_6_6
	MACOS_PLATFORM_VERSION_13_7
	MACOS_PLATFORM_VERSION_14_4_1
	MACOS_PLATFORM_VERSION_14_6_1
	MACOS_PLATFORM_VERSION_14_7
	MACOS_PLATFORM_VERSION_15_0
	END_MACOS_PLATFORM_VERSION

	START_WINDOWS_PLATFORM_VERSION
	WINDOWS_PLATFORM_VERSION_10_0_0
	WINDOWS_PLATFORM_VERSION_14_0_0
	END_WINDOWS_PLATFORM_VERSION

	// Architecture token types
	START_ARCH
	ARCH_X86
	ARCH_X64
	ARCH_ARM
	END_ARCH

	// Bitness token types
	START_BITNESS
	BIT_64
	END_BITNESS

	// Browser identifier token types
	MOZILLA_5_BROWSER_IDENTIFIER

	// Window system token types
	X11_WINDOW_SYSTEM

	// Device type token types
	MACINTOSH_DEVICE

	// OS token types
	LINUX

	START_MACOS
	MACOS_13_6_6
	MACOS_13_7
	MACOS_14_4_1
	MACOS_14_6_1
	MACOS_14_7
	MACOS_15_0
	END_MACOS

	START_WINDOWS
	WINDOWS_NT_10_0
	END_WINDOWS

	// OS bitness token types
	WIN64_ARCH

	// Processor architecture token types
	START_PROC_ARCH
	X64_PROC_ARCH
	X86_64_PROC_ARCH
	END_PROC_ARCH

	// Rendering engine token types
	START_APPLE_WEBKIT
	APPLE_WEBKIT_537_36
	END_APPLE_WEBKIT

	START_SAFARI_WEBKIT
	SAFARI_WEBKIT_537_36
	END_SAFARI_WEBKIT

	// Additional info token types
	KHTML_ADDITIONAL_INFO

	// Browser name token types
	START_CHROME
	CHROME_120_0
	CHROME_121_0
	CHROME_122_0
	CHROME_123_0
	CHROME_124_0
	CHROME_125_0
	CHROME_126_0
	CHROME_127_0
	CHROME_128_0
	CHROME_129_0
	END_CHROME
)

type TokenType int

func (t TokenType) String() string {
	switch t {
	case PLATFORM_LINUX:
		return "Linux"
	case PLATFORM_MACOS:
		return "macOS"
	case PLATFORM_WINDOWS:
		return "Windows"
	case LINUX_PLATFORM_VERSION_5_18_11:
		return "5.18.11"
	case LINUX_PLATFORM_VERSION_5_19_15:
		return "5.19.15"
	case LINUX_PLATFORM_VERSION_6_7_11:
		return "6.7.11"
	case LINUX_PLATFORM_VERSION_6_8_12:
		return "6.8.12"
	case LINUX_PLATFORM_VERSION_6_9_10:
		return "6.9.10"
	case LINUX_PLATFORM_VERSION_6_10_11:
		return "6.10.11"
	case MACOS_PLATFORM_VERSION_13_6_6:
		return "13.6.6"
	case MACOS_PLATFORM_VERSION_13_7:
		return "13.7"
	case MACOS_PLATFORM_VERSION_14_4_1:
		return "14.4.1"
	case MACOS_PLATFORM_VERSION_14_6_1:
		return "14.6.1"
	case MACOS_PLATFORM_VERSION_14_7:
		return "14.7"
	case MACOS_PLATFORM_VERSION_15_0:
		return "15.0"
	case WINDOWS_PLATFORM_VERSION_10_0_0:
		return "10.0.0"
	case WINDOWS_PLATFORM_VERSION_14_0_0:
		return "14.0.0"
	case ARCH_X86:
		return "x86"
	case ARCH_X64:
		return "x64"
	case ARCH_ARM:
		return "arm"
	case BIT_64:
		return "64"
	case MOZILLA_5_BROWSER_IDENTIFIER:
		return "Mozilla/5.0"
	case X11_WINDOW_SYSTEM:
		return "(X11;"
	case MACINTOSH_DEVICE:
		return "(Macintosh;"
	case LINUX:
		return "Linux"
	case MACOS_13_6_6:
		return "Intel Mac OS X 13_6_6)"
	case MACOS_13_7:
		return "Intel Mac OS X 13_7)"
	case MACOS_14_4_1:
		return "Intel Mac OS X 14_4_1)"
	case MACOS_14_6_1:
		return "Intel Mac OS X 14_6_1)"
	case MACOS_14_7:
		return "Intel Mac OS X 14_7)"
	case MACOS_15_0:
		return "Intel Mac OS X 15_0)"
	case WINDOWS_NT_10_0:
		return "(Windows NT 10.0;"
	case WIN64_ARCH:
		return "Win64;"
	case X64_PROC_ARCH:
		return "x64)"
	case X86_64_PROC_ARCH:
		return "x86_64)"
	case APPLE_WEBKIT_537_36:
		return "AppleWebKit/537.36"
	case KHTML_ADDITIONAL_INFO:
		return "(KHTML, like Gecko)"
	case SAFARI_WEBKIT_537_36:
		return "Safari/537.36"
	case CHROME_120_0:
		return "Chrome/120.0.0.0"
	case CHROME_121_0:
		return "Chrome/121.0.0.0"
	case CHROME_122_0:
		return "Chrome/122.0.0.0"
	case CHROME_123_0:
		return "Chrome/123.0.0.0"
	case CHROME_124_0:
		return "Chrome/124.0.0.0"
	case CHROME_125_0:
		return "Chrome/125.0.0.0"
	case CHROME_126_0:
		return "Chrome/126.0.0.0"
	case CHROME_127_0:
		return "Chrome/127.0.0.0"
	case CHROME_128_0:
		return "Chrome/128.0.0.0"
	case CHROME_129_0:
		return "Chrome/129.0.0.0"
	default:
		return ""
	}
}

type options struct {
	AllowedTokens []TokenType
	Condition     func(TokenType) bool
}

type Option func(*options)

// WithAllowedTokens allows to limit the possible token types to the given list.
func WithAllowedTokens(tokens ...TokenType) Option {
	return func(o *options) {
		o.AllowedTokens = tokens
	}
}

// WithCondition allows to filter the possible token types based on a condition.
func WithCondition(c func(TokenType) bool) Option {
	return func(o *options) {
		o.Condition = c
	}
}

type Token struct {
	Possibilities []TokenType
	rand          *rand.Rand
}

func NewToken(seed int64, opts ...Option) *Token {
	o := options{}
	for _, opt := range opts {
		opt(&o)
	}

	possibilities := make([]TokenType, 0, END_CHROME)
	if len(o.AllowedTokens) > 0 {
		possibilities = make([]TokenType, len(o.AllowedTokens))
		copy(possibilities, o.AllowedTokens)
	} else {
		for i := TokenType(0); i < END_CHROME; i++ {
			possibilities = append(possibilities, TokenType(i))
		}
	}
	if o.Condition != nil {
		filtered := make([]TokenType, 0, len(possibilities))
		for _, token := range possibilities {
			if o.Condition(token) {
				filtered = append(filtered, token)
			}
		}
		possibilities = filtered
	}

	return &Token{
		Possibilities: possibilities,
		rand:          rand.New(rand.NewSource(seed)),
	}
}

type UserAgent struct {
	Headers map[string]string
	tokens  []*Token
}

// NewUserAgent generates a new user agent headers with the given length and seed.
// The allowedTokens parameter is used to limit the possible token types.
func NewUserAgent(length int, seed int64, opts ...Option) *UserAgent {
	tokens := make([]*Token, length)
	for i := range tokens {
		tokens[i] = NewToken(seed, opts...)
	}
	tokens[0].Possibilities = filterTokens(tokens[0].Possibilities, START_PLATFORM, END_PLATFORM)

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
		return In(token, START_PLATFORM, END_PLATFORM)
	}

	isPlatformVersion := func(token TokenType) bool {
		return In(token, START_LINUX_PLATFORM_VERSION, END_LINUX_PLATFORM_VERSION) ||
			In(token, START_MACOS_PLATFORM_VERSION, END_MACOS_PLATFORM_VERSION) ||
			In(token, START_WINDOWS_PLATFORM_VERSION, END_WINDOWS_PLATFORM_VERSION)
	}

	isArch := func(token TokenType) bool {
		return In(token, START_ARCH, END_ARCH)
	}

	isBitness := func(token TokenType) bool {
		return In(token, START_BITNESS, END_BITNESS)
	}

	isChromeVersion := func(token TokenType) bool {
		return In(token, START_CHROME, END_CHROME)
	}

	// First token must be a platform
	if prev == 0 {
		return isPlatform(current)
	}

	// Second token must be a platform version corresponding to the platform
	if isPlatform(prev) {
		switch prev {
		case PLATFORM_LINUX:
			return In(current, START_LINUX_PLATFORM_VERSION, END_LINUX_PLATFORM_VERSION)
		case PLATFORM_MACOS:
			return In(current, START_MACOS_PLATFORM_VERSION, END_MACOS_PLATFORM_VERSION)
		case PLATFORM_WINDOWS:
			return In(current, START_WINDOWS_PLATFORM_VERSION, END_WINDOWS_PLATFORM_VERSION)
		}
		return false
	}

	// Third token must be an architecture
	if isPlatformVersion(prev) {
		switch collapsed {
		case PLATFORM_LINUX:
			return current == ARCH_X86
		case PLATFORM_WINDOWS:
			return current == ARCH_X64
		case PLATFORM_MACOS:
			return current == ARCH_ARM
		}
		return true
	}

	// Fourth token must be a bitness
	if isArch(prev) {
		return isBitness(current)
	}

	// Fifth token must be MOZILLA_5_BROWSER_IDENTIFIER
	if isBitness(prev) {
		return current == MOZILLA_5_BROWSER_IDENTIFIER
	}

	// After MOZILLA_5_BROWSER_IDENTIFIER, we expect the window system or device type or OS
	if prev == MOZILLA_5_BROWSER_IDENTIFIER {
		switch collapsed {
		case PLATFORM_LINUX:
			return current == X11_WINDOW_SYSTEM
		case PLATFORM_MACOS:
			return current == MACINTOSH_DEVICE
		case PLATFORM_WINDOWS:
			return current == WINDOWS_NT_10_0
		}
		return true
	}

	// After window system or device type, we expect the OS
	if prev == X11_WINDOW_SYSTEM {
		return current == LINUX
	}

	if prev == MACINTOSH_DEVICE {
		if In(collapsed, START_MACOS_PLATFORM_VERSION, END_MACOS_PLATFORM_VERSION) {
			platformVersionOffset := int(collapsed) - int(START_MACOS_PLATFORM_VERSION)
			return int(current) == int(START_MACOS)+platformVersionOffset
		}
		return true
	}

	if prev == WINDOWS_NT_10_0 {
		return current == WIN64_ARCH
	}

	// After OS, we expect the processor architecture
	if prev == LINUX {
		return current == X86_64_PROC_ARCH
	}

	if prev == WIN64_ARCH {
		return current == X64_PROC_ARCH
	}

	// After processor architecture, we expect AppleWebKit
	if In(prev, START_PROC_ARCH, END_PROC_ARCH) || In(prev, START_MACOS, END_MACOS) {
		return current == APPLE_WEBKIT_537_36
	}

	// After AppleWebKit, we expect KHTML additional info
	if prev == APPLE_WEBKIT_537_36 {
		return current == KHTML_ADDITIONAL_INFO
	}

	// After KHTML, we expect Chrome version
	if prev == KHTML_ADDITIONAL_INFO {
		return isChromeVersion(current)
	}

	// After Chrome version, we expect SafariWebKit
	if isChromeVersion(prev) {
		if collapsed == APPLE_WEBKIT_537_36 {
			return current == SAFARI_WEBKIT_537_36
		}
		return true
	}

	// No more tokens expected after SafariWebKit
	return false
}

func In(token, start, end TokenType) bool {
	return token > start && token < end
}

func filterTokens(tokens []TokenType, start, end TokenType) []TokenType {
	filtered := make([]TokenType, 0, len(tokens))
	for _, token := range tokens {
		if In(token, start, end) {
			filtered = append(filtered, token)
		}
	}

	return filtered
}
