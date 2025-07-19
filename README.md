# My Wallet - Investment Tracker

This is a simple web application designed to keep track of personal investments.

## Tech Stack

* **Backend:** Go
* **Database:** PostgreSQL
* **Frontend:** SvelteKit

## Project Structure

The project is organized as a monorepo with two main directories:

* `/backend`: Contains the Go API and business logic.
* `/frontend`: Contains the SvelteKit user interface.

## Running the Application (with Docker)

This project uses Docker Compose to manage the application stack. Three separate configuration files are provided for different development scenarios.

### Accessing the Database GUI (pgAdmin)

When you run any of the Docker Compose configurations, a pgAdmin service will also be started, providing a web-based GUI for the database.

* **URL:** `http://localhost:5050`
* **Login Email:** `admin@example.com`
* **Login Password:** `admin`

**First-Time Setup:**
When you log in for the first time, you will need to add a server connection to the `fire` database:

1. Right-click on "Servers" and select "Register" -> "Server...".
2. In the "General" tab, give it a name (e.g., `my-wallet-db`).
3. In the "Connection" tab, use the following details:
    * **Host name/address:** `db` (This is the service name from Docker Compose)
    * **Port:** `5432`
    * **Maintenance database:** `fire`
    * **Username:** `user`
    * **Password:** `password`
4. Click "Save".

This project uses Docker Compose to manage the application stack. Three separate configuration files are provided for different development scenarios.

### Backend Only (Backend + Database)

Useful for backend development and testing.

```bash
docker-compose -f docker-compose.backend.yml up --build
```

### Frontend and Database

Useful for frontend development. This starts the frontend container and the database. You would typically run the backend on your local machine, which would then connect to the database running in Docker.

```bash
docker-compose -f docker-compose.frontend.yml up --build
```
