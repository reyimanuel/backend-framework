package migrations

import "database/sql"

type createTeamTable struct{}

func getCreateTeamTable() migration {
	return &createTeamTable{}
}

func (c *createTeamTable) SkipProduction() bool {
	return false
}

func (c *createTeamTable) Name() string {
	return "create-team"
}

func (c *createTeamTable) Up(conn *sql.Tx) error {
	_, err := conn.Exec(`CREATE TABLE team (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL UNIQUE,
		role VARCHAR(50) NOT NULL,
		division VARCHAR(100) NOT NULL,
		year VARCHAR(4) NOT NULL,
		status VARCHAR(50) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	)`)

	return err
}

func (c *createTeamTable) Down(conn *sql.Tx) error {
	_, err := conn.Exec(`DROP TABLE IF EXISTS team`)
	return err
}
