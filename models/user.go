package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
	UserType  int    `json:"userType"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

func (user *User) CreateAdmin(userInput UserInput) {
	user.Name = userInput.Name
	user.Email = userInput.Email
	user.Password = []byte(userInput.Password)
	user.UserType = 1
}

func (user *User) CreateConsumer(userInput UserInput) {
	user.Name = userInput.Name
	user.Email = userInput.Email
	user.Password = []byte(userInput.Password)
	user.UserType = 2
}

func (userinput *UserInput) IsValidUser() (bool, string) {
	if userinput.Name == "" {
		return false, "Name is required"
	}
	if userinput.Email == "" {
		return false, "Email is required"
	}
	if userinput.Password == "" {
		return false, "Password is required"
	}
	if len(userinput.Password) < 6 {
		return false, "Password must be atleast 6 characters"
	}
	return true, ""
}
