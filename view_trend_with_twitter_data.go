package marketwatcher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	//. "github.com/sclevine/agouti/matchers"
	"time"
)

// Configuration
var _ = Describe("ViewTrend", func() {
  alarmTitleTimeStamp := time.Now().Format("0102 150405")
	var page *agouti.Page

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())

	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

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
      Expect(page.Find(AlarmPage["alertname"]).Fill("Alarm 3")).To(Succeed())
      Expect(page.Find(AlarmPage["required"]).Fill(" TW, ThoughtWorks, Thought Works, Thoughtworks")).To(Succeed())
      Expect(page.Find(AlarmPage["niceToHave"]).Fill("good, best office")).To(Succeed())
      Expect(page.Find(AlarmPage["threshold"]).Fill("1000")).To(Succeed())
      Expect(page.Find(AlarmPage["saveButton"]).Click()).To(Succeed())
    })

    It("should user can see alarm header in alarm detail page", func() {
      trendTitles := page.AllByXPath("//h5[text()='Test Otomasyon" + alarmTitleTimeStamp + "']")
      Eventually(trendTitles).Should(BeFound())
      Expect(trendTitles.Click()).To(Succeed())
      Eventually(page.Find(ViewAlarmPage["alarmTitle"])).Should(HaveText("Test Otomasyon" + alarmTitleTimeStamp))
      Eventually(page.Find(ViewAlarmPage["alarmTitle"])).Should(HaveText("Dashed Straight Right Arrow"))
    })

  })
}
