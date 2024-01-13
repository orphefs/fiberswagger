package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	// docs are generated by Swag CLI, you have to import them.
	// replace with your own docs folder, usually "github.com/username/reponame/docs"
	_ "github.com/orphefs/testserver/docs"
	"github.com/orphefs/testserver/models"
	
	"gopkg.in/go-playground/validator.v9"
)

// ResponseModel represents a response model.
// [StatusCode] is the status code of the response. i.e. 200, 400, 500
// [Success] is the status of the response, either true or false
// [Message] is the message of the response.
// [Data] is the data of the response.
type ResponseModel struct {
	StatusCode int         `json:"statusCode"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

// ValidateProduct represents a product in the shopping list.
// [Name] is the name of the product.
// [Description] is the description of the product.
// [Image] is the image of the product, a base64 string.
// [Price] is the price of the product.
type ValidateProduct struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Image       string `json:"image" validate:"required"`
	Price       int    `json:"price" validate:"required"`
}

// ValidateShopping represents a shopping list.
// [Name] is the name of the shopping list.
// [Description] is the description of the shopping list.
type ValidateShopping struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {

	defer models.CloseDatabase()
	
	// Custom config
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       "shoppingList",
	})

	models.SetUpDatabase()

	app.Get("/", HealthCheck)

	app.Post("/api/product", PostProduct)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))

	app.Listen(":8080")
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}

// Create a new product Route
func PostProduct(c *fiber.Ctx) error {

	// TODO: this should go as a global variable
	// Create a new response model, with default values
	responseModel := ResponseModel{
		Success:    true,
		StatusCode: 200,
		Message:    "",
		Data:       nil,
	}

	// Declare a new validateProduct struct
	var validateProduct ValidateProduct

	// Parse body into validateProduct
	if err := c.BodyParser(&validateProduct); err != nil {
		responseModel.Data = err
		responseModel.Message = "Failed to parse JSON body"
		responseModel.StatusCode = 422
		responseModel.Success = false
		return c.Status(responseModel.StatusCode).JSON(responseModel)
	}

	// validate request body using validator
	error := validator.New().Struct(validateProduct)
	if error != nil {
		responseModel.Data = error
		responseModel.Message = "Invalid parameter"
		responseModel.StatusCode = 422
		responseModel.Success = false
		return c.Status(fiber.StatusUnauthorized).JSON(responseModel)
	}

	// Create product
	product := models.CreateProduct(
		validateProduct.Name,
		validateProduct.Description,
		validateProduct.Image,
		validateProduct.Price)

	responseModel.Data = product
	responseModel.Message = "Product created successfully"
	responseModel.StatusCode = 200
	responseModel.Success = true

	return c.Status(responseModel.StatusCode).JSON(responseModel)
}
