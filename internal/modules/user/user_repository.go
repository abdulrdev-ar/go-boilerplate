package user

import (
	"errors"

	"github.com/inienam06/go-boilerplate/internal/model"
	"github.com/inienam06/go-boilerplate/internal/util"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user *model.User) (*model.User, error)
	GetByID(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	List(pagination util.Pagination) (*util.Pagination, error)
	Update(user *model.User) error
	Delete(id uint) error
}

type UserRepository struct {
	db *gorm.DB
}

func InitUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) (*model.User, error) {
	model := r.db.Create(user)
	if model.Error != nil {
		return nil, model.Error
	}

	return user, nil
}

func (r *UserRepository) GetByID(id uint) (*model.User, error) {
	var u model.User
	if err := r.db.Select("id", "name", "email", "created_at", "updated_at").First(&u, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	var u model.User
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) List(pagination util.Pagination) (*util.Pagination, error) {
	var users []model.User

	r.db.Scopes(util.Paginate(users, &pagination, r.db)).Select("id", "name", "email", "created_at", "updated_at").Order("created_at asc").Find(&users)
	pagination.Data = users

	return &pagination, nil
}

func (r *UserRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}
