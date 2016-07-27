package marketwatcher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	"testing"
)

func TestMarketwatcherTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MarketwatcherTests Suite")
}

var agoutiDriver *agouti.WebDriver

var LoginPage = map[string]string{
	"url":          "http://marketwatcher.tech",
	"directedUrl":  "http://marketwatcher.tech/landing",
	"email":        "#email",
	"password":     "#password",
	"submit":       "Submit",
	"information":  "#intro-text > :nth-child(2)",
	"errorMessage": "#error-text ",
}

var HomePage = map[string]string{
	"userName":       "#user-email > p > b",
	"createButton":   "#create-alert",
	"successMessage": ".site-notification-text",
	"viewAlarmTitle": "",
	"logout":         "#logout",
}

var AlarmPage = map[string]string{
	"alertname":       "#alert-title",
	"alertnameError":  "#name-error",
	"required":        "#must-include",
	"requiredError":   "#requiredCriteria-error",
	"niceToHave":      "#can-include",
	"niceToHaveError": "#niceTohaveCriteria-error",
	"excluded":        "#exclude",
	"excludedError":   "#excludedCriteria-error",
	"threshold":       "#threshold",
	"thresholdError":  "#threshold-error",
	"saveButton":      "#save-alert",
	"cancelButton":    "#cancel",
}

var ViewAlarmPage = map[string]string{
	"deleteAlarm": "",
	"popUp":       "",
	"alarmTitle":  "div.card h4.title",
}

var _ = BeforeSuite(func() {
	// agoutiDriver = agouti.PhantomJS()
	//agoutiDriver = agouti.Selenium()
	agoutiDriver = agouti.ChromeDriver()
	Expect(agoutiDriver.Start()).To(Succeed())
})

var _ = AfterSuite(func() {
	Expect(agoutiDriver.Stop()).To(Succeed())
})
