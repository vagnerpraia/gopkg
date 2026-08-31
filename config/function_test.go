package gpconfig

import (
	"os"
	"path/filepath"
	"testing"
)

type testConfig struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
	} `yaml:"database"`

	Simulator struct {
		Latency  string `yaml:"latency"`
		Slippage int    `yaml:"slippage"`
	} `yaml:"simulator"`

	Asset string `yaml:"asset"`
}

func TestUnmarshal(t *testing.T) {

	tests := []struct {
		name    string
		files   map[string]string
		path    string
		want    testConfig
		wantErr bool
	}{
		{
			name: "single file",
			files: map[string]string{
				"config.yaml": `
asset: WDO

database:
  host: localhost
  port: 27017
  username: quantos

simulator:
  latency: 10ms
  slippage: 1
`,
			},
			path: "config.yaml",
			want: func() testConfig {
				var config testConfig

				config.Asset = "WDO"
				config.Database.Host = "localhost"
				config.Database.Port = 27017
				config.Database.Username = "quantos"
				config.Simulator.Latency = "10ms"
				config.Simulator.Slippage = 1

				return config
			}(),
		},
		{
			name: "merge multiple files",
			files: map[string]string{
				"default.yaml": `
asset: WDO

database:
  host: localhost
  port: 27017
  username: quantos

simulator:
  latency: 10ms
  slippage: 1
`,
				"production.yaml": `
database:
  host: mongo.production

simulator:
  latency: 5ms
`,
			},
			path: "default.yaml,production.yaml",
			want: func() testConfig {
				var config testConfig

				config.Asset = "WDO"
				config.Database.Host = "mongo.production"
				config.Database.Port = 27017
				config.Database.Username = "quantos"
				config.Simulator.Latency = "5ms"
				config.Simulator.Slippage = 1

				return config
			}(),
		},
		{
			name: "semicolon separator",
			files: map[string]string{
				"default.yaml": `
asset: WDO

database:
  host: localhost
  port: 27017
`,
				"local.yaml": `
database:
  host: 127.0.0.1
`,
			},
			path: "default.yaml;local.yaml",
			want: func() testConfig {
				var config testConfig

				config.Asset = "WDO"
				config.Database.Host = "127.0.0.1"
				config.Database.Port = 27017

				return config
			}(),
		},
		{
			name: "later file overrides scalar value",
			files: map[string]string{
				"first.yaml": `
asset: WDO
`,
				"second.yaml": `
asset: WIN
`,
			},
			path: "first.yaml,second.yaml",
			want: func() testConfig {
				var config testConfig
				config.Asset = "WIN"
				return config
			}(),
		},
		{
			name: "nested maps are merged recursively",
			files: map[string]string{
				"first.yaml": `
database:
  host: localhost
  port: 27017
  username: quantos
`,
				"second.yaml": `
database:
  host: mongo.production
`,
			},
			path: "first.yaml,second.yaml",
			want: func() testConfig {
				var config testConfig

				config.Database.Host = "mongo.production"
				config.Database.Port = 27017
				config.Database.Username = "quantos"

				return config
			}(),
		},
		{
			name:    "missing file",
			files:   map[string]string{},
			path:    "does-not-exist.yaml",
			wantErr: true,
		},
		{
			name: "invalid yaml",
			files: map[string]string{
				"invalid.yaml": `
database:
  host: [invalid
`,
			},
			path:    "invalid.yaml",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()

			for name, content := range tt.files {
				path := filepath.Join(dir, name)

				if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
					t.Fatalf("failed to create test file %q: %v", name, err)
				}
			}

			paths := tt.path

			for name := range tt.files {
				paths = replacePath(paths, name, filepath.Join(dir, name))
			}

			var got testConfig

			err := Unmarshal(paths, &got)

			if (err != nil) != tt.wantErr {
				t.Fatalf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				return
			}

			if got != tt.want {
				t.Errorf("Unmarshal() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func replacePath(path, name, replacement string) string {

	result := path

	for _, separator := range []string{",", ";"} {
		result = replacePathPart(result, separator, name, replacement)
	}

	return result
}

func replacePathPart(path, separator, name, replacement string) string {

	parts := splitPreservingSeparator(path, separator)

	for i, part := range parts {
		if part == name {
			parts[i] = replacement
		}
	}

	return joinPreservingSeparator(parts, separator)
}

func splitPreservingSeparator(path, separator string) []string {

	var result []string

	start := 0

	for i := 0; i <= len(path)-len(separator); i++ {
		if path[i:i+len(separator)] == separator {
			result = append(result, path[start:i])
			start = i + len(separator)
			i += len(separator) - 1
		}
	}

	result = append(result, path[start:])

	return result
}

func joinPreservingSeparator(parts []string, separator string) string {

	result := ""

	for i, part := range parts {
		if i > 0 {
			result += separator
		}

		result += part
	}

	return result
}
