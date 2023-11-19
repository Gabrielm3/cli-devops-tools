package utils

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func GetChromeScreenShot(site string, quality int) {
	screenShotURL := fmt.Sprintf("https://%s/", site)

	var buf []byte

	var ext string = "png"

	if quality < 100 {
		// ext = "jpg"
		ext = "jpeg"
	}

	log.Printf("Capturando print %s", screenShotURL)

	var options []chromedp.ExecAllocatorOption
	options = append(options, chromedp.WindowSize(1400, 2000))

	options = append(options, chromedp.DefaultExecAllocatorOptions[:]...)
	
	actx, acancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer acancel()

	ctx, cancel := chromedp.NewContext(actx)
	defer cancel()

	tasks := chromedp.Tasks {
		chromedp.Navigate(screenShotURL),
		chromedp.Sleep(3 * time.Second),
		chromedp.CaptureScreenshot(&buf),
	}

	if err := chromedp.Run(ctx, tasks); err != nil {
		log.Fatal(err)
	}

	filename := fmt.Sprintf("%s-%d-standard.%s", strings.Replace(site, "/", "-", -1), time.Now().UTC().Unix(), ext)
	if err := ioutil.WriteFile(filename, buf, 0644); err != nil {
		log.Fatal(err)
	}
	log.Printf("Captura armazenada em %s", filename)
}