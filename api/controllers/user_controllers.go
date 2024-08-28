package controllers

import (
	"net/http"

	"github.com/hunderaweke/spher/api/utils"
	"github.com/hunderaweke/spher/domain"
)

type UserControllers struct {
	usecase domain.UserUsecase
}

func NewUserControllers(usecase domain.UserUsecase) UserControllers {
	return UserControllers{usecase: usecase}
}

func (c *UserControllers) Register(w http.ResponseWriter, r *http.Request) {
	user, err := utils.Decode[domain.User](r.Body)
	if err != nil {
		utils.PostJSON(w, map[string]string{"error": "invalid data format"}, http.StatusNotAcceptable)
		return
	}
	user, err = c.usecase.Register(&user)
	if err != nil {
		utils.PostJSON(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}
	utils.PostJSON(w, user, http.StatusOK)
}
