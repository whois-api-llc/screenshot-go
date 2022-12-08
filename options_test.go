package screenshotapi

import (
	"net/url"
	"reflect"
	"testing"
)

// TestOptions tests the Options functions.
func TestOptions(t *testing.T) {
	tests := []struct {
		name    string
		values  url.Values
		option  Option
		want    string
		wantErr string
	}{
		{
			name:    "errorsOutputFormat1",
			values:  url.Values{},
			option:  OptionErrorsOutputFormat("JSON"),
			want:    "errorsOutputFormat=JSON",
			wantErr: "",
		},
		{
			name:    "errorsOutputFormat2",
			values:  url.Values{},
			option:  OptionErrorsOutputFormat("xml"),
			want:    "errorsOutputFormat=XML",
			wantErr: "",
		},
		{
			name:    "errorsOutputFormat3",
			values:  url.Values{},
			option:  OptionErrorsOutputFormat("CSV"),
			want:    "",
			wantErr: "invalid argument: \"errorsOutputFormat\" must be JSON | XML",
		},
		{
			name:   "imageOutputFormat1",
			values: url.Values{},
			option: OptionImageOutputFormat("image"),
			want:   "imageOutputFormat=image",
		},
		{
			name:   "imageOutputFormat2",
			values: url.Values{},
			option: OptionImageOutputFormat("BASE64"),
			want:   "imageOutputFormat=base64",
		},
		{
			name:    "imageOutputFormat3",
			values:  url.Values{},
			option:  OptionImageOutputFormat("bmp"),
			want:    "",
			wantErr: "invalid argument: \"imageOutputFormat\" must be image | base64",
		},
		{
			name:   "credits1",
			values: url.Values{},
			option: OptionCredits("SA"),
			want:   "credits=SA",
		},
		{
			name:   "credits2",
			values: url.Values{},
			option: OptionCredits("drs"),
			want:   "credits=DRS",
		},
		{
			name:    "credits3",
			values:  url.Values{},
			option:  OptionCredits("whois"),
			want:    "",
			wantErr: "invalid argument: \"credits\" must be SA | DRS",
		},
		{
			name:   "type1",
			values: url.Values{},
			option: OptionType("jpg"),
			want:   "type=jpg",
		},
		{
			name:   "type2",
			values: url.Values{},
			option: OptionType("PNG"),
			want:   "type=png",
		},
		{
			name:    "type3",
			values:  url.Values{},
			option:  OptionType("BMP"),
			want:    "",
			wantErr: "invalid argument: \"imageType\" must be jpg | png | pdf",
		},
		{
			name:   "quality1",
			values: url.Values{},
			option: OptionQuality(85),
			want:   "quality=85",
		},
		{
			name:    "quality2",
			values:  url.Values{},
			option:  OptionQuality(0),
			want:    "",
			wantErr: "invalid argument: \"quality\" must be between 40 and 99",
		},
		{
			name:    "quality3",
			values:  url.Values{},
			option:  OptionQuality(185),
			want:    "",
			wantErr: "invalid argument: \"quality\" must be between 40 and 99",
		},
		{
			name:   "width1",
			values: url.Values{},
			option: OptionWidth(800),
			want:   "width=800",
		},
		{
			name:    "width2",
			values:  url.Values{},
			option:  OptionWidth(0),
			want:    "",
			wantErr: "invalid argument: \"width\" must be between 100 and 3000",
		},
		{
			name:    "width3",
			values:  url.Values{},
			option:  OptionWidth(4000),
			want:    "",
			wantErr: "invalid argument: \"width\" must be between 100 and 3000",
		},
		{
			name:   "height1",
			values: url.Values{},
			option: OptionHeight(600),
			want:   "height=600",
		},
		{
			name:    "height2",
			values:  url.Values{},
			option:  OptionHeight(0),
			want:    "",
			wantErr: "invalid argument: \"height\" must be between 100 and 3000",
		},
		{
			name:    "height3",
			values:  url.Values{},
			option:  OptionHeight(6000),
			want:    "",
			wantErr: "invalid argument: \"height\" must be between 100 and 3000",
		},
		{
			name:   "thumbWidth1",
			values: url.Values{},
			option: OptionThumbWidth(100),
			want:   "thumbWidth=100",
		},
		{
			name:    "thumbWidth2",
			values:  url.Values{},
			option:  OptionThumbWidth(0),
			want:    "",
			wantErr: "invalid argument: \"thumbWidth\" must be between 50 and width param value",
		},
		{
			name:    "thumbWidth3",
			values:  url.Values{},
			option:  OptionThumbWidth(4000),
			want:    "",
			wantErr: "invalid argument: \"thumbWidth\" must be between 50 and width param value",
		},
		{
			name:   "mode1",
			values: url.Values{},
			option: OptionMode("fast"),
			want:   "mode=fast",
		},
		{
			name:   "mode2",
			values: url.Values{},
			option: OptionMode("SLOW"),
			want:   "mode=slow",
		},
		{
			name:    "mode3",
			values:  url.Values{},
			option:  OptionMode("fastest"),
			want:    "",
			wantErr: "invalid argument: \"mode\" must be fast | slow",
		},
		{
			name:   "scroll1",
			values: url.Values{},
			option: OptionScroll(true),
			want:   "scroll=true",
		},
		{
			name:   "scroll2",
			values: url.Values{},
			option: OptionScroll(false),
			want:   "",
		},
		{
			name:   "scrollPosition1",
			values: url.Values{},
			option: OptionScrollPosition("top"),
			want:   "scrollPosition=top",
		},
		{
			name:   "scrollPosition2",
			values: url.Values{},
			option: OptionScrollPosition("BOTTOM"),
			want:   "scrollPosition=bottom",
		},
		{
			name:    "scrollPosition3",
			values:  url.Values{},
			option:  OptionScrollPosition("middle"),
			want:    "",
			wantErr: "invalid argument: \"scrollPosition\" must be top | bottom",
		},
		{
			name:   "fullPage1",
			values: url.Values{},
			option: OptionFullPage(true),
			want:   "fullPage=true",
		},
		{
			name:   "fullPage2",
			values: url.Values{},
			option: OptionFullPage(false),
			want:   "",
		},
		{
			name:   "noJs1",
			values: url.Values{},
			option: OptionNoJs(true),
			want:   "noJs=true",
		},
		{
			name:   "noJs2",
			values: url.Values{},
			option: OptionNoJs(false),
			want:   "",
		},
		{
			name:   "delay1",
			values: url.Values{},
			option: OptionDelay(100),
			want:   "delay=100",
		},
		{
			name:    "delay2",
			values:  url.Values{},
			option:  OptionDelay(20000),
			want:    "",
			wantErr: "invalid argument: \"delay\" must be between 0 and 10000",
		},
		{
			name:    "delay3",
			values:  url.Values{},
			option:  OptionDelay(0),
			want:    "delay=0",
			wantErr: "",
		},
		{
			name:   "timeout1",
			values: url.Values{},
			option: OptionTimeout(2000),
			want:   "timeout=2000",
		},
		{
			name:    "timeout2",
			values:  url.Values{},
			option:  OptionTimeout(200),
			want:    "",
			wantErr: "invalid argument: \"timeout\" must be between 1000 and 30000",
		},
		{
			name:    "timeout3",
			values:  url.Values{},
			option:  OptionTimeout(40000),
			want:    "",
			wantErr: "invalid argument: \"timeout\" must be between 1000 and 30000",
		},
		{
			name:   "scale1",
			values: url.Values{},
			option: OptionScale(1.0),
			want:   "scale=1.000000",
		},
		{
			name:    "scale2",
			values:  url.Values{},
			option:  OptionScale(0.4),
			want:    "",
			wantErr: "invalid argument: \"scale\" must be between 0.5 and 4.0",
		},
		{
			name:    "scale3",
			values:  url.Values{},
			option:  OptionScale(10),
			want:    "",
			wantErr: "invalid argument: \"scale\" must be between 0.5 and 4.0",
		},
		{
			name:   "retina1",
			values: url.Values{},
			option: OptionRetina(true),
			want:   "retina=true",
		},
		{
			name:   "retina2",
			values: url.Values{},
			option: OptionRetina(false),
			want:   "",
		},
		{
			name:   "UA1",
			values: url.Values{},
			option: OptionUA("UA"),
			want:   "ua=UA",
		},
		{
			name:   "UA2",
			values: url.Values{},
			option: OptionUA(""),
			want:   "ua=",
		},
		{
			name:   "Cookies1",
			values: url.Values{},
			option: OptionCookies(Cookies{"name1": "value1"}),
			want:   "cookies=name1%3Dvalue1",
		},
		{
			name:   "Cookies2",
			values: url.Values{},
			option: OptionCookies(Cookies{"name1": "value1", "name2": "value2"}),
			want:   "cookies=name1%3Dvalue1%3Bname2%3Dvalue2",
		},
		{
			name:   "Cookies3",
			values: url.Values{},
			option: OptionCookies(Cookies{}),
			want:   "cookies=",
		},
		{
			name:   "mobile1",
			values: url.Values{},
			option: OptionMobile(true),
			want:   "mobile=true",
		},
		{
			name:   "mobile2",
			values: url.Values{},
			option: OptionMobile(false),
			want:   "",
		},
		{
			name:   "touchScreen1",
			values: url.Values{},
			option: OptionTouchScreen(true),
			want:   "touchScreen=true",
		},
		{
			name:   "touchScreen2",
			values: url.Values{},
			option: OptionTouchScreen(false),
			want:   "",
		},
		{
			name:   "landscape1",
			values: url.Values{},
			option: OptionLandscape(true),
			want:   "landscape=true",
		},
		{
			name:   "landscape2",
			values: url.Values{},
			option: OptionLandscape(false),
			want:   "",
		},
		{
			name:   "failOnHostnameChange1",
			values: url.Values{},
			option: OptionFailOnHostnameChange(true),
			want:   "failOnHostnameChange=true",
		},
		{
			name:   "failOnHostnameChange2",
			values: url.Values{},
			option: OptionFailOnHostnameChange(false),
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.option(tt.values)
			if (err != nil || tt.wantErr != "") && (err == nil || err.Error() != tt.wantErr) {
				t.Errorf("Option() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got := tt.values.Encode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Option() = %v, want %v", got, tt.want)
			}
		})
	}
}
