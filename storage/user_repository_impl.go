package storage

import (
	"database/sql"
	"errors"
)

type userRepositoryImpl struct {
	DB *sql.DB
}

func (r *userRepositoryImpl) Create(user *User) error {
	res, err := r.DB.Exec(`
        INSERT INTO users (username, password, preferred_language, target_language)
        VALUES (?, ?, ?, ?)`,
		user.Username, user.Password, user.PreferredLanguage, user.TargetLanguage)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	user.Id = int(id)
	return nil
}

func (r *userRepositoryImpl) Update(user *User) error {
	_, err := r.DB.Exec(`
        UPDATE users
        SET username = ?,
		    password = ?,
		    preferred_language = ?,
		    target_language = ?
		WHERE id = ?;
    `, user.Username, user.Password, user.PreferredLanguage, user.TargetLanguage, user.Id)
	return err
}

func (r *userRepositoryImpl) GetByID(id int) (*User, error) {
	var u User
	row := r.DB.QueryRow(`
        SELECT id, username, password, preferred_language, target_language
        FROM users WHERE id = ?`, id)

	if err := row.Scan(&u.Id, &u.Username, &u.Password, &u.PreferredLanguage, &u.TargetLanguage); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *userRepositoryImpl) GetByUsername(username string) (*User, error) {
	var u User
	row := r.DB.QueryRow(`
        SELECT id, username, password, preferred_language, target_language
        FROM users WHERE username = ?`, username)
	if err := row.Scan(&u.Id, &u.Username, &u.Password, &u.PreferredLanguage, &u.TargetLanguage); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}
