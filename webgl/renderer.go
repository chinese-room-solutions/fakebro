package webgl

import (
	"embed"
	"fmt"
	"math/rand"
	"strings"

	"github.com/hashicorp/go-version"
	"gopkg.in/yaml.v3"
)

var (
	ErrUnsupportedPlatform    = fmt.Errorf("unsupported platform")
	ErrInvalidPlatformVersion = fmt.Errorf("invalid platform version")
	ErrNoCompatibleRenderer   = fmt.Errorf("no compatible renderer found")
)

//go:embed data.yml
var dataFile embed.FS

type RendererData struct {
	MacOS   map[string][]string `yaml:"macOS"`
	Linux   map[string][]string `yaml:"Linux"`
	Windows map[string][]string `yaml:"Windows"`
}

var data RendererData

func init() {
	yamlData, err := dataFile.ReadFile("data.yml")
	if err != nil {
		panic(fmt.Sprintf("Failed to read data.yml: %v", err))
	}

	err = yaml.Unmarshal(yamlData, &data)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal YAML data: %v", err))
	}
}

func GenerateRenderer(seed int64, platform string, platformVersion string) (string, error) {
	r := rand.New(rand.NewSource(seed))

	switch strings.ToLower(platform) {
	case "macos":
		return generateVersionedRenderer(r, data.MacOS, platformVersion)
	case "linux":
		return generateVersionedRenderer(r, data.Linux, platformVersion)
	case "windows":
		return generateVersionedRenderer(r, data.Windows, platformVersion)
	default:
		return "", fmt.Errorf("%w: %s", ErrUnsupportedPlatform, platform)
	}
}

func generateVersionedRenderer(r *rand.Rand, versionedRenderers map[string][]string, platformVersion string) (string, error) {
	var compatibleRenderers []string

	for versionStr, renderers := range versionedRenderers {
		v1, err := version.NewVersion(versionStr)
		if err != nil {
			panic(fmt.Sprintf("invalid version in data.yml: %s", versionStr))
		}

		v2, err := version.NewVersion(platformVersion)
		if err != nil {
			return "", fmt.Errorf("%w: %s", ErrInvalidPlatformVersion, platformVersion)
		}

		if v2.GreaterThanOrEqual(v1) {
			compatibleRenderers = append(compatibleRenderers, renderers...)
		}
	}

	if len(compatibleRenderers) == 0 {
		return "", fmt.Errorf("%w: %s", ErrNoCompatibleRenderer, platformVersion)
	}

	return compatibleRenderers[r.Intn(len(compatibleRenderers))], nil
}
