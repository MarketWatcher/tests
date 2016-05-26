package marketwatcher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

var _ = Describe("UserLogin", func() {
	var page *agouti.Page

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	It("should manage user authentication", func() {
		By("redirecting the user to the login form from the home page", func() {
			Expect(page.Navigate("http://localhost:3000")).To(Succeed())
			Expect(page).To(HaveURL("http://localhost:3000/landing?redirect=%2Fdashboard"))
		})

		By("allowing the user to fill out the login form and submit it", func() {
			Eventually(page.FindByLabel("Email address")).Should(BeFound())
			Expect(page.FindByLabel("Email address").Fill("user1@mail.com")).To(Succeed())
			Expect(page.FindByLabel("Password").Fill("passOne")).To(Succeed())
			Expect(page.FindByButton("Submit").Submit()).To(Succeed())
			Expect(page).To(HaveURL("http://localhost:3000/dashboard"))
		})

	})

})
