package main

import (
	"math/rand"
	"time"
)

type TokenType int

const (
	// Browser identifier token types
	Mozilla5BrowserIdentifier TokenType = iota
	// Window system token types
	X11WindowSystem
	// Device type token types
	MacintoshDevice
	IPhoneDevice
	IPadDevice
	// OS token types
	MacOSX
	WindowsNT
	Linux
	IPhoneOS
	IPadOS
	Android
	// MacOS version number token types
	MacOS_10_12_6
	MacOS_10_13_6
	MacOS_11_7_9
	MacOS_12_6_8
	MacOS_13_5_1
	MacOSLatest
	// IOS version number token types
	IOS_13_7
	IOS_14_8_1
	IOS_15_7_8
	IOS_16_6
	IOSLatest
	// OS architecture token types
	Win64Arc
	// Processor architecture token types
	X64ProcArc
	X86_64ProcArc
	// Revision number token types
	FirefoxRevision_99_0
	FirefoxRevision_102_0
	FirefoxRevision_105_0
	FirefoxRevisionLatest
	// Rendering engine token types
	GeckoRenderingEngine
	AppleWebKitRenderingEngine
	// Render version token types
	Gecko_20100101
	GeckoLatest
	AppleWebKit_604_1
	AppleWebKit_605_1_15
	AppleWebKit_604_1_38
	AppleWebKitLatest
	// Additional info token types
	KHTMLAdditionalInfo
	// Browser name token types
	FirefoxBrowser
	FirefoxMobileBrowser
	SafariBrowser
	// Browser version token types
	Firefox_99_0
	Firefox_102_0
	Firefox_105_0
	FirefoxLatest
	SafariBrowserVersion
	Safari_16_5_2
	Safari_15_6_1
	SafariLatest
	// Mobile token types
	Mobile
	// Mobile version token types
	Mobile_15E148
	MobileLatest
	TotalTokens
)

func (t TokenType) String() string {
	switch t {
	case Mozilla5BrowserIdentifier:
		return "Mozilla/5.0"
	case X11WindowSystem:
		return "X11"
	case MacintoshDevice:
		return "Macintosh"
	case IPhoneDevice:
		return "iPhone"
	case IPadDevice:
		return "iPad"
	case MacOSX:
		return "Intel Mac OS X"
	case WindowsNT:
		return "Windows NT 10.0"
	case Linux:
		return "Linux"
	case IPhoneOS:
		return "CPU iPhone OS %s like Mac OS X"
	case IPadOS:
		return "CPU OS %s like Mac OS X"
	case Android:
		return "Android"
	case MacOS_10_12_6:
		return "10_12_6"
	case MacOS_10_13_6:
		return "10_13_6"
	case MacOS_11_7_9:
		return "11_7_9"
	case MacOS_12_6_8:
		return "12_6_8"
	case MacOS_13_5_1, MacOSLatest:
		return "13_5_1"
	case IOS_13_7:
		return "13_7"
	case IOS_14_8_1:
		return "14_8_1"
	case IOS_15_7_8:
		return "15_7_8"
	case IOS_16_6, IOSLatest:
		return "16_6"
	case Win64Arc:
		return "Win64"
	case X64ProcArc:
		return "x64"
	case X86_64ProcArc:
		return "x86_64"
	case FirefoxRevision_99_0:
		return "rv:99.0"
	case FirefoxRevision_102_0:
		return "rv:102.0"
	case FirefoxRevision_105_0, FirefoxRevisionLatest:
		return "rv:105.0"
	case GeckoRenderingEngine:
		return "Gecko"
	case AppleWebKitRenderingEngine:
		return "AppleWebKit"
	case Gecko_20100101, GeckoLatest:
		return "20100101"
	case AppleWebKit_604_1:
		return "604.1"
	case AppleWebKit_605_1_15:
		return "605.1.15"
	case AppleWebKit_604_1_38, AppleWebKitLatest:
		return "604.1.38"
	case KHTMLAdditionalInfo:
		return "(KHTML, like Gecko)"
	case FirefoxBrowser:
		return "Firefox"
	case FirefoxMobileBrowser:
		return "FxiOS"
	case SafariBrowser:
		return "Safari"
	case Firefox_99_0:
		return "99.0"
	case Firefox_102_0:
		return "102.0"
	case Firefox_105_0, FirefoxLatest:
		return "105.0"
	case Safari_16_5_2:
		return "16.5.2"
	case Safari_15_6_1, SafariBrowserVersion:
		return "15.6.1"
	case Mobile:
		return "Mobile"
	case Mobile_15E148, MobileLatest:
		return "15E148"
	default:
		return ""
	}
}

type Token struct {
	Possibilities []TokenType
	rand          *rand.Rand
}

type UserAgent struct {
	Header  string
	Client  string
	Version string
}

func NewToken() *Token {
	possibilities := []TokenType{}
	for i := TokenType(0); i < TotalTokens; i++ {
		possibilities = append(possibilities, TokenType(i))
	}

	return &Token{
		Possibilities: possibilities,
		rand:          rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func NewUserAgent(length int) *UserAgent {
	tokens := make([]*Token, length)
	for i := 0; i < len(tokens); i++ {
		tokens[i] = NewToken()
	}
	tokens[0].Possibilities = []TokenType{Mozilla5BrowserIdentifier}
	Header, Client, Version := "", "", ""

	for i, t := range tokens {
		tt := t.Collapse()
		if tt == nil {
			break
		}
		Header += tt.String() + " "

		if *tt == FirefoxBrowser || *tt == SafariBrowser {
			Client = tt.String()
		}

		if (*tt >= Firefox_99_0 && *tt <= FirefoxLatest) || (*tt >= Safari_15_6_1 && *tt <= SafariLatest) {
			Version = tt.String()
		}

		for j := i + 1; j < len(tokens); j++ {
			tokens[j].Observe(t, tokens[j-1], j-i)
		}
	}

	return &UserAgent{
		Header:  Header,
		Client:  Client,
		Version: Version,
	}
}

func (t *Token) Collapse() *TokenType {
	if len(t.Possibilities) == 0 {
		return nil
	}
	t.Possibilities = []TokenType{
		t.Possibilities[t.rand.Intn(len(t.Possibilities))],
	}
	return &t.Possibilities[0]
}

func (t *Token) Observe(collapsed, prev *Token, depth int) {
	if prev == nil {
		return
	}

	reduced := []TokenType{}
	for _, currentType := range t.Possibilities {
		for _, prevType := range prev.Possibilities {
			if isCompatible(collapsed.Possibilities[0], prevType, currentType, depth) {
				reduced = append(reduced, currentType)
				break
			}
		}
	}
	t.Possibilities = reduced
}

func isCompatible(collapsed, prev, current TokenType, depth int) bool {
	// Browser identifier must be followed by Window system, Device Type, or OS type
	if prev == Mozilla5BrowserIdentifier {
		return current == X11WindowSystem ||
			(current >= MacintoshDevice && current <= IPadDevice) ||
			(current == Android || current == WindowsNT)
	}

	// Window system must be followed by OS (Linux)
	if prev == X11WindowSystem {
		return current == Linux
	}

	// Macintosh devices should be followed by MacOS
	if prev == MacintoshDevice {
		return current == MacOSX
	}

	// iPhone or iPad should be followed by corresponding OS
	if prev == IPhoneDevice {
		return current == IPhoneOS
	}
	if prev == IPadDevice {
		return current == IPadOS
	}

	// MacOSX, WindowsNT, or other OS should be followed by corresponding version numbers or architecture
	if prev == MacOSX {
		return current >= MacOS_10_12_6 && current <= MacOSLatest
	}
	if prev == WindowsNT {
		return current == Win64Arc
	}
	if prev == Linux {
		return current == X86_64ProcArc
	}
	if prev == IPhoneOS {
		return current >= IOS_13_7 && current <= IOSLatest
	}
	if prev == IPadOS {
		return current >= IOS_13_7 && current <= IOSLatest
	}
	if prev == Android {
		return current == Mobile
	}

	// Processor architecture and OS Architecture (e.g., Win64 should be followed by x64)
	if prev == Win64Arc {
		return current == X64ProcArc
	}

	// Processor architecture should be followed by Revision number or Rendering engine
	if prev == X64ProcArc || prev == X86_64ProcArc {
		return current >= FirefoxRevision_99_0 && current <= FirefoxRevisionLatest ||
			current == AppleWebKitRenderingEngine
	}

	// OS version should be followed by Rendering engine or Revision number
	if prev >= MacOS_10_12_6 && prev <= MacOSLatest {
		return current >= FirefoxRevision_99_0 && current <= FirefoxRevisionLatest || current == AppleWebKitRenderingEngine
	}

	// iOS version should be followed by Rendering engine
	if prev >= IOS_13_7 && prev <= IOSLatest {
		return current == AppleWebKitRenderingEngine
	}

	// Revision number should be followed by Rendering engine
	if prev >= FirefoxRevision_99_0 && prev <= FirefoxRevisionLatest {
		return current == GeckoRenderingEngine
	}

	// Rendering engine should be followed by Render version
	if prev == GeckoRenderingEngine {
		return current >= Gecko_20100101 && current <= GeckoLatest
	}
	if prev == AppleWebKitRenderingEngine {
		return current >= AppleWebKit_604_1_38 && current <= AppleWebKitLatest
	}

	// Gecko render version should be followed by Browser name
	if prev >= Gecko_20100101 && prev <= GeckoLatest {
		return current == FirefoxBrowser
	}

	// Apple WebKit version should be followed by additional info
	if prev >= AppleWebKit_604_1_38 && prev <= AppleWebKitLatest {
		return current == KHTMLAdditionalInfo
	}

	// Browser name should be followed by Browser version
	if prev == FirefoxBrowser || prev == FirefoxMobileBrowser {
		return current >= Firefox_99_0 && current <= FirefoxLatest
	}
	if prev == SafariBrowser {
		if collapsed >= AppleWebKit_604_1 && collapsed <= AppleWebKitLatest {
			return current == collapsed
		}
		return current >= AppleWebKit_604_1 && current <= AppleWebKitLatest
	}

	// Additional info should be followed by mobile Browser or Mobile token
	if prev == KHTMLAdditionalInfo {
		return current == FirefoxMobileBrowser || current == SafariBrowserVersion
	}

	// Safari version token should be followed by version token
	if prev == SafariBrowserVersion {
		return current >= Safari_15_6_1 && current <= SafariLatest
	}

	// Safari version should be followed by Safari browser
	if prev >= Safari_15_6_1 && prev <= SafariLatest {
		return current == SafariBrowser
	}

	// Mobile token should be followed by Mobile version
	if prev == Mobile {
		if collapsed == Android {
			return current <= Mobile_15E148 && current >= MobileLatest
		}
		return current >= Mobile_15E148 && current <= MobileLatest
	}

	// Mobile version should be followed by Browser token
	if prev >= Mobile_15E148 && prev <= MobileLatest {
		return current == SafariBrowser
	}

	// Mobile Firefox should be followed by Mobile token
	if depth == 2 && collapsed == FirefoxMobileBrowser {
		return current == Mobile
	}

	return false
}

func contain(tt TokenType, tts []TokenType) bool {
	for _, t := range tts {
		if t == tt {
			return true
		}
	}
	return false
}

func isInterpolatable(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '%' && s[i+1] == 's' {
			return true
		}
	}
	return false
}

func main() {
	ua := NewUserAgent(20)
	println(ua.Header)
}
