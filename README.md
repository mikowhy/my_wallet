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

## Database Migrations

This project uses `golang-migrate/migrate` to manage database schema changes. Migration files are located in the `/backend/migrations` directory.

The migrations are run automatically when you start the application stack with `docker-compose up`.

### Creating a New Migration

To create a new migration, you can use the `migrate` CLI tool installed in the backend container. First, ensure your containers are running. Then, execute the following command, replacing `migration_name` with a descriptive name (e.g., `add_investments_table`):

```bash
docker-compose exec backend /go/bin/migrate create -ext sql -dir /app/migrations -seq <migration_name>
```

This will create new `up` and `down` SQL files in the `/backend/migrations` directory with the next sequence number.

## Running the Application (with Docker)

This project uses Docker Compose to manage the application stack. To run the entire application, use the following command:

```bash
docker-compose up --build
```

This will start the backend, frontend, and database services.
