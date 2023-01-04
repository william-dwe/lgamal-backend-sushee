// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "final-project-backend/entity"

	mock "github.com/stretchr/testify/mock"
)

// CouponRepository is an autogenerated mock type for the CouponRepository type
type CouponRepository struct {
	mock.Mock
}

// AddCoupon provides a mock function with given fields: coupon
func (_m *CouponRepository) AddCoupon(coupon *entity.Coupon) (*entity.Coupon, error) {
	ret := _m.Called(coupon)

	var r0 *entity.Coupon
	if rf, ok := ret.Get(0).(func(*entity.Coupon) *entity.Coupon); ok {
		r0 = rf(coupon)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Coupon)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Coupon) error); ok {
		r1 = rf(coupon)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddUserCoupon provides a mock function with given fields: userCoupon
func (_m *CouponRepository) AddUserCoupon(userCoupon *entity.UserCoupon) (*entity.UserCoupon, error) {
	ret := _m.Called(userCoupon)

	var r0 *entity.UserCoupon
	if rf, ok := ret.Get(0).(func(*entity.UserCoupon) *entity.UserCoupon); ok {
		r0 = rf(userCoupon)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.UserCoupon)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.UserCoupon) error); ok {
		r1 = rf(userCoupon)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCouponById provides a mock function with given fields: couponId
func (_m *CouponRepository) DeleteCouponById(couponId int) (*entity.Coupon, error) {
	ret := _m.Called(couponId)

	var r0 *entity.Coupon
	if rf, ok := ret.Get(0).(func(int) *entity.Coupon); ok {
		r0 = rf(couponId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Coupon)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(couponId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCoupon provides a mock function with given fields:
func (_m *CouponRepository) GetCoupon() (*[]entity.Coupon, error) {
	ret := _m.Called()

	var r0 *[]entity.Coupon
	if rf, ok := ret.Get(0).(func() *[]entity.Coupon); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entity.Coupon)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCouponById provides a mock function with given fields: couponId
func (_m *CouponRepository) GetCouponById(couponId int) (*entity.Coupon, error) {
	ret := _m.Called(couponId)

	var r0 *entity.Coupon
	if rf, ok := ret.Get(0).(func(int) *entity.Coupon); ok {
		r0 = rf(couponId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Coupon)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(couponId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserCouponByCouponCode provides a mock function with given fields: userId, couponCode
func (_m *CouponRepository) GetUserCouponByCouponCode(userId int, couponCode string) (*entity.UserCoupon, error) {
	ret := _m.Called(userId, couponCode)

	var r0 *entity.UserCoupon
	if rf, ok := ret.Get(0).(func(int, string) *entity.UserCoupon); ok {
		r0 = rf(userId, couponCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.UserCoupon)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(userId, couponCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserCouponByUsername provides a mock function with given fields: username
func (_m *CouponRepository) GetUserCouponByUsername(username string) (*[]entity.UserCoupon, int, error) {
	ret := _m.Called(username)

	var r0 *[]entity.UserCoupon
	if rf, ok := ret.Get(0).(func(string) *[]entity.UserCoupon); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entity.UserCoupon)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(string) int); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(username)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateCouponById provides a mock function with given fields: couponId, newCoupon
func (_m *CouponRepository) UpdateCouponById(couponId int, newCoupon *entity.Coupon) error {
	ret := _m.Called(couponId, newCoupon)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, *entity.Coupon) error); ok {
		r0 = rf(couponId, newCoupon)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCouponRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewCouponRepository creates a new instance of CouponRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCouponRepository(t mockConstructorTestingTNewCouponRepository) *CouponRepository {
	mock := &CouponRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}