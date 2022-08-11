package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/renaldid/hotel_booking_management.git/helper"
	"github.com/renaldid/hotel_booking_management.git/model/domain"
)

type RoleRepositoryImpl struct {
}

func NewRoleRepository() RoleRepository {
	return &RoleRepositoryImpl{}
}

func (repository RoleRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role {
	SQL := "insert into roles(role_name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, role.RoleName)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	role.Id = int(id)
	return role
}

func (repository RoleRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role {
	SQL := "update roles set role_name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, role.RoleName, role.Id)
	helper.PanicIfError(err)

	return role
}

func (repository RoleRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.Role) {
	SQL := "delete from roles where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicIfError(err)
}

func (repository RoleRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, roleId int) (domain.Role, error) {
	SQL := "select id, role_name from roles where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, roleId)
	helper.PanicIfError(err)
	defer rows.Close()

	role := domain.Role{}
	if rows.Next() {
		err := rows.Scan(&role.Id, &role.RoleName)
		helper.PanicIfError(err)
		return role, nil
	} else {
		return role, errors.New("role is not found")
	}
}

func (repository RoleRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Role {
	SQL := "select id, role_name from roles"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var roles []domain.Role
	for rows.Next() {
		role := domain.Role{}
		err := rows.Scan(&role.Id, &role.RoleName)
		helper.PanicIfError(err)
		roles = append(roles, role)
	}
	return roles
}
