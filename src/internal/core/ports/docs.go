package ports

import (
	"bytes"

	"github.com/isaias-dgr/stori-balance/src/internal/core/domain"
)

type Docs interface {
	GetDoc(acc *domain.Account) (*bytes.Buffer, error)
}
