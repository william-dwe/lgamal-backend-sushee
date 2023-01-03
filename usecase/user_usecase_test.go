package usecase_test

import (
	"errors"
	"final-project-backend/entity"
	"final-project-backend/mocks"
	"final-project-backend/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserUsecaseImpl_Register(t *testing.T) {
	t.Run("should return no error", func(t *testing.T) {
		reqBody := &entity.UserRegisterReqBody{}
		
		
		v1:= &entity.User{}
		repo := mocks.NewUserRepository(t)
		repo.On("CheckUsernameExistence", mock.Anything).Return(0, nil)
		repo.On("CheckEmailExistence", mock.Anything).Return("", 0, nil)
		repo.On("CheckPhoneExistence", mock.Anything).Return("", 0, nil)
		repo.On("AddNewUser", mock.Anything).Return(v1, nil)
		
		ucConfig := usecase.UserUsecaseConfig{
			UserRepository: repo,
		}

		uc := usecase.NewUserUsecase(ucConfig)
		_, err := uc.Register(reqBody)

		a := assert.New(t)
		a.Nil(err)
	})
}

func TestUserUsecaseImpl_Login(t *testing.T) {
	t.Run("should return no error", func(t *testing.T) {
		identifier := ""
		password := ""
		
		
		v1:= &entity.User{}
		repo := mocks.NewUserRepository(t)
		repo.On("GetUserByEmailOrUsername", mock.Anything).Return(v1, nil)
		repo.On("GetDetailRole", mock.Anything).Return(nil, errors.New("e"))
		
		ucConfig := usecase.UserUsecaseConfig{
			UserRepository: repo,
		}

		uc := usecase.NewUserUsecase(ucConfig)
		_,_, err := uc.Login(identifier, password)

		a := assert.New(t)
		a.Error(err, "e")
	})
}