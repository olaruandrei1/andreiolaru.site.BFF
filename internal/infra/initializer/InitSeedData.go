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
			ImageURL:    "personal_photos/photo.png",
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
			CVURL:        "andrei_olaru_se.pdf",
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
		entries := []modeldb.EducationDB{
			{
				Institution: "Ovidius University of Constanta",
				Degree:      "Bachelor degree",
				Period:      "Oct 2019 – July 2022",
				Description: "Studied core concepts of Computer Science, including algorithms, software architecture, databases, and distributed systems.",
				Variant:     "PRIMARY",
				PhotoPath:   "photos/education/ovidius-university-of-constanta.jpeg",
			},
			{
				Institution: "Dometrain",
				Degree:      "Design Patterns in C#: Singleton",
				Period:      "Oct 2024",
				Description: "Covered the Singleton pattern in .NET Core, including design theory, implementation, and anti-pattern discussions.",
				Variant:     "SECONDARY",
				PhotoPath:   "photos/education/dometrain-logo.webp",
			},
			{
				Institution: "Udemy",
				Degree:      "C# .NET with MS SQL – Beginner to Master 2025",
				Period:      "Jan 2025",
				Description: "Complete course on building .NET applications using Dapper, Microsoft SQL Server, T-SQL and Visual Studio. Covered debugging and database integration.",
				Variant:     "SECONDARY",
				PhotoPath:   "photos/education/udemy-logo.png",
			},
			{
				Institution: "Udemy",
				Degree:      "The Complete JavaScript Course 2025: From Zero to Expert!",
				Period:      "Jan 2025",
				Description: "In-depth coverage of JavaScript ES6+, including DOM manipulation, async/await, OOP principles, and building real-world apps.",
				Variant:     "SECONDARY",
				PhotoPath:   "photos/education/udemy-logo.png",
			},
			{
				Institution: "Dometrain",
				Degree:      "Getting Started: C#",
				Period:      "Mar 2025",
				Description: "Focused on C# fundamentals including .NET Core basics, Visual Studio usage, debugging tools, and basic syntax.",
				Variant:     "SECONDARY",
				PhotoPath:   "photos/education/dometrain-logo.webp",
			},
			{
				Institution: "Dometrain",
				Degree:      "Deep Dive: C#",
				Period:      "Mar 2025",
				Description: "Advanced course covering software design principles, design patterns, and building robust C# applications with .NET Core.",
				Variant:     "SECONDARY",
				PhotoPath:   "photos/education/dometrain-logo.webp",
			},
			{
				Institution: "Dometrain",
				Degree:      "Cloud Fundamentals: AWS Services for C# Developers",
				Period:      "Apr 2025",
				Description: "Explored AWS fundamentals relevant to C# devs: Lambda, S3, SQS, SNS, DynamoDB, and API integration with .NET Core.",
				Variant:     "SECONDARY",
				PhotoPath:   "photos/education/dometrain-logo.webp",
			},
		}

		for _, edu := range entries {
			if err := db.Create(&edu).Error; err != nil {
				log.Printf("Failed to seed entry: %v", err)
			}
		}
		log.Println("Seeded EducationDB with full list")
	}
}

func initExperience(db *gorm.DB) {
	var count int64
	db.Model(&modeldb.ExperienceDB{}).Count(&count)
	if count == 0 {
		experiences := []modeldb.ExperienceDB{
			{
				Company:   "EXE Software",
				Title:     "Software Engineer",
				Period:    "Mar 2025 – Present",
				PhotoPath: "photos/experience/exe-software.png",
				Description: `• Working across multiple client-facing projects, focusing on backend, web, and cloud integration — building APIs, developing web applications, and leveraging Azure services to deliver scalable, production-ready solutions;
• Experience includes developing RESTful APIs, background services, and workers using .NET Core 8 and SQL Server, as well as creating web applications with Blazor (Server-Side) and Vue.js (Vue 2 and Vue 3);
• Handled database operations with Microsoft SQL Server and MySQL;
• Integrated external systems through APIs and authentication services such as Azure Active Directory;
• Worked extensively with Azure Functions, Service Bus, Blob Storage, Key Vault, and API Apps, ensuring cloud-based scalability and resilience;
• Deployment processes and version control were managed through Azure DevOps, including automated builds and releases to Azure Web Apps;
• Monitoring and diagnostics were performed using Application Insights and KQL;
• Collaborated with cross-functional teams in client sync meetings and provided support during UAT and production phases.`,
			},
			{
				Company:   "Banca Transilvania SA",
				Title:     "Junior Software Engineer",
				Period:    "Oct 2022 – Feb 2025",
				PhotoPath: "photos/experience/banca-transilvania.png",
				Description: `• Developed and deployed end-to-end web applications, including front-end interfaces, back-end services, and database architecture;
• Maintained and enhanced existing applications by identifying, tracking, and resolving bugs;
• Provided production support, monitoring and resolving issues to ensure smooth operation of applications;
• Collaborated closely with cross-functional teams, including Business Analysts, Project Managers, and business stakeholders, to gather requirements and deliver high-quality solutions;
• Optimized application performance, security, and database queries through code reviews and best practices.`,
			},
			{
				Company:   "Banca Transilvania SA",
				Title:     "Software Engineer Intern",
				Period:    "Sep 2022 – Sep 2022",
				PhotoPath: "photos/experience/banca-transilvania.png",
				Description: `• Assisted in the development and maintenance of web applications by implementing features and fixing bugs;
• Collaborated with senior developers to improve application performance, security, and code quality;
• Gained hands-on experience with front-end, back-end, and database technologies;
• Participated in code reviews, team meetings, and agile ceremonies to better understand software development processes;
• Continuously learned and adapted to new tools and frameworks to meet project needs.`,
			},
		}

		if err := db.Create(&experiences).Error; err != nil {
			log.Printf("Failed to seed ExperienceDB: %v", err)
		} else {
			log.Println("Seeded ExperienceDB")
		}
	}
}
