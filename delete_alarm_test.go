package marketwatcher
//
// import (
// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"
// 	"github.com/sclevine/agouti"
// 	. "github.com/sclevine/agouti/matchers"
// )
//
// // Configuration
// var _ = FDescribe("DeleteAlarm", func() {
// 	var page *agouti.Page
//
// 	BeforeEach(func() {
// 		var err error
// 		page, err = agoutiDriver.NewPage()
// 		Expect(err).NotTo(HaveOccurred())
//
// 	})
//
// 	AfterEach(func() {
// 		Expect(page.Destroy()).To(Succeed())
// 	})
//
// 	//Test Feature
//
// 	Describe("When user logged in and user try to delete alarm succesfully", func() {
// 		BeforeEach(func() {
// 			//login
// 			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
// 			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
// 			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
// 			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
// 		})
//
// 		It("should user create alarm succesfully", func() {
// 			//alarm 1
// 			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
// 			Expect(page.Find(AlarmPage["alertname"]).Fill("Alarm 1")).To(Succeed())
// 			Expect(page.Find(AlarmPage["required"]).Fill("\"TW\", \"ThoughtWorks\", \"Thought Works\", \"Thoughtworks\"")).To(Succeed())
// 			Expect(page.Find(AlarmPage["niceToHave"]).Fill("\"good\", \"best office\"")).To(Succeed())
// 			Expect(page.Find(AlarmPage["excluded"]).Fill("\"sucks\", \"not good enough\"")).To(Succeed())
// 			Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
// 			Expect(page.FindByButton(LoginPage["saveButton"]).Submit()).To(Succeed())
// 		})
//
// 		It("should user delete alarm successfully", func() {
// 			Expect(page.Find(HomePage["viewAlarm"]).Click()).To(Succeed())
// 			//Eventually(page.Find("#message")).Should(HaveText("Alarm 1"))
// 			Expect(page.Find(ViewAlarmPage["deleteAlarm"]).Click()).To(Succeed())
// 			Expect(page.Find(ViewAlarmPage["popUp"]).Click()).To(Succeed())
// 			Expect(page.Find("deletePopUp").Click()).To(Succeed())
// 			//Expect(page).To(HavePopupText("some alert"))
// 			Expect(page.ConfirmPopup()).To(Succeed())
// 		})
//
// 		It("should user can see error message on home page", func() {
// 			Eventually(page.Find("#errprmessage")).Should(HaveText("Alarm 1 has been removed from your alarm list"))
// 		})
// 	})
//
// 	Describe("When user logged in and user try to delete alarm and then give up", func() {
// 		BeforeEach(func() {
// 			//login
// 			Expect(page.Navigate(LoginPage["url"])).To(Succeed())
// 			Expect(page.Find(LoginPage["email"]).Fill("user1@mail.com")).To(Succeed())
// 			Expect(page.Find(LoginPage["password"]).Fill("passOne")).To(Succeed())
// 			Expect(page.FindByButton(LoginPage["submit"]).Submit()).To(Succeed())
// 		})
//
// 		It("should user create alarm succesfully", func() {
// 			//alarm 1
// 			Expect(page.Find(HomePage["createButton"]).Click()).To(Succeed())
// 			Expect(page.Find(AlarmPage["alertname"]).Fill("Alarm 1")).To(Succeed())
// 			Expect(page.Find(AlarmPage["required"]).Fill("\"TW\", \"ThoughtWorks\", \"Thought Works\", \"Thoughtworks\"")).To(Succeed())
// 			Expect(page.Find(AlarmPage["niceToHave"]).Fill("\"good\", \"best office\"")).To(Succeed())
// 			Expect(page.Find(AlarmPage["excluded"]).Fill("\"sucks\", \"not good enough\"")).To(Succeed())
// 			Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
// 			Expect(page.FindByButton(LoginPage["saveButton"]).Submit()).To(Succeed())
// 		})
//
// 		It("should user give up to delete alarm", func() {
// 			Expect(page.Find(HomePage["viewAlarm"]).Click()).To(Succeed())
// 			//Eventually(page.Find("#message")).Should(HaveText("Alarm 1"))
// 			Expect(page.Find(ViewAlarmPage["deleteAlarm"]).Click()).To(Succeed())
// 			Expect(page.Find(ViewAlarmPage["popUp"]).Click()).To(Succeed())
// 			Expect(page.Find("cancelPopUp").Click()).To(Succeed())
// 			//Expect(page).To(HavePopupText("some alert"))
// 			Expect(page.ConfirmPopup()).To(Succeed())
// 		})
//
// 		It("should user can see error message on home page", func() {
// 			Eventually(page.Find("div.has-error")).Should(HaveText("Alarm 1"))
// 		})
//
// 	})
// })
