package ports

import (
	"github.com/isaias-dgr/stori-balance/src/internal/core/domain"
)

type Repository interface {
	Save(account, product string, transaction *domain.Transaction) error
	GetUser(id string) (*domain.User, error)
}
