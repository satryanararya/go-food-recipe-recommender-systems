// Code generated by mockery v2.43.0. DO NOT EDIT.

package repositories

import (
	context "context"

	entities "github.com/satryanararya/go-chefbot/entities"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockUserAllergyRepository is an autogenerated mock type for the UserAllergyRepository type
type MockUserAllergyRepository struct {
	mock.Mock
}

// GetAllergies provides a mock function with given fields: ctx, userID
func (_m *MockUserAllergyRepository) GetAllergies(ctx context.Context, userID uuid.UUID) ([]*entities.UserAllergies, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetAllergies")
	}

	var r0 []*entities.UserAllergies
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]*entities.UserAllergies, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []*entities.UserAllergies); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.UserAllergies)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, userAllergy
func (_m *MockUserAllergyRepository) Save(ctx context.Context, userAllergy *entities.UserAllergies) error {
	ret := _m.Called(ctx, userAllergy)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.UserAllergies) error); ok {
		r0 = rf(ctx, userAllergy)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockUserAllergyRepository creates a new instance of MockUserAllergyRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserAllergyRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserAllergyRepository {
	mock := &MockUserAllergyRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
