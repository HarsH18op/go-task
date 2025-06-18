package models

type GetUserResponseModel struct {
	ID        uint    `json:"id" example:"1"`
	Age       uint    `json:"age" example:"25"`
	Name      string  `json:"name" example:"Harsh"`
	Email     string  `json:"email" example:"abd@gmail.com"`
	Pancard   string  `json:"pancard" example:"ABCDE12345"`
	Mobile    string  `json:"mobileNumber" example:"9898009898"`
	Bio       string  `json:"bio" example:"I'm Harsh"`
	Birthday  *string `json:"birthday" example:"2001-04-30"`
	CreatedAt string  `json:"createdAt" example:"2025-06-15 11:00:00"`
	UpdatedAt string  `json:"updatedAt" example:"2025-06-15 11:00:00"`
}
