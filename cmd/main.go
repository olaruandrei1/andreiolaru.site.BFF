package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"andreiolaru.site.bff/internal/adapter/rest"
	"andreiolaru.site.bff/internal/app"
	"andreiolaru.site.bff/internal/infra/repository"
	"andreiolaru.site.bff/pkg/middleware"
)

func main() {
	// .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Build DSN din variabilele .env
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, pass, host, name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	meRepo := repository.NewGormMeRepository(db)
	aboutRepo := repository.NewGormAboutRepository(db)
	experienceRepo := repository.NewGormExperienceRepository(db)
	skillRepo := repository.NewGormSkillRepository(db)
	projectRepo := repository.NewGormProjectRepository(db)
	educationRepo := repository.NewGormEducationRepository(db)
	contactRepo := repository.NewGormContactRepository(db)

	meService := app.NewMeService(meRepo)
	aboutService := app.NewAboutService(aboutRepo)
	experienceService := app.NewExperienceService(experienceRepo)
	skillService := app.NewSkillService(skillRepo)
	projectService := app.NewProjectService(projectRepo)
	educationService := app.NewEducationService(educationRepo)
	contactService := app.NewContactService(contactRepo)

	router := rest.NewRouter(
		meService, aboutService, experienceService,
		skillService, projectService, educationService, contactService,
	)

	mux := router.RegisterRoutes()
	secured := middleware.ApiKeyAuthMiddleware(mux)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", secured))
}
