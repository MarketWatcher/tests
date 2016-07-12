package marketwatcher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

// Configuration
var _ = FDescribe("UserLogin", func() {
	Describe("should manage user Authentication", func() {
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

        It("should redirect the user to the login form from the home page", func() {
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Eventually(page).Should(HaveURL(LoginPage["directedUrl"]))
			// Eventually(page.Find(LoginPage["email"])).Should(BeFound())
		})

        It("should log in user with correct username and password", func() {
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
			Eventually(page.Find(HomePage["userName"])).Should(HaveText("user1@mail.com"))

		})
		It("should log in user with correct username and wrong password", func() {
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("wrongPassword")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
			Eventually(page.Find(LoginPage["errorMessage"])).Should(HaveText("Incorrect email or password"))
		})

        It("should log in user with wrong username and correct password", func() {
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("wrong-username")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
			Eventually(page.Find(LoginPage["errorMessage"])).Should(HaveText("Incorrect email or password"))
		})

        It("should log in user with wrong username and wrong password", func() {
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("userone@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("wrong-password")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
			Eventually(page.Find(LoginPage["errorMessage"])).Should(HaveText("Incorrect email or password"))

		})

        It("should log in user with empty username and password", func() {
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passone")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
			Eventually(page.Find(LoginPage["errorMessage"])).Should(HaveText("There was an error logging in"))
		})

        It("should log in user with invalid username and password", func() {
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("abc")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passone")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
			Eventually(page.Find(LoginPage["errorMessage"])).Should(HaveText("Incorrect email or password"))

		})

        It("should log in user with any valid username and empty password", func() {
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
			Eventually(page.Find(LoginPage["errorMessage"])).Should(HaveText("There was an error logging in"))

		})

        It("should log in user with any empty username and correct password", func() {
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("    ")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
			Eventually(page.Find(LoginPage["errorMessage"])).Should(HaveText("There was an error logging in"))

		})
	})

})
