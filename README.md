
# Anywhere API

This project is part of the Mobile Programming course. It is called Anywhere, a mobile app that provides parking spaces for everyone.

The ***Anywhere API***, built with the Go-based Fiber web framework and PostgreSQL database, serves as the backend for the app. It uses JWT solely for authentication purposes.
## Tech Stack

[![My Skills](https://skillicons.dev/icons?i=go,postgres,docker)](https://skillicons.dev)
# Anywhere API



## Prerequisites

Before you begin, make sure you have the following installed:

- **Go**: Download and install Go from [https://golang.org/dl/](https://golang.org/dl/)
- **PostgreSQL**: You need a PostgreSQL instance running. You can install PostgreSQL from [https://www.postgresql.org/download/](https://www.postgresql.org/download/)
- **Git**: Ensure Git is installed for cloning the repository. Install from [https://git-scm.com/](https://git-scm.com/)

## Installation

Follow these steps to get the project running on your local machine.

### Step 1: Clone the Repository

Clone the project repository to your local machine:

```bash
git clone https://github.com/yourusername/anywhere-api.git
cd anywhere-api
```

### Step 2: Setup Environment Variables
Create a **.env** file in the root of your project directory and add the following database configuration:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_database_user
DB_PASSWORD=your_database_password
DB_NAME=your_database_name
JWT_SECRET=your_jwt_secret_key

```

### Step 3: Install Dependencies
Ensure you have Go installed on your system. Install the required dependencies using go mod:

```bash
go mod tidy
```
This will download all the required dependencies listed in the go.mod file.

### Step 4: Set Up PostgreSQL Database

```sql
CREATE DATABASE your_database_name;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password TEXT NOT NULL
);
```

### Step 5: Running the Application

```bash
go run cmd/main.go
```
## Authors

- [@attmhd](https://github.com/attnmhd/)