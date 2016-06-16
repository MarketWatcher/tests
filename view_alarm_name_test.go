package marketwatcher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

// Configuration
var _ = Describe("ViewAlarmPage", func() {
	var page *agouti.Page

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())

	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	Describe("When user logged in and user try to create alarm succesfully", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
		})

		It("should user create alarm succesfully", func() {
			//alarm 1
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("Alarm 1")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill("\"TW\", \"ThoughtWorks\", \"Thought Works\", \"Thoughtworks\"")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("\"good\", \"best office\"")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("\"sucks\", \"not good enough\"")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.FindByButton(LoginPage["saveButton"]).Submit()).To(Succeed())
		})

		It("should user can see alarm header in alarm detail page", func() {
			Eventually(page.Find("div.has-error")).Should(HaveText("Alarm 1"))
		})

	})
})
