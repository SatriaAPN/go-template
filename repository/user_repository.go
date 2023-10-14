package repository

import (
	"context"
	dtorepository "go-template/dto/repository"
	"go-template/entity"
	errorapp "go-template/share/error"

	"github.com/go-errors/errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, u entity.User) (entity.User, error)
	FindByEmail(ctx context.Context, email string) (entity.User, error)
	FindById(ctx context.Context, userId int) (entity.User, error)
	FindUserWalletAccountByUserId(ctx context.Context, userId int) (dtorepository.UserWalletDataResponse, error)
	CreateUserForgetPassword(ctx context.Context, urp entity.UserResetPassword) (entity.UserResetPassword, error)
	GetResetPasswordTokenByToken(ctx context.Context, token string) (entity.UserResetPassword, error)
	UpdateUser(ctx context.Context, u entity.User) (entity.User, error)
	DeleteUsedResetPassword(ctx context.Context, urp entity.UserResetPassword) (entity.UserResetPassword, error)
	DeletePreviousResetPassword(ctx context.Context, userId int) error
}

type userRepository struct {
	db *gorm.DB
}

type UserRepositoryConfig struct {
	Db *gorm.DB
}

func NewUserRepository(config UserRepositoryConfig) UserRepository {
	br := userRepository{
		db: config.Db,
	}

	return &br
}

func (ur *userRepository) Create(ctx context.Context, u entity.User) (entity.User, error) {

	err := ur.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&u).Error

		if err != nil {
			return errors.New(errorapp.ErrEmailAlreadyExist)
		}

		if err != nil {
			return err
		}

		return nil
	})

	return u, err
}

func (ur *userRepository) FindByEmail(ctx context.Context, email string) (entity.User, error) {
	u := entity.User{}

	err := ur.db.WithContext(ctx).Where("email = ?", email).First(&u).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return u, errors.New(errorapp.ErrEmailNotFound)
	}

	return u, err
}

func (ur *userRepository) FindById(ctx context.Context, userId int) (entity.User, error) {
	u := entity.User{}

	err := ur.db.WithContext(ctx).Where("id = ?", userId).First(&u).Error

	return u, err
}

func (ur *userRepository) FindUserWalletAccountByUserId(ctx context.Context, userId int) (dtorepository.UserWalletDataResponse, error) {
	res := dtorepository.UserWalletDataResponse{}

	err := ur.db.WithContext(ctx).Raw(`
		select 
			w."number" as wallet_number,
			sum(wh.amount) as wallet_balance
		from users u 
		inner join wallets w 
			on w.user_id = u.id 
		left join wallet_histories wh 
			on wh.wallet_id = w.id 
		where u.id = 1
		group by w."number" 
	`).Scan(&res).Error

	return res, err
}

func (ur *userRepository) CreateUserForgetPassword(ctx context.Context, urp entity.UserResetPassword) (entity.UserResetPassword, error) {
	err := ur.db.WithContext(ctx).Create(&urp).Error

	return urp, err
}

func (uu *userRepository) GetResetPasswordTokenByToken(ctx context.Context, token string) (entity.UserResetPassword, error) {
	urp := entity.UserResetPassword{}

	err := uu.db.WithContext(ctx).Where("token = ?", token).First(&urp).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return urp, errors.New(errorapp.ErrResetCodeNotFound)
	}

	return urp, err
}

func (uu *userRepository) UpdateUser(ctx context.Context, u entity.User) (entity.User, error) {
	err := uu.db.WithContext(ctx).Save(&u).Error

	return u, err
}

func (uu *userRepository) DeleteUsedResetPassword(ctx context.Context, urp entity.UserResetPassword) (entity.UserResetPassword, error) {
	err := uu.db.WithContext(ctx).Delete(&urp).Error

	return urp, err
}

func (uu *userRepository) DeletePreviousResetPassword(ctx context.Context, userId int) error {
	err := uu.db.WithContext(ctx).Where("user_id = ?", userId).Delete(&entity.UserResetPassword{}).Error

	return err
}
