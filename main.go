package main

import (
	"MiniProject-MuhammadRendyansyahDwiAnggara/config"
	"MiniProject-MuhammadRendyansyahDwiAnggara/handlers"
	"MiniProject-MuhammadRendyansyahDwiAnggara/middlewares"
	"MiniProject-MuhammadRendyansyahDwiAnggara/repositories"
	"MiniProject-MuhammadRendyansyahDwiAnggara/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("File .env tidak ditemukan, menggunakan environment system")
	}

	config.ConnectDatabase()
	config.MigrateDatabase()

	app := fiber.New()
	api := app.Group("/api/v1")
	app.Use("/uploads", static.New("./uploads"))

	// ... kode Dependency Injection & Routes di bawahnya tetap sama

	// Dependency Injection
	userRepo := repositories.NewUserRepository(config.DB)
	provCityService := services.NewProvCityService()
	authService := services.NewAuthService(userRepo, provCityService)
	userService := services.NewUserService(userRepo, provCityService)

	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	provCityHandler := handlers.NewProvCityHandler(provCityService)

	// Routes Auth
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	// Routes User
	user := api.Group("/user", middlewares.JWTMiddleware)
	user.Get("/", userHandler.GetMyProfile)
	user.Put("/", userHandler.UpdateProfile)

	// Routes ProvCity
	provcity := api.Group("/provcity")
	provcity.Get("/listprovincies", provCityHandler.GetListProvinces)
	provcity.Get("/listcities/:prov_id", provCityHandler.GetListCities)
	provcity.Get("/detailprovince/:prov_id", provCityHandler.GetDetailProvince)
	provcity.Get("/detailcity/:city_id", provCityHandler.GetDetailCity)

    addressRepo := repositories.NewAddressRepository(config.DB)
	categoryRepo := repositories.NewCategoryRepository(config.DB)
	storeRepo := repositories.NewStoreRepository(config.DB)

	addressService := services.NewAddressService(addressRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	storeService := services.NewStoreService(storeRepo)

	addressHandler := handlers.NewAddressHandler(addressService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	storeHandler := handlers.NewStoreHandler(storeService)

	// Routes Alamat (di bawah grup user yang sudah pakai JWTMiddleware)
	alamat := user.Group("/alamat")
	alamat.Get("/", addressHandler.GetMyAddresses)
	alamat.Get("/:id", addressHandler.GetByID)
	alamat.Post("/", addressHandler.Create)
	alamat.Put("/:id", addressHandler.Update)
	alamat.Delete("/:id", addressHandler.Delete)

	// Routes Toko (Wajib JWT)
	toko := api.Group("/toko", middlewares.JWTMiddleware)
	toko.Get("/my", storeHandler.GetMyStore)
	toko.Put("/:id_toko", storeHandler.UpdateStore)
	toko.Get("/:id_toko", storeHandler.GetByID)
	toko.Get("/", storeHandler.GetAll)

	// Routes Category
	category := api.Group("/category")
	category.Get("/", categoryHandler.GetAll)
	category.Get("/:id", middlewares.JWTMiddleware, categoryHandler.GetByID)
	category.Post("/", middlewares.JWTMiddleware, middlewares.AdminMiddleware, categoryHandler.Create)
	category.Put("/:id", middlewares.JWTMiddleware, middlewares.AdminMiddleware, categoryHandler.Update)
	category.Delete("/:id", middlewares.JWTMiddleware, middlewares.AdminMiddleware, categoryHandler.Delete)

	// Dependency Injection Product & Transaction
	productRepo := repositories.NewProductRepository(config.DB)
	trxRepo := repositories.NewTransactionRepository(config.DB)

	productService := services.NewProductService(productRepo, storeRepo, categoryRepo)
	trxService := services.NewTransactionService(trxRepo)

	productHandler := handlers.NewProductHandler(productService)
	trxHandler := handlers.NewTransactionHandler(trxService)

	// Routes Product
	product := api.Group("/product")
	product.Get("/", productHandler.GetAll)
	product.Get("/:id", productHandler.GetByID)
	product.Post("/", middlewares.JWTMiddleware, productHandler.Create)
	product.Put("/:id", middlewares.JWTMiddleware, productHandler.Update)
	product.Delete("/:id", middlewares.JWTMiddleware, productHandler.Delete)

	// Routes Transaction (Wajib JWT)
	trx := api.Group("/trx", middlewares.JWTMiddleware)
	trx.Get("/", trxHandler.GetAll)
	trx.Get("/:id", trxHandler.GetByID)
	trx.Post("/", trxHandler.Create)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}

	log.Fatal(app.Listen(":" + port))
}