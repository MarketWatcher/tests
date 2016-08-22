package marketwatcher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
	"time"
)

// Configuration
var _ = Describe("CreateAlarm", func() {
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
			time.Sleep(1000 * time.Millisecond)
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("Alarm 1")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill(" TW, ThoughtWorks, Thought Works, Thoughtworks")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("good, best office")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("sucks, not good enough")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.Find(AlarmPage["saveButton"]).Click()).To(Succeed())
			Eventually(page.Find(HomePage["successMessage"])).Should(HaveText("Alert was created"))
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
			time.Sleep(1000 * time.Millisecond)
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill("TW, ThoughtWorks, Thought Works, Thoughtworks")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("good, best office")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("sucks, not good enough")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.Find(AlarmPage["saveButton"]).Click()).To(Succeed())
			Eventually(page.Find(AlarmPage["alertnameError"])).Should(HaveText("\"Alert title\" length must be at least 3 characters long"))
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
			time.Sleep(1000 * time.Millisecond)
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("Alarm 2")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill(" TW, ThoughtWorks, Thought Works, Thoughtworks")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.Find(AlarmPage["saveButton"]).Click()).To(Succeed())
			Eventually(page.Find(HomePage["successMessage"])).Should(HaveText("Alert was created"))
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
			time.Sleep(1000 * time.Millisecond)
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("Alarm 3")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill("")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("good, best office")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("sucks, not good enough")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.Find(AlarmPage["saveButton"]).Click()).To(Succeed())
			Eventually(page.Find(AlarmPage["requiredError"])).Should(HaveText("\"Must include\" is not allowed to be empty"))
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

		It("should user can not create alarm succesfully with negative threshold", func() {
			//alarm 3 without threshold
			time.Sleep(1000 * time.Millisecond)
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("Alarm 3")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill(" TW, ThoughtWorks, Thought Works, Thoughtworks")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("good, best office")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("sucks, not good enough")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("0")).To(Succeed())
			Expect(page.Find(AlarmPage["saveButton"]).Click()).To(Succeed())
			Eventually(page.Find(AlarmPage["thresholdError"])).Should(HaveText("\"Threshold\" must be larger than or equal to 1"))
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

		It("should user can not create alarm succesfully without number", func() {
			//alarm 3 without threshold
			time.Sleep(1000 * time.Millisecond)
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("Alarm 3")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill(" TW, ThoughtWorks, Thought Works, Thoughtworks")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("good, best office")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("sucks, not good enough")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("-")).To(Succeed())
			Expect(page.Find(AlarmPage["saveButton"]).Click()).To(Succeed())
			Eventually(page.Find(AlarmPage["thresholdError"])).Should(HaveText("\"Threshold\" must be a number"))
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

		It("should user can not create alarm succesfully without empty threshold", func() {
			//alarm 3 without threshold
			time.Sleep(1000 * time.Millisecond)
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("Alarm 3")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill(" TW, ThoughtWorks, Thought Works, Thoughtworks")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("good, best office")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("sucks, not good enough")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("")).To(Succeed())
			Expect(page.Find(AlarmPage["saveButton"]).Click()).To(Succeed())
			Eventually(page.Find(AlarmPage["thresholdError"])).Should(HaveText("\"Threshold\" must be a number"))
		})

	})

	Describe("When user logged in and user try to create alarm and cancel it", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
		})

		It("should user can not create alarm with cancel button", func() {
			time.Sleep(1000 * time.Millisecond)
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("Alarm 3")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill(" TW, ThoughtWorks, Thought Works, Thoughtworks")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("good, best office")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("sucks, not good enough")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("")).To(Succeed())
			Expect(page.Find(AlarmPage["cancelButton"]).Click()).To(Succeed())
			Eventually(page.Find(HomePage["createButton"]))
		})

	})

	Describe("When user logged in and usertry to create alarm succesfully", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
		})

		It("should user can not create alarm with nonalphanumaric alert header", func() {
			//alarm 1
			time.Sleep(1000 * time.Millisecond)
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("%&^*^&*&*8")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill(" TW, ThoughtWorks, Thought Works, Thoughtworks")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("good, best office")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("sucks, not good enough")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.Find(AlarmPage["saveButton"]).Click()).To(Succeed())
			Eventually(page.Find(AlarmPage["alertnameError"])).Should(HaveText("\"Alert title\" must be alphanumeric"))
		})

	})

	Describe("When user logged in and usertry to create alarm succesfully", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
		})

		It("should user can not create alarm with more than 32 unicode characters alert header", func() {
			//alarm 1
			time.Sleep(1000 * time.Millisecond)
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("fnserkjgfnsekrgjnvkesjrngvkserjngvs")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill(" TW, ThoughtWorks, Thought Works, Thoughtworks")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("good, best office")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("sucks, not good enough")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.Find(AlarmPage["saveButton"]).Click()).To(Succeed())
			Eventually(page.Find(AlarmPage["alertnameError"])).Should(HaveText("\"Alert title\" length must be less than or equal to 32 characters long"))
		})

	})
	Describe("When user logged in and user try to create alarm succesfully", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
		})

		It("should user create alarm succesfully with turkish characters in header", func() {
			//alarm 1
			time.Sleep(1000 * time.Millisecond)
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("ğüğççççööş")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill(" TW, ThoughtWorks, Thought Works, Thoughtworks")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("good, best office")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("sucks, not good enough")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.Find(AlarmPage["saveButton"]).Click()).To(Succeed())
			Eventually(page.Find(HomePage["successMessage"])).Should(HaveText("Alert was created"))
		})
	})

	Describe("When user logged in and user try to create alarm succesfully", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
		})

		It("should user can not create alarm succesfully with more than 140 characterts for required field", func() {
			//alarm 1
			time.Sleep(1000 * time.Millisecond)
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("Alert 1")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill("fnserkjgfnsekrgjnvkesjrngvkserjngvsdfjcdnsfkjsndfknsdvndjvbdjfbdjsbfgdjshgbejshbgjfnkksjfnekjsfnejrfnenjfneefjnehrbfejrhgberhjgwdfjebdjeshfbe")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("good, best office")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("sucks, not good enough")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.Find(AlarmPage["saveButton"]).Click()).To(Succeed())
			Eventually(page.Find(AlarmPage["requiredError"])).Should(HaveText("\"Must include\" length must be less than or equal to 140 characters long"))
		})
	})
	Describe("When user logged in and user try to create alarm succesfully", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
		})

		It("should user can not create alarm succesfully with non unicode required", func() {
			//alarm 1
			time.Sleep(1000 * time.Millisecond)
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("Alert 1")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill("%$^%^&^*&^*^")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("good, best office")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("sucks, not good enough")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.Find(AlarmPage["saveButton"]).Click()).To(Succeed())
			Eventually(page.Find(AlarmPage["requiredError"])).Should(HaveText("\"Must include\" must be alphanumeric"))
		})
	})

	Describe("When user logged in and user try to create alarm succesfully", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
		})

		It("should user can not create alarm succesfully with non unicode characterts for niceToHave", func() {
			//alarm 1
			time.Sleep(1000 * time.Millisecond)
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("Alarm 1")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill(" TW, ThoughtWorks, Thought Works, Thoughtworks")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("%^&^&*&^*")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("sucks, not good enough")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.Find(AlarmPage["saveButton"]).Click()).To(Succeed())
			Eventually(page.Find(AlarmPage["niceToHaveError"])).Should(HaveText("\"Can include\" must be alphanumeric"))
		})

	})

	Describe("When user logged in and user try to create alarm succesfully", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
		})

		It("should user can not create alarm succesfully with 140 characterts for niceToHave", func() {
			//alarm 1
			time.Sleep(1000 * time.Millisecond)
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("Alarm 1")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill(" TW, ThoughtWorks, Thought Works, Thoughtworks")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("fnserkjgfnsekrgjnvkesjrngvkserjngvsdfjcdnsfkjsndfknsdvndjvbdjfbdjsbfgdjshgbejshbgjfnkksjfnekjsfnejrfnenjfneefjnehrbfejrhgberhjgwdfjebdjeshfbe")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("sucks, not good enough")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.Find(AlarmPage["saveButton"]).Click()).To(Succeed())
			Eventually(page.Find(AlarmPage["niceToHaveError"])).Should(HaveText("\"Can include\" length must be less than or equal to 140 characters long"))
		})

	})
	Describe("When user logged in and user try to create alarm succesfully", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
		})

		It("should user can not create alarm succesfully with 140 characterts for excluded", func() {
			//alarm 1
			time.Sleep(1000 * time.Millisecond)
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("Alarm 1")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill(" TW, ThoughtWorks, Thought Works, Thoughtworks")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("nice, good")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("fnserkjgfnsekrgjnvkesjrngvkserjngvsdfjcdnsfkjsndfknsdvndjvbdjfbdjsbfgdjshgbejshbgjfnkksjfnekjsfnejrfnenjfneefjnehrbfejrhgberhjgwdfjebdjeshfbe")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.Find(AlarmPage["saveButton"]).Click()).To(Succeed())
			Eventually(page.Find(AlarmPage["excludedError"])).Should(HaveText("\"Exclude criteria\" length must be less than or equal to 140 characters long"))

		})
	})
	Describe("When user logged in and user try to create alarm succesfully", func() {
		BeforeEach(func() {
			//login
			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
		})

		It("should user can not create alarm succesfully with non unicode characterts for excluded", func() {
			//alarm 1
			time.Sleep(1000 * time.Millisecond)
			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
			Expect(page.Find(AlarmPage["alertname"]).Fill("Alarm 1")).To(Succeed())
			Expect(page.Find(AlarmPage["required"]).Fill(" TW, ThoughtWorks, Thought Works, Thoughtworks")).To(Succeed())
			Expect(page.Find(AlarmPage["niceToHave"]).Fill("nice, good")).To(Succeed())
			Expect(page.Find(AlarmPage["excluded"]).Fill("%^&%&%^&*^&*")).To(Succeed())
			Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
			Expect(page.Find(AlarmPage["saveButton"]).Click()).To(Succeed())
			Eventually(page.Find(AlarmPage["excludedError"])).Should(HaveText("\"Exclude criteria\" must be alphanumeric"))

		})
	})

})
