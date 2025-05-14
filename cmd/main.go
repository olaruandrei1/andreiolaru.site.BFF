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
	"andreiolaru.site.bff/internal/app/helpers/httpclient"
	seed "andreiolaru.site.bff/internal/infra/initializer"
	"andreiolaru.site.bff/internal/infra/repository"
	"andreiolaru.site.bff/internal/infra/repository/migration"
	"andreiolaru.site.bff/pkg/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	externalAPIURL := os.Getenv("EXTERNAL_CONTACT_API")

	if externalAPIURL == "" {
		log.Fatal("EXTERNAL_CONTACT_API not set in .env")
	}

	externalAPIKey := os.Getenv("EXTERNAL_CONTACT_API_KEY")
	if externalAPIKey == "" {
		log.Fatal("❌ EXTERNAL_CONTACT_API_KEY not set")
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8mb4&loc=Local", user, pass, host, name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	migration.DropAllTables(db)

	migration.AutoMigrateAll(db)

	seed.InitSeedData(db)
	seed.SeedBackendSkills(db)
	seed.SeedCICDSkills(db)
	seed.SeedCloudSkills(db)
	seed.SeedDatabaseSkills(db)
	seed.SeedFrontendSkills(db)
	seed.SeedQueueSkills(db)
	seed.SeedToolSkills(db)
	seed.SeedMobileSkills(db)
	seed.SeedOtherTopics(db)

	client := httpclient.NewHTTPClientService()

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
	contactService := app.NewContactService(contactRepo, client, externalAPIURL, externalAPIKey)

	router := rest.NewRouter(
		meService, aboutService, experienceService,
		skillService, projectService, educationService, contactService,
	)

	mux := router.RegisterRoutes()

	apiSecured := middleware.ApiKeyAuthMiddleware(mux)
	corsHandler := middleware.CORS(apiSecured)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
