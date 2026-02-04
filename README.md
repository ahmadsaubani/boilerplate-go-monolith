# Go Boilerplate API 🚀

A clean and modular Go REST API boilerplate built with Gin framework, featuring structured logging, database connections, and clean architecture principles.

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.25-blue.svg)](https://golang.org/)
[![Gin Framework](https://img.shields.io/badge/gin-v1.11.0-green.svg)](https://gin-gonic.com/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

---

## 🌟 Features

- **🏗️ Clean Architecture** - Organized in layers with clear separation of concerns
- **🚀 Gin Framework** - Fast HTTP web framework with middleware support
- **📊 Structured Logging** - Integration with [go-logging-lib](https://github.com/ahmadsaubani/go-logging-lib) for comprehensive logging
- **💾 Database Support** - PostgreSQL and Redis connections with connection pooling
- **📧 Email Service** - Built-in email functionality with SMTP support
- **⚙️ Environment-based Configuration** - YAML-based configuration for different environments
- **🔧 Dependency Injection** - Clean dependency management with DTO pattern
- **🛡️ Middleware Stack** - Request logging, error handling, and panic recovery
- **📁 Daily Log Rotation** - Automatic log file rotation with structured formats

---

## 📂 Project Structure

```
boilerplate-go/
├── 📁 Config/                     # Configuration management
│   ├── 📁 DTO/                    # Data Transfer Objects
│   │   ├── config.go              # DTO package documentation
│   │   ├── module_config.go       # Repository configuration
│   │   └── utilities.go           # Service utilities
│   ├── app_config.go              # Application & API configuration
│   ├── database_config.go         # Database configuration structs
│   ├── database_connection.go     # Database connection interfaces
│   ├── connection_pool.go         # Connection pool implementation
│   ├── email_config.go            # Email service configuration
│   ├── engine.go                  # Environment loading & HTTP server
│   ├── environment_config.go      # Main environment config container
│   ├── interfaces.go              # Configuration interfaces
│   └── logging_config.go          # Logging configuration & setup
├── 📁 Controller/                 # HTTP request handlers
│   └── 📁 HealthCheckController/  # Health check endpoints
│       └── HealthCheckController.go
├── 📁 Environment/                # Environment-specific configurations
│   ├── Development.yml            # Development environment config
│   └── Production.yml             # Production environment config
├── 📁 Libraries/                  # Shared libraries and utilities
│   └── 📁 Helpers/               # Helper functions
│       ├── log.go                # Logging helper functions
│       └── response.go           # HTTP response helpers
├── 📁 Logs/                      # Application logs (gitignored)
│   ├── app.access-YYYY-MM-DD.log  # Access logs
│   ├── app.error-YYYY-MM-DD.log   # Error logs
│   └── app.error-loki-YYYY-MM-DD.log # Structured JSON logs
├── 📁 Middlewares/               # Gin middleware
│   └── log_middleware.go         # Logging middleware integration
├── 📁 Repositories/              # Data access layer
├── 📁 Routes/                    # Route definitions
├── 📁 Services/                  # Business logic layer
│   ├── 📁 Emails/               # Email service implementation
│   │   ├── email.go
│   │   └── email_model.go
│   └── services.go              # Service initialization
├── .gitignore                   # Git ignore file
├── go.mod                       # Go module definition
├── go.sum                       # Go module checksums
├── main.go                      # Application entry point
└── README.md                    # This file
```

---

## 🚀 Quick Start

### Prerequisites

- **Go 1.25+** - [Download Go](https://golang.org/dl/)
- **PostgreSQL** - [Download PostgreSQL](https://www.postgresql.org/download/)
- **Redis** - [Download Redis](https://redis.io/download/)

### Installation

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd boilerplate-go
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Configure environment:**
   ```bash
   # Copy and modify the development configuration
   cp Environment/Development.yml Environment/Development.local.yml
   # Edit the configuration file with your database credentials
   ```

4. **Set up databases:**
   ```bash
   # PostgreSQL
   createdb boilerplate
   
   # Redis (start the server)
   redis-server
   ```

### Running the Application

#### Development Mode
```bash
go run main.go -env=development
```

#### Production Mode
```bash
go build -o app .
./app -env=production
```

#### Available Environments
- `development` - Local development with debug logging
- `production` - Production with optimized logging
- `staging` - Staging environment
- `test` - Testing environment

---

## ⚙️ Configuration

Configuration is managed through YAML files in the `Environment/` directory.

### Development Configuration (`Environment/Development.yml`)

```yaml
app:
  name: "Boilerplate"
  debug: true
  port: "3000"
  host: "localhost"
  service: "http"

logging:
  service_name: "boilerplate-go-dev"
  log_path: "./Logs/app"
  enable_stdout: true
  enable_file: true
  enable_loki: false

databases:
  - name: "boilerplate"
    engine: "postgres"
    username: "root"
    password: "password"
    port: "5432"
    host: "localhost"
    maximum_connection: 20
    maximum_idle_time: 30s
    usage: "Main database"
    connection: "mainDB"
    
  - name: ""
    engine: "redis"
    username: ""
    password: ""
    port: "6379"
    host: "localhost"
    maximum_connection: 20
    usage: "Cache database"
    connection: "mainDBRedis"

api:
  read_timeout: 30
  write_timeout: 30
  idle_timeout: 30

email:
  host: "smtp.example.com"
  port: 587
  user: "your-email@example.com"
  password: "your-password"
  email_from: "noreply@example.com"
```

---

## 📊 Logging System

This project uses [go-logging-lib](https://github.com/ahmadsaubani/go-logging-lib) for comprehensive logging with:

- **📅 Daily Log Rotation** - Automatic dated log files
- **🎯 Multi-format Output** - Console, file, and JSON formats
- **🚀 Gin Integration** - Native middleware with anti-duplication
- **📡 Structured Logging** - Request metadata injection

### Log Files

```
Logs/
├── app.access-2024-01-31.log      # Access logs (HTTP requests)
├── app.error-2024-01-31.log       # Error logs (human-readable)
└── app.error-loki-2024-01-31.log  # JSON logs (for monitoring)
```

### Logging in Code

```go
// Using context-aware logging
func ExampleHandler(c *gin.Context) {
    // Manual error logging with anti-duplication
    err := someService.DoSomething()
    if err != nil {
        // LogErrorWithMark logs to error log, loki, and marks as logged
        Config.AppLogger.LogErrorWithMark(c, err)
        c.JSON(500, gin.H{"error": "Something went wrong"})
        return
    }
    
    // Success response
    c.JSON(200, gin.H{"message": "Success"})
}
```

---

## 🛣️ API Endpoints

### Health Check

| Method | Endpoint | Description | Response |
|--------|----------|-------------|----------|
| `GET` | `/ping` | Health check with email test | `200 OK` |

Example response:
```json
{
  "status": "OK",
  "data": {
    "message": "pong"
  },
  "code": 200,
  "accessTime": "31-01-2024 15:45:30"
}
```

---

## 🔧 Development

### Adding New Controllers

1. **Create controller file:**
   ```go
   // Controller/ExampleController/ExampleController.go
   package ExampleController
   
   import "github.com/gin-gonic/gin"
   
   type ExampleInterface interface {
       GetExample(c *gin.Context)
   }
   
   type controller struct{}
   
   func NewController() ExampleInterface {
       return &controller{}
   }
   
   func (ctrl *controller) GetExample(c *gin.Context) {
       c.JSON(200, gin.H{"message": "example"})
   }
   ```

2. **Register routes:**
   ```go
   // In Routes package
   router.GET("/example", exampleController.GetExample)
   ```

### Database Operations

```go
// Get database connection
db := connection.PostgreMainCon()

// Get Redis connection  
redis := connection.RedisMainCon(ctx)
```

### Error Handling

```go
// Use helpers for consistent error logging
if err != nil {
    Helpers.LogError(ctx, err)
    Helpers.HttpResponseError(c, "Error message", http.StatusInternalServerError)
    return
}
```

---

## 📦 Dependencies

### Core Dependencies
- **[Gin](https://github.com/gin-gonic/gin)** `v1.11.0` - HTTP web framework
- **[go-logging-lib](https://github.com/ahmadsaubani/go-logging-lib)** `v1.0.1` - Structured logging
- **[go-yaml](https://github.com/goccy/go-yaml)** `v1.19.2` - YAML configuration
- **[PostgreSQL Driver](https://github.com/lib/pq)** `v1.11.1` - Database driver
- **[Redis](https://github.com/redis/go-redis)** `v9.17.3` - Redis client
- **[UUID](https://github.com/google/uuid)** `v1.6.0` - UUID generation
- **[Gomail](https://gopkg.in/gomail.v2)** - Email functionality

### Development Dependencies
- **[CORS](https://github.com/gin-contrib/cors)** `v1.7.6` - CORS middleware

---

## 🚦 Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `APP_ENV` | Application environment | `development` | No |

---

### Code Style Guidelines

- Follow Go conventions and `gofmt`
- Use meaningful variable and function names
- Add comments for exported functions
- Write tests for new functionality
- Use the structured logging system

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🆘 Support

- 📚 [Go Documentation](https://golang.org/doc/)
- 🌐 [Gin Framework Guide](https://gin-gonic.com/docs/)
- 📊 [Logging Library](https://github.com/ahmadsaubani/go-logging-lib)
- 🐛 [Report Issues](https://github.com/your-username/boilerplate-go/issues)

---

**Made with ❤️ for Go developers**

⭐ **If this boilerplate helps your project, please give it a star!** ⭐