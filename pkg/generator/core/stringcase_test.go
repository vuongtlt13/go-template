package core

import "testing"

func TestSnakeCase(t *testing.T) {
	cases := map[string]string{
		"UserName":  "user_name",
		"userName":  "user_name",
		"user_name": "user_name",
		"user-name": "user_name",
		"User":      "user",
	}
	for input, want := range cases {
		if got := SnakeCase(input); got != want {
			t.Errorf("SnakeCase(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestCamelCase(t *testing.T) {
	cases := map[string]string{
		"user_name": "userName",
		"UserName":  "userName",
		"user-name": "userName",
		"user":      "user",
	}
	for input, want := range cases {
		if got := CamelCase(input); got != want {
			t.Errorf("CamelCase(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestPascalCase(t *testing.T) {
	cases := map[string]string{
		"user_name": "UserName",
		"userName":  "UserName",
		"user-name": "UserName",
		"user":      "User",
	}
	for input, want := range cases {
		if got := PascalCase(input); got != want {
			t.Errorf("PascalCase(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestKebabCase(t *testing.T) {
	cases := map[string]string{
		"user_name": "user-name",
		"UserName":  "user-name",
		"userName":  "user-name",
		"user":      "user",
	}
	for input, want := range cases {
		if got := KebabCase(input); got != want {
			t.Errorf("KebabCase(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestHumanCase(t *testing.T) {
	cases := map[string]string{
		"user_name": "User Name",
		"UserName":  "User Name",
		"userName":  "User Name",
		"user":      "User",
	}
	for input, want := range cases {
		if got := HumanCase(input); got != want {
			t.Errorf("HumanCase(%q) = %q, want %q", input, got, want)
		}
	}
}
