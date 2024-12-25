package services

import (
	"anywhere-api/internal/models"
	"anywhere-api/internal/repositories"
	"anywhere-api/pkg/jwt"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// UserService mendefinisikan interface untuk operasi terkait pengguna
type UserService interface {
	CreateUser(user *models.User) error
	AuthenticateUser(username, password string) (string, error)
	GetUserByID(id int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}

// service adalah implementasi dari UserService
type service struct {
	repo repositories.UserRepository
}

// NewUserService menginisialisasi service dengan repository
func NewUserService(repo repositories.UserRepository) UserService {
	return &service{repo: repo}
}

// AuthenticateUser mengotentikasi pengguna berdasarkan username dan password
func (s *service) AuthenticateUser(username, password string) (string, error) {
	user, err := s.GetUserByUsername(username)
	if err != nil {
		return "", fmt.Errorf("service: error getting user by username: %v", err)
	}

	// Membandingkan password yang dihash dengan password yang diberikan
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("service: invalid password: %v", err)
	}

	// Generate JWT token
	token, err := jwt.GenerateJWT(user.Username)
	if err != nil {
		return "", fmt.Errorf("service: error generating token: %v", err)
	}

	return token, nil
}

// CreateUser menyimpan user baru ke dalam database
func (s *service) CreateUser(user *models.User) error {
	// Validasi atau logika bisnis lainnya dapat ditambahkan di sini
	err := s.repo.CreateUser(user)
	if err != nil {
		return fmt.Errorf("service: error creating user: %v", err)
	}
	return nil
}

// GetUserByID mengambil user berdasarkan ID
func (s *service) GetUserByID(id int) (*models.User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("service: error getting user by ID: %v", err)
	}
	return user, nil
}

// GetUserByUsername mencari user berdasarkan username
func (s *service) GetUserByUsername(username string) (*models.User, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("service: error getting user by username: %v", err)
	}
	return user, nil
}

// GetUserByEmail mencari user berdasarkan email
func (s *service) GetUserByEmail(email string) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("service: error getting user by email: %v", err)
	}
	return user, nil
}

// GetAllUsers mengambil semua user dari database
func (s *service) GetAllUsers() ([]*models.User, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("service: error getting all users: %v", err)
	}
	return users, nil
}

// UpdateUser memperbarui data user dalam database
func (s *service) UpdateUser(user *models.User) error {
	err := s.repo.UpdateUser(user)
	if err != nil {
		return fmt.Errorf("service: error updating user: %v", err)
	}
	return nil
}

// DeleteUser menghapus user berdasarkan ID
func (s *service) DeleteUser(id int) error {
	err := s.repo.DeleteUser(id)
	if err != nil {
		return fmt.Errorf("service: error deleting user: %v", err)
	}
	return nil
}
