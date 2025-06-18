package models

type CreateUserRequestModel struct {
	Name         string `json:"name" validate:"required" example:"Harsh"`
	Email        string `json:"email" validate:"required,email" example:"abd@gmail.com"`
	Age          uint   `json:"age" validate:"gte=0,lte=120" example:"18"` // optional
	Pancard      string `json:"pancard" validate:"required,len=10" example:"ABCDE12345"`
	MobileNumber string `json:"mobile_number" validate:"required,len=10,numeric" example:"9898009898"`
	Bio          string `json:"bio" example:"I am Harsh"`                                               // optional
	Birthday     string `json:"birthday" validate:"omitempty,datetime=2006-01-02" example:"2001-04-30"` // optional, omitempty omits validations if it has no value
}

type CreateUserResponseModel struct {
	ID      uint   `json:"id" example:"1"`
	Message string `json:"message" example:"User created successfully"`
}
