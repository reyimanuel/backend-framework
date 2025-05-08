package migrations

import "database/sql"

type insertUsersTable struct{}

func getInsertUsersTable() migration {
	return &insertUsersTable{}
}

func (i *insertUsersTable) SkipProduction() bool {
	return true
}

func (i *insertUsersTable) Name() string {
	return "insert-users"
}

func (i *insertUsersTable) Up(conn *sql.Tx) error {
	_, err := conn.Exec(`
		INSERT INTO users (email, username, password, nim) VALUES
		('reyimanuel32@gmail.com', 'reyimanuel32', '$2a$10$tCPHIA36fL5O2Fss23Y1eOZhOj8RiaJQyl1Ls6nA3..Q2SpQILJJC', '220211060171')
	`)

	return err
}

func (i *insertUsersTable) Down(conn *sql.Tx) error {
	_, err := conn.Exec(`DELETE FROM users WHERE email IN ('reyimanuel32@gmail.com')`)
	return err
}
