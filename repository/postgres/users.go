package postgres

import (
	"database/sql"

	"github.com/ahsansaif47/cdc-app/models"
	"github.com/ahsansaif47/cdc-app/repository/postgres/schema/sqlc"
)

type IUserRepository interface {
	CheckExistingEmail(email string) (bool, error)
	CreateUser(user *models.User) error
	FindAll() ([]models.User, error)
	FindByID(id string) (*models.User, error)
	GetAllVendors() ([]models.User, error)
	GetAllUsers() ([]models.User, error)
	SetNewPassword(email, newPassword string) (bool, error)
	ValidateUserCredentials(email, password string) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
}

type UserRepository struct {
	db *generated.Queries
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CheckExistingEmail(email string) (bool, error) {
	panic("unimplemented")
}

func (r *UserRepository) CreateUser(user *models.User) error {
	panic("unimplemented")
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	panic("unimplemented")
}

func (r *UserRepository) FindByID(id string) (*models.User, error) {
	panic("unimplemented")
}

func (r *UserRepository) GetAllVendors() ([]models.User, error) {
	panic("unimplemented")
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	panic("unimplemented")
}

func (r *UserRepository) SetNewPassword(email, newPassword string) (bool, error) {
	panic("unimplemented")
}

func (r *UserRepository) ValidateUserCredentials(email, password string) (*models.User, error) {
	panic("unimplemented")
}

func (r *UserRepository) FindUserByEmail(email string) (*models.User, error) {
	panic("unimplemented")
}
