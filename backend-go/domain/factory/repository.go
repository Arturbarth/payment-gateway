package factory

import "github.com/Arturbarth/payment-gateway/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
