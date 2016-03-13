package tracers

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/tracer/mock/mux"
)

func TestRegister(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRouter := mock_mux.NewMockRouter(mockController)
	mockRouter.EXPECT().POST("/tracers", gomock.Any())
	Register(mockRouter)
}
