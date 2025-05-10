package migrations

import "database/sql"

type CreateTeamTable struct{}

func GetCreateTeamTable() migration {
	return &CreateTeamTable{}
}

func (c *CreateTeamTable) SkipProduction() bool {
	return false
}

func (c *CreateTeamTable) Name() string {
	return "create-team"
}

func (c *CreateTeamTable) Up(conn *sql.Tx) error {
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

func (c *CreateTeamTable) Down(conn *sql.Tx) error {
	_, err := conn.Exec(`DROP TABLE IF EXISTS team`)
	return err
}
