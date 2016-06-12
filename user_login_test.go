package marketwatcher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	//. "github.com/sclevine/agouti/matchers"
)

var _ = Describe("UserLogin", func() {
	var page *agouti.Page

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())

		var loginPage = map[string]string{
			"email":    "loginEmailAddress",
			"password": "exampleInputPassword1",
			"submit":   "Submit",
		}

	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	Describe("User Login Authentication", func() {

		It("should allow the user to dashboard with correct username and password and display username", func() {

			Expect(page.Find(loginPage["email"]).Fill("user1@mail.com"))
			Expect(page.Find(loginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(loginPage["submit"]).Submit()).To(Succeed())
			Eventually(profile.Find(".greeting")).Should(HaveText("user 1"))

		})

		It("should not allowing the user to dashboard with correct username and wrong password", func() {

			Expect(page.Find(loginPage["email"]).Fill("user1@mail.com"))
			Expect(page.Find(loginPage["password"]).Fill("wrongPassword")).To(Succeed())
			Expect(page.FindByButton(loginPage["submit"]).Submit()).To(Succeed())
		})

		It("should not allowing the user to dashboard with wrong username and correct password", func() {

			Expect(page.Find(loginPage["email"]).Fill("notauser@mail.com"))
			Expect(page.Find(loginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(loginPage["submit"]).Submit()).To(Succeed())
		})

		It("should not allowing the user to dashboard with null username and password", func() {

			Expect(page.Find(loginPage["email"]).Fill(" "))
			Expect(page.Find(loginPage["password"]).Fill(" ")).To(Succeed())
			Expect(page.FindByButton(loginPage["submit"]).Submit()).To(Succeed())
		})

	})

})
