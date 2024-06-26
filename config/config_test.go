package config

import (
	"runtime"
	"testing"
)

func TestSetupConfig(t *testing.T) {
	type args struct {
		env string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test SetupConfig Development",
			args: args{
				env: "development",
			},
			want: "debug",
		},
		{
			name: "Test SetupConfig Testing",
			args: args{
				env: "testing",
			},
			want: "test",
		},
		{
			name: "Test SetupConfig Production",
			args: args{
				env: "production",
			},
			want: "release",
		},
		{
			name: "Test SetupConfig Default",
			args: args{
				env: "",
			},
			want: "debug",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetupConfig(tt.args.env); got != tt.want {
				t.Errorf("SetupConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuntimeSetsMaxProcsToNumCPU(t *testing.T) {
	oldMaxProcs := runtime.GOMAXPROCS(0)
	defer runtime.GOMAXPROCS(oldMaxProcs)

	Runtime()

	numCPU := runtime.NumCPU()
	maxProcs := runtime.GOMAXPROCS(0)

	if maxProcs != numCPU {
		t.Errorf("Expected GOMAXPROCS to be set to %d, but got %d", numCPU, maxProcs)
	}
}
