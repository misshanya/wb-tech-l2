package event

import (
	"log/slog"
	"os"
	"testing"
)

func TestService_Delete(t *testing.T) {
	tests := []struct {
		Name       string
		InputID    int
		SetUpMocks func(repo *mockrepo, id int)
	}{
		{
			Name:    "Successfully deleted",
			InputID: 6,
			SetUpMocks: func(repo *mockrepo, id int) {
				repo.On("Delete", id).Once()
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

			tt.SetUpMocks(repo, tt.InputID)

			svc.Delete(tt.InputID)

			repo.AssertExpectations(t)
		})
	}
}
