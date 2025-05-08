package seed

import (
	"log"

	"andreiolaru.site.bff/internal/infra/repository/modeldb"
	"gorm.io/gorm"
)

func InitSeedData(db *gorm.DB) {
	initMe(db)
	initAbout(db)
	initEducation(db)
	initExperience(db)
}

func initMe(db *gorm.DB) {
	var count int64
	db.Model(&modeldb.MeDB{}).Count(&count)
	if count == 0 {
		me := modeldb.MeDB{
			Title:       "Hi, I'm Andrei 👋",
			Job:         "Software Engineer",
			Description: "I'm a results-driven engineer with a strong passion for designing full-stack software that is fast, secure, and scalable. I combine technical depth with creative vision to build elegant solutions that matter.",
			ImageURL:    "https://media.licdn.com/dms/image/v2/D5603AQGd_2T5wWJEjQ/profile-displayphoto-shrink_800_800/B56ZZ9ukfhGUAg-/0/1745866084244?e=1752105600&v=beta&t=PK27lNh8Y2RcGBggaKcw7Bk0vHgQ-rBx4LW5fH3xgbo",
		}
		if err := db.Create(&me).Error; err != nil {
			log.Printf("Failed to seed MeDB: %v", err)
		} else {
			log.Println("Seeded MeDB")
		}
	}
}

func initAbout(db *gorm.DB) {
	var count int64
	db.Model(&modeldb.AboutDB{}).Count(&count)
	if count == 0 {
		about := modeldb.AboutDB{
			WhoIAmTitle:  "Who I Am",
			WhoIAmText:   "I'm a passionate Software Engineer focused on building scalable, high-performance, and user-centric solutions across the full stack. I adapt quickly to new technologies and frameworks, with hands-on experience in front-end, back-end, and cloud services. Always eager to learn and evolve, I thrive in collaborative environments where I can deliver efficient, secure, and reliable software that drives real impact.",
			MindsetTitle: "Mindset & Vision",
			MindsetText:  "Driven by innovation and purpose, I aim to craft impactful solutions through collaboration, adaptability, and a growth mindset.",
			CVURL:        "https://andreiolaru.com/cv.pdf",
		}
		if err := db.Create(&about).Error; err != nil {
			log.Printf("Failed to seed AboutDB: %v", err)
		} else {
			log.Println("Seeded AboutDB")
		}
	}
}

func initEducation(db *gorm.DB) {
	var count int64
	db.Model(&modeldb.EducationDB{}).Count(&count)
	if count == 0 {
		edu := modeldb.EducationDB{
			Institution: "Ovidius University of Constanta",
			Degree:      "Bachelor degree",
			Period:      "Oct 2019 – July 2022",
			Description: "Studied core concepts of Computer Science, including algorithms, software architecture, databases, and distributed systems.",
			Variant:     "Computer Science",
		}
		if err := db.Create(&edu).Error; err != nil {
			log.Printf("Failed to seed EducationDB: %v", err)
		} else {
			log.Println("Seeded EducationDB")
		}
	}
}

func initExperience(db *gorm.DB) {
	var count int64
	db.Model(&modeldb.ExperienceDB{}).Count(&count)
	if count == 0 {
		exp := modeldb.ExperienceDB{
			Company: "Banca Transilvania SA",
			Title:   "Junior Software Engineer",
			Period:  "Oct 2022 – Feb 2025",
			Description: `• Developed and deployed end-to-end web applications, including front-end interfaces, back-end services, and database architecture;
• Maintained and enhanced existing applications by identifying, tracking, and resolving bugs;
• Provided production support, monitoring and resolving issues to ensure smooth operation of applications;
• Collaborated closely with cross-functional teams, including Business Analysts, Project Managers, and business stakeholders, to gather requirements and deliver high-quality solutions;
• Optimized application performance, security, and database queries through code reviews and best practices.`,
		}
		if err := db.Create(&exp).Error; err != nil {
			log.Printf("Failed to seed ExperienceDB: %v", err)
		} else {
			log.Println("Seeded ExperienceDB")
		}
	}
}
