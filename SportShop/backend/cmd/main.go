package main

import (
	"log"
	"net/http"
	"github.com/joho/godotenv"

	"sportshop/backend/internal/database"
	"sportshop/backend/internal/handler"
	"sportshop/backend/internal/repository"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = database.RunMigration(db, "backend/internal/database/migration.sql")
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	err = database.RunSQLFile(db, "backend/internal/database/data.sql")
	if err != nil {
		log.Fatal("Data seeding failed:", err)
	}

	userRepo := repository.NewUserRepository(db)
	userHandler := handler.NewUserHandler(userRepo)

	productRepo := repository.NewProductRepository(db)
	productHandler := handler.NewProductHandler(productRepo)


	http.HandleFunc("/users", userHandler.GetAllUsers)

http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		productHandler.GetAllProducts(w, r)
	} else if r.Method == http.MethodPost {
		productHandler.CreateProduct(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
})

	http.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			productHandler.GetProductByID(w, r)
		} else if r.Method == http.MethodPut {
			productHandler.UpdateProduct(w, r)
		} else if r.Method == http.MethodDelete {
			productHandler.DeleteProduct(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server is running on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}