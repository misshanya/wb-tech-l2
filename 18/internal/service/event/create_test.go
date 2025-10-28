package event

import (
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/misshanya/wb-tech-l2/18/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestService_Create(t *testing.T) {
	tests := []struct {
		Name          string
		InputEvent    *models.Event
		ExpectedEvent *models.Event
		SetUpMocks    func(repo *mockrepo, inputEvent *models.Event, expectedEvent *models.Event)
	}{
		{
			Name: "Successfully created",
			InputEvent: &models.Event{
				UserID: 2,
				Title:  "Some event",
				Date:   time.Date(2025, time.October, 28, 0, 0, 0, 0, time.UTC),
			},
			ExpectedEvent: &models.Event{
				ID:    6,
				Title: "Some event",
				Date:  time.Date(2025, time.October, 28, 0, 0, 0, 0, time.UTC),
			},
			SetUpMocks: func(repo *mockrepo, inputEvent *models.Event, expectedEvent *models.Event) {
				repo.On("Create", inputEvent).
					Return(expectedEvent).Once()
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

			tt.SetUpMocks(repo, tt.InputEvent, tt.ExpectedEvent)

			res := svc.Create(tt.InputEvent)
			assert.Equal(t, tt.ExpectedEvent, res)

			repo.AssertExpectations(t)
		})
	}
}
