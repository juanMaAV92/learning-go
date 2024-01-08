package database

import (
	"context"
	"database/sql"

	"api.com/go/rest-ws/models"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}

func (repo *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)", user.Id, user.Email, user.Password)
	return err
}

func (repo *PostgresRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	user := &models.User{}
	err := repo.db.QueryRowContext(ctx, "SELECT id, email, password FROM users WHERE id = $1", id).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	err := repo.db.QueryRowContext(ctx, "SELECT id, email, password FROM users WHERE email = $1", email).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
