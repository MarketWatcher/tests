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
	}

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	//Test Feature

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

		It("should user can see alarm message and alarm header", func() {
			Eventually(page.Find("#message")).Should(HaveText("Your new alarm Alarm 1 is created successfully"))
			Eventually(page.Find("#message_header")).Should(HaveText("Alarm 1"))
		})

	})

	Describe("When user logged in and user try to create alarm without header", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate("http://marketwatcher.tech/")).To(Succeed())
			Expect(page.Find(loginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(loginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(loginPage["submit"]).Submit()).To(Succeed())
		})

		It("should user can not create alarm successfully without header", func() {
			//empty alarm
			Expect(page.Find(homePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(alarmPage["alertname"]).Fill("")).To(Succeed())
			Expect(page.Find(alarmPage["required"]).Fill("\"TW\", \"ThoughtWorks\", \"Thought Works\", \"Thoughtworks\"")).To(Succeed())
			Expect(page.Find(alarmPage["niceToHave"]).Fill("\"good\", \"best office\"")).To(Succeed())
			Expect(page.Find(alarmPage["excluded"]).Fill("\"sucks\", \"not good enough\"")).To(Succeed())
			Expect(page.Find(alarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.FindByButton(loginPage["saveButton"]).Submit()).To(Succeed())
		})

		It("should user can see error message on create alarm page", func() {
			Eventually(page.Find("#message")).Should(HaveText("Alarm name cannot be null"))
		})
	})

	Describe("When user logged in and user try to create another alarm without nice to have and excluded", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate("http://marketwatcher.tech/")).To(Succeed())
			Expect(page.Find(loginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(loginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(loginPage["submit"]).Submit()).To(Succeed())
		})

		It("should user create alarm successfully without nice to have and excluded ", func() {
			//alarm 2
			Expect(page.Find(homePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(alarmPage["alertname"]).Fill("Alarm 2")).To(Succeed())
			Expect(page.Find(alarmPage["required"]).Fill("\"TW\", \"ThoughtWorks\", \"Thought Works\", \"Thoughtworks\"")).To(Succeed())
			Expect(page.Find(alarmPage["niceToHave"]).Fill("")).To(Succeed())
			Expect(page.Find(alarmPage["excluded"]).Fill("")).To(Succeed())
			Expect(page.Find(alarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.FindByButton(loginPage["saveButton"]).Submit()).To(Succeed())
		})
		It("should user can see alarm message and alarm header", func() {
			Eventually(page.Find("#message")).Should(HaveText("Your new alarm Alarm 2 is created successfully"))
			Eventually(page.Find("#message_header")).Should(HaveText("Alarm 2"))
		})

	})

	Describe("When user logged in and user try to create another alarm without required", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate("http://marketwatcher.tech/")).To(Succeed())
			Expect(page.Find(loginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(loginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(loginPage["submit"]).Submit()).To(Succeed())
		})

		It("should user can not create alarm successfully without required", func() {
			//alarm 3 without required
			Expect(page.Find(homePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(alarmPage["alertname"]).Fill("Alarm 3")).To(Succeed())
			Expect(page.Find(alarmPage["required"]).Fill("")).To(Succeed())
			Expect(page.Find(alarmPage["niceToHave"]).Fill("\"good\", \"best office\"")).To(Succeed())
			Expect(page.Find(alarmPage["excluded"]).Fill("\"sucks\", \"not good enough\"")).To(Succeed())
			Expect(page.Find(alarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.FindByButton(loginPage["saveButton"]).Submit()).To(Succeed())
		})

		It("should user can see error message on create alarm page", func() {
			Eventually(page.Find("#message")).Should(HaveText("Must included criteria cannot be null"))
		})

	})

	Describe("When user logged in and user try to create alarm without threshold", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate("http://marketwatcher.tech/")).To(Succeed())
			Expect(page.Find(loginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(loginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(loginPage["submit"]).Submit()).To(Succeed())
		})

		It("should user create alarm succesfully", func() {
			//alarm 3 without threshold
			Expect(page.Find(homePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(alarmPage["alertname"]).Fill("Alarm 3")).To(Succeed())
			Expect(page.Find(alarmPage["required"]).Fill("\"TW\", \"ThoughtWorks\", \"Thought Works\", \"Thoughtworks\"")).To(Succeed())
			Expect(page.Find(alarmPage["niceToHave"]).Fill("\"good\", \"best office\"")).To(Succeed())
			Expect(page.Find(alarmPage["excluded"]).Fill("\"sucks\", \"not good enough\"")).To(Succeed())
			Expect(page.Find(alarmPage["threshold"]).Fill("0")).To(Succeed())
			Expect(page.FindByButton(loginPage["saveButton"]).Submit()).To(Succeed())
		})

		It("should user can see error message on create alarm page", func() {
			Eventually(page.Find("#message")).Should(HaveText("Threshold must be a positivite number"))
		})
	})

	Describe("When user logged in and user try to create alarm with negative threshold", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate("http://marketwatcher.tech/")).To(Succeed())
			Expect(page.Find(loginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(loginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(loginPage["submit"]).Submit()).To(Succeed())
		})

		It("should user create alarm succesfully", func() {
			//alarm 3 without threshold
			Expect(page.Find(homePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(alarmPage["alertname"]).Fill("Alarm 3")).To(Succeed())
			Expect(page.Find(alarmPage["required"]).Fill("\"TW\", \"ThoughtWorks\", \"Thought Works\", \"Thoughtworks\"")).To(Succeed())
			Expect(page.Find(alarmPage["niceToHave"]).Fill("\"good\", \"best office\"")).To(Succeed())
			Expect(page.Find(alarmPage["excluded"]).Fill("\"sucks\", \"not good enough\"")).To(Succeed())
			Expect(page.Find(alarmPage["threshold"]).Fill("-1000")).To(Succeed())
			Expect(page.FindByButton(loginPage["saveButton"]).Submit()).To(Succeed())
		})

		It("should user can see error message on create alarm page", func() {
			Eventually(page.Find("#message")).Should(HaveText("Threshold must be a positivite number"))
		})
	})

	Describe("When user logged in and user try to create alarm with empty threshold", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate("http://marketwatcher.tech/")).To(Succeed())
			Expect(page.Find(loginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(loginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(loginPage["submit"]).Submit()).To(Succeed())
		})

		It("should user create alarm succesfully", func() {
			//alarm 3 without threshold
			Expect(page.Find(homePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(alarmPage["alertname"]).Fill("Alarm 3")).To(Succeed())
			Expect(page.Find(alarmPage["required"]).Fill("\"TW\", \"ThoughtWorks\", \"Thought Works\", \"Thoughtworks\"")).To(Succeed())
			Expect(page.Find(alarmPage["niceToHave"]).Fill("\"good\", \"best office\"")).To(Succeed())
			Expect(page.Find(alarmPage["excluded"]).Fill("\"sucks\", \"not good enough\"")).To(Succeed())
			Expect(page.Find(alarmPage["threshold"]).Fill("")).To(Succeed())
			Expect(page.FindByButton(loginPage["saveButton"]).Submit()).To(Succeed())
		})

		It("should user can see error message on create alarm page", func() {
			Eventually(page.Find("#message")).Should(HaveText("Threshold cannot be null"))
		})
	})

})
