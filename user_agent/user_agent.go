package user_agent

import (
	"math/rand"
	"strings"

	"golang.org/x/exp/slices"
)

type TokenType int

const (
	_ TokenType = iota
	// Browser identifier token types
	Mozilla5BrowserIdentifier
	// Window system token types
	X11WindowSystem
	// Device type token types
	MacintoshDevice
	IPhoneDevice
	IPadDevice
	// OS token types
	StartMacOS
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
	EndMacOS
	StartMacOSFirefox // MacOS version displayed in Firefox on desktop
	MacOS_F_10_14
	MacOS_F_10_15
	MacOS_F_11_5
	MacOS_F_11_6
	MacOS_F_11_7
	MacOS_F_12_3
	MacOS_F_12_5
	MacOS_F_12_6
	MacOS_F_13_1
	MacOS_F_13_2
	MacOS_F_13_3
	MacOS_F_13_5
	EndMacOSFirefox
	StartWindows
	WindowsNT_10_0
	EndWindows
	Linux
	StartIPhoneOS
	IPhoneOS_13_7
	IPhoneOS_14_8_1
	IPhoneOS_15_7_9
	IPhoneOS_16_6
	IPhoneOS_16_6_1
	IPhoneOS_17_0
	EndIPhoneOS
	StartIPadOS
	IPadOS_13_7
	IPadOS_14_8_1
	IPadOS_15_7_9
	IPadOS_16_6
	IPadOS_16_6_1
	IPadOS_17_0
	EndIPadOS
	StartAndroid
	Android_11
	Android_12
	Android_13
	EndAndroid
	// Android device models
	StartAndroidDeviceModel
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
	Redmi_Note_8_Pro
	Redmi_Note_9_Pro
	M2101K6G   // Redmi Note 10 Pro
	M2102J20SG // Xiaomi Poco X3 Pro
	DE2118     // OnePlus Nord N200 5G
	EndAndroidDeviceModel
	// OS architecture token types
	Win64Arc
	// Processor architecture token types
	X64ProcArc
	X86_64ProcArc
	// Revision number token types
	StartRevision
	Revision_99_0
	Revision_102_0
	Revision_105_0
	EndRevision
	// Rendering engine token types
	StartGeckoMobile
	GeckoMobile_99_0
	GeckoMobile_102_0
	GeckoMobile_105_0
	EndGeckoMobile
	StartGeckoPC
	Gecko_20100101
	EndGeckoPC
	StartAppleWebKit
	AppleWebKit_605_1_15
	EndAppleWebKit
	// Additional info token types
	KHTMLAdditionalInfo
	// Browser name token types
	StartFirefox
	Firefox_99_0
	Firefox_102_0
	Firefox_105_0
	EndFirefox
	StartFirefoxMobile
	FirefoxMobile_99_0
	FirefoxMobile_102_0
	FirefoxMobile_105_0
	EndFirefoxMobile
	StartSafari
	Safari_15_6
	Safari_16_5
	EndSafari
	StartSafariWebKit
	SafariWebKit_605_1_15
	EndSafariWebKit
	// Mobile token types
	Mobile
	StartMobile
	Mobile_15E148
	EndMobile

	TotalTokens
)

func (t TokenType) String() string {
	switch t {
	case Mozilla5BrowserIdentifier:
		return "Mozilla/5.0"
	case X11WindowSystem:
		return "(X11;"
	case MacintoshDevice:
		return "(Macintosh;"
	case IPhoneDevice:
		return "(iPhone;"
	case IPadDevice:
		return "(iPad;"
	case MacOS_10_14_6:
		return "Intel Mac OS X 10_14_6)"
	case MacOS_10_15_7:
		return "Intel Mac OS X 10_15_7)"
	case MacOS_11_5:
		return "Intel Mac OS X 11_5)"
	case MacOS_11_6_8:
		return "Intel Mac OS X 11_6_8)"
	case MacOS_11_7_9:
		return "Intel Mac OS X 11_7_9)"
	case MacOS_12_3_1:
		return "Intel Mac OS X 12_3_1)"
	case MacOS_12_5_1:
		return "Intel Mac OS X 12_5_1)"
	case MacOS_12_6_8:
		return "Intel Mac OS X 12_6_8)"
	case MacOS_13_1:
		return "Intel Mac OS X 13_1)"
	case MacOS_13_2_1:
		return "Intel Mac OS X 13_2_1)"
	case MacOS_13_3_1:
		return "Intel Mac OS X 13_3_1)"
	case MacOS_13_5:
		return "Intel Mac OS X 13_5)"
	case MacOS_13_5_1:
		return "Intel Mac OS X 13_5_1)"
	case MacOS_F_10_14:
		return "Intel Mac OS X 10.14)"
	case MacOS_F_10_15:
		return "Intel Mac OS X 10.15)"
	case MacOS_F_11_5:
		return "Intel Mac OS X 11.5)"
	case MacOS_F_11_6:
		return "Intel Mac OS X 11.6)"
	case MacOS_F_11_7:
		return "Intel Mac OS X 11.7)"
	case MacOS_F_12_3:
		return "Intel Mac OS X 12.3)"
	case MacOS_F_12_5:
		return "Intel Mac OS X 12.5)"
	case MacOS_F_12_6:
		return "Intel Mac OS X 12.6)"
	case MacOS_F_13_1:
		return "Intel Mac OS X 13.1)"
	case MacOS_F_13_2:
		return "Intel Mac OS X 13.2)"
	case MacOS_F_13_3:
		return "Intel Mac OS X 13.3)"
	case MacOS_F_13_5:
		return "Intel Mac OS X 13.5)"
	case WindowsNT_10_0:
		return "(Windows NT 10.0;"
	case Linux:
		return "Linux"
	case IPhoneOS_13_7:
		return "CPU iPhone OS 13_7 like Mac OS X)"
	case IPhoneOS_14_8_1:
		return "CPU iPhone OS 14_8_1 like Mac OS X)"
	case IPhoneOS_15_7_9:
		return "CPU iPhone OS 15_7_9 like Mac OS X)"
	case IPhoneOS_16_6:
		return "CPU iPhone OS 16_6 like Mac OS X)"
	case IPhoneOS_16_6_1:
		return "CPU iPhone OS 16_6_1 like Mac OS X)"
	case IPhoneOS_17_0:
		return "CPU iPhone OS 17_0 like Mac OS X)"
	case IPadOS_13_7:
		return "CPU OS 13_7 like Mac OS X)"
	case IPadOS_14_8_1:
		return "CPU OS 14_8_1 like Mac OS X)"
	case IPadOS_15_7_9:
		return "CPU OS 15_7_9 like Mac OS X)"
	case IPadOS_16_6:
		return "CPU OS 16_6 like Mac OS X)"
	case IPadOS_16_6_1:
		return "CPU OS 16_6_1 like Mac OS X)"
	case IPadOS_17_0:
		return "CPU OS 17_0 like Mac OS X)"
	case Android_11:
		return "(Android 11;"
	case Android_12:
		return "(Android 12;"
	case Android_13:
		return "(Android 13;"
	case SM_G973F:
		return "SM-G973F;"
	case SM_G973U:
		return "SM-G973U;"
	case SM_A515F:
		return "SM-A515F;"
	case SM_A515U:
		return "SM-A515U;"
	case SM_A536B:
		return "SM-A536B;"
	case SM_A536U:
		return "SM-A536U;"
	case SM_G991B:
		return "SM-G991B;"
	case SM_G991U:
		return "SM-G991U;"
	case SM_G998B:
		return "SM-G998B;"
	case SM_G998U:
		return "SM-G998U;"
	case SM_S901B:
		return "SM-S901B;"
	case SM_S901U:
		return "SM-S901U;"
	case SM_S908B:
		return "SM-S908B;"
	case SM_S908U:
		return "SM-S908U;"
	case Pixel_6:
		return "Pixel 6;"
	case Pixel_6a:
		return "Pixel 6a;"
	case Pixel_6_Pro:
		return "Pixel 6 Pro;"
	case Pixel_7:
		return "Pixel 7;"
	case Pixel_7_Pro:
		return "Pixel 7 Pro;"
	case Moto_G_Pure:
		return "Moto G Pure;"
	case Moto_G_Stylus_5G:
		return "Moto G Stylus 5G;"
	case Moto_G_Stylus_5G_2022:
		return "Moto G Stylus 5G (2022);"
	case Moto_G_5G_2022:
		return "Moto G 5G (2022);"
	case Moto_G_Power_2021:
		return "Moto G Power (2021);"
	case Moto_G_Power_2022:
		return "Moto G Power (2022);"
	case Redmi_Note_8_Pro:
		return "Redmi Note 8 Pro;"
	case Redmi_Note_9_Pro:
		return "Redmi Note 9 Pro;"
	case M2101K6G:
		return "M2101K6G;"
	case M2102J20SG:
		return "M2102J20SG;"
	case DE2118:
		return "DE2118;"
	case Win64Arc:
		return "Win64;"
	case X64ProcArc:
		return "x64;"
	case X86_64ProcArc:
		return "x86_64;"
	case Revision_99_0:
		return "rv:99.0)"
	case Revision_102_0:
		return "rv:102.0)"
	case Revision_105_0:
		return "rv:105.0)"
	case GeckoMobile_99_0:
		return "Gecko/99.0"
	case GeckoMobile_102_0:
		return "Gecko/102.0"
	case GeckoMobile_105_0:
		return "Gecko/105.0"
	case Gecko_20100101:
		return "Gecko/20100101"
	case AppleWebKit_605_1_15:
		return "AppleWebKit/605.1.15"
	case KHTMLAdditionalInfo:
		return "(KHTML, like Gecko)"
	case Firefox_99_0:
		return "Firefox/99.0"
	case Firefox_102_0:
		return "Firefox/102.0"
	case Firefox_105_0:
		return "Firefox/105.0"
	case FirefoxMobile_99_0:
		return "FxiOS/99.0"
	case FirefoxMobile_102_0:
		return "FxiOS/102.0"
	case FirefoxMobile_105_0:
		return "FxiOS/105.0"
	case Safari_15_6:
		return "Version/15.6"
	case Safari_16_5:
		return "Version/16.5"
	case SafariWebKit_605_1_15:
		return "Safari/605.1.15"
	case Mobile:
		return "Mobile"
	case Mobile_15E148:
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

func NewToken(seed int64, allowedTokens ...TokenType) *Token {
	possibilities := make([]TokenType, 0, TotalTokens)
	if len(allowedTokens) > 0 {
		possibilities = make([]TokenType, len(allowedTokens))
		copy(possibilities, allowedTokens)
	} else {
		for i := TokenType(0); i < TotalTokens; i++ {
			possibilities = append(possibilities, TokenType(i))
		}
	}

	return &Token{
		Possibilities: possibilities,
		rand:          rand.New(rand.NewSource(seed)),
	}
}

func NewUserAgent(length int, seed int64, allowedTokens ...TokenType) *UserAgent {
	tokens := make([]*Token, length)
	for i := range tokens {
		tokens[i] = NewToken(seed, allowedTokens...)
	}
	tokens[0].Possibilities = []TokenType{Mozilla5BrowserIdentifier}
	header, client, version := "", "", ""

	for i, t := range tokens {
		tt := t.Collapse()
		if tt == 0 {
			break
		}
		if i > 0 {
			header += " "
		}
		header += tt.String()

		switch {
		case in(tt, StartFirefox, EndFirefox) || in(tt, StartFirefoxMobile, EndFirefoxMobile):
			split := strings.Split(tt.String(), "/")
			client = split[0]
			version = split[1]
		case tt >= StartSafari && tt <= EndSafari:
			client = "Safari"
			version = strings.Split(tt.String(), "/")[1]
		}

		for j := i + 1; j < len(tokens); j++ {
			tokens[j].Observe(t, tokens[j-1])
		}
	}

	return &UserAgent{
		Header:  header,
		Client:  client,
		Version: version,
	}
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
	androidDevicesLimit := func(collapsed, current TokenType) bool {
		switch collapsed {
		case Android_11:
			return slices.Contains(
				[]TokenType{
					SM_G973F, SM_G973U, SM_A515F, SM_A515U, SM_G991B, SM_G991U,
					SM_G998B, SM_G998U, Moto_G_Pure, Moto_G_Stylus_5G, Moto_G_Power_2021,
					Moto_G_Power_2022, Redmi_Note_8_Pro, Redmi_Note_9_Pro, M2101K6G, M2102J20SG,
					DE2118,
				},
				current,
			)
		case Android_12:
			return slices.Contains(
				[]TokenType{
					SM_G973F, SM_G973U, SM_A515F, SM_A515U, SM_A536B, SM_A536U,
					SM_G991B, SM_G991U, SM_G998B, SM_G998U, SM_S901B, SM_S901U,
					SM_S908B, SM_S908U, Pixel_6, Pixel_6a, Pixel_6_Pro, Moto_G_Pure,
					Moto_G_Stylus_5G, Moto_G_Stylus_5G_2022, Moto_G_5G_2022, Moto_G_Power_2022,
					Redmi_Note_9_Pro, M2101K6G, M2102J20SG, DE2118,
				},
				current,
			)
		case Android_13:
			return slices.Contains(
				[]TokenType{
					SM_A515F, SM_A515U, SM_A536B, SM_A536U,
					SM_G991B, SM_G991U, SM_G998B, SM_G998U, SM_S901B, SM_S901U,
					SM_S908B, SM_S908U, Pixel_6, Pixel_6a, Pixel_6_Pro, Pixel_7,
					Pixel_7_Pro, Moto_G_Stylus_5G_2022, Moto_G_5G_2022, M2101K6G,
				},
				current,
			)
		default:
			return false
		}
	}
	geckoVersionLimit := func(prev, current TokenType) bool {
		// Gecko version must correspond to revision number
		switch prev {
		case Revision_99_0:
			return current == GeckoMobile_99_0
		case Revision_102_0:
			return current == GeckoMobile_102_0
		case Revision_105_0:
			return current == GeckoMobile_105_0
		default:
			return false
		}
	}
	firefoxMobileVersionLimit := func(prev, current TokenType) bool {
		// Firefox version must correspond to Gecko version
		switch prev {
		case GeckoMobile_99_0:
			return current == Firefox_99_0
		case GeckoMobile_102_0:
			return current == Firefox_102_0
		case GeckoMobile_105_0:
			return current == Firefox_105_0
		default:
			return false
		}
	}
	firefoxPCVersionLimit := func(collapsed, current TokenType) bool {
		// Firefox version must correspond to revision version
		switch collapsed {
		case Revision_99_0:
			return current == Firefox_99_0
		case Revision_102_0:
			return current == Firefox_102_0
		case Revision_105_0:
			return current == Firefox_105_0
		default:
			return false
		}
	}

	switch {
	// Browser identifier must be followed by Window system, Device Type, or OS type
	case prev == Mozilla5BrowserIdentifier:
		return current == X11WindowSystem ||
			iin(current, MacintoshDevice, IPadDevice) ||
			in(current, StartWindows, EndWindows) ||
			in(current, StartAndroid, EndAndroid)

	// Window system must be followed by OS (Linux)
	case prev == X11WindowSystem:
		return current == Linux

	// Macintosh devices should be followed by MacOS
	case prev == MacintoshDevice:
		return in(current, StartMacOS, EndMacOS) || in(current, StartMacOSFirefox, EndMacOSFirefox)

	// iPhone or iPad should be followed by corresponding OS
	case prev == IPhoneDevice:
		return in(current, StartIPhoneOS, EndIPhoneOS)
	case prev == IPadDevice:
		return in(current, StartIPadOS, EndIPadOS)

	// Linux should be followed by processor architecture
	case prev == Linux:
		return current == X86_64ProcArc

	// Android should be followed by device model
	case in(prev, StartAndroid, EndAndroid):
		return current == Mobile

	// Mobile should be followed by device model or a specific set of devices based on the collapsed token
	case prev == Mobile:
		return androidDevicesLimit(collapsed, current) || in(current, StartAndroidDeviceModel, EndAndroidDeviceModel)

	// Device model should be followed by revision number or rendering engine
	case in(prev, StartAndroidDeviceModel, EndAndroidDeviceModel):
		return in(current, StartRevision, EndRevision)

	// Windows should be followed by OS architecture
	case prev == WindowsNT_10_0:
		return current == Win64Arc

	// Processor architecture and OS Architecture (e.g., Win64 should be followed by x64)
	case prev == Win64Arc:
		return current == X64ProcArc

	// Processor architecture should be followed by revision or rendering engine
	case prev == X64ProcArc, prev == X86_64ProcArc:
		return in(current, StartRevision, EndRevision)

	// OS should be followed by rendering engine or revision number
	case in(prev, StartMacOS, EndMacOS):
		return in(current, StartAppleWebKit, EndAppleWebKit)
	case in(prev, StartMacOSFirefox, EndMacOSFirefox):
		return in(current, StartRevision, EndRevision)
	case in(prev, StartIPhoneOS, EndIPhoneOS), in(prev, StartIPadOS, EndIPadOS):
		return in(current, StartAppleWebKit, EndAppleWebKit)

	// Revision should be followed by rendering engine
	case in(prev, StartRevision, EndRevision):
		if collapsed == Mobile {
			return geckoVersionLimit(prev, current)
		}
		if slices.Contains([]TokenType{X11WindowSystem, WindowsNT_10_0, MacintoshDevice}, collapsed) {
			return in(current, StartGeckoPC, EndGeckoPC)
		}
		return in(current, StartGeckoMobile, EndGeckoPC)

	// Gecko renderer should be followed by browser
	case in(prev, StartGeckoPC, EndGeckoPC):
		return firefoxPCVersionLimit(collapsed, current) || in(current, StartFirefox, EndFirefox)
	case in(prev, StartGeckoMobile, EndGeckoMobile):
		return firefoxMobileVersionLimit(prev, current)

	// Apple WebKit version should be followed by additional info
	case in(prev, StartAppleWebKit, EndAppleWebKit):
		return current == KHTMLAdditionalInfo

	// Additional info should be followed by mobile browser or Safari
	case prev == KHTMLAdditionalInfo:
		if collapsed == MacintoshDevice {
			return in(current, StartSafari, EndSafari)
		}
		return in(current, StartFirefoxMobile, EndFirefoxMobile) || in(current, StartSafari, EndSafari)

	// Mobile Browser should be followed by mobile build
	case in(prev, StartFirefoxMobile, EndFirefoxMobile):
		return in(current, StartMobile, EndMobile)

	// Safari should be followed by Safari-WebKit version and mobile build should be followed by browser token
	case in(prev, StartSafari, EndSafari):
		switch collapsed {
		case IPhoneDevice, IPadDevice:
			return in(current, StartMobile, EndMobile)
		case MacintoshDevice:
			return !in(current, StartMobile, EndMobile)
		default:
			return in(current, StartSafariWebKit, EndSafariWebKit) || in(current, StartMobile, EndMobile)
		}
	case in(prev, StartMobile, EndMobile):
		return in(current, StartSafariWebKit, EndSafariWebKit)

	default:
		return false
	}
}

func iin(val, start, end TokenType) bool {
	return val >= start && val <= end
}

func in(val, start, end TokenType) bool {
	return val > start && val < end
}
