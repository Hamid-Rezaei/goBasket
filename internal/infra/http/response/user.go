package response

import (
	"github.com/Hamid-Rezaei/goBasket/internal/infra/repository"
	"github.com/Hamid-Rezaei/goBasket/internal/utils"
)

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func NewUserResponse(u *repository.UserDTO) *UserResponse {
	r := new(UserResponse)
	r.Username = u.Username
	r.Email = u.Email
	r.Token = utils.GenerateJWT(u.ID)
	return r
}
