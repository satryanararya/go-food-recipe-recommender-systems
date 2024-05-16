// Code generated by mockery v2.43.0. DO NOT EDIT.

package repositories

import (
	context "context"

	entities "github.com/satryanararya/go-chefbot/entities"
	mock "github.com/stretchr/testify/mock"
)

// MockUserCookingSkillRepository is an autogenerated mock type for the UserCookingSkillRepository type
type MockUserCookingSkillRepository struct {
	mock.Mock
}

// AddCookingSkill provides a mock function with given fields: ctx, userCookingSkill
func (_m *MockUserCookingSkillRepository) AddCookingSkill(ctx context.Context, userCookingSkill *entities.UserCookingSkill) error {
	ret := _m.Called(ctx, userCookingSkill)

	if len(ret) == 0 {
		panic("no return value specified for AddCookingSkill")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.UserCookingSkill) error); ok {
		r0 = rf(ctx, userCookingSkill)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditCookingSkill provides a mock function with given fields: ctx, userCookingSkill
func (_m *MockUserCookingSkillRepository) EditCookingSkill(ctx context.Context, userCookingSkill *entities.UserCookingSkill) error {
	ret := _m.Called(ctx, userCookingSkill)

	if len(ret) == 0 {
		panic("no return value specified for EditCookingSkill")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.UserCookingSkill) error); ok {
		r0 = rf(ctx, userCookingSkill)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockUserCookingSkillRepository creates a new instance of MockUserCookingSkillRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserCookingSkillRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserCookingSkillRepository {
	mock := &MockUserCookingSkillRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
