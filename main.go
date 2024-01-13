// Command screenshot is a chromedp example demonstrating how to take a
// screenshot of a specific element and of the entire browser viewport.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chromedp/chromedp"
)

func main() {
	fmt.Println(1)
	// create context
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()
	fmt.Println(2)

	// capture screenshot of an element
	var buf []byte
	if err := chromedp.Run(ctx, chromedp.Navigate(`https://pkg.go.dev/`)); err != nil {
		log.Fatal(err)
	}
	fmt.Println(22)
	if err := chromedp.Run(ctx, chromedp.Screenshot(`img.Homepage-logo`, &buf, chromedp.NodeVisible)); err != nil {
		log.Fatal(err)
	}
	fmt.Println(3)
	if err := os.WriteFile("elementScreenshot.png", buf, 0o644); err != nil {
		log.Fatal(err)
	}

	fmt.Println(4)
	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, fullScreenshot(`https://brank.as/`, 90, &buf)); err != nil {
		log.Fatal(err)
	}
	fmt.Println(5)
	if err := os.WriteFile("fullScreenshot.png", buf, 0o644); err != nil {
		log.Fatal(err)
	}

	log.Printf("wrote elementScreenshot.png and fullScreenshot.png")
}

// elementScreenshot takes a screenshot of a specific element.
func elementScreenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	}
}

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Note: chromedp.FullScreenshot overrides the device's emulation settings. Use
// device.Reset to reset the emulation and viewport settings.
func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.FullScreenshot(res, quality),
	}
}
