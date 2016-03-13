package metadata

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/tracer/mock/mux"
	"github.com/tracer/mock/tracers/metadata"
)

func TestRegister(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRouter := mock_mux.NewMockRouter(mockController)
	mockRouter.EXPECT().POST("/tracers/{id}/metadata", gomock.Any())
	Register(mockRouter, nil)
}

func Test_handleMetadata(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	// The parameters to pass to handleMetadata.
	req, err := http.NewRequest("POST", "/tracers/0/metadata", strings.NewReader("{}"))
	if err != nil {
		t.Errorf("Request failed.")
	}
	w := httptest.NewRecorder()
	m := make(map[string]string)
	m["id"] = "0"

	mockRouter := mock_mux.NewMockRouter(mockController)
	mockRouter.EXPECT().Vars(req).Return(m)

	mockStorage := mock_metadata.NewMockStorage(mockController)
	mockStorage.EXPECT().Store("0", "{}")
	// Create a private router instance defined in metadata.go.
	r := router{mockRouter, mockStorage}

	r.handleMetadata(w, req)
}
