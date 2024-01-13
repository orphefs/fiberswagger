package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Product represents a product in the shopping list.
// [Name] is the name of the product.
// [Description] is the description of the product.
// [Image] is the image of the product.
// [Price] is the price of the product.
type Product struct {
	gorm.Model
	Name        string `gorm:"unique" json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Price       int    `gorm:"default:0" json:"price"`
}

// ShoppingList represents a category for organizing product lists.
// [Name] is the name of the ShoppingList.
// [Description] is the description of the ShoppingList.
type ShoppingList struct {
	gorm.Model
	Name        string    `gorm:"unique" json:"name"`
	Description string    `json:"description"`
	Products    []Product `gorm:"many2many:shopping_list_products"`
}

// SetUpDatabase sets up the database with models and auto migrate.
func SetUpDatabase() {
	// Connect to sqlite database
	// you can use any database driver you wish
	db, err := gorm.Open(sqlite.Open("testserver.db"), &gorm.Config{})

	// Connect to Postgres database
	// db, err := gorm.Open(postgres.Open("host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"), &gorm.Config{})

	// Connect to MySQL database
	// db, err := gorm.Open(mysql.Open("gorm:gorm@tcp(localhost:9920)/gorm?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	// Check if there is an error when connecting to database
	if err != nil {
		panic("failed to connect database")
	}

	// Auto migrate models schema to database
	db.AutoMigrate(&Product{}, &ShoppingList{})
}

// getDatabase returns a database connection.
func getDatabase() *gorm.DB {
	// Connect to sqlite database
	db, err := gorm.Open(sqlite.Open("testserver.db"), &gorm.Config{})
	// Check if there is an error when connecting to database
	if err != nil {
		panic("failed to connect database")
	}
	// return database
	return db
}

// GetProduct returns a product with the given ID.
// If ID is nil, it returns all products.
func GetProducts(id int) []Product {
	// declare products as an array
	var products []Product
	// get database
	db := getDatabase()
	if id == 0 {
		// get all products
		db.Find(&products)
	} else {
		// get product with the given id
		db.Find(&products, id)
	}

	return products
}

// CreateProduct creates a product with the given name, description, image and price.
func CreateProduct(name string, description string, image string, price int) Product {
	// create product
	db := getDatabase()
	// declare product
	product := Product{
		Name:        name,
		Description: description,
		Image:       image,
		Price:       price,
	}
	// create product
	db.Create(&product)
	// return product
	return product
}

// UpdateProduct updates a product with the given ID.
// if name, description or price is nil, it will not be updated.
func UpdateProduct(id uint, newProduct Product) Product {
	// create product
	db := getDatabase()
	// declare product
	product := newProduct
	// update product
	db.Model(&product).Where("id = ?", id).Updates(&product)
	// return product
	return product
}

// DeleteProduct deletes a product with the given ID.
func DeleteProduct(id uint) bool {
	// create product
	db := getDatabase()
	// delete product
	db.Delete(&Product{}, id)
	return true
}

// GetShoppingList returns a shopping list with the given ID.
// If ID is nil, it returns all shopping lists.
func FetchShoppingList(shoppingListID int) ([]ShoppingList, error) {
	// Fetch shopping list
	var shoppingList []ShoppingList
	// get database
	db := getDatabase()
	// check shoppingListID
	if shoppingListID == 0 {
		// get all shopping lists
		db.Preload("Products").Find(&shoppingList)

	} else {
		// get shopping list
		db.Preload("Products").First(&shoppingList, shoppingListID)
	}
	// return shopping list
	return shoppingList, nil
}

func GetShoppingList(id int) []ShoppingList {
	// declare products as an array
	var shoppingList []ShoppingList
	// get database
	db := getDatabase()
	if id == 0 {
		// get all products
		db.Find(&shoppingList)
	}
	// get product with the given id
	db.First(&shoppingList, id)
	return shoppingList
}

func CreateShoppingList(name string, description string) ShoppingList {
	// create shopping list
	db := getDatabase()
	// declare shopping list
	shoppingList := ShoppingList{
		Name:        name,
		Description: description,
	}
	// create shopping list
	db.Create(&shoppingList)
	// return shopping list
	return shoppingList
}

func UpdateShoppingList(id uint, shoppingList ShoppingList) ShoppingList {
	// create product
	db := getDatabase()
	// declare product
	shopping := shoppingList
	// update product
	db.Model(&shopping).Where("id = ?", id).Updates(&shopping)
	// return product
	return shopping
}

// Delete Shopping List
func DeleteShoppingList(id uint) bool {
	// create product
	db := getDatabase()
	// delete product
	db.Delete(&ShoppingList{}, id)
	return true
}

// add product to shopping list
func AddProductToShoppingList(shoppingListID int, productID int) bool {
	// create shopping list
	db := getDatabase()
	// add product to shopping list
	db.Model(&ShoppingList{Model: gorm.Model{ID: uint(shoppingListID)}}).Association("Products").Append(&Product{Model: gorm.Model{ID: uint(productID)}})
	return true
}

// add product to shopping list
func RemoveProductToShoppingList(shoppingListID int, productID int) bool {
	// create shopping list
	db := getDatabase()
	// remove product from shopping list
	db.Model(&ShoppingList{Model: gorm.Model{ID: uint(shoppingListID)}}).Association("Products").Delete(&Product{Model: gorm.Model{ID: uint(productID)}})
	return true
}
