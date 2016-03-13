package metadata

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/tracer/mock/mux"
)

func TestRegister(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRouter := mock_mux.NewMockRouter(mockController)
	mockRouter.EXPECT().POST("/tracers/{id}/metadata", gomock.Any())
	Register(mockRouter)
}
