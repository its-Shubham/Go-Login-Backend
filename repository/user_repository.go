package repository

import (
	"backend/models"
	"database/sql"
	"errors"
)

type UserRepository struct {
	DB *sql.DB
}

// NewUserRepository initializes and returns a new UserRepository instance.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

// Register adds a new user to the database.
func (ur *UserRepository) Register(user *models.User) error {
	// Check if the user already exists
	var count int
	err := ur.DB.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", user.Email).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("user already exists")
	}

	// Insert the user into the database
	_, err = ur.DB.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

// Login verifies user credentials and returns the user from the database.
func (ur *UserRepository) Login(email, password string) (*models.User, error) {
	// Retrieve the user from the database
	var user models.User
	err := ur.DB.QueryRow("SELECT id, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	// Verify the password
	if user.Password != password {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}

func (ur *UserRepository) GetUserList() ([]*models.User, error) {
	// Prepare the query to retrieve users from the database
	rows, err := ur.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create a slice to hold the users
	var users []*models.User

	// Iterate over the rows and scan each row into a User struct
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		// Append the user to the slice
		users = append(users, &user)
	}

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
