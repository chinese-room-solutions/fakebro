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
	MacOS_10_14_6
	MacOS_10_15_7
	MacOS_11_5
	MacOS_11_6_8
	MacOS_11_7_9
	MacOS_12_3_1
	MacOS_12_5_1
	MacOS_12_6_8
	MacOS_13_1
	MacOS_13_2_1
	MacOS_13_3_1
	MacOS_13_5
	MacOS_13_5_1
	MacOSLatest
	WindowsNT_10_0
	WindowsLatest
	Linux
	IPhoneOS_13_7
	IPhoneOS_14_8_1
	IPhoneOS_15_7_8
	IPhoneOS_16_6
	IPhoneOSLatest
	IPadOS_13_7
	IPadOS_14_8_1
	IPadOS_15_7_8
	IPadOS_16_6
	IPadOSLatest
	Android_11
	Android_12
	Android_13
	AndroidLatest
	// Android device models
	SM_G973F // Galaxy S10
	SM_G973U
	SM_A515F // Galaxy A51
	SM_A515U
	SM_A536B // Galaxy A53
	SM_A536U
	SM_G991B // Galaxy S21
	SM_G991U
	SM_G998B
	SM_G998U
	SM_S901B // Galaxy S22
	SM_S901U
	SM_S908B
	SM_S908U
	Pixel_6
	Pixel_6a
	Pixel_6_Pro
	Pixel_7
	Pixel_7_Pro
	Moto_G_Pure
	Moto_G_Stylus_5G
	Moto_G_Stylus_5G_2022
	Moto_G_5G_2022
	Moto_G_Power_2021
	Moto_G_Power_2022
	// OS architecture token types
	Win64Arc
	// Processor architecture token types
	X64ProcArc
	X86_64ProcArc
	// Revision number token types
	Revision_99_0
	Revision_102_0
	Revision_105_0
	RevisionLatest
	// Rendering engine token types
	Gecko_99_0
	Gecko_102_0
	Gecko_105_0
	Gecko_20100101
	AppleWebKit_604_1
	AppleWebKit_605_1_15
	AppleWebKit_604_1_38
	AppleWebKitLatest
	// Additional info token types
	KHTMLAdditionalInfo
	// Browser name token types
	Firefox_99_0
	Firefox_102_0
	Firefox_105_0
	FirefoxLatest
	FirefoxMobile_99_0
	FirefoxMobile_102_0
	FirefoxMobile_105_0
	FirefoxMobileLatest
	Safari_16_5_2
	Safari_15_6_1
	SafariLatest
	SafariWebKit_604_1
	SafariWebKit_605_1_15
	SafariWebKit_604_1_38
	SafariWebKitLatest
	// Mobile token types
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
	case MacOS_11_5:
		return "Intel Mac OS X 11_5"
	case MacOS_11_6_8:
		return "Intel Mac OS X 11_6_8"
	case MacOS_11_7_9:
		return "Intel Mac OS X 11_7_9"
	case MacOS_12_3_1:
		return "Intel Mac OS X 12_3_1"
	case MacOS_12_5_1:
		return "Intel Mac OS X 12_5_1"
	case MacOS_12_6_8:
		return "Intel Mac OS X 12_6_8"
	case MacOS_13_1:
		return "Intel Mac OS X 13_1"
	case MacOS_13_2_1:
		return "Intel Mac OS X 13_2_1"
	case MacOS_13_3_1:
		return "Intel Mac OS X 13_3_1"
	case MacOS_13_5:
		return "Intel Mac OS X 13_5"
	case MacOS_13_5_1, MacOSLatest:
		return "Intel Mac OS X 13_5_1"
	case WindowsNT_10_0, WindowsLatest:
		return "Windows NT 10.0"
	case Linux:
		return "Linux"
	case IPhoneOS_13_7:
		return "CPU iPhone OS 13_7 like Mac OS X"
	case IPhoneOS_14_8_1:
		return "CPU iPhone OS 14_8_1 like Mac OS X"
	case IPhoneOS_15_7_8:
		return "CPU iPhone OS 15_7_8 like Mac OS X"
	case IPhoneOS_16_6, IPhoneOSLatest:
		return "CPU iPhone OS 16_6 like Mac OS X"
	case IPadOS_13_7:
		return "CPU OS 13_7 like Mac OS X"
	case IPadOS_14_8_1:
		return "CPU OS 14_8_1 like Mac OS X"
	case IPadOS_15_7_8:
		return "CPU OS 15_7_8 like Mac OS X"
	case IPadOS_16_6, IPadOSLatest:
		return "CPU OS 16_6 like Mac OS X"
	case Android_11:
		return "Android 11"
	case Android_12:
		return "Android 12"
	case Android_13, AndroidLatest:
		return "Android 13"
	case SM_G973F:
		return "SM-G973F"
	case SM_G973U:
		return "SM-G973U"
	case SM_A515F:
		return "SM-A515F"
	case SM_A515U:
		return "SM-A515U"
	case SM_A536B:
		return "SM-A536B"
	case SM_A536U:
		return "SM-A536U"
	case SM_G991B:
		return "SM-G991B"
	case SM_G991U:
		return "SM-G991U"
	case SM_G998B:
		return "SM-G998B"
	case SM_G998U:
		return "SM-G998U"
	case SM_S901B:
		return "SM-S901B"
	case SM_S901U:
		return "SM-S901U"
	case SM_S908B:
		return "SM-S908B"
	case SM_S908U:
		return "SM-S908U"
	case Pixel_6:
		return "Pixel 6"
	case Pixel_6a:
		return "Pixel 6a"
	case Pixel_6_Pro:
		return "Pixel 6 Pro"
	case Pixel_7:
		return "Pixel 7"
	case Pixel_7_Pro:
		return "Pixel 7 Pro"
	case Moto_G_Pure:
		return "Moto G Pure"
	case Moto_G_Stylus_5G:
		return "Moto G Stylus 5G"
	case Moto_G_Stylus_5G_2022:
		return "Moto G Stylus 5G (2022)"
	case Moto_G_5G_2022:
		return "Moto G 5G (2022)"
	case Moto_G_Power_2021:
		return "Moto G Power (2021)"
	case Moto_G_Power_2022:
		return "Moto G Power (2022)"
	case Win64Arc:
		return "Win64"
	case X64ProcArc:
		return "x64"
	case X86_64ProcArc:
		return "x86_64"
	case Revision_99_0:
		return "rv:99.0"
	case Revision_102_0:
		return "rv:102.0"
	case Revision_105_0, RevisionLatest:
		return "rv:105.0"
	case Gecko_99_0:
		return "Gecko/99.0"
	case Gecko_102_0:
		return "Gecko/102.0"
	case Gecko_105_0:
		return "Gecko/105.0"
	case Gecko_20100101:
		return "Gecko/20100101"
	case AppleWebKit_604_1:
		return "AppleWebKit/604.1"
	case AppleWebKit_605_1_15:
		return "AppleWebKit/605.1.15"
	case AppleWebKit_604_1_38, AppleWebKitLatest:
		return "AppleWebKit/604.1.38"
	case KHTMLAdditionalInfo:
		return "(KHTML, like Gecko)"
	case Firefox_99_0:
		return "Firefox/99.0"
	case Firefox_102_0:
		return "Firefox/102.0"
	case Firefox_105_0, FirefoxLatest:
		return "Firefox/105.0"
	case FirefoxMobile_99_0:
		return "FxiOS/99.0"
	case FirefoxMobile_102_0:
		return "FxiOS/102.0"
	case FirefoxMobile_105_0, FirefoxMobileLatest:
		return "FxiOS/105.0"
	case Safari_16_5_2:
		return "Version/16.5.2"
	case Safari_15_6_1, SafariLatest:
		return "Version/15.6.1"
	case SafariWebKit_604_1:
		return "Safari/604.1"
	case SafariWebKit_605_1_15:
		return "Safari/605.1.15"
	case SafariWebKit_604_1_38, SafariWebKitLatest:
		return "Safari/604.1.38"
	case Mobile_15E148, MobileLatest:
		return "Mobile/15E148"
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

		// if *tt == FirefoxBrowser || *tt == SafariBrowser {
		// 	Client = tt.String()
		// }

		if (*tt >= Firefox_99_0 && *tt <= FirefoxLatest) || (*tt >= Safari_15_6_1 && *tt <= SafariLatest) {
			Version = tt.String()
		}

		for j := i + 1; j < len(tokens); j++ {
			tokens[j].Observe(t, tokens[j-1])
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

func (t *Token) Observe(collapsed, prev *Token) {
	if prev == nil {
		return
	}

	reduced := []TokenType{}
	for _, currentType := range t.Possibilities {
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
	safariVersionLimit := func(collapsed, current TokenType) bool {
		// Safari WebKit version must correspond to Apple WebKit version
		if collapsed == AppleWebKit_604_1 {
			return current == SafariWebKit_604_1
		}
		if collapsed == AppleWebKit_604_1_38 {
			return current == SafariWebKit_604_1_38
		}
		if collapsed == AppleWebKit_605_1_15 {
			return current == SafariWebKit_605_1_15
		}
		return false
	}
	geckoVersionLimit := func(collapsed, current TokenType) bool {
		// Gecko version must correspond to revision number
		if collapsed == Revision_99_0 {
			return current == Gecko_99_0
		}
		if collapsed == Revision_102_0 {
			return current == Gecko_102_0
		}
		if collapsed == Revision_105_0 {
			return current == Gecko_105_0
		}
		return false
	}

	// Browser identifier must be followed by Window system, Device Type, or OS type
	if prev == Mozilla5BrowserIdentifier {
		return current == X11WindowSystem ||
			(current >= MacintoshDevice && current <= IPadDevice) ||
			(current == WindowsNT_10_0 || current == Linux)
	}

	// Window system must be followed by OS (Linux)
	if prev == X11WindowSystem {
		return current == Linux
	}

	// Macintosh devices should be followed by MacOS
	if prev == MacintoshDevice {
		return current >= MacOS_10_14_6 && current <= MacOSLatest
	}

	// iPhone or iPad should be followed by corresponding OS
	if prev == IPhoneDevice {
		return current >= IPhoneOS_13_7 && current <= IPhoneOSLatest
	}
	if prev == IPadDevice {
		return current >= IPadOS_13_7 && current <= IPadOSLatest
	}

	// Linux should be followed by processor architecture
	if prev == Linux {
		if collapsed == X11WindowSystem {
			return current == X86_64ProcArc
		}
		return current == X86_64ProcArc || current >= Android_11 && current <= AndroidLatest
	}

	// Android should be followed by device model
	if prev >= Android_11 && prev <= AndroidLatest {
		return current >= SM_G973F && current <= Moto_G_Power_2022
	}

	if prev >= SM_G973F && prev <= Moto_G_Power_2022 {
		return current >= AppleWebKit_604_1 && current <= AppleWebKitLatest || current >= Revision_99_0 && current <= RevisionLatest
	}

	// Windows should be followed by OS architecture
	if prev == WindowsNT_10_0 {
		return current == Win64Arc
	}

	// Processor architecture and OS Architecture (e.g., Win64 should be followed by x64)
	if prev == Win64Arc {
		return current == X64ProcArc
	}

	// Processor architecture should be followed by revision or rendering engine
	generalCond := current >= Revision_99_0 && current <= RevisionLatest
	if prev == X64ProcArc || prev == X86_64ProcArc {
		if collapsed >= WindowsNT_10_0 && collapsed <= WindowsLatest {
			return generalCond || current >= AppleWebKit_604_1 && current <= AppleWebKitLatest
		}
		return generalCond
	}

	// OS should be followed by revision number or rendering engine
	generalCond = current >= AppleWebKit_604_1 && current <= AppleWebKitLatest
	if prev >= MacOS_10_14_6 && prev <= MacOSLatest {
		return current >= Revision_99_0 && current <= RevisionLatest || generalCond
	}
	if prev >= IPhoneOS_13_7 && prev <= IPhoneOSLatest || prev >= IPadOS_13_7 && prev <= IPadOSLatest {
		return generalCond
	}

	// Revision should be followed by rendering engine
	if prev >= Revision_99_0 && prev <= RevisionLatest {
		if collapsed >= Android_11 && collapsed <= AndroidLatest {
			return current >= Gecko_99_0 && current <= Gecko_105_0
		}
		return current >= Gecko_20100101 && current <= Gecko_20100101
	}

	// Gecko renderer should be followed by browser
	if prev >= Gecko_20100101 && prev <= Gecko_20100101 {
		return current >= Firefox_99_0 && current <= FirefoxLatest
	}

	// Apple WebKit version should be followed by additional info
	if prev >= AppleWebKit_604_1 && prev <= AppleWebKitLatest {
		return current == KHTMLAdditionalInfo
	}

	// Additional info should be followed by mobile browser or Safari
	if prev == KHTMLAdditionalInfo {
		return current >= FirefoxMobile_99_0 && current <= FirefoxMobileLatest || current >= Safari_15_6_1 && current <= SafariLatest
	}

	// Mobile Browser should be followed by mobile build
	if prev >= FirefoxMobile_99_0 && prev <= FirefoxMobileLatest {
		return current >= Mobile_15E148 && current <= MobileLatest
	}

	// Safari should be followed by Safari-WebKit version and mobile build should be followed by browser token
	if (prev >= Safari_15_6_1 && prev <= SafariLatest) || (prev >= Mobile_15E148 && prev <= MobileLatest) {
		if collapsed >= AppleWebKit_604_1 && collapsed <= AppleWebKitLatest {
			return safariVersionLimit(collapsed, current)
		}
		return current >= SafariWebKit_604_1 && current <= SafariWebKitLatest
	}

	return false
}

func main() {
	ua := NewUserAgent(20)
	println(ua.Header)
}
