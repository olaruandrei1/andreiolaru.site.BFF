package mapping

import "strings"

var githubTopicAliases = map[string]string{
	"csharp":                    "C#",
	"dotnet":                    ".NET",
	"netcore":                   ".NET",
	"java":                      "Java",
	"spring-boot":               "Spring Boot",
	"go":                        "Go",
	"golang":                    "Go",
	"python":                    "Python",
	"django":                    "Django",
	"git":                       "Git",
	"github":                    "GitHub",
	"azure-devops":              "Azure DevOps",
	"jenkins":                   "Jenkins",
	"jira":                      "Jira",
	"confluence":                "Confluence",
	"agile":                     "Agile",
	"scrum":                     "Scrum",
	"kanban":                    "Kanban",
	"waterfall":                 "Waterfall",
	"azure-functions":           "Azure Functions",
	"azure-storage":             "Azure Storage Container",
	"azure-key-vault":           "Azure Key Vault",
	"azure-active-directory":    "Azure AD",
	"aws-lambda":                "AWS Lambda",
	"aws-dynamodb":              "AWS DynamoDB",
	"aws-s3":                    "AWS S3",
	"mssql":                     "Microsoft SQL Server (T-SQL)",
	"sqlserver":                 "Microsoft SQL Server (T-SQL)",
	"oracle":                    "Oracle (PL/SQL)",
	"plsql":                     "Oracle (PL/SQL)",
	"postgres":                  "PostgreSQL (PL/pgSQL)",
	"postgresql":                "PostgreSQL (PL/pgSQL)",
	"mysql":                     "MySQL",
	"html":                      "HTML",
	"css":                       "CSS",
	"javascript":                "JavaScript",
	"typescript":                "TypeScript",
	"react":                     "React",
	"angular":                   "Angular",
	"vue":                       "Vue",
	"vuejs":                     "Vue",
	"jquery":                    "jQuery",
	"blazor":                    "Blazor",
	"azure-service-bus":         "Azure ServiceBus",
	"aws-sns":                   "AWS SNS",
	"aws-sqs":                   "AWS SQS",
	"kafka":                     "Kafka",
	"ibm-mq":                    "IBM-Queues",
	"grafana":                   "Grafana",
	"elasticsearch":             "Elasticsearch",
	"application-insights":      "Azure Application Insights",
	"kql":                       "KQL",
	"cpp":                       "C++",
	"c++":                       "C++",
	"api-rest":                  "API",
	"xcode":                     "Xcode",
	"swift":                     "Swift",
	"swiftui":                   "SwiftUI",
	"ios":                       "Apple",
	"ios-app":                   "Apple",
	"apple":                     "Apple",
	"net-http":                  "API",
	"mui":                       "Material UI",
	"material-ui":               "Material UI",
	"mediatr":                   "MediatR",
	"mediator":                  "MediatR",
	"mediator-pattern":          "MediatR",
	"ef":                        "Entity Framework",
	"entity-framework":          "Entity Framework",
	"entity-framework-core":     "Entity Framework",
	"stl":                       "STL",
	"standard-template-library": "STL",
	"plpgsql":                   "PostgreSQL (PL/pgSQL)",
	"plpg":                      "PostgreSQL (PL/pgSQL)",
	"plpg-sql":                  "PostgreSQL (PL/pgSQL)",
	"rbac":                      "RABC",
	"adonet":                    "ADO NET",
	"machine-learning":          "Machine Learning",
	"bachelor-thesis":           "Bachelor Thesis",
	"crispy-forms":              "Crispy Forms",
	"discordbot":                "Discord BOT",
}

func NormalizeTopic(topic string) string {
	t := strings.ToLower(strings.TrimSpace(topic))

	if t == "" {
		return ""
	}

	if mapped, exists := githubTopicAliases[t]; exists {
		return mapped
	}

	return strings.ToUpper(t[:1]) + t[1:]
}
