package main

import (
	"math/rand"
	"time"
)

type BrowserIdentifierTokenType uint8

const (
	Mozilla5BrowserIdentifier BrowserIdentifierTokenType = iota
	TotalBrowserIdentifiers
)

func (t *BrowserIdentifierTokenType) String() string {
	names := []string{
		"Mozilla/5.0",
	}
	if *t < BrowserIdentifierTokenType(len(names)) {
		return names[*t]
	}
	return ""
}

type WindowSystemTokenType uint8

const (
	X11WindowSystem WindowSystemTokenType = iota
	TotalWindowSystems
)

func (t *WindowSystemTokenType) String() string {
	names := []string{
		"X11",
	}
	if *t < WindowSystemTokenType(len(names)) {
		return names[*t]
	}
	return ""
}

type DeviceTypeTokenType uint8

const (
	MacintoshDevice DeviceTypeTokenType = iota
	IPhoneDevice
	IPadDevice
	TotalDeviceTypes
)

func (t *DeviceTypeTokenType) String() string {
	names := []string{
		"Macintosh",
		"iPhone",
		"iPad",
	}
	if *t < DeviceTypeTokenType(len(names)) {
		return names[*t]
	}
	return ""
}

type OSTokenType uint8

const (
	MacOSX OSTokenType = iota
	WindowsNT
	IPhoneOS
	IPadOS
	Linux
	Android
	TotalOSs
)

func (t *OSTokenType) String() string {
	names := []string{
		"Mac OS X",
		"Windows NT",
		"iPhone OS",
		"iPad OS",
		"Linux",
		"Android",
	}
	if *t < OSTokenType(len(names)) {
		return names[*t]
	}
	return ""
}

type OSVersionNumberTokenType uint16

const (
	// Mac OS
	MacOS_10_12_6 OSVersionNumberTokenType = iota // Sierra
	MacOSSierraLatest
	MacOS_10_13_6 // High Sierra
	MacOSHighSierraLatest
	MacOS_11_7_9 // Big Sur
	MacOSBigSurLatest
	MacOS_12_6_8 // Monterey
	MacOSMontereyLatest
	MacOS_13_5_1 // Ventura
	MacOSVenturaLatest
	TotalOSVersionNumbers
)

func (t *OSVersionNumberTokenType) String() string {
	names := []string{
		"10_12_6",
		"10_12_6",
		"10_13_6",
		"10_13_6",
		"11_7_9",
		"11_7_9",
		"12_6_8",
		"12_6_8",
		"13_5_1",
		"13_5_1",
	}
	if *t < OSVersionNumberTokenType(len(names)) {
		return names[*t]
	}
	return ""
}

type OSArcTokenType uint8

const (
	Win64Arc OSArcTokenType = iota
	TotalOSArcs
)

func (t *OSArcTokenType) String() string {
	names := []string{
		"Win64",
	}
	if *t < OSArcTokenType(len(names)) {
		return names[*t]
	}
	return ""
}

type ProcArcTokenType uint8

const (
	X64ProcArc ProcArcTokenType = iota
	X86_64ProcArc
	TotalProcArcs
)

func (t *ProcArcTokenType) String() string {
	names := []string{
		"x64",
		"x86_64",
	}
	if *t < ProcArcTokenType(len(names)) {
		return names[*t]
	}
	return ""
}

type RevisionTokenType uint8

const (
	FirefoxRevision RevisionTokenType = iota
	TotalRevisions
)

func (t *RevisionTokenType) String() string {
	names := []string{
		"rv:",
	}
	if *t < RevisionTokenType(len(names)) {
		return names[*t]
	}
	return ""
}

type RevisionNumberTokenType uint8

const (
	FirefoxRevision_99_0 RevisionNumberTokenType = iota
	FirefoxRevision_102_0
	FirefoxRevision_105_0
	TotalRevisionNumbers
)

func (t *RevisionNumberTokenType) String() string {
	names := []string{
		"99.0",
		"102.0",
		"105.0",
	}
	if *t < RevisionNumberTokenType(len(names)) {
		return names[*t]
	}
	return ""
}

type RenderingEngineTokenType uint8

const (
	GeckoRenderingEngine RenderingEngineTokenType = iota
	AppleWebKitRenderingEngine
	TotalRenderingEngines
)

func (t *RenderingEngineTokenType) String() string {
	names := []string{
		"Gecko",
		"AppleWebKit",
	}
	if *t < RenderingEngineTokenType(len(names)) {
		return names[*t]
	}
	return ""
}

type RenderingEngineVersionTokenType uint8

const (
	Gecko_20100101 RenderingEngineVersionTokenType = iota
	AppleWebKit_537_36
	AppleWebKit_605_1_15
	TotalRenderingEngineVersions
)

func (t *RenderingEngineVersionTokenType) String() string {
	names := []string{
		"20100101",
		"537.36",
		"605.1.15",
	}
	if *t < RenderingEngineVersionTokenType(len(names)) {
		return names[*t]
	}
	return ""
}

type AdditionalInfoTokenType uint8

const (
	KHTMLAdditionalInfo AdditionalInfoTokenType = iota
	TotalAdditionalInfos
)

func (t *AdditionalInfoTokenType) String() string {
	names := []string{
		"KHTML",
	}
	if *t < AdditionalInfoTokenType(len(names)) {
		return names[*t]
	}
	return ""
}

type BrowserNameTokenType uint8

const (
	FirefoxBrowser BrowserNameTokenType = iota
	FxiOSBrowser
	SafariBrowser
	BraveBrowser
	TotalBrowsers
)

func (t *BrowserNameTokenType) String() string {
	names := []string{
		"Firefox",
		"FxiOS",
		"Safari",
		"Brave",
	}
	if *t < BrowserNameTokenType(len(names)) {
		return names[*t]
	}
	return ""
}

type BrowserVersionTokenType uint8

const (
	Firefox_99_0 BrowserVersionTokenType = iota
	Firefox_102_0
	Firefox_105_0
	Safari_604_1_38
	Safari_604_1
	Brave_1_29_81
	TotalBrowserVersions
)

func (t *BrowserVersionTokenType) String() string {
	names := []string{
		"99.0",
		"102.0",
		"105.0",
		"604.1.38",
		"604.1",
		"1.29.81",
	}
	if *t < BrowserVersionTokenType(len(names)) {
		return names[*t]
	}
	return ""
}

type MobileTokenType uint8

const (
	MobileMobile MobileTokenType = iota
	TotalMobiles
)

func (t *MobileTokenType) String() string {
	names := []string{
		"Mobile",
	}
	if *t < MobileTokenType(len(names)) {
		return names[*t]
	}
	return ""
}

type MobileBuildNumberTokenType uint8

const (
	Mobile_15E148 MobileBuildNumberTokenType = iota
	TotalMobileBuildNumbers
)

func (t *MobileBuildNumberTokenType) String() string {
	names := []string{
		"15E148",
	}
	if *t < MobileBuildNumberTokenType(len(names)) {
		return names[*t]
	}
	return ""
}

type Token struct {
	Possibilities []interface{}
	rand          *rand.Rand
}

func NewToken() *Token {
	possibilities := []interface{}{}
	for i := BrowserIdentifierTokenType(0); i < TotalBrowserIdentifiers; i++ {
		possibilities = append(possibilities, BrowserIdentifierTokenType(i))
	}
	for i := WindowSystemTokenType(0); i < TotalWindowSystems; i++ {
		possibilities = append(possibilities, WindowSystemTokenType(i))
	}
	for i := DeviceTypeTokenType(0); i < TotalDeviceTypes; i++ {
		possibilities = append(possibilities, DeviceTypeTokenType(i))
	}
	for i := OSTokenType(0); i < TotalOSs; i++ {
		possibilities = append(possibilities, OSTokenType(i))
	}
	for i := OSVersionNumberTokenType(0); i < TotalOSVersionNumbers; i++ {
		possibilities = append(possibilities, OSVersionNumberTokenType(i))
	}
	for i := OSArcTokenType(0); i < TotalOSArcs; i++ {
		possibilities = append(possibilities, OSArcTokenType(i))
	}
	for i := ProcArcTokenType(0); i < TotalProcArcs; i++ {
		possibilities = append(possibilities, ProcArcTokenType(i))
	}
	for i := RevisionTokenType(0); i < TotalRevisions; i++ {
		possibilities = append(possibilities, RevisionTokenType(i))
	}
	for i := RevisionNumberTokenType(0); i < TotalRevisionNumbers; i++ {
		possibilities = append(possibilities, RevisionNumberTokenType(i))
	}
	for i := RenderingEngineTokenType(0); i < TotalRenderingEngines; i++ {
		possibilities = append(possibilities, RenderingEngineTokenType(i))
	}
	for i := RenderingEngineVersionTokenType(0); i < TotalRenderingEngineVersions; i++ {
		possibilities = append(possibilities, RenderingEngineVersionTokenType(i))
	}
	for i := AdditionalInfoTokenType(0); i < TotalAdditionalInfos; i++ {
		possibilities = append(possibilities, AdditionalInfoTokenType(i))
	}
	for i := BrowserNameTokenType(0); i < TotalBrowsers; i++ {
		possibilities = append(possibilities, BrowserNameTokenType(i))
	}
	for i := BrowserVersionTokenType(0); i < TotalBrowserVersions; i++ {
		possibilities = append(possibilities, BrowserVersionTokenType(i))
	}
	for i := MobileTokenType(0); i < TotalMobiles; i++ {
		possibilities = append(possibilities, MobileTokenType(i))
	}
	for i := MobileBuildNumberTokenType(0); i < TotalMobileBuildNumbers; i++ {
		possibilities = append(possibilities, MobileBuildNumberTokenType(i))
	}

	return &Token{
		Possibilities: possibilities,
		rand:          rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (t *Token) Observe(in *Token) {
	for _, possibility := range in.Possibilities {
		switch v := possibility.(type) {
		case DeviceTypeTokenType:
			switch v {
			case IPhoneDevice:
				t.removePossibilities(OSTokenType(0), TotalOSs, IPhoneOS)
			case IPadDevice:
				t.removePossibilities(OSTokenType(0), TotalOSs, IPadOS)
			case MacintoshDevice:
				t.removePossibilities(OSTokenType(0), TotalOSs, MacOSX)
			}

		case OSTokenType:
			switch v {
			case IPhoneOS:
				t.removePossibilities(DeviceTypeTokenType(0), TotalDeviceTypes, IPhoneDevice)
			case IPadOS:
				t.removePossibilities(DeviceTypeTokenType(0), TotalDeviceTypes, IPadDevice)
			case MacOSX:
				t.removePossibilities(DeviceTypeTokenType(0), TotalDeviceTypes, MacintoshDevice)
			case WindowsNT:
				t.removePossibilities(OSArcTokenType(0), TotalOSArcs, Win64Arc)
			}

		case BrowserNameTokenType:
			switch v {
			case SafariBrowser:
				// Removing non-Apple possibilities
				t.removePossibilities(DeviceTypeTokenType(0), TotalDeviceTypes, MacintoshDevice, IPhoneDevice, IPadDevice)
				t.removePossibilities(OSTokenType(0), TotalOSs, MacOSX, IPhoneOS, IPadOS)
			case FirefoxBrowser:
				t.removePossibility(FxiOSBrowser)
			case FxiOSBrowser:
				t.removePossibility(FirefoxBrowser)
			}
		}
	}
}

// removePossibilities is a helper function that removes all possibilities within the range [start, end], except those in the `keep` list
func (t *Token) removePossibilities(start, end interface{}, keep ...interface{}) {
	toRemove := []interface{}{}
	for _, possibility := range t.Possibilities {
		if isBetween(possibility, start, end) && !contains(keep, possibility) {
			toRemove = append(toRemove, possibility)
		}
	}

	// Actually remove the values
	for _, rem := range toRemove {
		t.removePossibility(rem)
	}
}

// removePossibility is a helper function that removes a specific possibility
func (t *Token) removePossibility(val interface{}) {
	index := -1
	for i, possibility := range t.Possibilities {
		if possibility == val {
			index = i
			break
		}
	}
	if index != -1 {
		t.Possibilities = append(t.Possibilities[:index], t.Possibilities[index+1:]...)
	}
}

// isBetween checks if a value is between two given boundaries
func isBetween(val, start, end interface{}) bool {
	return val.(int) >= start.(int) && val.(int) < end.(int)
}

// contains checks if a list contains a specific value
func contains(list []interface{}, val interface{}) bool {
	for _, item := range list {
		if item == val {
			return true
		}
	}
	return false
}

// func (t *Token) Collapse() (*TokenType, *string) {
// 	if len(t.Possibilities) == 0 {
// 		return nil, nil
// 	}
// 	p := t.Possibilities[t.rand.Intn(len(t.Possibilities))]
// 	p.Values = []string{p.Values[t.rand.Intn(len(p.Values))]}
// 	t.Possibilities = []*Possibility{p}
// 	return p.Type, &p.Values[0]
// }

// type UserAgent struct {
// 	Header  string
// 	Client  string
// 	Version string
// }

// func NewUserAgent() *UserAgent {
// 	t := NewToken()
// 	tt, tv := BrowserIdentifierToken.Collapse()
// 	Tokens := []*Token{
// 		NewToken(BrowserIdentifierToken, BrowserIdentifiers, nil),
// 	}
// }

// var UAParts = map[TokenType][]string{
// 	BrowserIdentifierToken: {
// 		"Mozilla/5.0",
// 	},
// 	WindowSystemToken: {
// 		"X11",
// 	},
// 	DeviceTypeToken: {
// 		"Macintosh", "Windows NT", "iPhone", "iPad", "Linux", "Android",
// 	},
// 	OperatingSystemToken: {
// 		"Intel Mac OS X", "CPU iPhone OS",
// 	},
// }

var (
	BrowserIdentifiers = []string{"Mozilla/5.0"}
	WindowSystems      = []string{"X11"}
	DeviceTypes        = []string{"Macintosh", "Windows NT", "iPhone", "iPad", "Linux", "Android"}
	OperatingSystems   = []string{
		"Intel Mac OS X 10_12_6", "Intel Mac OS X 11_0_1",
		"6.1; WOW64", "10.0; Win64; x64",
		"CPU iPhone OS 10_3_2 like Mac OS X",
		"CPU OS 10_3_2 like Mac OS X",
		"Android 4.3; GT-I9300 Build/JSS15J",
		"4.3; Mobile",
	}
	MacOSVersions = []string{"10_12_6", "11_0_1"}
	BrowserNames  = []string{"Firefox", "Safari", "Brave"}

	RenderingEngineMapping = map[string]string{
		"Firefox": "Gecko",
		"Safari":  "AppleWebKit",
		"Brave":   "AppleWebKit",
	}

	RenderVersionMapping = map[string][]string{
		"Firefox": {"rv:54.0", "rv:99.0"},
		"Safari":  {},
		"Brave":   {},
	}

	AdditionalInfos = map[string][]string{
		"AppleWebKit": {"(KHTML, like Gecko)", "(KHTML, like Gecko) Chrome/59.0.3071.115", "(KHTML, like Gecko) Chrome/59.0.3071.125 Mobile"},
		"Gecko":       {"Firefox/54.0", "Firefox/99.0"},
	}
)

// func (t *Token) Observe() string {
// 	return t.Options[rand.Intn(len(t.Options))]
// }

// func NewToken(TokenType TokenType, observedValue string) *Token {
// 	t := &Token{Type: TokenType}

// 	switch TokenType {
// 	case BrowserIdentifierToken:
// 		t.Options = BrowserIdentifiers
// 	case DeviceTypeToken:
// 		t.Options = DeviceTypes
// 	case OperatingSystemToken:
// 		t.Options = OperatingSystems[observedValue]
// 	case RenderVersionToken:
// 		t.Options = RenderVersionMapping[observedValue]
// 	case RenderingEngineToken:
// 		t.Options = []string{RenderingEngineMapping[observedValue]}
// 	case AdditionalInfoToken:
// 		t.Options = AdditionalInfos[observedValue]
// 	case BrowserNameToken:
// 		t.Options = BrowserNames
// 	case BrowserVersionToken:
// 		t.Options = []string{"54.0", "99.0", "603.3.8", "537.36"}
// 	}

// 	return t
// }

// func GenerateUserAgent() string {
// 	browserIdentifierToken := NewToken(BrowserIdentifierToken, "")
// 	browserIdentifierValue := browserIdentifierToken.Observe()

// 	deviceTypeToken := NewToken(DeviceTypeToken, "")
// 	deviceTypeValue := deviceTypeToken.Observe()

// 	operatingSystemToken := NewToken(OperatingSystemToken, deviceTypeValue)
// 	operatingSystemValue := operatingSystemToken.Observe()

// 	browserNameToken := NewToken(BrowserNameToken, "")
// 	browserNameValue := browserNameToken.Observe()

// 	renderingEngineToken := NewToken(RenderingEngineToken, browserNameValue)
// 	renderingEngineValue := renderingEngineToken.Observe()

// 	additionalInfoToken := NewToken(AdditionalInfoToken, renderingEngineValue)
// 	additionalInfoValue := additionalInfoToken.Observe()

// 	renderVersionToken := NewToken(RenderVersionToken, browserNameValue)
// 	renderVersionValue := renderVersionToken.Observe()

// 	browserVersionToken := NewToken(BrowserVersionToken, browserNameValue)
// 	browserVersionValue := browserVersionToken.Observe()

// 	return fmt.Sprintf("%s (%s; %s; %s) %s %s %s/%s", browserIdentifierValue, deviceTypeValue, operatingSystemValue, renderVersionValue, renderingEngineValue, additionalInfoValue, browserNameValue, browserVersionValue)
// }

func isInterpolatable(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '%' && s[i+1] == 's' {
			return true
		}
	}
	return false
}

func main() {
}
