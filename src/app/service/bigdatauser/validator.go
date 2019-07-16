package user

type ValidateUser struct {
	ID int `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	AvatarUrl string `json:"avatar_url" binding:"required"`
}