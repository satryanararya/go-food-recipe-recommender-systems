// Code generated by mockery v2.43.0. DO NOT EDIT.

package repositories

import (
	context "context"

	entities "github.com/satryanararya/go-chefbot/entities"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockUserFoodPreferencesRepository is an autogenerated mock type for the UserFoodPreferencesRepository type
type MockUserFoodPreferencesRepository struct {
	mock.Mock
}

// AddFoodPreference provides a mock function with given fields: ctx, userFoodPref
func (_m *MockUserFoodPreferencesRepository) AddFoodPreference(ctx context.Context, userFoodPref *entities.UserFoodPreference) error {
	ret := _m.Called(ctx, userFoodPref)

	if len(ret) == 0 {
		panic("no return value specified for AddFoodPreference")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.UserFoodPreference) error); ok {
		r0 = rf(ctx, userFoodPref)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteFoodPreference provides a mock function with given fields: ctx, userID
func (_m *MockUserFoodPreferencesRepository) DeleteFoodPreference(ctx context.Context, userID uuid.UUID) error {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteFoodPreference")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditFoodPreference provides a mock function with given fields: ctx, userFoodPref
func (_m *MockUserFoodPreferencesRepository) EditFoodPreference(ctx context.Context, userFoodPref *entities.UserFoodPreference) error {
	ret := _m.Called(ctx, userFoodPref)

	if len(ret) == 0 {
		panic("no return value specified for EditFoodPreference")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.UserFoodPreference) error); ok {
		r0 = rf(ctx, userFoodPref)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockUserFoodPreferencesRepository creates a new instance of MockUserFoodPreferencesRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserFoodPreferencesRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserFoodPreferencesRepository {
	mock := &MockUserFoodPreferencesRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
