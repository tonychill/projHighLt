package service

import (
	"errors"
	"fmt"

	"github.com/highlight/fit/data"
	"github.com/highlight/fit/entities"
)

var (
	repo data.SessionRepository
)

type SchedulerService interface {
	CreateSession(request string) ([]entities.Session, error)
	GetSession(id string) (entities.Session, error)
	CreateReminder(id string, time string) error
}

type Service struct {
	repo data.SessionRepository
}

func NewSessionService(sessionRepo data.SessionRepository) *Service {

	return &Service{
		repo: sessionRepo,
	}

}

func (s *Service) CreateSession(session *entities.Session) error {

	fmt.Println("createing sessiong with id: %s", session.Id)

	if session == nil {
		return errors.New("the session is empty")

	}

	s.repo.Save(session)

	// session.Id = fmt.Sprint(rand.Int63())
	return nil
}

func (s *Service) GetSession(id string) (*entities.Session, error) {
	ids := []string{id}
	sessions, err := s.repo.Get(ids)
	if err != nil {

		return nil, err
	}

	if len(sessions) == 0 {
		return nil, nil
	}
	return sessions[0], nil
}

func (*Service) CreateReminder(id string, time string) error {
	return nil
}
