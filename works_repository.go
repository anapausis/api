package database

import (
	"database/sql"
	"internal/domain"

	"log"

	_ "github.com/mattn/go-sqlite3" //追加した
)

type WorkRepository struct {
	DB *sql.DB
}

func (r *WorkRepository) Create(Works *domain.Works) error {
	_, err := r.DB.Exec("INSERT INTO works (user_id, title,statement, description) VALUES (?, ?, ?)", Works.UserID, Works.Title, Works.Statement, Works.Description)
	return err
}

func (r *WorkRepository) GetByID(id int) (*domain.Works, error) {
	row := r.DB.QueryRow("SELECT id, user_id, title, description FROM works WHERE id = ?", id)
	var work domain.Works
	err := row.Scan(&work.ID, &work.UserID, &work.Title, &work.Description)
	if err != nil {
		return nil, err
	}
	return &work, nil
}

// ここからは追加したもの
func InitDB(dataSourceName string) *sql.DB {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE
    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
