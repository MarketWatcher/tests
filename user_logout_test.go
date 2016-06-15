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

	var logoutPage = map[string]string{
		"logout": ".fa.fa-sign-out",
	}

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})
	// Test Feature

	Describe("When user logged in and user try to logout", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate("http://marketwatcher.tech/")).To(Succeed())
			Expect(page.Find(loginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(loginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(loginPage["submit"]).Submit()).To(Succeed())
		})

		It("should user logout succesfully", func() {
			Expect(page.Find(logoutPage["logout"]).Click()).To(Succeed())
		})

		It("should user can not see logout button after logged out", func() {
			Expect(page.Find(logoutPage["logout"]).Click()).To(Succeed())
			Eventually(page.Find(logoutPage["logout"])).ShouldNot(BeFound())
		})
	})
})
