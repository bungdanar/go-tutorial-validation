package tutorialvalidation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	validate := validator.New()
	if validate == nil {
		t.Error("validator is nil")
	}
}

func TestValidationVariable(t *testing.T) {
	validate := validator.New()
	user := ""

	err := validate.Var(user, "required")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationTwoVariables(t *testing.T) {
	validate := validator.New()

	password := "rahasia"
	confirmPassword := "salah"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMultipleTag(t *testing.T) {
	validate := validator.New()
	user := "bung123"

	err := validate.Var(user, "required,numeric")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestTagParameter(t *testing.T) {
	validate := validator.New()
	user := "1234567890"

	err := validate.Var(user, "required,numeric,min=5,max=10")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestStruct(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()
	loginReq := LoginRequest{
		Username: "bung@gmail.com",
		Password: "12345",
	}

	err := validate.Struct(loginReq)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationErrors(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()
	loginReq := LoginRequest{
		Username: "bung",
		Password: "1234",
	}

	err := validate.Struct(loginReq)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

func TestStructCrossField(t *testing.T) {
	type RegisterUser struct {
		Username        string `validate:"required,email"`
		Password        string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
	}

	validate := validator.New()
	registerReq := RegisterUser{
		Username:        "bung@gmail.com",
		Password:        "12345",
		ConfirmPassword: "12345",
	}

	err := validate.Struct(registerReq)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestNestedStruct(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id      string  `validate:"required"`
		Name    string  `valdiate:"required"`
		Address Address `validate:"required"`
	}

	validate := validator.New()
	request := User{
		Id:      "1",
		Name:    "Bung",
		Address: Address{},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string    `validate:"required"`
		Name      string    `valdiate:"required"`
		Addresses []Address `validate:"required,min=1,dive"`
	}

	validate := validator.New()
	request := User{
		Id:   "1",
		Name: "Bung",
		Addresses: []Address{
			{
				City: "Jaktim",
			},
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestBasicCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string    `validate:"required"`
		Name      string    `valdiate:"required"`
		Addresses []Address `validate:"required,min=1,dive"`
		Hobbies   []string  `validate:"required,min=1,dive,required,min=3"`
	}

	validate := validator.New()
	request := User{
		Id:   "1",
		Name: "Bung",
		Addresses: []Address{
			{
				City:    "Jaktim",
				Country: "Indo",
			},
		},
		Hobbies: []string{"Coding"},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMap(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id        string            `validate:"required"`
		Name      string            `valdiate:"required"`
		Addresses []Address         `validate:"required,min=1,dive"`
		Hobbies   []string          `validate:"required,min=1,dive,required,min=3"`
		Schools   map[string]School `validate:"required,min=1,dive,keys,required,min=2,endkeys,dive"`
	}

	validate := validator.New()
	request := User{
		Id:   "1",
		Name: "Bung",
		Addresses: []Address{
			{
				City:    "Jaktim",
				Country: "Indo",
			},
		},
		Hobbies: []string{"Coding"},
		Schools: map[string]School{
			"sd": {Name: "Bandar"},
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestBasicMap(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id        string            `validate:"required"`
		Name      string            `valdiate:"required"`
		Addresses []Address         `validate:"required,min=1,dive"`
		Hobbies   []string          `validate:"required,min=1,dive,required,min=3"`
		Schools   map[string]School `validate:"required,min=1,dive,keys,required,min=2,endkeys,dive"`
		Wallets   map[string]int    `validate:"required,min=1,dive,keys,required,min=2,endkeys,required,gt=1"`
	}

	validate := validator.New()
	request := User{
		Id:   "1",
		Name: "Bung",
		Addresses: []Address{
			{
				City:    "Jaktim",
				Country: "Indo",
			},
		},
		Hobbies: []string{"Coding"},
		Schools: map[string]School{
			"sd": {Name: "Bandar"},
		},
		Wallets: map[string]int{
			"bc": 2,
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}
