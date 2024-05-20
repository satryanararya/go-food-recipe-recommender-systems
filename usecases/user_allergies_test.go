package usecases

import (
	"context"
	"testing"

	"github.com/google/uuid"
	dto "github.com/satryanararya/go-chefbot/dto/user"

	mock_driver "github.com/satryanararya/go-chefbot/mocks/drivers/spoonacular_api/ingredients"
	mock_repo "github.com/satryanararya/go-chefbot/mocks/repositories"

	"github.com/satryanararya/go-chefbot/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewUserAllergies(t *testing.T) {
	// Inisialisasi mock repository
	mockIngredientRepo := new(mock_repo.MockIngredientRepository)
	mockUserAllergiesRepo := new(mock_repo.MockUserAllergyRepository)
	mockClient := new(mock_driver.MockIngredientClient)

	// Inisialisasi use case
	uac := NewUserAllergiesUseCase(mockClient, mockIngredientRepo, mockUserAllergiesRepo)

	// Verifikasi bahwa use case berhasil diinisialisasi
	assert.NotNil(t, uac)
}

func TestGetIngredientInfo(t *testing.T) {
	// Inisialisasi mock repositories dan client
	mockIngredientRepo := new(mock_repo.MockIngredientRepository)
	mockUserAllergiesRepo := new(mock_repo.MockUserAllergyRepository)
	mockClient := new(mock_driver.MockIngredientClient)

	// Inisialisasi use case dengan mock repositories dan client
	uac := &userAllergiesUseCase{
		ingredientRepo:    mockIngredientRepo,
		userAllergiesRepo: mockUserAllergiesRepo,
		client:            mockClient,
	}

	// Buat contoh request dan response
	userID := uuid.New()
	dto := &dto.UserAllergiesRequest{IngredientName: "Garlic"}
	ingredient := &entities.Ingredient{ID: int64(1), Name: "Garlic"}

	// Atur perilaku dari mock repositories dan client
	mockIngredientRepo.On("ExistsByName", mock.Anything, dto.IngredientName).Return(true, nil)
	mockIngredientRepo.On("GetByName", mock.Anything, dto.IngredientName).Return(ingredient, nil)
	mockUserAllergiesRepo.On("Save", mock.Anything, userID, ingredient.ID).Return(nil)

	// Panggil fungsi yang ingin diuji
	result, err := uac.AddAllergies(context.Background(), userID, dto)

	// Verifikasi hasil
	assert.NoError(t, err)
	assert.Equal(t, ingredient, result)

	// Verifikasi bahwa semua fungsi mock dipanggil
	mockIngredientRepo.AssertExpectations(t)
	mockUserAllergiesRepo.AssertExpectations(t)
}
