package marketwatcher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

// Configuration
var _ = FDescribe("CreateAlarm", func() {
	var page *agouti.Page

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())

	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	//Test Feature

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
			Expect(page.Find(AlarmPage["required"]).Fill(" TW, ThoughtWorks, Thought Works, Thoughtworks")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("good, best office")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("sucks, not good enough")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.Find(AlarmPage["saveButton"]).Click()).To(Succeed())
			Eventually(page.Find(HomePage["successMessage"])).Should(HaveText("Alert was created"))
		})

	})

	Describe("When user logged in and user try to create alarm without header", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
		})

		It("should user can not create alarm successfully without header", func() {
			//empty alarm
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill(" TW, ThoughtWorks, Thought Works, Thoughtworks")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("good, best office")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("sucks, not good enough")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.Find(AlarmPage["saveButton"]).Click()).To(Succeed())
			Eventually(page.Find(AlarmPage["alertnameError"])).Should(HaveText("\"Alert title\" length must be at least 3 characters long"))
		})

	})

	Describe("When user logged in and user try to create another alarm without nice to have and excluded", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
		})

		It("should user create alarm successfully without nice to have and excluded ", func() {
			//alarm 2
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("Alarm 2")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill(" TW, ThoughtWorks, Thought Works, Thoughtworks")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.Find(AlarmPage["saveButton"]).Click()).To(Succeed())
			Eventually(page.Find(HomePage["successMessage"])).Should(HaveText("Alert was created"))
		})

	})

	Describe("When user logged in and user try to create another alarm without required", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
		})

		It("should user can not create alarm successfully without required", func() {
			//alarm 3 without required
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("Alarm 3")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill("")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("good, best office")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("sucks, not good enough")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.Find(AlarmPage["saveButton"]).Click()).To(Succeed())
			Eventually(page.Find(AlarmPage["requiredError"])).Should(HaveText("\"Must include\" is not allowed to be empty"))
		})

	})

	Describe("When user logged in and user try to create alarm without threshold", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
		})

		It("should user can not create alarm succesfully with negative threshold", func() {
			//alarm 3 without threshold
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("Alarm 3")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill(" TW, ThoughtWorks, Thought Works, Thoughtworks")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("good, best office")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("sucks, not good enough")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("0")).To(Succeed())
			Expect(page.Find(AlarmPage["saveButton"]).Click()).To(Succeed())
			Eventually(page.Find(AlarmPage["thresholdError"])).Should(HaveText("\"Threshold\" must be larger than or equal to 1"))
		})

	})

	Describe("When user logged in and user try to create alarm with negative threshold", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
		})

		It("should user can not create alarm succesfully without number", func() {
			//alarm 3 without threshold
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("Alarm 3")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill(" TW, ThoughtWorks, Thought Works, Thoughtworks")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("good, best office")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("sucks, not good enough")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("-")).To(Succeed())
			Expect(page.Find(AlarmPage["saveButton"]).Click()).To(Succeed())
			Eventually(page.Find(AlarmPage["thresholdError"])).Should(HaveText("\"Threshold\" must be a number"))
		})
	})

	Describe("When user logged in and user try to create alarm with empty threshold", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
		})

		It("should user can not create alarm succesfully without empty threshold", func() {
			//alarm 3 without threshold
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("Alarm 3")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill(" TW, ThoughtWorks, Thought Works, Thoughtworks")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("good, best office")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("sucks, not good enough")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("")).To(Succeed())
			Expect(page.Find(AlarmPage["saveButton"]).Click()).To(Succeed())
			Eventually(page.Find(AlarmPage["thresholdError"])).Should(HaveText("\"Threshold\" must be a number"))
		})

	})

})
