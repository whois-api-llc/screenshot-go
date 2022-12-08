package screenshotapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

const (
	pathScreenshotAPIResponseOK         = "/ScreenshotAPI/ok"
	pathScreenshotAPIResponseError      = "/ScreenshotAPI/error"
	pathScreenshotAPIResponse500        = "/ScreenshotAPI/500"
	pathScreenshotAPIResponsePartial1   = "/ScreenshotAPI/partial"
	pathScreenshotAPIResponsePartial2   = "/ScreenshotAPI/partial2"
	pathScreenshotAPIResponseUnparsable = "/ScreenshotAPI/unparsable"
)

const apiKey = "at_LoremIpsumDolorSitAmetConsect"

// dummyServer is the sample of the Screenshot API server for testing.
func dummyServer(resp, respUnparsable string, respErr string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var response string

		response = resp

		switch req.URL.Path {
		case pathScreenshotAPIResponseOK:
		case pathScreenshotAPIResponseError:
			w.WriteHeader(499)
			response = respErr
		case pathScreenshotAPIResponse500:
			w.WriteHeader(500)
			response = respUnparsable
		case pathScreenshotAPIResponsePartial1:
			response = response[:len(response)-10]
		case pathScreenshotAPIResponsePartial2:
			w.Header().Set("Content-Length", strconv.Itoa(len(response)))
			response = response[:len(response)-10]
		case pathScreenshotAPIResponseUnparsable:
			response = respUnparsable
		default:
			panic(req.URL.Path)
		}
		_, err := w.Write([]byte(response))
		if err != nil {
			panic(err)
		}
	}))

	return server
}

// newAPI returns new Screenshot API client for testing.
func newAPI(apiServer *httptest.Server, link string) *Client {
	apiURL, err := url.Parse(apiServer.URL)
	if err != nil {
		panic(err)
	}

	apiURL.Path = link

	params := ClientParams{
		HTTPClient:           apiServer.Client(),
		ScreenshotAPIBaseURL: apiURL,
	}

	return NewClient(apiKey, params)
}

// TestScreenshotAPIGet tests the Get function.
func TestScreenshotAPIGet(t *testing.T) {
	ctx := context.Background()

	const resp = `data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBA`

	const respUnparsable = `<?xml version="1.0" encoding="utf-8"?><>`

	const errResp = `{"code":499,"messages":"Test error message."}`

	server := dummyServer(resp, respUnparsable, errResp)
	defer server.Close()

	type options struct {
		mandatory1 string
		mandatory2 string
		option     Option
	}

	type args struct {
		ctx     context.Context
		options options
	}

	tests := []struct {
		name    string
		path    string
		args    args
		want    bool
		wantErr string
	}{
		{
			name: "successful request",
			path: pathScreenshotAPIResponseOK,
			args: args{
				ctx: ctx,
				options: options{
					"whoisxmlapi.com",
					"/tmp/filename.jpg",
					OptionImageOutputFormat("base64"),
				},
			},
			want:    true,
			wantErr: "",
		},
		{
			name: "non 200 status code",
			path: pathScreenshotAPIResponse500,
			args: args{
				ctx: ctx,
				options: options{
					"whoisxmlapi.com",
					"/tmp/filename.jpg",
					OptionImageOutputFormat("base64"),
				},
			},
			want:    false,
			wantErr: "API failed with status code: 500",
		},
		{
			name: "partial response 1",
			path: pathScreenshotAPIResponsePartial1,
			args: args{
				ctx: ctx,
				options: options{
					"whoisxmlapi.com",
					"/tmp/filename.jpg",
					OptionImageOutputFormat("base64"),
				},
			},
			want:    true,
			wantErr: "",
		},
		{
			name: "partial response 2",
			path: pathScreenshotAPIResponsePartial2,
			args: args{
				ctx: ctx,
				options: options{
					"whoisxmlapi.com",
					"/tmp/filename.jpg",
					OptionImageOutputFormat("base64"),
				},
			},
			want:    false,
			wantErr: "cannot read response: unexpected EOF",
		},
		{
			name: "could not process request",
			path: pathScreenshotAPIResponseError,
			args: args{
				ctx: ctx,
				options: options{
					"whoisxmlapi.com",
					"/tmp/filename.jpg",
					OptionImageOutputFormat("base64"),
				},
			},
			want:    false,
			wantErr: "API error: [499] Test error message.",
		},
		{
			name: "unparsable response",
			path: pathScreenshotAPIResponseUnparsable,
			args: args{
				ctx: ctx,
				options: options{
					"whoisxmlapi.com",
					"/tmp/filename.jpg",
					OptionImageOutputFormat("base64"),
				},
			},
			want:    true,
			wantErr: "",
		},
		{
			name: "invalid argument1",
			path: pathScreenshotAPIResponseOK,
			args: args{
				ctx: ctx,
				options: options{
					"",
					"/tmp/filename.jpg",
					OptionImageOutputFormat("base64"),
				},
			},
			want:    false,
			wantErr: `invalid argument: "URL" can not be empty`,
		},
		{
			name: "invalid argument2",
			path: pathScreenshotAPIResponseOK,
			args: args{
				ctx: ctx,
				options: options{
					"whoisxmlapi.com",
					"",
					OptionImageOutputFormat("base64"),
				},
			},
			want:    false,
			wantErr: `invalid argument: "filename" can not be empty`,
		},
		{
			name: "invalid argument3",
			path: pathScreenshotAPIResponseOK,
			args: args{
				ctx: ctx,
				options: options{
					"whoisxmlapi.com",
					"/nonexistent/path/filename.jpg",
					OptionImageOutputFormat("base64"),
				},
			},
			want:    false,
			wantErr: `open /nonexistent/path/filename.jpg: no such file or directory`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := newAPI(server, tt.path)

			err := api.Get(tt.args.ctx, tt.args.options.mandatory1, tt.args.options.mandatory2, tt.args.options.option)
			if (err != nil || tt.wantErr != "") && (err == nil || err.Error() != tt.wantErr) {
				t.Errorf("ScreenshotAPI.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

// TestScreenshotAPIGetRaw tests the GetRaw function.
func TestScreenshotAPIGetRaw(t *testing.T) {
	checkResultRaw := func(res []byte) bool {
		return len(res) != 0
	}

	ctx := context.Background()

	const resp = `data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBA`

	const respUnparsable = `<?xml version="1.0" encoding="utf-8"?><>`

	const errResp = `{"code":499,"messages":"Test error message."}`

	server := dummyServer(resp, respUnparsable, errResp)
	defer server.Close()

	type options struct {
		mandatory string
		option    Option
	}

	type args struct {
		ctx     context.Context
		options options
	}

	tests := []struct {
		name    string
		path    string
		args    args
		wantErr string
	}{
		{
			name: "successful request",
			path: pathScreenshotAPIResponseOK,
			args: args{
				ctx: ctx,
				options: options{
					"whoisxmlapi.com",
					OptionImageOutputFormat("base64"),
				},
			},
			wantErr: "",
		},
		{
			name: "non 200 status code",
			path: pathScreenshotAPIResponse500,
			args: args{
				ctx: ctx,
				options: options{
					"whoisxmlapi.com",
					OptionImageOutputFormat("base64"),
				},
			},
			wantErr: "API failed with status code: 500",
		},
		{
			name: "partial response 1",
			path: pathScreenshotAPIResponsePartial1,
			args: args{
				ctx: ctx,
				options: options{
					"whoisxmlapi.com",
					OptionImageOutputFormat("base64"),
				},
			},
			wantErr: "",
		},
		{
			name: "partial response 2",
			path: pathScreenshotAPIResponsePartial2,
			args: args{
				ctx: ctx,
				options: options{
					"whoisxmlapi.com",
					OptionImageOutputFormat("base64"),
				},
			},
			wantErr: "cannot read response: unexpected EOF",
		},
		{
			name: "unparsable response",
			path: pathScreenshotAPIResponseUnparsable,
			args: args{
				ctx: ctx,
				options: options{
					"whoisxmlapi.com",
					OptionImageOutputFormat("base64"),
				},
			},
			wantErr: "",
		},
		{
			name: "could not process request",
			path: pathScreenshotAPIResponseError,
			args: args{
				ctx: ctx,
				options: options{
					"whoisxmlapi.com",
					OptionImageOutputFormat("base64"),
				},
			},
			wantErr: "API failed with status code: 499",
		},
		{
			name: "invalid argument",
			path: pathScreenshotAPIResponseError,
			args: args{
				ctx: ctx,
				options: options{
					"",
					OptionImageOutputFormat("base64"),
				},
			},
			wantErr: `invalid argument: "URL" can not be empty`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := newAPI(server, tt.path)

			resp, err := api.GetRaw(tt.args.ctx, tt.args.options.mandatory)
			if (err != nil || tt.wantErr != "") && (err == nil || err.Error() != tt.wantErr) {
				t.Errorf("ScreenshotAPI.GetRaw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp != nil && !checkResultRaw(resp.Body) {
				t.Errorf("ScreenshotAPI.GetRaw() got = %v, expected something else", string(resp.Body))
			}
		})
	}
}
