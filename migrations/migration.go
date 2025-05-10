package migrations

import (
	"backend/config"
	"database/sql"
)

// The migration interface defines methods that each migration should implement.
type migration interface {
	Name() string            // Returns the name of the migration.
	Up(conn *sql.Tx) error   // Defines the logic for applying the migration.
	Down(conn *sql.Tx) error // Defines the logic for rolling back the migration.
	SkipProduction() bool    // Indicates whether the migration should be skipped in production.
}

// getMigrations returns a list of migrations that need to be executed.
func getMigrations() []migration {
	return []migration{
		GetCreateUsersTable(), // Implementing the users table migration
		GetInsertUsersTable(), // Implementing the users table data migration
		GetCreateTeamTable(),
		GetCreateEventTable(),
		// Add your migrations here
	}
}

// checkDuplicateMigration ensures there are no duplicate migration names in the list.
func checkDuplicateMigration(migrations []migration) {
	nameSet := make(map[string]bool)
	for _, m := range migrations {
		if nameSet[m.Name()] {
			panic("Duplicate migration name: " + m.Name()) // Panic if a duplicate migration name is found.
		}
		nameSet[m.Name()] = true
	}
}

// Up applies all pending migrations to the database.
func Up(db *sql.DB) {
	// Get all migrations
	migrations := getMigrations()

	// Load configuration
	cfg := config.Get()

	// Ensure there are no duplicate migration names
	checkDuplicateMigration(migrations)

	// Create the migrations table if it doesn't exist
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			name VARCHAR(255) PRIMARY KEY,
			applied_at TIMESTAMP NOT NULL DEFAULT NOW()
		)
	`)

	if err != nil {
		panic(err)
	}

	// Begin transaction to apply migrations
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback() // Ensure rollback if migration fails

	for _, m := range migrations {
		// Check if the migration has already been applied
		var count int
		err := tx.QueryRow("SELECT COUNT(*) FROM migrations WHERE name = $1", m.Name()).Scan(&count)
		if err != nil {
			panic(err)
		}

		if count == 0 { // Apply migration only if it's not already applied
			if cfg.IsProduction && m.SkipProduction() {
				continue // Skip migration if production environment and marked as skippable
			}

			// Execute the migration's Up method
			if err := m.Up(tx); err != nil {
				panic(err)
			}

			// Record the migration as applied in the database
			_, err = tx.Exec("INSERT INTO migrations (name) VALUES ($1)", m.Name())
			if err != nil {
				panic(err)
			}

			println("Applied migration:", m.Name())
		}
	}

	// Commit transaction to finalize migration
	if err := tx.Commit(); err != nil {
		panic(err)
	}
}

// Down reverts the last applied migration.
func Down(db *sql.DB) {
	// Get all migrations
	migrations := getMigrations()

	// Ensure no duplicate migration names
	checkDuplicateMigration(migrations)

	// Begin transaction for rolling back the migration
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	// Retrieve the last applied migration
	var lastMigration string
	err = tx.QueryRow("SELECT name FROM migrations ORDER BY applied_at DESC LIMIT 1").Scan(&lastMigration)
	if err != nil {
		if err == sql.ErrNoRows {
			println("No migrations to revert") // No migration to rollback
			return
		}
		panic(err)
	}

	// Find the migration in the list
	var migrationToRevert migration
	for i := len(migrations) - 1; i >= 0; i-- {
		if migrations[i].Name() == lastMigration {
			migrationToRevert = migrations[i]
			break
		}
	}

	if migrationToRevert == nil {
		panic("Last applied migration not found in migration list")
	}

	// Execute the Down method to revert the migration
	if err := migrationToRevert.Down(tx); err != nil {
		panic(err)
	}

	// Remove the migration record from the migrations table
	_, err = tx.Exec("DELETE FROM migrations WHERE name = $1", lastMigration)
	if err != nil {
		panic(err)
	}

	println("Reverted migration:", lastMigration)

	// Commit the rollback transaction
	if err := tx.Commit(); err != nil {
		panic(err)
	}
}

// DownAll reverts all applied migrations.
func DownAll(db *sql.DB) {
	// Get all migrations
	migrations := getMigrations()

	// Ensure no duplicate migration names
	checkDuplicateMigration(migrations)

	// Begin transaction to revert all migrations
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	// Iterate over all migrations in reverse order to revert them
	for i := len(migrations) - 1; i >= 0; i-- {
		m := migrations[i]

		// Check if the migration has been applied
		var count int
		err := tx.QueryRow("SELECT COUNT(*) FROM migrations WHERE name = $1", m.Name()).Scan(&count)
		if err != nil {
			panic(err)
		}

		if count > 0 {
			// Execute the Down method to revert the migration
			if err := m.Down(tx); err != nil {
				panic(err)
			}

			// Remove the migration record from the database
			_, err = tx.Exec("DELETE FROM migrations WHERE name = $1", m.Name())
			if err != nil {
				panic(err)
			}

			println("Reverted migration:", m.Name())
		}
	}

	// Commit transaction after all migrations have been reverted
	if err := tx.Commit(); err != nil {
		panic(err)
	}
}
