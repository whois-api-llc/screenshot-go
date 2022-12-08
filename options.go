package screenshotapi

import (
	"fmt"
	"net/url"
	"strings"
)

// Option adds parameters to the query.
type Option func(v url.Values) error

var _ = []Option{
	OptionErrorsOutputFormat("JSON"),
	OptionImageOutputFormat("image"),
	OptionCredits("DRS"),
	OptionType("jpg"),
	OptionQuality(50),
	OptionWidth(800),
	OptionHeight(600),
	OptionThumbWidth(51),
	OptionMode("fast"),
	OptionScroll(true),
	OptionScrollPosition("top"),
	OptionFullPage(true),
	OptionNoJs(true),
	OptionDelay(250),
	OptionTimeout(15000),
	OptionScale(1.0),
	OptionRetina(true),
	OptionUA("UA"),
	OptionCookies(Cookies{"name": "value"}),
	OptionMobile(true),
	OptionTouchScreen(true),
	OptionLandscape(true),
	OptionFailOnHostnameChange(true),
}

const (
	MinDelay      = 0
	MaxDelay      = 10000
	MinJPGQuality = 40
	MaxJPGQuality = 99
	MinScale      = 0.5
	MaxScale      = 4.0
	MinSize       = 100
	MaxSize       = 3000
	MinThumbWidth = 50
	MinTimeout    = 1000
	MaxTimeout    = 30000
	DefaultWidth  = 800
)

var (
	errorsOutputFormats = map[string]bool{
		"JSON": true,
		"XML":  true,
	}
	imageOutputFormats = map[string]bool{
		"image":  true,
		"base64": true,
	}
	creditTypes = map[string]bool{
		"SA":  true,
		"DRS": true,
	}
	imageTypes = map[string]bool{
		"jpg": true,
		"png": true,
		"pdf": true,
	}
	modes = map[string]bool{
		"fast": true,
		"slow": true,
	}
	scrollPositions = map[string]bool{
		"top":    true,
		"bottom": true,
	}
)

// OptionErrorsOutputFormat sets the Errors output format. Acceptable values: JSON | XML. Default: JSON.
func OptionErrorsOutputFormat(errorsOutputFormat string) Option {
	return func(v url.Values) error {
		if !errorsOutputFormats[strings.ToUpper(errorsOutputFormat)] {
			return &ArgError{"errorsOutputFormat", "must be JSON | XML"}
		}
		v.Set("errorsOutputFormat", strings.ToUpper(errorsOutputFormat))
		return nil
	}
}

// OptionImageOutputFormat sets the Response output format. Acceptable values: image | base64. Default: image.
func OptionImageOutputFormat(imageOutputFormat string) Option {
	return func(v url.Values) error {
		if !imageOutputFormats[strings.ToLower(imageOutputFormat)] {
			return &ArgError{"imageOutputFormat", "must be image | base64"}
		}
		v.Set("imageOutputFormat", strings.ToLower(imageOutputFormat))
		return nil
	}
}

// OptionCredits sets the type of credits used.
// SA — Screenshot API credits will be taken into account when the API is called.
// DRS — Domain Research Suite credits will be taken into account when the API is called.
// Acceptable values: SA | DRS. Default: SA.
func OptionCredits(credits string) Option {
	return func(v url.Values) error {
		if !creditTypes[strings.ToUpper(credits)] {
			return &ArgError{"credits", "must be SA | DRS"}
		}
		v.Set("credits", strings.ToUpper(credits))
		return nil
	}
}

// OptionType sets the image output type. Acceptable values: jpg | png | pdf. Default: jpg.
func OptionType(imageType string) Option {
	return func(v url.Values) error {
		if !imageTypes[strings.ToLower(imageType)] {
			return &ArgError{"imageType", "must be jpg | png | pdf"}
		}
		v.Set("type", strings.ToLower(imageType))
		return nil
	}
}

// OptionQuality sets the image quality. (only for jpg type).Acceptable values: 40 < quality < 99 Default: 85.
func OptionQuality(quality int) Option {
	return func(v url.Values) error {
		if quality < MinJPGQuality || quality > MaxJPGQuality {
			return &ArgError{"quality", "must be between 40 and 99"}
		}
		v.Set("quality", fmt.Sprintf("%d", quality))
		return nil
	}
}

// OptionWidth sets the image width (px). Acceptable values: 100 < width < 3000. Default: 800.
func OptionWidth(width int) Option {
	return func(v url.Values) error {
		if width < MinSize || width > MaxSize {
			return &ArgError{"width", "must be between 100 and 3000"}
		}
		v.Set("width", fmt.Sprintf("%d", width))
		return nil
	}
}

// OptionHeight sets the image height (px). Acceptable values: 100 < width < 3000. Default: 600.
func OptionHeight(height int) Option {
	return func(v url.Values) error {
		if height < MinSize || height > MaxSize {
			return &ArgError{"height", "must be between 100 and 3000"}
		}
		v.Set("height", fmt.Sprintf("%d", height))
		return nil
	}
}

// OptionThumbWidth sets the image thumb width (px). Acceptable values: 50 < thumbWidth < width param value. Default: 0.
func OptionThumbWidth(thumbWidth int) Option {
	return func(v url.Values) error {
		if thumbWidth < MinThumbWidth || thumbWidth > MaxSize {
			return &ArgError{"thumbWidth", "must be between 50 and width param value"}
		}
		v.Set("thumbWidth", fmt.Sprintf("%d", thumbWidth))
		return nil
	}
}

// OptionMode sets the retrieving mode. fast - waiting for the document.load event.
// slow - waiting for network idle event. Acceptable values: fast | slow. Default: fast.
func OptionMode(mode string) Option {
	return func(v url.Values) error {
		if !modes[strings.ToLower(mode)] {
			return &ArgError{"mode", "must be fast | slow"}
		}
		v.Set("mode", strings.ToLower(mode))
		return nil
	}
}

// OptionScroll If true, the API scrolls down and to the scrollPosition (useful for full-page screenshots).
func OptionScroll(scroll bool) Option {
	return func(v url.Values) error {
		if scroll {
			v.Set("scroll", fmt.Sprint(scroll))
		}
		return nil
	}
}

// OptionScrollPosition specifies scroll's behavior. Acceptable values: top | bottom. Default: top.
func OptionScrollPosition(scrollPosition string) Option {
	return func(v url.Values) error {
		if !scrollPositions[strings.ToLower(scrollPosition)] {
			return &ArgError{"scrollPosition", "must be top | bottom"}
		}
		v.Set("scrollPosition", strings.ToLower(scrollPosition))
		return nil
	}
}

// OptionFullPage if true, the API makes full-page screenshot.
func OptionFullPage(fullPage bool) Option {
	return func(v url.Values) error {
		if fullPage {
			v.Set("fullPage", fmt.Sprint(fullPage))
		}
		return nil
	}
}

// OptionNoJs if true, the API disables JS.
func OptionNoJs(noJs bool) Option {
	return func(v url.Values) error {
		if noJs {
			v.Set("noJs", fmt.Sprint(noJs))
		}
		return nil
	}
}

// OptionDelay sets the custom delay (ms) before screen capture. Acceptable values: 0 < delay < 10000 ms. Default: 250.
func OptionDelay(delay int) Option {
	return func(v url.Values) error {
		if delay < MinDelay || delay >= MaxDelay {
			return &ArgError{"delay", "must be between 0 and 10000"}
		}
		v.Set("delay", fmt.Sprintf("%d", delay))
		return nil
	}
}

// OptionTimeout sets the custom timeout (ms) for page loading. API will respond with an error if our server can't load
// the page within the specified timeout. Acceptable values: 1000 < timeout < 30000 ms. Default: 15000.
func OptionTimeout(timeout int) Option {
	return func(v url.Values) error {
		if timeout < MinTimeout || timeout > MaxTimeout {
			return &ArgError{"timeout", "must be between 1000 and 30000"}
		}
		v.Set("timeout", fmt.Sprintf("%d", timeout))
		return nil
	}
}

// OptionScale sets the deviceScaleFactor value for the emulator. Acceptable values: 0.5 < scale < 4.0 Default: 1.0.
func OptionScale(scale float64) Option {
	return func(v url.Values) error {
		if scale < MinScale || scale > MaxScale {
			return &ArgError{"scale", "must be between 0.5 and 4.0"}
		}
		v.Set("scale", fmt.Sprintf("%f", scale))
		return nil
	}
}

// OptionRetina if true, the API emulates retina display.
func OptionRetina(retina bool) Option {
	return func(v url.Values) error {
		if retina {
			v.Set("retina", fmt.Sprint(retina))
		}
		return nil
	}
}

// OptionUA sets the 'User-Agent' header string.
func OptionUA(ua string) Option {
	return func(v url.Values) error {
		v.Set("ua", ua)
		return nil
	}
}

// OptionCookies sets the 'Cookie' header string in the following format: name1=value1;name2=value2.
func OptionCookies(cookies Cookies) Option {
	return func(v url.Values) error {
		v.Set("cookies", cookies.toString())
		return nil
	}
}

// OptionMobile if true, the API emulates mobile device.
func OptionMobile(mobile bool) Option {
	return func(v url.Values) error {
		if mobile {
			v.Set("mobile", fmt.Sprint(mobile))
		}
		return nil
	}
}

// OptionTouchScreen if true, the API emulates device with a touch screens.
func OptionTouchScreen(touchScreen bool) Option {
	return func(v url.Values) error {
		if touchScreen {
			v.Set("touchScreen", fmt.Sprint(touchScreen))
		}
		return nil
	}
}

// OptionLandscape if true, the API renders page in landscape mode (useful for smartphone emulation).
func OptionLandscape(landscape bool) Option {
	return func(v url.Values) error {
		if landscape {
			v.Set("landscape", fmt.Sprint(landscape))
		}
		return nil
	}
}

// OptionFailOnHostnameChange if true the API responds with 422 HTTP error code when the target domain name
// is changed due to redirects.
func OptionFailOnHostnameChange(failOnHostnameChange bool) Option {
	return func(v url.Values) error {
		if failOnHostnameChange {
			v.Set("failOnHostnameChange", fmt.Sprint(failOnHostnameChange))
		}
		return nil
	}
}
