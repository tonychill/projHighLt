package data

import (
	"github.com/highlight/fit/entities"
)

type SessionRepository interface {
	Save(session *entities.Session) error
	Get(ids []string) ([]*entities.Session, error)
	Update(session *entities.Session) (*entities.Session, error)
	Delete(id string) error
}
