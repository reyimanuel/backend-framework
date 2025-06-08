package migrations

import "database/sql"

type CreateEventTable struct{}

func GetCreateEventTable() migration {
	return &CreateEventTable{}
}

func (c *CreateEventTable) SkipProduction() bool {
	return false
}

func (c *CreateEventTable) Name() string {
	return "create-event"
}

func (c *CreateEventTable) Up(conn *sql.Tx) error {
	_, err := conn.Exec(`
		CREATE TABLE event (
			id SERIAL PRIMARY KEY,
			event_name VARCHAR(255) NOT NULL,
			event_date DATE NOT NULL,
			event_time TIME NOT NULL,
			event_location VARCHAR(255) NOT NULL,
			event_description TEXT,
			event_organizer VARCHAR(255) NOT NULL,
			event_status VARCHAR(50) NOT NULL,
			event_image_url VARCHAR(255),
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		)
	`)
	return err
}

func (c *CreateEventTable) Down(conn *sql.Tx) error {
	_, err := conn.Exec(`DROP TABLE IF EXISTS event`)
	return err
}
