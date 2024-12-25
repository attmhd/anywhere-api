package repositories

import (
	"anywhere-api/internal/models"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Driver PostgreSQL
)

// UserRepository mendefinisikan interface untuk operasi CRUD pada user
type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}

// repository adalah implementasi dari UserRepository
type repository struct {
	db *sql.DB
}

// NewUserRepository menginisialisasi repository dengan koneksi database
func NewUserRepository(db *sql.DB) UserRepository {
	return &repository{db: db}
}

// createQueryRow is a helper function to execute queries that return a single row
func (r *repository) createQueryRow(query string, args ...interface{}) *sql.Row {
	return r.db.QueryRow(query, args...)
}

// createQuery is a helper function to execute queries that return multiple rows
func (r *repository) createQuery(query string, args ...interface{}) (*sql.Rows, error) {
	return r.db.Query(query, args...)
}

// CreateUser menyimpan user baru ke dalam database
func (r *repository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
	err := r.createQueryRow(query, user.Username, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}
	return nil
}

// GetUserByID mengambil user berdasarkan ID
func (r *repository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, email FROM users WHERE id = $1`
	err := r.createQueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Tidak ditemukan user
		}
		return nil, fmt.Errorf("error getting user by ID: %v", err)
	}
	return &user, nil
}

// GetUserByUsername mencari user berdasarkan username
func (r *repository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, email, password FROM users WHERE username = $1`
	err := r.createQueryRow(query, username).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Tidak ditemukan user
		}
		return nil, fmt.Errorf("error getting user by username: %v", err)
	}
	return &user, nil
}

// GetUserByEmail mencari user berdasarkan email
func (r *repository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, email FROM users WHERE email = $1`
	err := r.createQueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Tidak ditemukan user
		}
		return nil, fmt.Errorf("error getting user by email: %v", err)
	}
	return &user, nil
}

// GetAllUsers mengambil semua user dari database
func (r *repository) GetAllUsers() ([]*models.User, error) {
	query := `SELECT id, username, email FROM users`
	rows, err := r.createQuery(query)
	if err != nil {
		return nil, fmt.Errorf("error getting all users: %v", err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			return nil, fmt.Errorf("error scanning user: %v", err)
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return users, nil
}

// UpdateUser memperbarui data user dalam database
func (r *repository) UpdateUser(user *models.User) error {
	query := `UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4`
	_, err := r.db.Exec(query, user.Username, user.Email, user.Password, user.ID)
	if err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}
	return nil
}

// DeleteUser menghapus user berdasarkan ID
func (r *repository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}
	return nil
}
