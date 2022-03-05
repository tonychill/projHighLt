package data

import (
	"github.com/highlight/fit/entities"
)

type repo struct {
	Sessions []*entities.Session
}

func NewSessionRepository() SessionRepository {
	return &repo{}
}

func (r *repo) Save(session *entities.Session) error {

	r.Sessions = append(r.Sessions, session)
	return nil
}

// Get(ids []string) ([]entities.Session, error)
//
//
func (r *repo) Get(ids []string) ([]*entities.Session, error) {

	return r.Sessions, nil

}

func (*repo) Update(session *entities.Session) (*entities.Session, error) {
	return nil, nil
}
func (*repo) Delete(id string) error {
	return nil
}
