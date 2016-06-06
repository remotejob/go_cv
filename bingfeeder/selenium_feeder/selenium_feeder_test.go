package selenium_feeder

import (
	gm "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	am "github.com/sclevine/agouti/matchers"
	"testing"
	"time"
)

func TestBing(t *testing.T) {

	gm.RegisterTestingT(t)
	driver := agouti.ChromeDriver()
	gm.Expect(driver.Start()).To(gm.Succeed())
	page, err := driver.NewPage(agouti.Browser("chrome"))
	gm.Expect(err).NotTo(gm.HaveOccurred())
	gm.Expect(page.Navigate("https://login.live.com/login.srf?wa=wsignin1.0&rpsnv=12&ct=1465178498&rver=6.7.6636.0&wp=MBI&wreply=https:%2F%2Fwww.bing.com%2Fsecure%2FPassport.aspx%3Frequrl%3Dhttps%253a%252f%252fwww.bing.com%252fwebmaster%252fWebmasterManageSitesPage.aspx%253frflid%253d1&lc=1033&id=264960")).To(gm.Succeed())
	//	gm.Expect(page.Navigate("https://stackoverflow.com/users/login?ssrc=head&returnurl=http%3a%2f%2fstackoverflow.com%2fjobs")).To(gm.Succeed())
	gm.Eventually(page.FindByClass("placeholder")).Should(am.BeFound())
	gm.Expect(page.FindByClass("placeholder").Fill("email")).To(gm.Succeed())

	time.Sleep(10000 * time.Millisecond)
	gm.Expect(driver.Stop()).To(gm.Succeed())

	//		driver := agouti.ChromeDriver()
	//		gm.Expect(driver.Start()).To(gm.Succeed())
	//		page, err := driver.NewPage(agouti.Browser("chrome"))
	//		gm.Expect(err).NotTo(gm.HaveOccurred())
	//		gm.Expect(page.Navigate("https://stackoverflow.com/users/login?ssrc=head&returnurl=http%3a%2f%2fstackoverflow.com%2fjobs")).To(gm.Succeed())
	//		gm.Expect(page).To(am.HaveURL("https://stackoverflow.com/users/login?ssrc=head&returnurl=http%3a%2f%2fstackoverflow.com%2fjobs"))

}
