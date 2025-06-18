package postgres

import (
	"database/sql"
	"errors"

	"github.com/evaizee/seat-arrangements/backend/models"
	"github.com/evaizee/seat-arrangements/backend/repositories"
)

// UserRepository is a PostgreSQL implementation of the UserRepository interface
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *sql.DB) repositories.UserRepository {
	return &UserRepository{db: db}
}

// Create creates a new user in the database
func (r *UserRepository) Create(user *models.User) error {
	query := `
		INSERT INTO users (id, email, password, first_name, last_name, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.Exec(
		query,
		user.ID,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.CreatedAt,
		user.UpdatedAt,
	)

	return err
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(id string) (*models.User, error) {
	query := `
		SELECT id, email, password, first_name, last_name, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	user := &models.User{}
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // User not found
		}
		return nil, err
	}

	return user, nil
}

// GetByEmail retrieves a user by email
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, email, password, first_name, last_name, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	user := &models.User{}
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // User not found
		}
		return nil, err
	}

	return user, nil
}

// Update updates a user in the database
func (r *UserRepository) Update(user *models.User) error {
	query := `
		UPDATE users
		SET email = $2, password = $3, first_name = $4, last_name = $5, updated_at = $6
		WHERE id = $1
	`

	_, err := r.db.Exec(
		query,
		user.ID,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.UpdatedAt,
	)

	return err
}

// Delete deletes a user from the database
func (r *UserRepository) Delete(id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}