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
	"url":         "http://marketwatcher.tech",
	"directedUrl": "http://marketwatcher.tech/landing?redirect=%2Fdashboard",
	"email":       "#email",
	"password":    "#password",
	"submit":      "Submit",
}

var LogoutPage = map[string]string{
	"logout": ".fa.fa-sign-out",
}

var HomePage = map[string]string{
	"userName":     "#user-email > p > b",
	"createButton": "div.i.fa-plus-square-o",
	"viewAlarm":    "",
}

var AlarmPage = map[string]string{
	"alertname":    "#alert-title",
	"required":     "#must-include",
	"niceToHave":   "#can-include",
	"excluded":     "#exclude",
	"threshold":    "#threshold",
	"saveButton":   "#save-alert",
	"cancelButton": "#cancel",
}

var ViewAlarmPage = map[string]string{
	"deleteAlarm": "",
	"popUp":       "",
}

var _ = BeforeSuite(func() {
	//agoutiDriver = agouti.PhantomJS()
	//agoutiDriver = agouti.Selenium()
	agoutiDriver = agouti.ChromeDriver()
	Expect(agoutiDriver.Start()).To(Succeed())
})

var _ = AfterSuite(func() {
	Expect(agoutiDriver.Stop()).To(Succeed())
})
