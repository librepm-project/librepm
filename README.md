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

Production installation instructions are coming soon.
