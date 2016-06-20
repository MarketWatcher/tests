package marketwatcher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

// Configuration
var _ = FDescribe("UserLogin", func() {
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
	It("should direct to login page", func() {
		By("redirecting the user to the login form from the home page", func() {
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Eventually(page).Should(HaveURL(LoginPage["directedUrl"]))
		})

	})

	It("should manage user Authentication", func() {
		By("loging in user with correct username and password", func() {
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
			//Expect(page).To(HaveTitle("div.collapse.navbar-collapse"))
			Eventually(page.Find(HomePage["userName"])).Should(HaveText("user1@mail.com"))

		})

	})

	It("should manage user Authentication", func() {
		By("loging in user with correct username and wrong password", func() {
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("wrongPassword")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
			Eventually(page.Find("div.has-error")).Should(HaveText("Username and password did not match"))

		})

	})

	It("should manage user Authentication", func() {
		By("loging in user with wrong username and correct password", func() {
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("userone@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
			Eventually(page.Find("div.has-error")).Should(HaveText("Username and password did not match"))
			//Expect(page).To(HaveTitle("div.collapse.navbar-collapse"))
		})
	})

	It("should manage user Authentication", func() {
		By("loging in user with wrong username and wrong password", func() {
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("userone@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passone")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
			Eventually(page.Find("div.has-error")).Should(HaveText("Username and password did not match"))

		})
	})

	It("should manage user Authentication", func() {
		By("loging in user with empty username and password", func() {
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passone")).To(Succeed())
			//Expect(page.Find(LoginPage["password"]).Fill("wrongPassword")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
			Eventually(page.Find("div.has-error")).Should(HaveText("Username and password did not match"))

		})
	})

	It("should manage user Authentication", func() {
		By("loging in user with invalid username and password", func() {
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("abc")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passone")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
			Eventually(page.Find("div.has-error")).Should(HaveText("Username and password did not match"))

		})
	})

	It("should manage user Authentication", func() {
		By("loging in user with any valid username and empty password", func() {
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
			Eventually(page.Find("div.has-error")).Should(HaveText("Username and password did not match"))

		})
	})

	It("should manage user Authentication", func() {
		By("loging in user with any empty username and correct password", func() {
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("    ")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
			Eventually(page.Find("div.has-error")).Should(HaveText("Username and password did not match"))

		})
	})
})
