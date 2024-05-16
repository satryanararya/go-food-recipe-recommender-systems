// Code generated by mockery v2.43.0. DO NOT EDIT.

package recipe

import (
	context "context"

	recipe "github.com/satryanararya/go-chefbot/drivers/spoonacular_api/recipe"
	mock "github.com/stretchr/testify/mock"
)

// MockRecipeClient is an autogenerated mock type for the RecipeClient type
type MockRecipeClient struct {
	mock.Mock
}

// GetMultipleRecipeInformation provides a mock function with given fields: ctx, recipeIDs
func (_m *MockRecipeClient) GetMultipleRecipeInformation(ctx context.Context, recipeIDs []int) ([]recipe.RecipeInformation, error) {
	ret := _m.Called(ctx, recipeIDs)

	if len(ret) == 0 {
		panic("no return value specified for GetMultipleRecipeInformation")
	}

	var r0 []recipe.RecipeInformation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []int) ([]recipe.RecipeInformation, error)); ok {
		return rf(ctx, recipeIDs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []int) []recipe.RecipeInformation); ok {
		r0 = rf(ctx, recipeIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]recipe.RecipeInformation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []int) error); ok {
		r1 = rf(ctx, recipeIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRecipeInformation provides a mock function with given fields: ctx, recipeID
func (_m *MockRecipeClient) GetRecipeInformation(ctx context.Context, recipeID int) (recipe.RecipeInformation, error) {
	ret := _m.Called(ctx, recipeID)

	if len(ret) == 0 {
		panic("no return value specified for GetRecipeInformation")
	}

	var r0 recipe.RecipeInformation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (recipe.RecipeInformation, error)); ok {
		return rf(ctx, recipeID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) recipe.RecipeInformation); ok {
		r0 = rf(ctx, recipeID)
	} else {
		r0 = ret.Get(0).(recipe.RecipeInformation)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, recipeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchRecipe provides a mock function with given fields: ctx, name
func (_m *MockRecipeClient) SearchRecipe(ctx context.Context, name string) (recipe.SearchRecipeResponse, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for SearchRecipe")
	}

	var r0 recipe.SearchRecipeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (recipe.SearchRecipeResponse, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) recipe.SearchRecipeResponse); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(recipe.SearchRecipeResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockRecipeClient creates a new instance of MockRecipeClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRecipeClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRecipeClient {
	mock := &MockRecipeClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}