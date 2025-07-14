package repository

import (
	"context"
	"errors"

	// "log"
	"user-service/internal/core/domain/entity"
	"user-service/internal/core/domain/model"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
}

type userRepository struct {
	db *gorm.DB
}

// Get User By Email implements UserRepositoryInterface
func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	modelUser := model.User{}

	if err := u.db.Where("email = ? && is_verified = ?", email, true).
		Preload("Roles").First(&modelUser).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					err = errors.New("404")
					log.Infof("[UserRepository-1] GetUserByEmail: User not found")
					return nil, err	
				}
			log.Errorf("[UserRepository-1] GetUserByEmail: %v", err)
			return nil, err
	}


	return &entity.UserEntity{
		ID:			modelUser.ID,
		Name:		modelUser.Name,
		Email:		email,
		Password:	modelUser.Password,
		RoleName:	modelUser.Roles[0].Name,
		Address:	modelUser.Address,
		Phone:		modelUser.Phone,
		Photo:		modelUser.Photo,
		Lat:		modelUser.Lat,
		Lng:		modelUser.Lng,
		IsVerified:	modelUser.IsVerified,
	}, nil
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepository{db: db}
}