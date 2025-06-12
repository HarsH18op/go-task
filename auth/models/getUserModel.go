package models

type GetUserResponseModel struct {
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
