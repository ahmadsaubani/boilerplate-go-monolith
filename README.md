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
boilerplate-go-monolith/
├── 📁 Config/                     # Configuration management
│   ├── 📁 DTO/                    # Data Transfer Objects
│   │   ├── 📁 ConfigStructs/      # Configuration structs
│   │   │   └── 📁 Alerts/         # Alert notification DTOs
│   │   │       └── alert.go       # LoggingConfig, AlertConfig, platform configs
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
│   └── logging_config.go          # Logging configuration & alert setup
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
│   ├── app.access.log            # Access logs (or app.access-YYYY-MM-DD.log with rotation)
│   ├── app.error.log             # Error logs (human-readable)
│   └── app.loki.log              # Structured JSON logs for Loki/Grafana
├── 📁 Middlewares/               # Gin middleware
│   └── log_middleware.go         # Logging middleware integration
├── 📁 Modules/                   # Feature modules
├── 📁 Repositories/              # Data access layer
├── 📁 Routes/                    # Route definitions
├── 📁 Services/                  # Business logic layer
│   ├── 📁 Emails/               # Email service implementation
│   │   ├── email.go
│   │   └── email_model.go
│   └── services.go              # Service initialization
├── 📁 Infrastructures/          # Infrastructure setup
├── .gitignore                   # Git ignore file
├── docker-compose.yml           # Docker configuration
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
  certificate:
  pem_key:

logging:
  service_name: "boilerplate-go-dev"
  log_path: "./Logs"
  file_prefix: "app"
  enable_stdout: true
  enable_file: true
  enable_loki: true
  enable_rotation: false
  alerts:
    enabled: true
    min_level: "ERROR"
    rate_limit_sec: 300
    discord:
      enabled: false
      webhook_url: "https://discord.com/api/webhooks/YOUR_WEBHOOK_ID/YOUR_WEBHOOK_TOKEN"
      username: "Error Bot"
    slack:
      enabled: false
      webhook_url: "https://hooks.slack.com/services/XXX/YYY/ZZZ"
      channel: "#alerts"
    telegram:
      enabled: false
      bot_token: "123456:ABC-DEF..."
      chat_id: "-1001234567890"
    email:
      enabled: true
      smtp_host: "smtp.gmail.com"
      smtp_port: 587
      username: "alerts@example.com"
      password: "app-password"
      from: "alerts@example.com"
      to:
        - "dev@example.com"
      use_tls: false

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

### Alert Configuration

| Platform | Required Fields | Optional Fields |
|----------|-----------------|-----------------|
| Discord | `webhook_url` | `username`, `avatar_url` |
| Slack | `webhook_url` | `channel`, `username`, `icon_emoji` |
| Telegram | `bot_token`, `chat_id` | - |
| Email | `smtp_host`, `smtp_port`, `from`, `to` | `username`, `password`, `use_tls`, `skip_verify` |

### Alert Levels

| Level | Priority | Trigger Condition |
|-------|----------|-------------------|
| WARN | 1 | Status 300-399 with error |
| ERROR | 2 | Status 400-499 with error |
| CRITICAL | 3 | Status 500+ or explicit critical |

Setting `min_level: "ERROR"` triggers alerts for ERROR and CRITICAL only.

---

## 📊 Logging System

This project uses [go-logging-lib](https://github.com/ahmadsaubani/go-logging-lib) for comprehensive logging with:

- **📅 Daily Log Rotation** - Automatic dated log files (optional)
- **🎯 Multi-format Output** - Console, file, and JSON formats
- **🚀 Gin Integration** - Native middleware with anti-duplication
- **📡 Structured Logging** - Request metadata injection
- **🚨 Alert Notifications** - Send errors to Discord, Slack, Telegram, Email

### Log Files

```
Logs/
├── app.access.log      # Access logs (HTTP requests)
├── app.error.log       # Error logs (human-readable)
└── app.loki.log        # JSON logs (for Loki/Grafana monitoring)
```

With `enable_rotation: true`:
```
Logs/
├── app.access-2026-02-04.log
├── app.error-2026-02-04.log
└── app.loki-2026-02-04.log
```

### Unified Loki JSON Format

All requests are logged in a consistent JSON structure for Grafana visualization:

**Success Response:**
```json
{
    "ts": "2026-02-04T22:13:29+07:00",
    "level": "INFO",
    "service": "boilerplate-go-dev",
    "request_id": "27fd79fe-1e04-47a9-8c56-683269a4c5f0",
    "status_code": 200,
    "latency_ms": 15,
    "http": {
        "ip": "127.0.0.1",
        "method": "GET",
        "path": "/ping",
        "ua": "Mozilla/5.0"
    },
    "errors": null
}
```

**Error Response:**
```json
{
    "ts": "2026-02-04T22:13:29+07:00",
    "level": "CRITICAL",
    "service": "boilerplate-go-dev",
    "request_id": "27fd79fe-1e04-47a9-8c56-683269a4c5f0",
    "status_code": 500,
    "latency_ms": 0,
    "http": {
        "ip": "127.0.0.1",
        "method": "POST",
        "path": "/users",
        "ua": "Mozilla/5.0"
    },
    "errors": {
        "error": "database connection failed",
        "source": {
            "file": "user_handler.go",
            "line": 45
        },
        "stack": [
            "user_handler.go:45 handlers.CreateUser",
            "context.go:192 gin.(*Context).Next"
        ]
    }
}
```

### Logging in Code

```go
// Using context-aware logging
func ExampleHandler(c *gin.Context) {
    // Manual error logging with anti-duplication
    err := someService.DoSomething()
    if err != nil {
        // LogErrorWithMark logs to error log, loki, triggers alert, and marks as logged
        Config.AppLogger.LogErrorWithMark(c, err)
        c.JSON(500, gin.H{"error": "Something went wrong"})
        return
    }
    
    // Success response
    c.JSON(200, gin.H{"message": "Success"})
}
```

### Triggering Alerts

```go
// Option 1: Manual error logging (recommended)
Config.AppLogger.LogErrorWithMark(c, err)

// Option 2: Set error in context (middleware handles alert)
logging.SetLoggedError(c, err)
c.JSON(500, gin.H{"error": "Internal error"})
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