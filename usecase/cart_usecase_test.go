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

func TestCartUsecaseImpl_GetCart(t *testing.T) {
	t.Run("should return nil & error", func(t *testing.T) {
		username := "user"
		repo := mocks.NewCartRepository(t)
		repo.On("GetCartByUsername", mock.Anything).Return(nil, errors.New("e"))
		ucConfig := usecase.CartUsecaseConfig{
			CartRepository: repo,
		}

		uc := usecase.NewCartUsecase(ucConfig)
		res, err := uc.GetCart(username)

		a := assert.New(t)
		a.Nil(res)
		a.Error(err, "e")
	})
	t.Run("should return cart nil", func(t *testing.T) {
		username := "user"
		carts := &[]entity.Cart{}
		repo := mocks.NewCartRepository(t)
		repo.On("GetCartByUsername", mock.Anything).Return(carts, nil)
		ucConfig := usecase.CartUsecaseConfig{
			CartRepository: repo,
		}

		uc := usecase.NewCartUsecase(ucConfig)
		res, err := uc.GetCart(username)

		a := assert.New(t)
		a.Equal(carts, res)
		a.Nil(err)
	})
}

func TestCartUsecaseImpl_AddCart(t *testing.T) {
	t.Run("should return nil & error", func(t *testing.T) {
		username := "user"
		repo := mocks.NewCartRepository(t)
		repo.On("GetCartByUsername", mock.Anything).Return(nil, errors.New("e"))
		ucConfig := usecase.CartUsecaseConfig{
			CartRepository: repo,
		}

		uc := usecase.NewCartUsecase(ucConfig)
		res, err := uc.GetCart(username)

		a := assert.New(t)
		a.Nil(res)
		a.Error(err, "e")
	})
	t.Run("should return nil & error", func(t *testing.T) {
		username := "user"
		c := &entity.CartReqBody{}
		v1 := &entity.User{}
		err := errors.New("e")
		repo1 := mocks.NewUserRepository(t)
		repo1.On("GetUserByEmailOrUsername", mock.Anything).Return(v1, nil)


		repo3 := mocks.NewCartRepository(t)
		repo3.On("AddItemToCart", mock.Anything).Return(nil, err)
		ucConfig := usecase.CartUsecaseConfig{
			UserRepository: repo1,
			CartRepository: repo3,
		}

		uc := usecase.NewCartUsecase(ucConfig)
		res, err := uc.AddCart(username, c)

		a := assert.New(t)
		a.Nil(res)
		a.Error(err, "e")
	})
	t.Run("should return cart nil", func(t *testing.T) {
		username := "user"
		c := &entity.CartReqBody{}
		v1 := &entity.User{}
		v3 := &entity.Cart{}
		v4 := &entity.CartResBody{}
		repo1 := mocks.NewUserRepository(t)
		repo1.On("GetUserByEmailOrUsername", mock.Anything).Return(v1, nil)


		repo3 := mocks.NewCartRepository(t)
		repo3.On("AddItemToCart", mock.Anything).Return(v3, nil)
		ucConfig := usecase.CartUsecaseConfig{
			UserRepository: repo1,
			CartRepository: repo3,
		}

		uc := usecase.NewCartUsecase(ucConfig)
		res, err := uc.AddCart(username, c)

		a := assert.New(t)
		a.Equal(v4, res)
		a.Nil(err)
	})
}

func TestCartUsecaseImpl_DeleteCart(t *testing.T) {
	t.Run("should return nil & error", func(t *testing.T) {
		username := "user"
		repo := mocks.NewCartRepository(t)
		repo.On("DeleteCart", mock.Anything).Return(nil)
		ucConfig := usecase.CartUsecaseConfig{
			CartRepository: repo,
		}

		uc := usecase.NewCartUsecase(ucConfig)
		err := uc.DeleteCart(username)

		a := assert.New(t)
		a.Nil(err)
	})
}


func TestCartUsecaseImpl_DeleteCartByCartId(t *testing.T) {
	t.Run("should return nil & error", func(t *testing.T) {
		username := "user"
		cartId := 1
		v1:= &entity.Cart{}
		v2:= &entity.User{}
		repo := mocks.NewCartRepository(t)
		repo.On("GetCartByCartId", mock.Anything).Return(v1, nil)
		repo.On("DeleteCartByCartId", mock.Anything).Return(nil)
		repo2 := mocks.NewUserRepository(t)
		repo2.On("GetUserByEmailOrUsername", mock.Anything).Return(v2, nil)
		ucConfig := usecase.CartUsecaseConfig{
			CartRepository: repo,
			UserRepository: repo2,
		}

		uc := usecase.NewCartUsecase(ucConfig)
		err := uc.DeleteCartByCartId(username, cartId)

		a := assert.New(t)
		a.Nil(err)
	})
}

func TestCartUsecaseImpl_UpdateCartByCartId(t *testing.T) {
	t.Run("should return nil & error", func(t *testing.T) {
		username := "user"
		cartId := 1
		reqbody := &entity.CartEditDetailsReqBody{}
		v1:= &entity.Cart{}
		v2:= &entity.User{}
		repo := mocks.NewCartRepository(t)
		repo.On("GetCartByCartId", mock.Anything).Return(v1, nil)
		repo.On("UpdateCartByCartId", mock.Anything, mock.Anything).Return(nil)
		repo.On("GetCartByCartId", mock.Anything).Return(v1, nil)
		repo2 := mocks.NewUserRepository(t)
		repo2.On("GetUserByEmailOrUsername", mock.Anything).Return(v2, nil)
		ucConfig := usecase.CartUsecaseConfig{
			CartRepository: repo,
			UserRepository: repo2,
		}

		uc := usecase.NewCartUsecase(ucConfig)
		_, err := uc.UpdateCartByCartId(username, cartId, reqbody)

		a := assert.New(t)
		a.Nil(err)
	})
}