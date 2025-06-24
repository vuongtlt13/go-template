package core

import (
	"os"
	"strings"
	"testing"
)

const testModel = "package testmodel\n\n" +
	"type User struct {\n" +
	"\tID        uint   `gorm:\"primaryKey\" json:\"id\" crud:\"showable\" validate:\"required\"`\n" +
	"\tEmail     string `gorm:\"type:varchar(255);unique\" json:\"email\" crud:\"showable,creatable,editable\" validate:\"required,email\"`\n" +
	"\tPassword  string `gorm:\"type:varchar(255)\" json:\"password\" crud:\"creatable,editable\" validate:\"required,min=6\"`\n" +
	"\tFullName  string `gorm:\"type:varchar(255)\" json:\"full_name\" crud:\"showable,creatable,editable\"`\n" +
	"\tIsAdmin   bool   `gorm:\"default:false\" json:\"is_admin\" crud:\"showable,creatable,editable\"`\n" +
	"\tProfileID uint   `gorm:\"foreignKey:ProfileID\" json:\"profile_id\" crud:\"creatable,editable\"`\n" +
	"}"

func writeTestModelFile(filename string) error {
	return os.WriteFile(filename, []byte(testModel), 0644)
}

func TestParseGoStructFile(t *testing.T) {
	filename := "user_test_model.go"
	err := writeTestModelFile(filename)
	if err != nil {
		t.Fatalf("failed to write test model: %v", err)
	}
	defer os.Remove(filename)

	entity, err := ParseGoStructFile(filename)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	if entity == nil {
		t.Fatalf("entity is nil")
	}
	if entity.Name != "User" {
		t.Errorf("expected struct name User, got %s", entity.Name)
	}
	if len(entity.Columns) != 6 {
		t.Errorf("expected 6 columns, got %d", len(entity.Columns))
	}

	for _, col := range entity.Columns {
		switch col.Name {
		case "ID":
			if !col.PrimaryKey {
				t.Errorf("ID should be primary key")
			}
			if !col.Required {
				t.Errorf("ID should be required")
			}
			if !col.CrudConfig["showable"] {
				t.Errorf("ID should be showable")
			}
			if col.JsonName != "id" {
				t.Errorf("ID json name wrong: %s", col.JsonName)
			}
			if got := rulesToVee(col.Rules); got != "required" {
				t.Errorf("ID rules wrong: got %s", got)
			}
		case "Email":
			if !col.Unique {
				t.Errorf("Email should be unique")
			}
			if !col.CrudConfig["showable"] || !col.CrudConfig["creatable"] || !col.CrudConfig["editable"] {
				t.Errorf("Email crud config wrong: %+v", col.CrudConfig)
			}
			if col.Validate["email"] != "true" {
				t.Errorf("Email validate should have email")
			}
			want := "required|max:255|unique|email"
			if got := rulesToVee(col.Rules); got != want {
				t.Errorf("Email rules wrong: got %s, want %s", got, want)
			}
		case "Password":
			if col.CrudConfig["showable"] {
				t.Errorf("Password should not be showable")
			}
			if col.Validate["min"] != "6" {
				t.Errorf("Password min validate wrong: %+v", col.Validate)
			}
			want := "required|max:255|min:6"
			if got := rulesToVee(col.Rules); got != want {
				t.Errorf("Password rules wrong: got %s, want %s", got, want)
			}
		case "ProfileID":
			if col.ForeignKey != "ProfileID" {
				t.Errorf("ProfileID foreign key wrong: %s", col.ForeignKey)
			}
		}
	}
}

// rulesToVee chuyển []string rules sang style vee-validate (required|max:255|unique...)
func rulesToVee(rules []string) string {
	veeRules := make([]string, 0, len(rules))
	for _, rule := range rules {
		if strings.HasPrefix(rule, "max=") {
			veeRules = append(veeRules, "max:"+strings.TrimPrefix(rule, "max="))
		} else if strings.HasPrefix(rule, "min=") {
			veeRules = append(veeRules, "min:"+strings.TrimPrefix(rule, "min="))
		} else if strings.Contains(rule, "=") {
			parts := strings.SplitN(rule, "=", 2)
			veeRules = append(veeRules, parts[0]+":"+parts[1])
		} else {
			veeRules = append(veeRules, rule)
		}
	}
	return strings.Join(veeRules, "|")
}
