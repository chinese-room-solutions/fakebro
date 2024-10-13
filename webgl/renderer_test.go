package webgl

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateRenderer(t *testing.T) {
	tests := []struct {
		name            string
		seed            int64
		platform        string
		platformVersion string
		expectedError   error
		expectedPrefix  string
	}{
		{
			name:            "valid macOS",
			seed:            12345,
			platform:        "macOS",
			platformVersion: "11.1",
			expectedPrefix:  "ANGLE (Apple, Apple M1",
		},
		{
			name:            "valid Linux",
			seed:            67890,
			platform:        "Linux",
			platformVersion: "5.10.0",
			expectedPrefix:  "ANGLE (",
		},
		{
			name:            "valid Windows 11",
			seed:            11111,
			platform:        "Windows",
			platformVersion: "14.0.0.0",
			expectedPrefix:  "ANGLE (",
		},
		{
			name:            "unsupported platform",
			seed:            22222,
			platform:        "Android",
			platformVersion: "11",
			expectedError:   ErrUnsupportedPlatform,
		},
		{
			name:            "invalid platform version",
			seed:            33333,
			platform:        "macOS",
			platformVersion: "invalid",
			expectedError:   ErrInvalidPlatformVersion,
		},
		{
			name:            "no compatible renderer",
			seed:            44444,
			platform:        "macOS",
			platformVersion: "1.0",
			expectedError:   ErrNoCompatibleRenderer,
		},
		{
			name:            "case insensitive platform",
			seed:            55555,
			platform:        "MACOS",
			platformVersion: "11.1",
			expectedPrefix:  "ANGLE (Apple, Apple M1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			renderer, err := GenerateRenderer(tt.seed, tt.platform, tt.platformVersion)

			if tt.expectedError != nil {
				require.ErrorIs(t, err, tt.expectedError)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, renderer)
				require.True(t, strings.HasPrefix(renderer, tt.expectedPrefix), "Renderer should start with expected prefix")
			}
		})
	}
}

func TestGenerateRendererDeterministic(t *testing.T) {
	seed := int64(99999)
	platform := "macOS"
	platformVersion := "12.5"

	renderer1, err := GenerateRenderer(seed, platform, platformVersion)
	require.NoError(t, err)

	renderer2, err := GenerateRenderer(seed, platform, platformVersion)
	require.NoError(t, err)

	require.Equal(t, renderer1, renderer2, "Renderers should be the same for the same seed")
}

func TestGenerateRendererDifferentSeeds(t *testing.T) {
	platform := "Linux"
	platformVersion := "5.10.0"

	renderer1, err := GenerateRenderer(11111, platform, platformVersion)
	require.NoError(t, err)

	renderer2, err := GenerateRenderer(22222, platform, platformVersion)
	require.NoError(t, err)

	require.NotEqual(t, renderer1, renderer2, "Renderers should be different for different seeds")
}
