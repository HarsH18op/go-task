package models

type CreateUserRequestModel struct {
	Name         string `json:"name" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Age          uint   `json:"age" validate:"gte=0,lte=120"` // optional
	Pancard      string `json:"pancard" validate:"required,len=10"`
	MobileNumber string `json:"mobile_number" validate:"required,len=10,numeric"`
	Bio          string `json:"bio"`                                               // optional
	Birthday     string `json:"birthday" validate:"omitempty,datetime=2006-01-02"` // optional, omitempty omits validations if it has no value
}

type CreateUserResponseModel struct {
	ID        uint    `json:"id"`
	Age       uint    `json:"age"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Pancard   string  `json:"pancard"`
	Mobile    string  `json:"mobile_number"`
	Bio       string  `json:"bio"`
	Birthday  *string `json:"birthday"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}