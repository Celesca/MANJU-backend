package controllers

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) CreateUser() {}

func (uc *UserController) ListUsers() {}

func (uc *UserController) GetUser() {}

func (uc *UserController) UpdateUser() {}

func (uc *UserController) DeleteUser() {}
