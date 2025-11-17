package postgres

import (
	"context"

	sqlc "github.com/ahsansaif47/cdc-app/repository/postgres/schema/sqlc/generated"
	tutorial "github.com/ahsansaif47/cdc-app/repository/postgres/schema/sqlc/generated"
	"github.com/jackc/pgx/v5/pgtype"
)

type IUserRepository interface {
	CheckExistingEmail(email string) (bool, error)
	CreateUser(user sqlc.CreateUserParams) error
	FindAll() ([]sqlc.User, error)
	FindByID(id pgtype.UUID) (*sqlc.User, error)
	SetNewPassword(newPassDetails sqlc.SetNewPasswordParams) (bool, error)
	ValidateUserCredentials(credentials sqlc.ValidateUserCredentialsParams) (*sqlc.User, error)
	FindUserByEmail(email string) (*sqlc.User, error)
}

type UserRepository struct {
	db  *sqlc.Queries
	ctx context.Context
}

func NewUserRepository(db sqlc.DBTX) IUserRepository {
	return &UserRepository{
		db:  tutorial.New(db),
		ctx: context.Background(),
	}
}

func (r *UserRepository) CheckExistingEmail(email string) (bool, error) {
	status, err := r.db.CheckExistingEmail(r.ctx, email)
	if err != nil {
		return status, err
	}

	return status, nil
}

func (r *UserRepository) CreateUser(user sqlc.CreateUserParams) error {
	if err := r.db.CreateUser(r.ctx, user); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindAll() ([]sqlc.User, error) {
	users, err := r.db.FindAll(r.ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) FindByID(id pgtype.UUID) (*sqlc.User, error) {
	// Move this to the service layer
	// parsedID, err := uuid.Parse(id)
	// if err != nil {
	// 	return nil, err
	// }
	// pgUUID := pgtype.UUID{}

	// err = pgUUID.Scan(parsedID)

	user, err := r.db.FindByID(r.ctx, id)
	if err != nil {
		return nil, err
	}

	return &user, err
}

func (r *UserRepository) SetNewPassword(newPassDetails sqlc.SetNewPasswordParams) (bool, error) {
	if err := r.db.SetNewPassword(r.ctx, newPassDetails); err != nil {
		return false, err
	}

	return true, nil
}

func (r *UserRepository) ValidateUserCredentials(credentials sqlc.ValidateUserCredentialsParams) (*sqlc.User, error) {
	user, err := r.db.ValidateUserCredentials(r.ctx, credentials)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindUserByEmail(email string) (*sqlc.User, error) {
	user, err := r.db.FindUserByEmail(r.ctx, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
