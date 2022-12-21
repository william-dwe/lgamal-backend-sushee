package usecase

import (
	"errors"
	"final-project-backend/errorlist"
	"final-project-backend/utils"
	"strconv"
	"time"

	"final-project-backend/config"
	"final-project-backend/entity"
	"final-project-backend/repository"

	"gorm.io/gorm"
)

type UserUsecase interface {
	Register(*entity.UserRegisterReqBody) (*entity.UserRegisterResBody, error)
	Login(string, string) (*entity.UserLoginResBody, string, error)
	Refresh(string) (*entity.UserLoginResBody, error)
	GetDetailUserByUsername(accessToken string) (*entity.UserContext, error)
	UpdateUserDetailsByUsername(username string, updatePremises entity.UserEditDetailsReqBody) (*entity.UserContext, error)
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

	nRow, err = u.userRepository.CheckPhoneExistence(reqBody.Phone)
	if nRow > 0 {
		return nil, errorlist.BadRequestError("phone already registered", "PHONE_EXISTED")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorlist.InternalServerError()
	}

	initialProfilePicture := ""
	initialPlayAttempt := 0
	defaultRoleId := 1
	hashedPass, _ := utils.HashAndSalt(reqBody.Password)
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

func (u *userUsecaseImpl) Login(identifier, password string) (*entity.UserLoginResBody, string, error) {
	var user *entity.User
	var err error
	user, err = u.userRepository.GetUserByEmailOrUsername(identifier)
	if errors.Is(err, gorm.ErrRecordNotFound){
		return nil, "", errorlist.UnauthorizedError()
	}
	if err != nil {
		return nil, "", err
	}

	a := utils.NewAuthUtil()
	if !a.ComparePassword(user.Password, password) {
		return nil, "", errorlist.UnauthorizedError()
	}

	accessTokenStr, err := a.GenerateAccessToken(user.Username)
	if err != nil {
		return nil, "", err
	}
	refreshTokenStr, err := a.GenerateRefreshToken()
	if err != nil {
		return nil, "", err
	}

	expirationLimit, _ := strconv.ParseInt(config.Config.AuthConfig.TimeLimitRefreshToken, 10, 64)
	session := entity.Session{
		RefreshToken: refreshTokenStr,
		UserId: int(user.ID),
		ExpiredAt: time.Now().Add(time.Second * time.Duration(expirationLimit)),
	}
	u.userRepository.AddNewUserSession(&session)

	token := entity.UserLoginResBody{
		AccessToken: accessTokenStr,
	}
	return &token, refreshTokenStr, err
}


func (u *userUsecaseImpl) Refresh(refreshToken string) (*entity.UserLoginResBody, error) {
	var err error
	a := utils.NewAuthUtil()
	_, err = a.ValidateToken(refreshToken, config.Config.AuthConfig.HmacSecretRefreshToken)
	if err != nil {
		return nil, err
	}

	session, err := u.userRepository.GetUserSessionByRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}
	if time.Now().After(session.ExpiredAt) {
		return nil, errorlist.UnauthorizedError()
	}

	user, err := u.userRepository.GetUserById(session.UserId)
	if err != nil {
		return nil, err
	}

	accessTokenStr, err := a.GenerateAccessToken(user.Username)
	if err != nil {
		return nil, err
	}
	accessToken := entity.UserLoginResBody{
		AccessToken: accessTokenStr,
	}
	return &accessToken, err
}

func (u *userUsecaseImpl) GetDetailUserByUsername(username string) (*entity.UserContext, error) {
	user, err := u.userRepository.GetUserByEmailOrUsername(username)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}

	userContext := entity.UserContext{
		Username: user.Username,
		FullName: user.FullName,
		Email: user.Email,
		Phone: user.Phone,
		ProfilePicture: user.ProfilePicture,
		PlayAttempt: user.PlayAttempt,
		RoleId: user.RoleId,
	}

	return &userContext, nil
}


func (u *userUsecaseImpl) UpdateUserDetailsByUsername(username string, reqBody entity.UserEditDetailsReqBody) (*entity.UserContext, error) {
	var nRow int
	var err error
	nRow, err = u.userRepository.CheckEmailExistence(reqBody.Email)
	if nRow > 0 {
		return nil, errorlist.BadRequestError("email already registered", "EMAIL_EXISTED")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorlist.InternalServerError()
	}

	nRow, err = u.userRepository.CheckPhoneExistence(reqBody.Phone)
	if nRow > 0 {
		return nil, errorlist.BadRequestError("phone already registered", "PHONE_EXISTED")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorlist.InternalServerError()
	}
	
	hashedPass, _ := utils.HashAndSalt(reqBody.Password)

	newUser := entity.User{
		FullName: reqBody.FullName,
		Phone: reqBody.Phone,
		Email: reqBody.Email,
		Password: hashedPass,
		ProfilePicture: reqBody.ProfilePicture,
	}

	err = u.userRepository.UpdateUserDetailsByUsername(username, &newUser)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}

	user, err := u.userRepository.GetUserByEmailOrUsername(username)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}
	userContext := entity.UserContext{
		Username: user.Username,
		FullName: user.FullName,
		Email: user.Email,
		Phone: user.Phone,
		ProfilePicture: user.ProfilePicture,
		PlayAttempt: user.PlayAttempt,
		RoleId: user.RoleId,
	}
	return &userContext, err
}
