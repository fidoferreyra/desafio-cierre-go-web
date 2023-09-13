package tickets

import (
	"context"
)

type Service interface {
	GetTotalTickets(c context.Context, destination string) (int, error)
	GetDestinationPercentage(c context.Context, destination string) (float64, error)
}

type service struct {
	repo Repository
}

// GetDestinationPercentage returns the percentage of the tickets that are destined to a certain destination
func (s *service) GetDestinationPercentage(c context.Context, destination string) (float64, error) {
	list, err := s.repo.GetAll(c)
	if err != nil {
		return 0, err
	}
	totalTickets := float64(len(list))

	listByCountry, err := s.repo.GetTicketByDestination(c, destination)
	if err != nil {
		return 0, err
	}
	ticketsByCountry := float64(len(listByCountry))

	return totalTickets / ticketsByCountry, nil
}

// GetTotalTickets return the amount of ticket
func (s *service) GetTotalTickets(c context.Context, destination string) (int, error) {
	list, err := s.repo.GetTicketByDestination(c, destination)
	if err != nil {
		return 0, err
	}
	return len(list), nil
}

func NewService(r Repository) Service {
	return &service{repo: r}
}
