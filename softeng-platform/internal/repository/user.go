package repository

import (
	"context"
	"database/sql"
	"errors"
	"softeng-platform/internal/model"
	"time"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id int) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	UpdatePassword(ctx context.Context, userID int, hashedPassword string) error
}

type userRepository struct {
	db *Database
}

func NewUserRepository(db *Database) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	query := `
		INSERT INTO users (username, nickname, email, password, avatar, description, face_photo, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id
	`

	err := r.db.QueryRowContext(ctx, query,
		user.Username,
		user.Nickname,
		user.Email,
		user.Password,
		user.Avatar,
		user.Description,
		user.FacePhoto,
		user.Role,
		time.Now(),
		time.Now(),
	).Scan(&user.ID)

	return err
}

func (r *userRepository) GetByID(ctx context.Context, id int) (*model.User, error) {
	query := `
		SELECT id, username, nickname, email, avatar, description, face_photo, role, created_at, updated_at
		FROM users WHERE id = $1
	`

	user := &model.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Nickname,
		&user.Email,
		&user.Avatar,
		&user.Description,
		&user.FacePhoto,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	query := `
		SELECT id, username, nickname, email, password, avatar, description, face_photo, role, created_at, updated_at
		FROM users WHERE username = $1
	`

	user := &model.User{}
	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Nickname,
		&user.Email,
		&user.Password,
		&user.Avatar,
		&user.Description,
		&user.FacePhoto,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `
		SELECT id, username, nickname, email, password, avatar, description, face_photo, role, created_at, updated_at
		FROM users WHERE email = $1
	`

	user := &model.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Nickname,
		&user.Email,
		&user.Password,
		&user.Avatar,
		&user.Description,
		&user.FacePhoto,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	query := `
		UPDATE users 
		SET nickname = $1, avatar = $2, description = $3, face_photo = $4, updated_at = $5
		WHERE id = $6
	`

	_, err := r.db.ExecContext(ctx, query,
		user.Nickname,
		user.Avatar,
		user.Description,
		user.FacePhoto,
		time.Now(),
		user.ID,
	)

	return err
}

func (r *userRepository) UpdatePassword(ctx context.Context, userID int, hashedPassword string) error {
	query := `UPDATE users SET password = $1, updated_at = $2 WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, hashedPassword, time.Now(), userID)
	return err
}
