package main

import (
	"fmt"
	"time"
	"github.com/tebeka/selenium"
	"github.com/remotejob/go_cv/domains"
	"gopkg.in/gcfg.v1"
	"log"		
)

var mlogin = ""
var mpass = ""

func init() {

	var cfg domains.ServerConfig
	if err := gcfg.ReadFileInto(&cfg, "/home/juno/neonworkspace/go_cv/config.gcfg"); err != nil {
		log.Fatalln(err.Error())

	} else {

		mlogin = cfg.Login.Mlogin
		mpass = cfg.Pass.Mpass

	}

}

// Errors are ignored for brevity.

func main() {
	// FireFox driver without specific version
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, _ := selenium.NewRemote(caps, "")
	defer wd.Quit()

	wd.Get("https://login.live.com/login.srf?wa=wsignin1.0&rpsnv=12&ct=1465178498&rver=6.7.6636.0&wp=MBI&wreply=https:%2F%2Fwww.bing.com%2Fsecure%2FPassport.aspx%3Frequrl%3Dhttps%253a%252f%252fwww.bing.com%252fwebmaster%252fWebmasterManageSitesPage.aspx%253frflid%253d1&lc=1033&id=264960")

	time.Sleep(time.Millisecond * 5000)

	elem, err := wd.FindElement(selenium.ByID, "i0116")
	if err != nil {
		fmt.Println(err.Error())
	}
	pass, err := wd.FindElement(selenium.ByID, "i0118")
	if err != nil {
		fmt.Println(err.Error())
	}	
	time.Sleep(time.Millisecond * 1000)

	err = elem.SendKeys(mlogin)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = pass.SendKeys(mpass)
	if err != nil {
		fmt.Println(err.Error())
	}	
	btm, err := wd.FindElement(selenium.ByID, "idSIButton9")
	if err != nil {
		fmt.Println(err.Error())
	}
	btm.Click()
	time.Sleep(time.Millisecond * 15000)

	//	// Enter code in textarea
	//	elem, _ := wd.FindElement(selenium.ByCSSSelector, "#code")
	//	elem.Clear()
	//	elem.SendKeys(code)
	//
	//	// Click the run button
	//	btn, _ := wd.FindElement(selenium.ByCSSSelector, "#run")
	//	btn.Click()
	//
	//	// Get the result
	//	div, _ := wd.FindElement(selenium.ByCSSSelector, "#output")
	//
	//	output := ""
	//	// Wait for run to finish
	//	for {
	//		output, _ = div.Text()
	//		if output != "Waiting for remote server..." {
	//			break
	//		}
	//		time.Sleep(time.Millisecond * 100)
	//	}
	//
	//	fmt.Printf("Got: %s\n", output)
}
