package core

import (
	"testing"
)

func TestSingularizePluralize(t *testing.T) {
	cases := []struct {
		Singular string
		Plural   string
	}{
		{"user", "users"},
		{"person", "people"},
		{"company", "companies"},
		{"bus", "buses"},
		{"child", "children"},
		{"sheep", "sheep"},
	}
	for _, c := range cases {
		if got := Pluralize(c.Singular); got != c.Plural {
			t.Errorf("Pluralize(%q) = %q, want %q", c.Singular, got, c.Plural)
		}
		if got := Singularize(c.Plural); got != c.Singular {
			t.Errorf("Singularize(%q) = %q, want %q", c.Plural, got, c.Singular)
		}
	}
}

func TestTextForm(t *testing.T) {
	tf := NewTextForm("user")
	if tf.SnakeCase != "user" {
		t.Errorf("SnakeCase = %q, want 'user'", tf.SnakeCase)
	}
	if tf.SnakeCasePlural != "users" {
		t.Errorf("SnakeCasePlural = %q, want 'users'", tf.SnakeCasePlural)
	}
	if tf.KebabCase != "user" {
		t.Errorf("KebabCase = %q, want 'user'", tf.KebabCase)
	}
	if tf.KebabCasePlural != "users" {
		t.Errorf("KebabCasePlural = %q, want 'users'", tf.KebabCasePlural)
	}
	if tf.CamelCase != "user" {
		t.Errorf("CamelCase = %q, want 'user'", tf.CamelCase)
	}
	if tf.CamelCasePlural != "users" {
		t.Errorf("CamelCasePlural = %q, want 'users'", tf.CamelCasePlural)
	}
}

func TestModelTextForm(t *testing.T) {
	mtf := NewModelTextForm("User")
	if mtf.TableName != "users" {
		t.Errorf("TableName = %q, want 'users'", mtf.TableName)
	}
	if mtf.ModelName != "User" {
		t.Errorf("ModelName = %q, want 'User'", mtf.ModelName)
	}
	if mtf.ModelNamePlural != "Users" {
		t.Errorf("ModelNamePlural = %q, want 'Users'", mtf.ModelNamePlural)
	}
	if mtf.ModelNameSnakeCase != "user" {
		t.Errorf("ModelNameSnakeCase = %q, want 'user'", mtf.ModelNameSnakeCase)
	}
	if mtf.ModelNameSnakeCasePlural != "users" {
		t.Errorf("ModelNameSnakeCasePlural = %q, want 'users'", mtf.ModelNameSnakeCasePlural)
	}
}
