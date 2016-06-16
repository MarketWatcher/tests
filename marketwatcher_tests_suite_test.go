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
	"userName":     "#user-email.b",
	"createButton": "div.i.fa-plus-square-o",
}

var AlarmPage = map[string]string{
	"alertname":  "#alert-name-input",
	"required":   "#required-input",
	"niceToHave": "#nice-to-have-input",
	"excluded":   "#excluded-input",
	"threshold":  "#threshold",
	"saveButton": "#save-button",
	"alarmOne":   "alarm1",
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
