package marketwatcher

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	//. "github.com/sclevine/agouti/matchers"
)

// Configuration
var _ = Describe("CreateAlarm", func() {

	Describe("When user logged in and user try to create alarm succesfully", func() {
		var page *agouti.Page

		BeforeEach(func() {
			var err error
			page, err = agoutiDriver.NewPage()
			Expect(err).NotTo(HaveOccurred())
			//TODO: fix login page on front end
			page.Size(1000, 800)
			page.Reset()
		})

		AfterEach(func() {
			Expect(page.Destroy()).To(Succeed())
		})

		//Test Feature

		It("should user can not create alarm succesfully with uppercase and lowercase", func() {
			//alarm 1
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			//Expect(page.Find(AlarmPage["alertname"]).Fill("ğüğççççööşOffLGDFV")).To(Succeed())
			//Expect(page.Find(AlarmPage["required"]).Fill(" TW, ThoughtWorks, Thought Works, Thoughtworks")).To(Succeed())
			//Expect(page.Find(AlarmPage["niceToHave"]).Fill("good, best office")).To(Succeed())
			//Expect(page.Find(AlarmPage["excluded"]).Fill("sucks, not good enough")).To(Succeed())
			//Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
			//Expect(page.Find(AlarmPage["saveButton"]).Click()).To(Succeed())
			//Eventually(page.Find(HomePage["errorMessage"])).Should(HaveText("There was an error creating the alert"))
		})
	})
})
