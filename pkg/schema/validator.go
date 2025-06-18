package schema

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// BaseModel is similar to Pydantic's BaseModel
type BaseModel struct {
	validator *validator.Validate
}

// NewBaseModel creates a new BaseModel
func NewBaseModel() *BaseModel {
	return &BaseModel{
		validator: validator.New(),
	}
}

// Validate validates a struct using validator
func Validate(s interface{}) error {
	if err := validate.Struct(s); err != nil {
		return formatError(err)
	}
	return nil
}

// formatError formats validation errors in a more readable way
func formatError(err error) error {
	if err == nil {
		return nil
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}

	var errMsgs []string
	for _, e := range validationErrors {
		field := e.Field()
		tag := e.Tag()
		param := e.Param()

		var msg string
		switch tag {
		case "required":
			msg = fmt.Sprintf("%s is required", field)
		case "email":
			msg = fmt.Sprintf("%s must be a valid email address", field)
		case "min":
			msg = fmt.Sprintf("%s must be at least %s characters long", field, param)
		case "max":
			msg = fmt.Sprintf("%s must not exceed %s characters", field, param)
		case "gt":
			msg = fmt.Sprintf("%s must be greater than %s", field, param)
		default:
			msg = fmt.Sprintf("%s failed validation: %s", field, tag)
		}
		errMsgs = append(errMsgs, msg)
	}

	return fmt.Errorf("validation failed: %s", strings.Join(errMsgs, "; "))
}

// MustValidate validates a struct and panics if validation fails
func MustValidate(s interface{}) {
	if err := Validate(s); err != nil {
		panic(err)
	}
}

// GetFieldValue gets the value of a field by name
func GetFieldValue(s interface{}, fieldName string) (interface{}, error) {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("not a struct")
	}
	return v.FieldByName(fieldName).Interface(), nil
}

// SetFieldValue sets the value of a field by name
func SetFieldValue(s interface{}, fieldName string, value interface{}) error {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("not a struct")
	}
	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("field %s not found", fieldName)
	}
	if !field.CanSet() {
		return fmt.Errorf("field %s cannot be set", fieldName)
	}
	field.Set(reflect.ValueOf(value))
	return nil
}
