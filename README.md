# LibrePM

This project is a web application designed for project management.

## Technology Stack

*   **Backend:** Go
*   **Frontend:** TypeScript, Vite
*   **Database:** MySQL
*   **Development Environment:** Docker, Nx
*   **Hot Reloading:** Air (for Go backend)

## Development Setup

To set up and run the project in development mode, follow these steps:

1.  **Clone the repository** (if you haven't already).
2.  **Ensure Docker is installed and running** on your system.
3.  **Navigate to the project's root directory** in your terminal.
4.  **Start the development environment** by running the following command:
    ```bash
    make dev
    ```
    This command will use Docker Compose to set up and launch all necessary services, including the backend API, frontend, and database.

5.  **Access the application** typically via `http://localhost:8081` (the `open` command in the Makefile might do this automatically).
6.  **(Optional) Seed the database** with initial data:
    ```bash
    make seed
    ```
    This command uses the `seed.yaml` file to populate the database.

7.  **(Advanced) Access the project's CLI environment:**
    ```bash
    make cli
    ```
    This command starts a CLI tool in a detached Docker container, allowing you to run various utility commands.

8.  **(Advanced) Purge all data from the database:**
    ```bash
    make purge
    ```
    This command will execute a purge operation, which is useful for resetting the database to an empty state. **Use with caution!**

For detailed commands or specific services, you can explore the `Makefile` in the project's root directory.

## Production Installation

The production image runs the Go API and the compiled Vue frontend in a single container behind nginx.

### Pull the image

```bash
docker pull ghcr.io/librepm-project/librepm:latest
```

### Example docker-compose.production.yaml

```yaml
services:
  app:
    image: ghcr.io/librepm-project/librepm:latest
    restart: unless-stopped
    ports:
      - "80:80"
    environment:
      DB_HOST: database
      DB_PORT: 3306
      DB_USER: librepm
      DB_PASSWORD: changeme
      DB_NAME: librepm
      JWT_SECRET: changeme-use-a-long-random-string
    volumes:
      - uploads:/app/uploads
    depends_on:
      database:
        condition: service_healthy

  database:
    image: mariadb:11.5
    restart: unless-stopped
    environment:
      MYSQL_DATABASE: librepm
      MYSQL_USER: librepm
      MYSQL_PASSWORD: changeme
      MYSQL_ROOT_PASSWORD: changeme-root
    volumes:
      - db-data:/var/lib/mysql
    healthcheck:
      test: ["CMD", "healthcheck.sh", "--connect", "--innodb_initialized"]
      interval: 10s
      timeout: 5s
      retries: 10

volumes:
  uploads:
  db-data:
```

### First run — seed the database

After the containers are up, seed the initial data:

```bash
docker compose -f docker-compose.production.yaml exec app \
  /usr/local/bin/api-server seed
```

### Environment variables

| Variable | Required | Description |
|----------|----------|-------------|
| `DB_HOST` | yes | MariaDB/MySQL hostname |
| `DB_PORT` | yes | Database port (default: `3306`) |
| `DB_USER` | yes | Database user |
| `DB_PASSWORD` | yes | Database password |
| `DB_NAME` | yes | Database name |
| `JWT_SECRET` | yes | Secret used to sign JWT tokens — use a long random string in production |
| `REGISTER_ALLOWED` | no | Allow new user registration via the `/register` page. Set to `true` to enable, omit or set to `false` to disable (recommended in production) |

### Reverse proxy (optional)

The container exposes port `80`. If you put it behind a reverse proxy (nginx, Caddy, Traefik), point the upstream at the container's port 80. No path prefix stripping is needed — the internal nginx already handles `/api/` routing.
