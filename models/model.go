package models

type User struct {
	UserID      string `json:"userid,omitempty"  validate:"required"`
	Name        string `json:"name" validate:"required"`
	Password    string `json:"password,omitempty" validate:"required"`
	Description string `json:"description"`
}
