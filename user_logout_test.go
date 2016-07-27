package marketwatcher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
	"time"
)

// Configuration
var _ = Describe("UserLogout", func() {
	Describe("should manage user Authentication", func() {
		var page *agouti.Page

		BeforeEach(func() {
			var err error
			page, err = agoutiDriver.NewPage()
			Expect(err).NotTo(HaveOccurred())

		})

		AfterEach(func() {
			Expect(page.Destroy()).To(Succeed())
		})
		// Test Feature

		Describe("When user logged in and user try to logout", func() {
			BeforeEach(func() {
				//login
				Expect(page.Navigate(LoginPage["url"])).To(Succeed())
				Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
				Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
				Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
			})

			It("should user logout succesfully", func() {
				time.Sleep(1000 * time.Millisecond)
				Expect(page.Find(HomePage["logout"]).Click()).To(Succeed())
			})

			It("should user can not see logout button after logged out", func() {
				time.Sleep(1000 * time.Millisecond)
				Expect(page.Find(HomePage["logout"]).Click()).To(Succeed())
				Eventually(page.Find(HomePage["logout"])).ShouldNot(BeFound())
			})
		})
	})
})
