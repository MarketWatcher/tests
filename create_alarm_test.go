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
			Expect(page.Find(AlarmPage["required"]).Fill("\"TW\", \"ThoughtWorks\", \"Thought Works\", \"Thoughtworks\"")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("\"good\", \"best office\"")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("\"sucks\", \"not good enough\"")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.FindByButton(LoginPage["saveButton"]).Submit()).To(Succeed())
		})

		It("should user can see alarm message and alarm header", func() {
			Eventually(page.Find("#message")).Should(HaveText("Your new alarm Alarm 1 is created successfully"))
			Eventually(page.Find("#message_header")).Should(HaveText("Alarm 1"))
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
			Expect(page.Find(AlarmPage["required"]).Fill("\"TW\", \"ThoughtWorks\", \"Thought Works\", \"Thoughtworks\"")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("\"good\", \"best office\"")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("\"sucks\", \"not good enough\"")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.FindByButton(LoginPage["saveButton"]).Submit()).To(Succeed())
		})

		It("should user can see error message on create alarm page", func() {
			Eventually(page.Find("#message")).Should(HaveText("Alarm name cannot be null"))
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
			Expect(page.Find(AlarmPage["required"]).Fill("\"TW\", \"ThoughtWorks\", \"Thought Works\", \"Thoughtworks\"")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.FindByButton(LoginPage["saveButton"]).Submit()).To(Succeed())
		})
		It("should user can see alarm message and alarm header", func() {
			Eventually(page.Find("#message")).Should(HaveText("Your new alarm Alarm 2 is created successfully"))
			Eventually(page.Find("#message_header")).Should(HaveText("Alarm 2"))
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
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("\"good\", \"best office\"")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("\"sucks\", \"not good enough\"")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.FindByButton(LoginPage["saveButton"]).Submit()).To(Succeed())
		})

		It("should user can see error message on create alarm page", func() {
			Eventually(page.Find("#message")).Should(HaveText("Must included criteria cannot be null"))
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

		It("should user create alarm succesfully", func() {
			//alarm 3 without threshold
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("Alarm 3")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill("\"TW\", \"ThoughtWorks\", \"Thought Works\", \"Thoughtworks\"")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("\"good\", \"best office\"")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("\"sucks\", \"not good enough\"")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("0")).To(Succeed())
			Expect(page.FindByButton(LoginPage["saveButton"]).Submit()).To(Succeed())
		})

		It("should user can see error message on create alarm page", func() {
			Eventually(page.Find("#message")).Should(HaveText("Threshold must be a positivite number"))
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

		It("should user create alarm succesfully", func() {
			//alarm 3 without threshold
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("Alarm 3")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill("\"TW\", \"ThoughtWorks\", \"Thought Works\", \"Thoughtworks\"")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("\"good\", \"best office\"")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("\"sucks\", \"not good enough\"")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("-1000")).To(Succeed())
			Expect(page.FindByButton(LoginPage["saveButton"]).Submit()).To(Succeed())
		})

		It("should user can see error message on create alarm page", func() {
			Eventually(page.Find("#message")).Should(HaveText("Threshold must be a positivite number"))
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

		It("should user create alarm succesfully", func() {
			//alarm 3 without threshold
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("Alarm 3")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill("\"TW\", \"ThoughtWorks\", \"Thought Works\", \"Thoughtworks\"")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("\"good\", \"best office\"")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("\"sucks\", \"not good enough\"")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("")).To(Succeed())
			Expect(page.FindByButton(LoginPage["saveButton"]).Submit()).To(Succeed())
		})

		It("should user can see error message on create alarm page", func() {
			Eventually(page.Find("#message")).Should(HaveText("Threshold cannot be null"))
		})
	})

})
