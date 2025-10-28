package event

import (
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/misshanya/wb-tech-l2/18/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestService_GetForMonth(t *testing.T) {
	tests := []struct {
		Name           string
		InputTime      time.Time
		ExpectedEvents []*models.Event
		AllEvents      []*models.Event
		SetUpMocks     func(repo *mockrepo, allEvents []*models.Event)
	}{
		{
			Name:      "2 events",
			InputTime: time.Date(2025, time.October, 28, 0, 0, 0, 0, time.UTC),
			ExpectedEvents: []*models.Event{
				{
					ID:     1,
					UserID: 4,
					Title:  "event 2",
					Date:   time.Date(2025, time.October, 10, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:     3,
					UserID: 7,
					Title:  "event 4",
					Date:   time.Date(2025, time.October, 29, 0, 0, 0, 0, time.UTC),
				},
			},
			AllEvents: []*models.Event{
				{
					ID:     0,
					UserID: 3,
					Title:  "event 1",
					Date:   time.Date(2024, time.February, 15, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:     1,
					UserID: 4,
					Title:  "event 2",
					Date:   time.Date(2025, time.October, 10, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:     2,
					UserID: 4,
					Title:  "event 3",
					Date:   time.Date(2025, time.September, 21, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:     3,
					UserID: 7,
					Title:  "event 4",
					Date:   time.Date(2025, time.October, 29, 0, 0, 0, 0, time.UTC),
				},
			},
			SetUpMocks: func(repo *mockrepo, allEvents []*models.Event) {
				repo.On("GetAll").
					Return(allEvents).Once()
			},
		},
		{
			Name:           "0 events",
			InputTime:      time.Date(2025, time.October, 28, 0, 0, 0, 0, time.UTC),
			ExpectedEvents: nil,
			AllEvents: []*models.Event{
				{
					ID:     0,
					UserID: 3,
					Title:  "event 1",
					Date:   time.Date(2024, time.October, 10, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:     1,
					UserID: 4,
					Title:  "event 2",
					Date:   time.Date(2025, time.February, 25, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:     2,
					UserID: 4,
					Title:  "event 3",
					Date:   time.Date(2025, time.July, 3, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:     3,
					UserID: 7,
					Title:  "event 4",
					Date:   time.Date(2025, time.December, 28, 0, 0, 0, 0, time.UTC),
				},
			},
			SetUpMocks: func(repo *mockrepo, allEvents []*models.Event) {
				repo.On("GetAll").
					Return(allEvents).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			repo := &mockrepo{}
			svc := New(slog.New(
				slog.NewTextHandler(
					os.Stdout,
					nil,
				),
			),
				repo,
			)

			tt.SetUpMocks(repo, tt.AllEvents)

			events := svc.GetForMonth(tt.InputTime)
			assert.Equal(t, tt.ExpectedEvents, events)

			repo.AssertExpectations(t)
		})
	}
}
