package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/renaldid/hotel_booking_management.git/helper"
	"github.com/renaldid/hotel_booking_management.git/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "insert into users(name, email, password, role_id) values (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, user.Name, user.Email, user.Password, user.RoleId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)
	return user
}

func (repository UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "update users set name = ?, email = ?, password = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Name, user.Email, user.Password, user.Id)
	helper.PanicIfError(err)

	return user
}

func (repository UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := "delete from users where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicIfError(err)
}

func (repository UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "select u.id, u.name, u.email, u.password, u.role_id, r.role_name from users u inner join roles r on u.role_id=r.id where u.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.RoleId, &user.RoleName)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

func (repository UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "select u.id, u.name, u.email, u.password, u.role_id, r.role_name from users u inner join roles r on u.role_id=r.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.RoleId, &user.RoleName)
		helper.PanicIfError(err)
		users = append(users, user)
	}
	return users
}
