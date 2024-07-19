package postgres

import (
	"database/sql"
	"fmt"
	"learning/env/config"
	stg "learning/storage"
	_ "github.com/lib/pq"
)

type Storage struct {
	Db    *sql.DB
	learn stg.LearningService
}

func NewPostgresStorage() (stg.InitRoot, error) {
	config := config.Load()
	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.PostgresUser, config.PostgresPassword,
		config.PostgresHost, config.PostgresPort,
		config.PostgresDatabase)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Storage{Db: db}, nil
}

func (s *Storage) Learning() stg.LearningService { // Implementing the required method
	if s.learn == nil {
		s.learn = &LearningStorage{db: s.Db}
	}
	return s.learn
}
