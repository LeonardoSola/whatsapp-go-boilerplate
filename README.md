# WhatsApp Go Boilerplate

A clean and extensible WhatsApp bot boilerplate built with Go, using the [whatsmeow](https://github.com/tulir/whatsmeow) library. This project provides a solid foundation for building WhatsApp bots with a well-structured architecture.

## 🚀 Features

- **Clean Architecture**: Follows clean architecture principles with clear separation of concerns
- **WhatsApp Integration**: Built on top of whatsmeow for reliable WhatsApp Web API integration
- **Command System**: Simple command-based message handling system
- **Admin Controls**: Admin-only command execution for security
- **Database Support**: Support any SQL database, like PostgreSQL, MySQL, etc.
- **Event-Driven**: Event-based message handling with extensible router
- **QR Code Authentication**: Terminal-based QR code scanning for WhatsApp login
- **Graceful Shutdown**: Proper signal handling for clean application termination

## 📋 Prerequisites

- Go 1.23.1 or higher
- WhatsApp account for bot authentication
- SQLite (Optional)

## 🛠️ Installation

### Option 1: Local Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/LeonardoSola/whatsapp-go-boilerplate.git
   cd whatsapp-go-boilerplate
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   export ADMIN_PHONE_NUMBER="5516999999999"
   export SQLITE_DB="file:store.db?_foreign_keys=on"
   ```

4. **Run the application**
   ```bash
   go run main.go
   ```

### Option 2: Docker Installation *(Recommended)*

1. **Clone the repository**
   ```bash
   git clone https://github.com/LeonardoSola/whatsapp-go-boilerplate.git
   cd whatsapp-go-boilerplate
   ```

2. **Set up environment variables**
   ```bash
   export ADMIN_PHONE_NUMBER="your-admin-phone-number@c.us"
   ```

3. **Build and run with Docker**
   ```bash
   # Using Makefile (recommended)
   make build-and-run
   
   # Or using docker-compose directly
   docker-compose up -d
   ```

4. **View logs**
   ```bash
   make logs
   # or
   docker-compose logs -f
   ```

## 🔐 First Time Setup

When you run the application for the first time:

1. A QR code will be displayed in your terminal
2. Open WhatsApp on your phone
3. Go to Settings > Linked Devices > Link a Device
4. Scan the QR code displayed in your terminal
5. The bot will connect to WhatsApp and start listening for messages

**For Docker users**: The QR code will appear in the Docker logs. Use `make logs` or `docker-compose logs -f` to view it.

## 📖 Usage

### Available Commands

Currently, the bot supports the following commands (admin-only):

- `PING` - Responds with "Pong" (useful for testing connectivity)

### Adding New Commands

To add new commands, follow these steps:

1. **Create a new business function** in `app/business/`:
   ```go
   func YourCommand(platform models.Platform, message models.Message) {
       platform.Reply(message, "Your response here")
   }
   ```

2. **Register the command** in `app/handlers/message.go`:
   ```go
   var commands = map[string]func(platform models.Platform, message models.Message){
       "PING": business.Ping,
       "YOUR_COMMAND": business.YourCommand, // Add your new command here
   }
   ```

### Message Handling

The bot processes messages through the following flow:

1. **Message Reception**: WhatsApp messages are received via whatsmeow
2. **Event Routing**: Messages are routed through the event handler
3. **Admin Check**: Only messages from the admin phone number are processed
4. **Command Execution**: Commands are executed based on the message text
5. **Response**: Responses are sent back to the user

## 🏗️ Project Structure

```
whatsapp-go-boilerplate/
├── adapters/                 # External adapters and infrastructure
│   ├── cmd/whastapp/        # WhatsApp client implementation
│   │   ├── client.go        # WhatsApp client setup
│   │   ├── login.go         # Authentication and QR code handling
│   │   ├── platform.go      # Platform-specific implementations
│   │   ├── storage.go       # Device storage management
│   │   └── events/          # Event handlers
│   │       ├── message.go   # Message event processing
│   │       └── router.go    # Event routing
│   ├── database/sqlite/     # Database adapters
│   │   └── config.go        # SQLite configuration
│   └── env/                 # Environment configuration
│       ├── config.go        # General configuration
│       └── database.go      # Database environment variables
├── app/                     # Application business logic
│   ├── business/            # Business logic implementations
│   │   └── basic.go         # Basic command implementations
│   ├── constants/           # Application constants
│   │   ├── type.go          # Type definitions
│   │   └── whatsapp.go      # WhatsApp-specific constants
│   ├── handlers/            # Message handlers
│   │   └── message.go       # Main message processing logic
│   └── models/              # Data models
│       ├── message.go       # Message model
│       └── platform.go      # Platform interface
├── main.go                  # Application entry point
├── go.mod                   # Go module file
└── go.sum                   # Go module checksums
```

## 🔧 Configuration

### Environment Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `ADMIN_PHONE_NUMBER` | Admin phone number for command execution | `5511999999999@c.us` |
| `SQLITE_DB` | SQLite database file path | `whatsapp_bot.db` |

### Database

The project uses SQLite for data persistence. The database file is automatically created when the application starts.

### Docker Configuration

The project includes comprehensive Docker support with the following files:

- `Dockerfile` - Production-ready multi-stage build
- `Dockerfile.dev` - Development environment with hot reloading
- `docker-compose.yml` - Production container orchestration
- `docker-compose.dev.yml` - Development overrides
- `Makefile` - Simplified Docker operations
- `.air.toml` - Hot reloading configuration for development

#### Docker Commands

```bash
# Build the image
make build

# Run in production mode
make run

# Run in development mode (with hot reloading)
make dev

# Stop the container
make stop

# View logs
make logs

# Access container shell
make shell

# Clean up everything
make clean
```

## 🚨 Security Considerations

- **Admin-Only Commands**: All commands are restricted to the admin phone number
- **Environment Variables**: Sensitive configuration is stored in environment variables
- **Device Storage**: WhatsApp session data is stored locally for persistence
- **Docker Security**: Container runs as non-root user with minimal permissions
- **Volume Persistence**: WhatsApp session data is persisted in Docker volumes

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [whatsmeow](https://github.com/tulir/whatsmeow) - WhatsApp Web API library
- [Go](https://golang.org/) - Programming language
- [SQLite](https://www.sqlite.org/) - Database engine

## 📞 Support

If you encounter any issues or have questions, please open an issue on GitHub.

---

**Note**: This is a boilerplate project. Make sure to customize it according to your specific needs and add proper error handling, logging, and security measures for production use.
