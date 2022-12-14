package usecase

import (
	"errors"
	"final-project-backend/errorlist"
	"final-project-backend/utils"

	"final-project-backend/entity"
	"final-project-backend/repository"

	"gorm.io/gorm"
)

type UserUsecase interface {
	// Register(string, string, string) (*entity.UserRegisterResBody, error)
	Login(string, string) (*entity.UserLoginResBody, error)
	// GetDetailUser(int) (*entity.User, error)
}

type userUsecaseImpl struct {
	userRepository   repository.UserRepository
}

type UserUsecaseConfig struct {
	UserRepository   repository.UserRepository
}

func NewUserUsecase(c UserUsecaseConfig) UserUsecase {
	return &userUsecaseImpl{
		userRepository:   c.UserRepository,
	}
}

// func (u *userUsecaseImpl) Register(name, email, pass string) (*entity.UserRegisterResBody, error) {
// 	nRow, err := u.userRepository.CheckEmailExistence(email)
// 	if nRow > 0 {
// 		return nil, errorlist.BadRequestError("email already registered", "EMAIL_EXISTED")
// 	}
// 	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
// 		return nil, errorlist.InternalServerError()
// 	}

// 	hashedPass, _ := utils.HashAndSalt(pass)
// 	validReqNewUser := entity.User{
// 		Name:     name,
// 		Email:    email,
// 		Password: hashedPass,
// 	}

// 	newUser, err := u.userRepository.AddNewUser(&validReqNewUser)
// 	if err != nil {
// 		return nil, errorlist.InternalServerError()
// 	}

// 	validResNewUser := entity.UserRegisterResBody{
// 		Id:    newUser.Id,
// 		Name:  newUser.Name,
// 		Email: newUser.Email,
// 	}

// 	var default_balance = float64(0)
// 	newWallet := entity.Wallet{
// 		UserId:  validResNewUser.Id,
// 		Balance: default_balance,
// 	}

// 	_, err = u.walletRepository.AddNewWallet(&newWallet)
// 	if err != nil {
// 		return nil, errorlist.InternalServerError()
// 	}

// 	return &validResNewUser, nil
// }

func (u *userUsecaseImpl) Login(identifier, password string) (*entity.UserLoginResBody, error) {
	user, err := u.userRepository.GetUserByEmailOrUsername(identifier)
	if errors.Is(err, gorm.ErrRecordNotFound){
		return nil, errorlist.UnauthorizedError()
	}

	if err != nil {
		return nil, err
	}

	a := utils.NewAuthUtil()
	if !a.ComparePassword(user.Password, password) {
		return nil, errorlist.UnauthorizedError()
	}
	tokenStr, err := a.GenerateAccessToken(user)
	token := entity.UserLoginResBody{
		AccessToken: tokenStr,
	}
	return &token, err
}

// func (u *userUsecaseImpl) GetDetailUser(id int) (*entity.User, error) {
// 	user, err := u.userRepository.GetUserDetailById(id)
// 	if err != nil {
// 		return nil, errorlist.InternalServerError()
// 	}

// 	return user, nil
// }
