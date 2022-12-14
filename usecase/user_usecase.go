package usecase

import (
	"errors"
	"final-project-backend/errorlist"
	"final-project-backend/utils"
	"time"

	"final-project-backend/entity"
	"final-project-backend/repository"

	"gorm.io/gorm"
)

type UserUsecase interface {
	Register(*entity.UserRegisterReqBody) (*entity.UserRegisterResBody, error)
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

func (u *userUsecaseImpl) Register(reqBody *entity.UserRegisterReqBody) (*entity.UserRegisterResBody, error) {
	var nRow int
	var err error
	nRow, err = u.userRepository.CheckEmailExistence(reqBody.Email)
	if nRow > 0 {
		return nil, errorlist.BadRequestError("email already registered", "EMAIL_EXISTED")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorlist.InternalServerError()
	}

	nRow, err = u.userRepository.CheckUsernameExistence(reqBody.Username)
	if nRow > 0 {
		return nil, errorlist.BadRequestError("username already registered", "USERNAME_EXISTED")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorlist.InternalServerError()
	}

	initialProfilePicture := ""
	initialPlayAttempt := 0
	defaultRoleId := 1
	hashedPass, _ := utils.HashAndSalt(reqBody.Email)
	validReqNewUser := entity.User{
		FullName:     reqBody.FullName,
		Phone: reqBody.Phone,
		Email:    reqBody.Email,
		Username: reqBody.Username,
		Password: hashedPass,
		RegisterDate: time.Now(),
		ProfilePicture: initialProfilePicture,
		PlayAttempt: initialPlayAttempt,
		RoleId: defaultRoleId,
		
	}

	newUser, err := u.userRepository.AddNewUser(&validReqNewUser)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}

	validResNewUser := entity.UserRegisterResBody{
		FullName:     newUser.FullName,
		Phone: newUser.Phone,
		Email:    newUser.Email,
		Username: newUser.Username,
		RegisterDate: newUser.RegisterDate,
	}

	return &validResNewUser, nil
}

func (u *userUsecaseImpl) Login(identifier, password string) (*entity.UserLoginResBody, error) {
	var user *entity.User
	var err error
	user, err = u.userRepository.GetUserByEmailOrUsername(identifier)
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
