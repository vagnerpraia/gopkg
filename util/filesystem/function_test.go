package gpfilesystem

import (
	"reflect"
	"testing"

	gpenum "github.com/vagnerpraia/gopkg/enum"
)

const os gpenum.OS = gpenum.OSWindows

func TestSubdirectories(t *testing.T) {

	tests := []struct {
		name         string
		finalPath    string
		relativePath string
		os           gpenum.OS
		want         []string
		wantErr      bool
	}{
		{
			name:         "multiple subdirectories",
			finalPath:    "/dev/shm/insaio/quantos/data/tick/wdo/2026/06",
			relativePath: "/dev/shm/insaio/quantos/data",
			os:           os,
			want: []string{
				NormalizePathOS("/dev/shm/insaio/quantos/data/tick", os),
				NormalizePathOS("/dev/shm/insaio/quantos/data/tick/wdo", os),
				NormalizePathOS("/dev/shm/insaio/quantos/data/tick/wdo/2026", os),
				NormalizePathOS("/dev/shm/insaio/quantos/data/tick/wdo/2026/06", os),
			},
		},
		{
			name:         "single subdirectory",
			finalPath:    "/dev/shm/insaio/quantos/data/tick",
			relativePath: "/dev/shm/insaio/quantos/data",
			want: []string{
				NormalizePathOS("/dev/shm/insaio/quantos/data/tick", os),
			},
		},
		{
			name:         "same path",
			finalPath:    "/dev/shm/insaio/quantos/data",
			relativePath: "/dev/shm/insaio/quantos/data",
			want:         []string{},
		},
		{
			name:         "paths with dot",
			finalPath:    "/dev/shm/insaio/quantos/data/tick/../tick/wdo",
			relativePath: "/dev/shm/insaio/quantos/data/.",
			want: []string{
				NormalizePathOS("/dev/shm/insaio/quantos/data/tick", os),
				NormalizePathOS("/dev/shm/insaio/quantos/data/tick/wdo", os),
			},
		},
		{
			name:         "final path outside relative path",
			finalPath:    "/dev/shm/insaio/quantos/other",
			relativePath: "/dev/shm/insaio/quantos/data",
			wantErr:      true,
		},
		{
			name:         "final path is parent of relative path",
			finalPath:    "/dev/shm/insaio/quantos",
			relativePath: "/dev/shm/insaio/quantos/data",
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Subdirectories(tt.finalPath, tt.relativePath, tt.os)

			if (err != nil) != tt.wantErr {
				t.Fatalf("Subdirectories() error = %v, wantErr = %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subdirectories() = %v, want %v", got, tt.want)
			}
		})
	}
}
