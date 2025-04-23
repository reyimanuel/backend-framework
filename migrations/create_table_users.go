package migrations

import "database/sql"

// Struct that implements the migration interface for creating the "users" table
type createUsersTable struct{}

// Returns a new instance of the createUsersTable migration
func getCreateUsersTable() migration {
	return &createUsersTable{}
}

// Indicates whether this migration should be skipped in production
// Returning false means this migration will also run in production
func (c *createUsersTable) SkipProduction() bool {
	return false
}

// Returns the unique name of the migration
func (c *createUsersTable) Name() string {
	return "create-users"
}

// Up is the logic for applying the migration
// It creates the "users" table with specified columns
func (c *createUsersTable) Up(conn *sql.Tx) error {
	_, err := conn.Exec(`CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) NOT NULL UNIQUE,
		username VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		nim VARCHAR(255) NOT NULL UNIQUE,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	)`)

	return err // Returns error if creation fails
}

// Down is the logic for reverting the migration
// It drops the "users" table if it exists
func (c *createUsersTable) Down(conn *sql.Tx) error {
	_, err := conn.Exec(`DROP TABLE users`) // Deletes the table

	return err // Returns error if deletion fails
}
