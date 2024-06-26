// Code generated by mockery v2.43.0. DO NOT EDIT.

package repositories

import (
	context "context"

	entities "github.com/satryanararya/go-chefbot/entities"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockFavoriteRecipeRepository is an autogenerated mock type for the FavoriteRecipeRepository type
type MockFavoriteRecipeRepository struct {
	mock.Mock
}

// AddToFavorites provides a mock function with given fields: ctx, favoriteRecipe
func (_m *MockFavoriteRecipeRepository) AddToFavorites(ctx context.Context, favoriteRecipe *entities.FavoriteRecipe) error {
	ret := _m.Called(ctx, favoriteRecipe)

	if len(ret) == 0 {
		panic("no return value specified for AddToFavorites")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.FavoriteRecipe) error); ok {
		r0 = rf(ctx, favoriteRecipe)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindFavoritesByUserID provides a mock function with given fields: ctx, userID
func (_m *MockFavoriteRecipeRepository) FindFavoritesByUserID(ctx context.Context, userID uuid.UUID) ([]entities.FavoriteRecipe, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for FindFavoritesByUserID")
	}

	var r0 []entities.FavoriteRecipe
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]entities.FavoriteRecipe, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []entities.FavoriteRecipe); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.FavoriteRecipe)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveFromFavorites provides a mock function with given fields: ctx, userID, recipeID
func (_m *MockFavoriteRecipeRepository) RemoveFromFavorites(ctx context.Context, userID uuid.UUID, recipeID int64) error {
	ret := _m.Called(ctx, userID, recipeID)

	if len(ret) == 0 {
		panic("no return value specified for RemoveFromFavorites")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, int64) error); ok {
		r0 = rf(ctx, userID, recipeID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockFavoriteRecipeRepository creates a new instance of MockFavoriteRecipeRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFavoriteRecipeRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFavoriteRecipeRepository {
	mock := &MockFavoriteRecipeRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
