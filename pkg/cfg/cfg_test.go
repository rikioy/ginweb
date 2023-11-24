package cfg

import (
	"testing"
)

func TestGetString(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test_GetString",
			args: args{"test_key"},
			want: "test_value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddConfigPath("./test_data")
			ReadInConfig()
			if got := GetString(tt.args.key); got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetEnv(t *testing.T) {
	type args struct {
		key string
		env string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "read dev env",
			args: args{env: "dev", key: "test_key"},
			want: "dev_env_test_value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetEnv(tt.args.env)
			AddConfigPath("./test_data")
			ReadInConfig()
			if got := GetString(tt.args.key); got != tt.want {
				t.Errorf("SetEnv() have no effect, got %v, want %v", got, tt.want)
			}
		})
	}
}
