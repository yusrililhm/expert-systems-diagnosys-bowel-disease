package user_repository

import (
	"database/sql"
	"log"
	"sync"

	"healthy-bowel/internal/domain"
	"healthy-bowel/internal/pkg/errors"
)

type userRepositoryImpl struct {
	db *sql.DB
	wg *sync.WaitGroup
}

const (
	addUserQuery = `insert into users (username, email, full_name, gender, role, password) values ($1, $2, $3, $4, $5, $6)`

	changePasswordQuery = `update on users set password = $2, updated_at = now() where id = $1`

	editUserQuery = `update on users set username = $2, email = $3, full_name = $4, gender = $5, updated_at = now() where id = $1`

	getAllUsersQuery = `select id, username, email, full_name, gender, role, created_at, updated_at`

	getUserByIdQuery = `select id, username, email, full_name, gender, role, created_at, updated_at where id = $1`

	getUserByUsernameQuery = `select id, username, email, full_name, gender, role, password, created_at, updated_at where username = $1`
)

// Add implements UserRepository.
func (ur *userRepositoryImpl) Add(user *domain.User) errors.Errors {
	tx, err := ur.db.Begin()

	if err != nil {
		log.Println("[warn]", err.Error())
		tx.Rollback()
		return errors.NewInternalServerError("Oops! Something went wrong!")
	}

	stmt, err := tx.Prepare(addUserQuery)

	if err != nil {
		log.Println("[warn]", err.Error())
		tx.Rollback()
		return errors.NewInternalServerError("Oops! Something went wrong!")
	}

	defer stmt.Close()

	if _, err := stmt.Exec(user.Username, user.Email, user.FullName, user.Gender, user.Role, user.Password); err != nil {
		log.Println("[warn]", err.Error())
		tx.Rollback()
		return errors.NewInternalServerError("Oops! Something went wrong!")
	}

	if err := tx.Commit(); err != nil {
		log.Println("[warn]", err.Error())
		tx.Rollback()
		return errors.NewInternalServerError("Oops! Something went wrong!")
	}

	return nil
}

// ChanngePassword implements UserRepository.
func (ur *userRepositoryImpl) ChanngePassword(user *domain.User) errors.Errors {
	tx, err := ur.db.Begin()

	if err != nil {
		log.Println("[warn]", err.Error())
		tx.Rollback()
		return errors.NewInternalServerError("Oops! Something went wrong!")
	}

	stmt, err := tx.Prepare(changePasswordQuery)

	if err != nil {
		log.Println("[warn]", err.Error())
		tx.Rollback()
		return errors.NewInternalServerError("Oops! Something went wrong!")
	}

	defer stmt.Close()

	if _, err := stmt.Exec(user.Id, user.Password); err != nil {
		log.Println("[warn]", err.Error())
		tx.Rollback()
		return errors.NewInternalServerError("Oops! Something went wrong!")
	}

	if err := tx.Commit(); err != nil {
		log.Println("[warn]", err.Error())
		tx.Rollback()
		return errors.NewInternalServerError("Oops! Something went wrong!")
	}

	return nil
}

// Edit implements UserRepository.
func (ur *userRepositoryImpl) Edit(user *domain.User) errors.Errors {
	tx, err := ur.db.Begin()

	if err != nil {
		log.Println("[warn]", err.Error())
		tx.Rollback()
		return errors.NewInternalServerError("Oops! Something went wrong!")
	}

	stmt, err := tx.Prepare(editUserQuery)

	if err != nil {
		log.Println("[warn]", err.Error())
		tx.Rollback()
		return errors.NewInternalServerError("Oops! Something went wrong!")
	}

	defer stmt.Close()

	if _, err := stmt.Exec(user.Id, user.Username, user.Email, user.FullName, user.Gender); err != nil {
		log.Println("[warn]", err.Error())
		tx.Rollback()
		return errors.NewInternalServerError("Oops! Something went wrong!")
	}

	if err := tx.Commit(); err != nil {
		log.Println("[warn]", err.Error())
		tx.Rollback()
		return errors.NewInternalServerError("Oops! Something went wrong!")
	}

	return nil
}

// GetAll implements UserRepository.
func (ur *userRepositoryImpl) GetAll() ([]*domain.User, errors.Errors) {
	users := []*domain.User{}

	rows, err := ur.db.Query(getAllUsersQuery)

	if err != nil {
		log.Println("[warn]", err.Error())
		return nil, errors.NewInternalServerError("Oops! Something went wrong!")
	}

	defer rows.Close()

	for rows.Next() {
		user := &domain.User{}

		if err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Email,
			&user.FullName,
			&user.Gender,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			log.Println("[warn]", err.Error())
			return nil, errors.NewInternalServerError("Oops! Something went wrong!")
		}

		users = append(users, user)
	}

	return users, nil
}

// GetById implements UserRepository.
func (ur *userRepositoryImpl) GetById(id uint) (*domain.User, errors.Errors) {
	user := &domain.User{}

	if err := ur.db.QueryRow(getUserByIdQuery, id).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.FullName,
		&user.Gender,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			log.Println("[warn]", err.Error())
			return nil, errors.NewNotFound("User not found")
		}

		log.Println("[warn]", err.Error())
		return nil, errors.NewInternalServerError("Oops! Something went wrong!")
	}

	return user, nil
}

// GetByUsername implements UserRepository.
func (ur *userRepositoryImpl) GetByUsername(username string) (*domain.User, errors.Errors) {
	user := &domain.User{}

	if err := ur.db.QueryRow(getUserByUsernameQuery, username).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.FullName,
		&user.Gender,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			log.Println("[warn]", err.Error())
			return nil, errors.NewNotFound("User not found")
		}

		log.Println("[warn]", err.Error())
		return nil, errors.NewInternalServerError("Oops! Something went wrong!")
	}

	return user, nil
}

func NewUserRepositoryImpl(db *sql.DB, wg *sync.WaitGroup) UserRepository {
	return &userRepositoryImpl{
		db: db,
		wg: wg,
	}
}
