package tutorialvalidation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

func TestAlias(t *testing.T) {
	validate := validator.New()
	validate.RegisterAlias("varchar", "required,max=255")

	type Seller struct {
		Id     string `validate:"varchar,min=5"`
		Name   string `validate:"varchar"`
		Owner  string `validate:"varchar"`
		Slogan string `validate:"varchar"`
	}

	seller := Seller{
		Id:     "123",
		Name:   "",
		Owner:  "",
		Slogan: "",
	}

	err := validate.Struct(seller)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func MustValidUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if value != strings.ToUpper(value) {
			return false
		}

		if len(value) < 5 {
			return false
		}

		return true
	} else {
		return false
	}
}

func TestCustomValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", MustValidUsername)

	type LoginRequest struct {
		Username string `validate:"required,username"`
		Password string `validate:"required"`
	}

	request := LoginRequest{
		Username: "ABCDE",
		Password: "12345",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

var regexNumber = regexp.MustCompile("^[0-9]+$")

func MustValidPin(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())
	if err != nil {
		panic(err)
	}

	value := field.Field().String()
	if !regexNumber.MatchString(value) {
		return false
	}

	return len(value) == length
}

func TestCustomValidationParameter(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("pin", MustValidPin)

	type Login struct {
		Phone string `validate:"required,number"`
		Pin   string `validate:"required,pin=6"`
	}

	request := Login{
		Phone: "1234",
		Pin:   "123456",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestOrRule(t *testing.T) {
	type Login struct {
		Username string `validate:"required,email|numeric"`
		Password string `validate:"required"`
	}

	validate := validator.New()

	request := Login{
		Username: "080989999",
		Password: "12345",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func MustEqualIgnoreCase(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()
	if !ok {
		panic("field not ok")
	}

	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())

	return firstValue == secondValue
}

func TestCustomValidationCrossField(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("equal_ignore_case", MustEqualIgnoreCase)

	type User struct {
		Username string `validate:"required,equal_ignore_case=Email|equal_ignore_case=Phone"`
		Email    string `validate:"required,email"`
		Phone    string `validate:"required,numeric"`
		Name     string `validate:"required"`
	}

	user := User{
		Username: "080989999",
		Email:    "abc@gmail.com",
		Phone:    "080989999",
		Name:     "Bung",
	}

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

type RegisterRequest struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Phone    string `validate:"required,numeric"`
	Password string `validate:"required"`
}

func MustValidRegisterSuccess(level validator.StructLevel) {
	registerReq := level.Current().Interface().(RegisterRequest)

	if registerReq.Username == registerReq.Email || registerReq.Username == registerReq.Phone {

	} else {
		level.ReportError(registerReq.Username, "Username", "Username", "username", "")
	}
}

func TestStructLevelValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterStructValidation(MustValidRegisterSuccess, RegisterRequest{})

	request := RegisterRequest{
		Username: "abcd@gmail.com",
		Email:    "abc@gmail.com",
		Phone:    "080989999",
		Password: "12345",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}
