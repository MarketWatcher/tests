package marketwatcher_test

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"os"
	"fmt"
	"net/http/httptest"
)

func TestMain(m *testing.M) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Success!")
	})

	// run server using httptest
	server := httptest.NewServer(handler)
	os.Setenv("DATA_INGESTION_URL", server.URL)
}

func TestAlertCreationWithEmptyRequestBody(t *testing.T) {
	e := httpexpect.New(t, os.Getenv("ALERT_SERVICE_URL"))

	e.POST("/api/alerts").
		Expect().
		Status(http.StatusBadRequest)
}

func TestAlertCreationWithEmptyTitle(t *testing.T) {

	alert := map[string]interface{}{
		"owner_id":1,
		"required_criteria":"aaa",
		"nice_to_have_criteria":"fff",
		"excluded_criteria":"ddd",
		"threshold":555,
		"status":1,

	}
	e := httpexpect.New(t, os.Getenv("ALERT_SERVICE_URL"))

	e.POST("/api/alerts").WithJSON(alert).
		Expect().
		Status(http.StatusBadRequest)
}

func TestSuccesfulAlertCreation(t *testing.T) {
	alert := map[string]interface{}{
		"owner_id":1,
		"name": "namee",
		"required_criteria":"aaa",
		"nice_to_have_criteria":"fff",
		"excluded_criteria":"ddd",
		"threshold":555,
		"status":1,
	}

	e := httpexpect.New(t, os.Getenv("ALERT_SERVICE_URL"))

	e.POST("/api/alerts").WithJSON(alert).
		Expect().
		Status(http.StatusOK)

}

func TestAlertCreationWithEmptyThreshold(t *testing.T) {
	alert := map[string]interface{}{
		"owner_id":1,
		"name": "namee",
		"required_criteria":"aaa",
		"nice_to_have_criteria":"fff",
		"excluded_criteria":"ddd",
		"status":1,
	}

	e := httpexpect.New(t, os.Getenv("ALERT_SERVICE_URL"))

	e.POST("/api/alerts").WithJSON(alert).
		Expect().
		Status(http.StatusBadRequest)
}

func TestAlertCreationWithEmptyMust(t *testing.T) {

	alert := map[string]interface{}{
		"owner_id":1,
		"name": "namee",
		"nice_to_have_criteria":"fff",
		"excluded_criteria":"ddd",
		"threshold":555,
		"status":1,
	}

	e := httpexpect.New(t, os.Getenv("ALERT_SERVICE_URL"))
	e.POST("/api/alerts").WithJSON(alert).
		Expect().
		Status(http.StatusBadRequest)

}
