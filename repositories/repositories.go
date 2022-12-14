package repositories

import (
	"ourgym/models"

	"github.com/midtrans/midtrans-go/snap"
)

type UserRepository interface {
	GetAll(name string) []models.User
	GetOneByFilter(key string, value any) models.User
	Create(userRequest models.User) models.User
	Update(id string, userRequest models.User) models.User
	ChangePassword(id string, newPassword string) bool
	Delete(id string) bool
	DeleteMany(ids string) bool
	CountUser() int64
}

type ClassRepository interface {
	GetAll(classType string, name string) []models.Class
	GetOneByFilter(key string, value any) models.Class
	Create(classRequest models.Class) models.Class
	Update(id string, userRequest models.Class) models.Class
	Delete(id string) bool
	DeleteMany(ids string) bool
	CountClass() int64
}

type TrainerRepository interface {
	GetAll(name string) []models.Trainer
	GetByID(id string) models.Trainer
	CountTrainer() int64
}

type PaymentMethodRepository interface {
	GetAll() []models.PaymentMethod
	GetByID(id string) models.PaymentMethod
}

type TransactionRepository interface {
	GetAll() []models.Transaction
	GetHistory() []models.Transaction
	GetByUserID(userID string) []models.Transaction
	GetByID(id string) models.Transaction
	Create(classRequest models.Transaction) (snap.Response, error)
	UpdatedByMidtransAPI(transactionRequest models.Transaction) error
	Update(id string, userRequest models.Transaction) models.Transaction
	Delete(id string) bool
	DeleteMany(ids string) bool
	CountTotalIncome() int64
}

type OtpRepository interface {
	GetOneByFilter(key string, value any) models.Otp
	Delete(id uint) bool
}
