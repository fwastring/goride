package dbstore

import (
	"database/sql"
	"fmt"
	"goride/store"

	"github.com/google/uuid"
)

type SessionStore struct {
	db *sql.DB
}

type NewSessionStoreParams struct {
	DB *sql.DB
}

 
func NewSessionStore(params NewSessionStoreParams) *SessionStore {
	return &SessionStore{
		db: params.DB,
	}
}

 
func (s *SessionStore) CreateSession(session *store.Session) (*store.Session, error) {
	session.SessionID = uuid.New().String()

	query := `
		INSERT INTO session (session_uuid, user_id, last_seen)
		VALUES ($1, $2, $3)
		RETURNING session_uuid, user_id, last_seen;
	`

	row := s.db.QueryRow(query, session.SessionID, session.UserID, session.LastSeen)

	err := row.Scan(&session.SessionID, &session.UserID, &session.LastSeen)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	return session, nil
}

 
func (s *SessionStore) GetUserFromSession(sessionID string, userID string) (*store.User, error) {
	var user store.User

	query := `
		SELECT u.user_id, u.username
		FROM session s
		JOIN users u ON s.user_id = u.user_id
		WHERE s.session_uuid = $1 AND s.user_id = $2;
	`

	row := s.db.QueryRow(query, sessionID, userID)

	err := row.Scan(&user.ID, &user.Username)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("no user associated with the session")
	} else if err != nil {
		return nil, fmt.Errorf("failed to get user from session: %w", err)
	}

	return &user, nil
}

