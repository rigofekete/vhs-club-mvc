# VHS Club MVC

<p align="center">
  <img src="docs/images/logo.png" alt="VHS Club Logo" width="200"/>
</p>

<p align="center">
  <strong>A retro-styled VHS movie rental service</strong>
</p>

<p align="center">
  <a href="#features">Features</a> •
  <a href="#tech-stack">Tech Stack</a> •
  <a href="#demo">Demo</a> •
  <a href="#installation">Installation</a> •
  <a href="#ci-cd">CI/CD</a>
</p>

---

## Overview

VHS Club is a full-stack web application for a fictional VHS movie rental service. Users can browse a catalog of classic VHS movies and rent their favorite titles. The application comes with pre-created users and an admin account loaded from SQL seed scripts.

> **Note:** User accounts are created via SQL scripts and cannot be registered through the web interface. The seed script provides an admin account and sample users for testing.

## Features

- **User Authentication**: Secure login with JWT-based authentication
- **VHS Catalog**: Browse a collection of VHS movies with details (title, director, genre, price)
- **Rental System**: Rent available tapes and track rental history
- **Role-Based Access**: Admin and user roles with different permissions
- **Responsive Design**: Modern UI that works across devices
- **RESTful API**: Clean API design following REST principles
- **Database Seeding**: SQL scripts provide initial data including admin account, sample users, and VHS tapes

## Tech Stack

### Backend

| Technology | Description |
|------------|-------------|
| **Go** | Backend language (v1.25) |
| **Gin Gonic** | High-performance HTTP web framework |
| **PostgreSQL** | Relational database for data persistence |
| **JWT** | Authentication tokens |
| **Argon2id** | Secure password hashing |
| **SQLC** | Type-safe SQL code generator |

### Frontend

| Technology | Description |
|------------|-------------|
| **React** | UI library (v19) |
| **Vite** | Build tool and dev server |
| **React Router** | Client-side routing |
| **Vitest** | Unit testing framework |
| **Nginx** | Production web server |

### DevOps & Tools

| Technology | Description |
|------------|-------------|
| **Docker** | Containerization |
| **Docker Compose** | Multi-container orchestration |
| **GitHub Actions** | CI/CD pipeline |
| **Nginx** | Reverse proxy and static file serving |

## Demo

### Screenshots

<p align="center">
  <img src="docs/images/screenshot-home.png" alt="Home Page" width="400"/>
  <img src="docs/images/screenshot-catalog.png" alt="Catalog Page" width="400"/>
</p>

<p align="center">
  <img src="docs/images/screenshot-login.png" alt="Login Page" width="400"/>
  <img src="docs/images/screenshot-rental.png" alt="Rental Page" width="400"/>
</p>

### Video Demo

<p align="center">
  <a href="docs/videos/demo.mp4">
    <img src="docs/images/video-thumbnail.png" alt="Watch Demo Video" width="600"/>
  </a>
</p>

## Installation

### Prerequisites

- **Docker** (v20.10+) and **Docker Compose** (v2.0+) - for containerized setup
- **Go** (v1.25+) - for local backend development
- **Node.js** (v24+) and **npm** - for local frontend development
- **PostgreSQL** (v16+) - for local database

### Option 1: Docker Compose (Recommended)

The easiest way to run the entire application stack is using Docker Compose. This will spin up all three services (database, backend, and frontend) with a single command.

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/vhs-club.git
   cd vhs-club
   ```

2. **Set up environment variables:**

   Create a `.env` file in the project root (or use the existing one):

   ```bash
   # Generate a secure JWT secret (optional - one is provided in the example .env)
   # JWT_SECRET=$(openssl rand -base64 32)
   ```

   The default `.env` file contains:

   ```
   DB_URL=postgres://postgres:postgres@localhost:5432/vhs_club?sslmode=disable
   JWT_SECRET=your-jwt-secret-here
   ```

3. **Build and run with Docker Compose:**

   ```bash
   docker compose up -d --build
   ```

   This command will:
   - Build the database image with initial schema and seed data (includes admin account and sample users)
   - Build the Go backend API
   - Build the React frontend with Nginx
   - Start all services with proper networking

   > **Important:** The database seeding happens automatically on first startup. The `Dockerfile.db` copies `sql/seed.sql` into the PostgreSQL initialization directory, which PostgreSQL executes when the container starts for the first time.

4. **Access the application:**

   | Service | URL | Description |
   |---------|-----|-------------|
   | Frontend | <http://localhost> | Main web application |
   | Backend API | <http://localhost:8080> | REST API endpoints |
   | Health Check | <http://localhost:8080/health> | API health status |

5. **Stop the services:**

   ```bash
   docker compose down
   ```

   To remove all data (including the database volume):

   ```bash
   docker compose down -v
   ```

### Option 2: Local Development Setup (Without Docker)

If you prefer to run the services directly on your machine for development purposes, follow these steps.

#### Step 1: Database Setup

1. **Install PostgreSQL** (v16+):

   On Arch Linux:

   ```bash
   sudo pacman -S postgresql
   sudo systemctl enable postgresql
   sudo systemctl start postgresql
   ```

   > **Note:** These instructions have been tested on Arch Linux. For other distributions, consult your package manager's documentation for PostgreSQL installation.

2. **Run the seed script:**

   The `sql/seed.sql` file handles everything: it creates the database, schema, tables, and populates them with initial data (admin account, sample users, and VHS tapes).

   ```bash
   # Run as postgres superuser - this creates the database and everything else
   sudo -u postgres psql -f sql/seed.sql
   ```

   > **Important:** The seed script creates the `vhs_club` database, all tables, and initial data in one step. After seeding, you can log in with the default credentials listed in the [Database Seeding](#database-seeding) section.

   > **Note:** The `sql/seed_dev.sql` file is not needed for local development - all essential data is in `sql/seed.sql` which is the same file used by Docker Compose.

#### Step 2: Backend Setup

1. **Install Go** (v1.25+):

   - Download from [https://go.dev/dl/](https://go.dev/dl/)
   - Or use your package manager

2. **Install dependencies:**

   ```bash
   go mod download
   ```

3. **Configure the database connection:**

   Create a `.env` file in the project root with your PostgreSQL credentials:

   ```bash
   # Replace with your actual PostgreSQL credentials
   DB_URL=postgres://your_username:your_password@localhost:5432/vhs_club?sslmode=disable
   JWT_SECRET=your-secret-key-here
   ```

   For example, if your PostgreSQL user is `postgres` with password `postgres`:

   ```bash
   DB_URL=postgres://postgres:postgres@localhost:5432/vhs_club?sslmode=disable
   JWT_SECRET=your-secret-key-here
   ```

   **Finding your PostgreSQL credentials:**

   - If you installed PostgreSQL locally, the default superuser is usually `postgres` with a password you set during installation
   - Check your PostgreSQL configuration in `pg_hba.conf` for authentication methods
   - If you seeded using `sudo -u postgres psql -f sql/seed.sql`, the database was created by the `postgres` superuser

   Or export them directly in your terminal:

   ```bash
   export DB_URL="postgres://your_username:your_password@localhost:5432/vhs_club?sslmode=disable"
   export JWT_SECRET="your-secret-key-here"
   ```

4. **Run the backend:**

   ```bash
   go run main.go
   ```

   The API will be available at `http://localhost:8080`

5. **Run tests:**

   ```bash
   go test ./...
   ```

#### Step 3: Frontend Setup

1. **Install Node.js** (v24+):

   - Download from [https://nodejs.org/](https://nodejs.org/)
   - Or use nvm: `nvm install 24`

2. **Navigate to the frontend directory:**

   ```bash
   cd react-app
   ```

3. **Install dependencies:**

   ```bash
   npm install
   ```

4. **Start the development server:**

   ```bash
   npm run dev
   ```

   The frontend will be available at `http://localhost:5173` (default Vite port)

5. **Build for production:**

   ```bash
   npm run build
   ```

6. **Run tests:**

   ```bash
   npx vitest run
   ```

### Development URLs Summary

| Service | Local Development | Docker |
|---------|-------------------|--------|
| Frontend | <http://localhost:5173> | <http://localhost> |
| Backend API | <http://localhost:8080> | <http://localhost:8080> |
| Database | localhost:5432 | localhost:5432 |

## API Documentation

### Authentication Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/users/login` | Authenticate and receive JWT token |

> **Note:** User accounts are pre-created via SQL seed scripts. There is no public registration endpoint. See the [Database Seeding](#database-seeding) section for default credentials.

### Tapes Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/tapes` | List all available tapes (public) |
| GET | `/api/tapes/:id` | Get a specific tape by ID (public) |
| POST | `/api/tapes` | Create a new tape (admin only) |
| POST | `/api/tapes/batch` | Create multiple tapes (admin only) |
| PATCH | `/api/tapes/:id` | Update a tape (admin only) |
| DELETE | `/api/tapes/:id` | Delete a tape (admin only) |
| DELETE | `/api/tapes` | Delete all tapes (admin only) |

### Rentals Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/rentals` | List all active rentals (public) |
| POST | `/api/rentals/:id` | Create a new rental (authenticated users) |
| PATCH | `/api/rentals/:id` | Return a rented tape (authenticated users) |
| DELETE | `/api/rentals` | Delete all rentals (admin only) |

### User Management Endpoints (Admin)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/users` | Create a new user (admin only) |
| POST | `/api/users/batch` | Create multiple users (admin only) |
| GET | `/api/users` | List all users (admin only) |
| GET | `/api/users/:id` | Get user by ID (admin only) |
| DELETE | `/api/users` | Delete all users (admin only) |

## Database Seeding

The application does not have a public user registration endpoint. Instead, users are pre-created via SQL seed scripts. The **same `sql/seed.sql` file** is used for both Docker and local development, ensuring consistency across environments. The database is automatically populated with sample data when using Docker Compose, or you can manually apply the seed script for local development.

### Default Credentials

The following accounts are available after seeding:

| Username | Email | Password | Role |
|----------|-------|----------|------|
| **Admin** | `admin@vhs-club.hu` | `12345678` | admin |
| **ArthurCClarke** | `thesentinel@space.odissey` | `12345678` | user |
| **MilesDavis** | `grumpy.genius@cool.com` | `12345678` | user |

### Sample Tapes

The seed script also populates the catalog with classic VHS movies:

| Title | Director | Genre | Quantity | Price |
|-------|----------|-------|----------|-------|
| Amarcord | Federico Fellini | Drama | 1 | 5999.99 |
| Taxi Driver | Martin Scorsese | Thriller | 2 | 5999.99 |
| Back to the Future | Robert Zemeckis | Adventure | 5 | 2999.99 |
| Alien | Ridley Scott | Horror | 10 | 5999.99 |
| A torinói ló | Béla Tarr | Drama | 3 | 5999.99 |
| Batman | Tim Burton | Action | 4 | 2999.99 |
| Fitzcarraldo | Werner Herzog | Drama | 11 | 5999.99 |

### Seeding with Docker Compose

When using Docker Compose, the database is automatically seeded on first startup using the same `sql/seed.sql` file used for local development. The `Dockerfile.db` copies the seed script into the PostgreSQL initialization directory, which PostgreSQL executes when the container starts for the first time.

```bash
# Start all services - seeding happens automatically
docker compose up -d --build
```

> **Note:** The database is only seeded on the **first** container startup. If you need to re-seed, you must remove the volume: `docker compose down -v` and then start again.

### Seeding Manually (Local Development)

For local development without Docker, run the **same seed script** used by Docker Compose:

```bash
# This creates the database, schema, tables, and all initial data in one step
sudo -u postgres psql -f sql/seed.sql
```

The `sql/seed.sql` file handles everything:

- Creates the `vhs_club` database
- Creates all tables (users, tapes, rentals)
- Inserts the admin account and sample users
- Inserts the VHS tape catalog

> **Note:** The seed script uses `CREATE DATABASE` and `\c` (connect) commands, so it must be run as a superuser (e.g., `postgres` user) and not as a regular database user.

### Additional Development Data (Optional)

If you have extra development data in `sql/seed_dev.sql`, you can apply it after the main seed:

```bash
psql -U postgres -d vhs_club -f sql/seed_dev.sql
```

### Customizing Seed Data

You can modify the seed file to add your own initial data:

- **`sql/seed.sql`** - Single file containing database creation, schema, tables, and all initial data

To add new users, add INSERT statements to the users section of `sql/seed.sql`:

```sql
INSERT INTO users (username, email, role, hashed_password) VALUES
  (
    'YourUser', 'email@example.com', 'user',
    '$argon2id$v=19$m=65536,t=1,p=24$...'
  );
```

> **Note:** Passwords must be hashed using Argon2id. You can generate hashed passwords using the Go application or an Argon2id tool.

### About `sql/seed_dev.sql`

The `sql/seed_dev.sql` file contains additional development/test data and uses `ON CONFLICT DO NOTHING` to avoid duplicate errors. It is optional and can be applied after `sql/seed.sql` if you need extra test data. This file is not used by Docker Compose.

## Project Structure

```
vhs-club-mvc/
├── .github/
│   └── workflows/
│       └── ci.yml              # GitHub Actions CI pipeline
├── config/
│   └── config.go             # Application configuration
├── handler/                    # HTTP request handlers (Controllers)
│   ├── user_handler.go
│   ├── tape_handler.go
│   ├── rental_handler.go
│   └── *_types.go            # Request/response DTOs
├── internal/
│   ├── apperror/             # Application error handling
│   ├── auth/                 # Authentication utilities
│   └── database/             # SQLC generated code
├── middleware/                 # Gin middleware
│   ├── auth.go
│   └── cors.go
├── model/                    # Domain models
│   ├── user.go
│   ├── tape.go
│   └── rental.go
├── repository/                 # Data access layer
│   ├── user_repository.go
│   ├── tape_repository.go
│   ├── rental_repository.go
│   └── repo_helpers.go
├── service/                    # Business logic layer
│   ├── user_service.go
│   ├── tape_service.go
│   ├── rental_service.go
│   └── *_test.go             # Unit tests
├── sql/
│   ├── schema/               # Database migrations (used by sqlc, not for seeding)
│   │   ├── 001_tapes.sql
│   │   ├── 002_users.sql
│   │   └── 003_rentals.sql
│   ├── queries/              # SQLC query definitions
│   │   ├── users.sql
│   │   ├── tapes.sql
│   │   └── rentals.sql
│   ├── seed.sql              # Complete database setup (CREATE DB, schema, tables, data)
│   └── seed_dev.sql          # Optional additional dev/test data (not required)
├── react-app/                # Frontend React application
│   ├── src/
│   │   ├── components/       # React components
│   │   ├── pages/            # Page components
│   │   ├── App.jsx
│   │   └── main.jsx
│   ├── public/
│   ├── index.html
│   ├── package.json
│   ├── vite.config.js
│   ├── nginx.conf
│   └── Dockerfile.frontend
├── .env                      # Environment variables
├── .dockerignore
├── docker-compose.yaml       # Docker Compose configuration
├── Dockerfile.backend        # Backend container definition
├── Dockerfile.db             # Database container definition
├── go.mod                    # Go module dependencies
├── go.sum                    # Go module checksums
├── main.go                   # Application entry point
├── sqlc.yaml                 # SQLC configuration
└── README.md                 # This file

```

## CI/CD

### GitHub Actions Pipeline

The project includes a comprehensive CI/CD pipeline defined in `.github/workflows/ci.yml` that runs on every push and pull request to the `master` branch.

### Pipeline Stages

```mermaid
graph LR
    A[Push/PR] --> B[Test Backend]
    A --> C[Test Frontend]
    B --> D[Build Images]
    C --> D
    D --> E[Smoke Test]
```

1. **Test Backend**: Runs all Go unit tests
2. **Test Frontend**: Runs all Vitest tests for the React application
3. **Build**: Builds Docker images for all three services (backend, frontend, database)
4. **Smoke Test**: Deploys the full stack and verifies the application is accessible

## Environment Variables

| Variable | Description | Required | Default |
|----------|-------------|----------|---------|
| `DB_URL` | PostgreSQL connection string | Yes | - |
| `JWT_SECRET` | Secret key for JWT signing | Yes | - |

### Generating a JWT Secret

For production deployments, generate a secure random secret:

```bash
# Generate a 32-byte base64-encoded secret
openssl rand -base64 32
```

### Database Connection String Format

The `DB_URL` follows the PostgreSQL connection string format:

```
postgres://username:password@host:port/database?sslmode=disable
```

**Components:**

- `username` - PostgreSQL username (e.g., `postgres`)
- `password` - PostgreSQL password for that user
- `host` - Database host (use `localhost` for local development, `vhs-db` for Docker)
- `port` - PostgreSQL port (default: `5432`)
- `database` - Database name (use `vhs_club`)
- `sslmode` - SSL mode (use `disable` for local development)

**Example connection strings:**

```bash
# Local development (after running seed.sql with sudo -u postgres)
DB_URL=postgres://postgres:postgres@localhost:5432/vhs_club?sslmode=disable

# Docker Compose (backend container connecting to db container)
DB_URL=postgres://postgres:postgres@vhs-db:5432/vhs_club?sslmode=disable

# With a custom user you created
DB_URL=postgres://myuser:mypassword@localhost:5432/vhs_club?sslmode=disable
```

## Development

### Project Architecture

The backend follows a layered architecture pattern:

```
┌─────────────────┐
│   Handler       │  ← HTTP request handlers (Controllers)
│   (Gin)         │
├─────────────────┤
│   Service       │  ← Business logic
│                 │
├─────────────────┤
│   Repository    │  ← Data access layer
│                 │
├─────────────────┤
│   Database      │  ← PostgreSQL
│   (PostgreSQL)  │
└─────────────────┘
```

### Backend Development

```bash
# Run the server with hot reload (requires air or similar)
go run main.go

# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Generate SQLC code (after modifying SQL queries)
sqlc generate
```

### Frontend Development

```bash
cd react-app

# Start development server
npm run dev

# Run tests
npx vitest run

# Run tests in watch mode
npx vitest

# Build for production
npm run build

# Preview production build
npm run preview
```

## Testing

### Unit Tests

The project includes comprehensive unit tests for the service layer:

```bash
# Run all Go tests with verbose output
go test -v ./...

# Run tests with coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run specific test file
go test -v ./service/rental_service_test.go ./service/rental_service.go
```

### Frontend Tests

```bash
cd react-app

# Run Vitest tests
npx vitest run

# Run with coverage
npx vitest run --coverage

# Run in UI mode
npx vitest --ui
```

## Requirements & Dependencies

### System Requirements

- **OS**: Linux (tested on Arch Linux)
- **Memory**: Minimum 4GB RAM (8GB recommended for Docker)
- **Disk**: 2GB free space

### Backend Dependencies (Go)

```go
// Core Framework
github.com/gin-gonic/gin v1.11.0

// Authentication
github.com/golang-jwt/jwt/v5 v5.3.1
github.com/alexedwards/argon2id v1.0.0

// Database
github.com/lib/pq v1.11.1

// Validation
github.com/go-playground/validator/v10 v10.27.0

// Utilities
github.com/google/uuid v1.6.0
github.com/joho/godotenv v1.5.1
```

### Frontend Dependencies (React)

```json
{
  "dependencies": {
    "react": "^19.2.0",
    "react-dom": "^19.2.0",
    "react-router-dom": "^7.13.1"
  },
  "devDependencies": {
    "vite": "^7.3.1",
    "vitest": "^4.1.1",
    "@testing-library/react": "^16.3.2",
    "eslint": "^9.39.1"
  }
}
```

### Docker Services

| Service | Image | Port | Purpose |
|---------|-------|------|---------|
| Database | postgres:16-alpine | 5432 | PostgreSQL database |
| Backend | golang:1.25-alpine | 8080 | Go REST API |
| Frontend | nginx:stable-alpine | 80 | React SPA with Nginx |

## CI/CD

The project uses **GitHub Actions** for continuous integration and deployment. The pipeline is defined in `.github/workflows/ci.yml`.

### Pipeline Workflow

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│   Push/PR       │────▶│  Test Backend   │────▶│  Test Frontend  │────▶│  Build Images   │
│   to master     │     │  (Go tests)     │     │  (Vitest)       │     │  (Docker)       │
└─────────────────┘     └─────────────────┘     └─────────────────┘     └────────┬────────┘
                                                                                │
                                                                                ▼
                                                                       ┌─────────────────┐
                                                                       │  Smoke Test     │
                                                                       │  (Integration)  │
                                                                       └─────────────────┘
```

### Jobs Description

| Job | Description | Technologies |
|-----|-------------|--------------|
| `test-backend` | Runs Go unit tests across all packages | Go 1.25, testify |
| `test-frontend` | Runs React component tests with Vitest | Node 24, Vitest, Testing Library |
| `build` | Builds Docker images for all services | Docker |
| `smoke-test` | Deploys the full stack and performs health checks | Docker Compose, curl |

### CI Configuration

```yaml
# .github/workflows/ci.yml
name: CI

on:
  push:
    branches: [master]
  pull_request:
    branches: [master]

jobs:
  test-backend:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.25'
      - name: Run Go tests
        run: go test ./...

  test-frontend:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: '24'
      - name: Install dependencies
        working-directory: ./react-app
        run: npm ci
      - name: Run Vitest
        working-directory: ./react-app
        run: npx vitest run

  build:
    runs-on: ubuntu-latest
    needs: [test-backend, test-frontend]
    steps:
      - uses: actions/checkout@v4
      - name: Build Docker images
        run: |
          docker build -f Dockerfile.backend -t vhs-backend .
          docker build -f react-app/Dockerfile.frontend -t vhs-frontend ./react-app
          docker build -f Dockerfile.db -t vhs-db .

  smoke-test:
    runs-on: ubuntu-latest
    needs: [build]
    env:
      JWT_SECRET: ci-test-secret
    steps:
      - uses: actions/checkout@v4
      - name: Start all services
        run: docker compose up -d --build
      - name: Wait for frontend
        run: |
          for i in $(seq 1 30); do
            curl -s http://localhost:80 && exit 0
            sleep 1
          done
          exit 1
      - name: Curl front page
        run: |
          STATUS=$(curl -o /dev/null -s -w "%{http_code}" http://localhost:80)
          [ "$STATUS" = "200" ] || exit 1
      - name: Tear down
        if: always()
        run: docker compose down
```

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

Please ensure:

- All tests pass before submitting
- Code follows the existing style
- New features include appropriate tests
- Documentation is updated if needed

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin) - Fast and lightweight web framework for Go
- [SQLC](https://sqlc.dev/) - Type-safe SQL code generator
- [Vite](https://vitejs.dev/) - Next generation frontend tooling
- [React](https://react.dev/) - A JavaScript library for building user interfaces
- [PostgreSQL](https://www.postgresql.org/) - The world's most advanced open source relational database

---

<p align="center">
  Made with ❤️ for VHS enthusiasts
</p>
