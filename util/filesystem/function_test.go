package gpfilesystem

import (
	"fmt"
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
			want:         []string{"/dev/shm/insaio/quantos/data"},
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

func TestParentDirs(t *testing.T) {

	tests := []struct {
		name string
		path string
		want []string
	}{
		{
			name: "absolute path",
			path: "/dev/shm/insaio/quantos/config/general.yaml",
			want: []string{
				NormalizePathOS("/dev", os),
				NormalizePathOS("/dev/shm", os),
				NormalizePathOS("/dev/shm/insaio", os),
				NormalizePathOS("/dev/shm/insaio/quantos", os),
				NormalizePathOS("/dev/shm/insaio/quantos/config", os),
			},
		},
		// {
		// 	name: "single directory",
		// 	path: "/tmp/file.txt",
		// 	want: []string{
		// 		NormalizePathOS("/tmp", os),
		// 	},
		// },
		// {
		// 	name: "relative path",
		// 	path: "config/general.yaml",
		// 	want: []string{
		// 		NormalizePathOS("config", os),
		// 	},
		// },
		// {
		// 	name: "file in current directory",
		// 	path: "general.yaml",
		// 	want: []string{},
		// },
		// {
		// 	name: "clean path",
		// 	path: "/dev/shm/insaio/quantos/./config/../config/general.yaml",
		// 	want: []string{
		// 		NormalizePathOS("/dev", os),
		// 		NormalizePathOS("/dev/shm", os),
		// 		NormalizePathOS("/dev/shm/insaio", os),
		// 		NormalizePathOS("/dev/shm/insaio/quantos", os),
		// 		NormalizePathOS("/dev/shm/insaio/quantos/config", os),
		// 	},
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParentDirs(tt.path, os)

			fmt.Println("==================================================")

			for _, p := range tt.want {
				fmt.Println(p)
			}

			fmt.Println("-------------------------")

			for _, p := range got {
				fmt.Println(p)
			}

			fmt.Println("==================================================")

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParentDirs(%q) = %#v, want %#v", tt.path, got, tt.want)
			}
		})
	}
}
