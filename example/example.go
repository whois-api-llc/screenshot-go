package example

import (
	"context"
	"errors"
	screenshotapi "github.com/whois-api-llc/screenshot-go"
	"log"
)

func GetData(apikey string) {
	client := screenshotapi.NewBasicClient(apikey)

	// Take a screenshot of the specified URL and save it to the file.
	err := client.Get(context.Background(),
		"whoisxmlapi.com",
		// specify the filename to save the screenshot.
		"/tmp/filename.jpg",
	)

	if err != nil {
		// Handle error message returned by server.
		var apiErr *screenshotapi.ErrorMessage
		if errors.As(err, &apiErr) {
			log.Println(apiErr.Code)
			log.Println(apiErr.Message)
		}
		log.Fatal(err)
	}

}

func GetRawData1(apikey string) {
	client := screenshotapi.NewBasicClient(apikey)

	cookies := screenshotapi.Cookies{
		"name1": "value1",
		"name2": "value2",
	}

	// Get raw API response with the following options:
	// Pass cookies, emulate mobile device, disable JS, wait for network idle
	// event, output API errors in XML, capture PDF screenshot in full page
	// mode and get image data in base64.
	resp, err := client.GetRaw(context.Background(),
		"whoisxmlapi.com",
		screenshotapi.OptionType("pdf"),
		screenshotapi.OptionMode("slow"),
		screenshotapi.OptionFullPage(true),
		screenshotapi.OptionMobile(true),
		screenshotapi.OptionErrorsOutputFormat("XML"),
		screenshotapi.OptionImageOutputFormat("base64"),
		screenshotapi.OptionNoJs(true),
		screenshotapi.OptionCookies(cookies),
	)

	if err != nil {
		// Handle error message returned by server.
		log.Fatal(err)
	}

	log.Println(len(resp.Body))
}

func GetRawData2(apikey string) {
	client := screenshotapi.NewBasicClient(apikey)

	// Get raw API response with the following options:
	// emulate touch screen device, wait 10s for page loading with delay of 100ms before capture,
	// renders page in landscape mode with scrolling to the bottom, capture JPG screenshot in full page mode
	// with the highest quality, use the DRS credits and return error on hostname change.
	resp, err := client.GetRaw(context.Background(),
		"whoisxmlapi.com",
		screenshotapi.OptionTouchScreen(true),
		screenshotapi.OptionTimeout(10000),
		screenshotapi.OptionFailOnHostnameChange(true),
		screenshotapi.OptionLandscape(true),
		screenshotapi.OptionDelay(100),
		screenshotapi.OptionScale(2),
		screenshotapi.OptionFullPage(true),
		screenshotapi.OptionScroll(true),
		screenshotapi.OptionScrollPosition("bottom"),
		screenshotapi.OptionWidth(1024),
		screenshotapi.OptionHeight(768),
		screenshotapi.OptionQuality(99),
		screenshotapi.OptionCredits("DRS"),
		screenshotapi.OptionType("jpg"),
	)

	if err != nil {
		// Handle error message returned by server.
		log.Fatal(err)
	}

	log.Println(len(resp.Body))
}
