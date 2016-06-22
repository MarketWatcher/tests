package marketwatcher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

// Configuration
var _ = Describe("HomePage", func() {
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
	Describe("When user try to visit marketwatcher website", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
		})
		It("should user can see webpage information on the landing page", func() {
			Eventually(page.Find(LoginPage["information"])).Should(HaveText("Market Watcher aims to automate the manual task of analyzing feedback patterns in twitter real-time streams. It minimizes your analysis time to save you time."))
			//Eventually(page.Find(LoginPage["information"])).Should(HaveText("Create your account, log in and build a profile to monitor your trend graphs quickly and easily. You will receive notifications when your search criteria's threshold is passed."))
		})
	})
})
