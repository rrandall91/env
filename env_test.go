package env

import (
	"os"
	"testing"
)

func TestSet(t *testing.T) {
	type args struct {
		key   string
		value string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Sets the value of the environment variable named by the key",
			args: args{
				key:   "TEST_KEY",
				value: "TEST_VALUE",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Set(tt.args.key, tt.args.value)
		})

		if got := os.Getenv(tt.args.key); got != tt.args.value {
			t.Errorf("Set() = %v, want %v", got, tt.args.value)
		}
	}
}

func TestGet(t *testing.T) {
	type args struct {
		key   string
		value string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Returns the value of the environment variable named by the key",
			args: args{
				key:   "TEST_KEY",
				value: "TEST_VALUE",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.args.key); got != tt.args.value {
				t.Errorf("Get() = %v, want %v", got, tt.args.value)
			}
		})
	}
}

func TestGetWithDefault(t *testing.T) {
	type args struct {
		key      string
		value    string
		fallback string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Returns the value of the environment variable named by the key",
			args: args{
				key:      "TEST_KEY",
				value:    "TEST_VALUE",
				fallback: "TEST_FALLBACK",
			},
			want: "TEST_VALUE",
		},
		{
			name: "Returns the fallback value if the variable is not set",
			args: args{
				key:      "TEST_KEY",
				fallback: "TEST_FALLBACK",
			},
			want: "TEST_FALLBACK",
		},
	}

	for _, tt := range tests {
		os.Setenv(tt.args.key, "")

		if tt.args.value != "" {
			os.Setenv(tt.args.key, tt.args.value)
		}

		t.Run(tt.name, func(t *testing.T) {
			if got := GetWithDefault(tt.args.key, tt.args.fallback); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.args.fallback)
			}
		})
	}
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get("TEST_KEY")
	}
}

func BenchmarkGetWithDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetWithDefault("TEST_KEY", "TEST_FALLBACK")
	}
}

func BenchmarkSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Set("TEST_KEY", "TEST_VALUE")
	}
}
