package usecases

import (
	// "context"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"

	// "time"
	// "errors"
	// "net/http"
	// "net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/satryanararya/go-chefbot/entities"
	mock_service "github.com/satryanararya/go-chefbot/mocks/drivers/cloudinary"
	mock_driver "github.com/satryanararya/go-chefbot/mocks/drivers/spoonacular_api/recipe"
	mock_repo "github.com/satryanararya/go-chefbot/mocks/repositories"

	// "github.com/satryanararya/go-chefbot/drivers/spoonacular_api/recipe"
	model "github.com/satryanararya/go-chefbot/drivers/spoonacular_api/recipe"

	dto "github.com/satryanararya/go-chefbot/dto/recipe"
	// "github.com/labstack/echo/v4"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewRecipeUseCase(t *testing.T) {
	// Inisialisasi mock client
	mockClient := new(mock_driver.MockRecipeClient)
	mockRecipeRepo := new(mock_repo.MockRecipeRepository)
	mockCloudinaryService := new(mock_service.MockCloudinaryService)

	// Inisialisasi use case
	ruc := NewRecipeUseCase(mockClient, mockRecipeRepo, mockCloudinaryService)

	// Verifikasi bahwa use case berhasil diinisialisasi
	assert.NotNil(t, ruc)
}

func TestSearchRecipe(t *testing.T) {
	// Initialize mock client
	mockClient := new(mock_driver.MockRecipeClient)

	// Initialize use case with mock client
	ruc := &recipeUseCase{
		client: mockClient,
	}

	// Create example request and response
	name := "Garlic"
	searchResponse := model.SearchRecipeResponse{
		Results: []model.Recipe{
			{
				ID:    1,
				Title: "Garlic Bread",
				Image: "https://example.com/garlic_bread.jpg",
			},
		},
		Offset:       0,
		Number:       1,
		TotalResults: 1,
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/search", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set behavior of mock client
	mockClient.On("SearchRecipe", mock.Anything, name).Return(searchResponse, nil)

	// Call the function to be tested
	result, err := ruc.SearchRecipe(c, name)

	// Verify the result
	assert.NoError(t, err)
	assert.Equal(t, searchResponse, result)

	// Verify that all mock functions were called
	mockClient.AssertExpectations(t)
}

func TestGetRecipeInformation(t *testing.T) {
	// Initialize mock client
	mockClient := new(mock_driver.MockRecipeClient)

	// Initialize use case with mock client
	ruc := &recipeUseCase{
		client: mockClient,
	}

	// Create example response
	recipeID := 1
	recipeInformationResponse := model.RecipeInformation{
		Title:               "Garlic Bread",
		Image:               "https://example.com/garlic_bread.jpg",
		SourceName:          "Source",
		CookingMinutes:      30,
		ExtendedIngredients: []model.ExtendedIngredient{{Name: "Garlic", Amount: 2, Unit: "cloves"}},
		PricePerServing:     1.5,
		ReadyInMinutes:      45,
		Servings:            4,
		HealthScore:         80,
		Diets:               []string{"vegetarian"},
		IsSustainable:       false,
		Instruction:         "Bake in oven",
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/info/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set behavior of mock client
	mockClient.On("GetRecipeInformation", mock.Anything, recipeID).Return(recipeInformationResponse, nil)

	// Call the function to be tested
	result, err := ruc.GetRecipeInformation(c, recipeID)

	// Verify the result
	assert.NoError(t, err)
	assert.Equal(t, recipeInformationResponse, result)

	// Verify that all mock functions were called
	mockClient.AssertExpectations(t)
}

func TestCreateRecipe(t *testing.T) {
    // Initialize mock client and mock cloudinary service
    mockClient := new(mock_driver.MockRecipeClient)
    mockCloudinaryService := new(mock_service.MockCloudinaryService)
    mockRecipeRepo := new(mock_repo.MockRecipeRepository)

    // Initialize use case with mock client and mock cloudinary service
    ruc := &recipeUseCase{
        client:            mockClient,
        cloudinaryService: mockCloudinaryService,
        recipeRepo:        mockRecipeRepo,
    }

    // Create example request and response
    id := int64(1)
    req := &dto.CreateRecipeRequest{
        Title:           "Garlic Bread",
        SourceName:      "Source",
        CookingMinutes:  30,
        PricePerServing: 1.5,
        ReadyInMinutes:  45,
        Servings:        4,
        HealthScore:     80,
        Diets:           "vegetarian",
        IsSustainable:   false,
        Instruction:     "Bake in oven",
    }
    imageURL := "https://example.com/garlic_bread.jpg"

    // Create a new buffer
    buf := new(bytes.Buffer)
    // Create a new multipart writer
    writer := multipart.NewWriter(buf)

    // Add the fields of the recipe to the form
    _ = writer.WriteField("Title", req.Title)
    _ = writer.WriteField("SourceName", req.SourceName)
    _ = writer.WriteField("CookingMinutes", strconv.Itoa(int(req.CookingMinutes)))
    _ = writer.WriteField("PricePerServing", fmt.Sprintf("%f", req.PricePerServing))
    _ = writer.WriteField("ReadyInMinutes", strconv.Itoa(int(req.ReadyInMinutes)))
    _ = writer.WriteField("Servings", strconv.Itoa(int(req.Servings)))
    _ = writer.WriteField("HealthScore", strconv.Itoa(int(req.HealthScore)))
    _ = writer.WriteField("Diets", req.Diets)
    _ = writer.WriteField("IsSustainable", strconv.FormatBool(req.IsSustainable))
    _ = writer.WriteField("Instruction", req.Instruction)

    dummyFile, err := os.Create("test.jpg")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove("test.jpg")

    // Add the image file to the form
    fileWriter, err := writer.CreateFormFile("image", "test.jpg")
    if err != nil {
        t.Fatal(err)
    }

    _, err = io.Copy(fileWriter, dummyFile)
    if err != nil {
        t.Fatal(err)
    }
    dummyFile.Close()

    // Close the multipart writer
    _ = writer.Close()

    e := echo.New()
    // Create the HTTP request with the multipart form data
    httpReq := httptest.NewRequest(http.MethodPost, "/create", buf)
    httpReq.Header.Set("Content-Type", writer.FormDataContentType())
    rec := httptest.NewRecorder()
    c := e.NewContext(httpReq, rec)

    // Set behavior of mock cloudinary service
    mockCloudinaryService.On("UploadImage", mock.Anything, mock.Anything).Return(imageURL, nil)

    // Set behavior of mock recipe repository
    mockRecipeRepo.On("CreateRecipe", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
        recipe := args.Get(1).(*entities.Recipe)
        recipe.ID = id
    })

    // Call the function to be tested
    result, err := ruc.CreateRecipe(c, id, req)

    // Verify the result
    assert.NoError(t, err)
    assert.Equal(t, &dto.CreateRecipeResponse{
        ID:    id,
        Title: req.Title,
        Image: imageURL,
    }, result)

    // Verify that all mock functions were called
    mockCloudinaryService.AssertExpectations(t)
    mockRecipeRepo.AssertExpectations(t)
}