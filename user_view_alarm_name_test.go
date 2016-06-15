package marketwatcher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

// Configuration
var _ = Describe("UserLogout", func() {
	var page *agouti.Page

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())

	})

	var loginPage = map[string]string{
		"email":    "#email",
		"password": "#password",
		"submit":   "Submit",
	}

	var homePage = map[string]string{
		"createButton": "div.i.fa-plus-square-o",
	}

	var alarmPage = map[string]string{
		"alertname":  "#alert-name-input",
		"required":   "#required-input",
		"niceToHave": "#nice-to-have-input",
		"excluded":   "#excluded-input",
		"threshold":  "#threshold",
		"saveButton": "#save-button",
		"alarmOne":   "alarm1",
	}

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	Describe("When user logged in and user try to create alarm succesfully", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate("http://marketwatcher.tech/")).To(Succeed())
			Expect(page.Find(loginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(loginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(loginPage["submit"]).Submit()).To(Succeed())
		})

		It("should user create alarm succesfully", func() {
			//alarm 1
			Expect(page.Find(homePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(alarmPage["alertname"]).Fill("Alarm 1")).To(Succeed())
			Expect(page.Find(alarmPage["required"]).Fill("\"TW\", \"ThoughtWorks\", \"Thought Works\", \"Thoughtworks\"")).To(Succeed())
			Expect(page.Find(alarmPage["niceToHave"]).Fill("\"good\", \"best office\"")).To(Succeed())
			Expect(page.Find(alarmPage["excluded"]).Fill("\"sucks\", \"not good enough\"")).To(Succeed())
			Expect(page.Find(alarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.FindByButton(loginPage["saveButton"]).Submit()).To(Succeed())
		})

		It("should user can see alarm header in alarm detail page", func() {
			Eventually(page.Find("div.has-error")).Should(HaveText("Alarm 1"))
		})

	})
})
