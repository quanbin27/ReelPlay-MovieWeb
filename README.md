
# ReelPlay-MovieWeb

ReelPlay-MovieWeb is an online movie web application where users can explore, search, and watch their favorite movies.

## Key Features

- Movie Browsing: Search and explore categorized movies.
- Movie Details: View cast, genre, release year, and more.
- Favorites & Comments: Save favorites and leave comments.
- Ratings: Rate movies and adjust anytime.
- Watch History: Resume from last watched point.
- Admin Controls: Manage users, movies, and view statistics.

## Demo

- https://youtu.be/lwsgQuDp18I


## Technologies Used

- Frontend: HTML, CSS, JavaScript
- Backend: Go, Echo Framework
- Database: My SQL



## Installation
Before setting up the project, ensure you have the following installed:

- **Go** (>=1.2) [Download here](https://go.dev/dl/)
- **MySQL** (>=5.7) [Download here](https://dev.mysql.com/downloads/)
- **golang-migrate** (for database migrations) [Installation Guide](https://github.com/golang-migrate/migrate)

📥 Clone the Repository

```sh
git clone https://github.com/quanbin27/ReelPlay-MovieWeb
cd yourproject
```
Open MySQL and create the database:
```sh
CREATE DATABASE IF NOT EXISTS yourdatabase;
```
Update the .env file with your database credentials:
```sh
DSN=root:12345678@tcp(127.0.0.1:3306)/yourdatabase?charset=utf8mb4&parseTime=True&loc=Local
```
Make sure golang-migrate is installed, then run:
```sh
 migrate -path migrate/migrations -database "mysql://root:12345678@tcp(localhost:3306)/yourdatabase" up
```
▶️ Run the Project
```sh
go run cmd/main.go
```

