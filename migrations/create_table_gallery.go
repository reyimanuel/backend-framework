package migrations

import "database/sql"

type CreateGalleryTable struct{}

func GetCreateGalleryTable() migration {
	return &CreateGalleryTable{}
}

func (c *CreateGalleryTable) SkipProduction() bool {
	return false
}

func (c *CreateGalleryTable) Name() string {
	return "create-gallery"
}

func (c *CreateGalleryTable) Up(conn *sql.Tx) error {
	_, err := conn.Exec(`
		CREATE TABLE gallery (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			description TEXT,
			image_url VARCHAR(255) NOT NULL,
			category VARCHAR(100) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		)
	`)
	return err
}

func (c *CreateGalleryTable) Down(conn *sql.Tx) error {
	_, err := conn.Exec(`DROP TABLE IF EXISTS gallery`)
	return err
}
