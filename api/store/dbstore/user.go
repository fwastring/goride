package dbstore

import (
	"database/sql"
	"errors"
	"fmt"
	"goride/hash"
	"goride/store"
)

type UserStore struct {
	db           *sql.DB
	passwordhash hash.PasswordHash
}

type NewUserStoreParams struct {
	DB           *sql.DB
	PasswordHash hash.PasswordHash
}

func NewUserStore(params NewUserStoreParams) *UserStore {
	return &UserStore{
		db:           params.DB,
		passwordhash: params.PasswordHash,
	}
}

 
func (s *UserStore) CreateUser(username string, password string) error {
	 
	hashedPassword, err := s.passwordhash.GenerateFromPassword(password)
	if err != nil {
		return err
	}

	 
	query := `INSERT INTO users (username, password_hash) VALUES ($1, $2)`
	_, err = s.db.Exec(query, username, hashedPassword)

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (s *UserStore) GetLoggedInUser(sessionID string) (*store.User, error) {
	var user store.User

	 
	query := `
		SELECT u.user_id, u.username, u.password_hash
		FROM users u
		INNER JOIN session s ON u.user_id = s.user_id
		WHERE s.session_uuid = $1
	`

	 
	row := s.db.QueryRow(query, sessionID)

	 
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return nil, errors.New("no user associated with this session")
	} else if err != nil {
		return nil, fmt.Errorf("failed to retrieve logged-in user: %w", err)
	}

	return &user, nil
}

 
func (s *UserStore) GetUser(username string) (*store.User, error) {
	var user store.User

	 
	query := `SELECT user_id, username, password_hash FROM users WHERE username = $1`
	row := s.db.QueryRow(query, username)

	 
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, fmt.Errorf("failed to retrieve user: %w", err)
	}

	return &user, nil
}
