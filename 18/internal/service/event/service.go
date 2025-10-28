package event

import (
	"log/slog"
	"time"

	"github.com/misshanya/wb-tech-l2/18/internal/models"
)

type repo interface {
	Create(e *models.Event) *models.Event
	GetForDay(date time.Time) []*models.Event
	GetForWeek(date time.Time) []*models.Event
	Update(e *models.Event) error
	Delete(id int)
}

type service struct {
	l    *slog.Logger
	repo repo
}

func New(l *slog.Logger, repo repo) *service {
	return &service{
		l:    l,
		repo: repo,
	}
}
