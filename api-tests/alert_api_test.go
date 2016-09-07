package api

import (
	"net/http"
	"testing"
	"github.com/gavv/httpexpect"
	"os"
	"time"
	"math/rand"
)

func init() {
	rand.Seed(time.Now().UnixNano())
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
		"name": RandomWord(10),
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
		"name": RandomWord(10),
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

func TestGetAlertWithAlertId(t *testing.T){
	name := RandomWord(10)
	alert := map[string]interface{}{
		"owner_id":2,
		"name": name,
		"required_criteria":"aab",
		"nice_to_have_criteria":"ffe",
		"excluded_criteria":"ddr",
		"threshold":555,
		"status":23,
	}
	e := httpexpect.New(t, os.Getenv("ALERT_SERVICE_URL"))

	savedAlert := e.POST("/api/alerts").WithJSON(alert).
		Expect().
		Status(http.StatusOK).
		JSON().Object()

	getById := 	"/api/alerts/id/" + savedAlert.Value("id").Raw().(string)
	e.GET(getById).
	Expect().
	Status(http.StatusOK).
	JSON().Object().
	ContainsKey("name").
	ValueEqual("name", name)
}

func TestAlertCreationWithEmptyMust(t *testing.T) {

	alert := map[string]interface{} {
		"owner_id":1,
		"name": RandomWord(10),
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

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomWord(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}
