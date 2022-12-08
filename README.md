[![screenshot-go license](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![screenshot-go made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://pkg.go.dev/github.com/whois-api-llc/screenshot-go)
[![screenshot-go test](https://github.com/whois-api-llc/screenshot-go/workflows/Test/badge.svg)](https://github.com/whois-api-llc/screenshot-go/actions/)

# Overview

The client library for
[Screenshot API](https://website-screenshot.whoisxmlapi.com/)
in Go language.

The minimum go version is 1.17.

# Installation

The library is distributed as a Go module

```bash
go get github.com/whois-api-llc/screenshot-go
```

# Examples

Full API documentation available [here](https://website-screenshot.whoisxmlapi.com/api/documentation/making-requests)

You can find all examples in `example` directory.

## Create a new client

To start making requests you need the API Key. 
You can find it on your profile page on [whoisxmlapi.com](https://whoisxmlapi.com/).
Using the API Key you can create Client.

Most users will be fine with `NewBasicClient` function. 
```go
client := screenshotapi.NewBasicClient(apiKey)
```

If you want to set custom `http.Client` to use proxy then you can use `NewClient` function.
```go
transport := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}

client := screenshotapi.NewClient(apiKey, screenshotapi.ClientParams{
    HTTPClient: &http.Client{
        Transport: transport,
        Timeout:   20 * time.Second,
    },
})
```

## Make basic requests

Screenshot API lets you get a screenshot of any web page as a jpg, png or pdf file.

```go

// Take a screenshot of the specified URL and save it to the file.
err := client.Get(ctx, "whoisxmlapi.com", "filename.jpg")
if err != nil {
    log.Fatal(err)
}

```

## Advanced usage
```go
cookies := screenshotapi.Cookies{
    "name1": "value1",
    "name2": "value2",
}

// Get raw API response with the following options:
// Pass cookies, emulate mobile device, disable JS, wait for network idle
// event, output API errors in XML, capture PDF screenshot in full page
// mode and get image data in base64.
resp, err := client.GetRaw(ctx,
    "whoisxmlapi.com",
    screenshotapi.OptionType("pdf"),
    screenshotapi.OptionMode("slow"),
    screenshotapi.OptionFullPage(true),
    screenshotapi.OptionMobile(true),
    screenshotapi.OptionErrorsOutputFormat("JSON"),
    screenshotapi.OptionImageOutputFormat("base64"),
    screenshotapi.OptionNoJs(true),
    screenshotapi.OptionCookies(cookies))

if err != nil {
    log.Fatal(err)
}

// body contains base64-encoded pdf file,
// either json with error message on failure
_ = resp.Body

```