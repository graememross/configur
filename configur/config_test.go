package configur

import (
	"reflect"
	"testing"

	"github.com/knadh/koanf"
	"github.com/spf13/pflag"
)

func Test_scanFiles(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
	}{
		{"test-1", args{"test"}}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanFiles(tt.args.name)
		})
	}
	GetConfigSet().Print()
}

func Test_readFile(t *testing.T) {
	type args struct {
		name ParsePair
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			readFile(tt.args.name)
		})
	}
}

func Test_scanEnvironment(t *testing.T) {
	type args struct {
		prefix string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanEnvironment(tt.args.prefix)
		})
	}
}

func TestConfigSet_BindValues(t *testing.T) {
	type fields struct {
		Koanf *koanf.Koanf
	}
	type args struct {
		aFlagset *pflag.FlagSet
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ConfigSet{
				Koanf: tt.fields.Koanf,
			}
			c.BindValues(tt.args.aFlagset)
		})
	}
}

func TestGetConfigSet(t *testing.T) {
	tests := []struct {
		name string
		want *ConfigSet
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetConfigSet(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConfigSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
