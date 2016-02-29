package tracers

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/tracer/mock"
)

func TestRegister(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRouter := mock_router.NewMockRouter(mockController)
	mockRouter.EXPECT().HandleFunc("/tracers/", gomock.Any())
	Register(mockRouter)
}
