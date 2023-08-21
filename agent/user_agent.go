package agent

import (
	"fmt"
	"math/rand"
)

type TileType int

const (
	DeviceTypeTile TileType = iota
	OperatingSystemTile
	RenderVersionTile
	RenderingEngineTile
	AdditionalInfoTile
	BrowserNameTile
	BrowserVersionTile
	BrowserIdentifierTile
)

var (
	BrowserIdentifiers = []string{"Mozilla/5.0"}

	WindowSystems = []string{"X11"}

	DeviceTypes = []string{"Macintosh", "Windows NT", "iPhone", "iPad", "Linux", "Android"}

	OperatingSystems = map[string][]string{
		"Macintosh":  {"Intel Mac OS X 10_12_6", "Intel Mac OS X 11_0_1"},
		"Windows NT": {"6.1; WOW64", "10.0; Win64; x64"},
		"iPhone":     {"CPU iPhone OS 10_3_2 like Mac OS X"},
		"iPad":       {"CPU OS 10_3_2 like Mac OS X"},
		"Linux":      {"Android 4.3; GT-I9300 Build/JSS15J"},
		"Android":    {"4.3; Mobile"},
	}

	BrowserNames = []string{"Firefox", "Safari", "Brave"}

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

type Tile struct {
	Type    TileType
	Options []string
}

func (t *Tile) Observe() string {
	return t.Options[rand.Intn(len(t.Options))]
}

func NewTile(tileType TileType, observedValue string) *Tile {
	t := &Tile{Type: tileType}

	switch tileType {
	case BrowserIdentifierTile:
		t.Options = BrowserIdentifiers
	case DeviceTypeTile:
		t.Options = DeviceTypes
	case OperatingSystemTile:
		t.Options = OperatingSystems[observedValue]
	case RenderVersionTile:
		t.Options = RenderVersionMapping[observedValue]
	case RenderingEngineTile:
		t.Options = []string{RenderingEngineMapping[observedValue]}
	case AdditionalInfoTile:
		t.Options = AdditionalInfos[observedValue]
	case BrowserNameTile:
		t.Options = BrowserNames
	case BrowserVersionTile:
		t.Options = []string{"54.0", "99.0", "603.3.8", "537.36"}
	}

	return t
}

func GenerateUserAgent() string {
	browserIdentifierTile := NewTile(BrowserIdentifierTile, "")
	browserIdentifierValue := browserIdentifierTile.Observe()

	deviceTypeTile := NewTile(DeviceTypeTile, "")
	deviceTypeValue := deviceTypeTile.Observe()

	operatingSystemTile := NewTile(OperatingSystemTile, deviceTypeValue)
	operatingSystemValue := operatingSystemTile.Observe()

	browserNameTile := NewTile(BrowserNameTile, "")
	browserNameValue := browserNameTile.Observe()

	renderingEngineTile := NewTile(RenderingEngineTile, browserNameValue)
	renderingEngineValue := renderingEngineTile.Observe()

	additionalInfoTile := NewTile(AdditionalInfoTile, renderingEngineValue)
	additionalInfoValue := additionalInfoTile.Observe()

	renderVersionTile := NewTile(RenderVersionTile, browserNameValue)
	renderVersionValue := renderVersionTile.Observe()

	browserVersionTile := NewTile(BrowserVersionTile, browserNameValue)
	browserVersionValue := browserVersionTile.Observe()

	return fmt.Sprintf("%s (%s; %s; %s) %s %s %s/%s", browserIdentifierValue, deviceTypeValue, operatingSystemValue, renderVersionValue, renderingEngineValue, additionalInfoValue, browserNameValue, browserVersionValue)
}
