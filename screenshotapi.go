package screenshotapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

// ScreenshotAPIService is an interface for Screenshot API.
type ScreenshotAPIService interface {
	// Get captures a screenshot to a file, or returns a parsed Screenshot API error.
	Get(ctx context.Context, url string, filename string, opts ...Option) error

	// GetRaw returns raw Screenshot API response as the Response struct with Body saved as a byte slice.
	GetRaw(ctx context.Context, URL string, opts ...Option) (*Response, error)
}

// Response is the http.Response wrapper with Body saved as a byte slice.
type Response struct {
	*http.Response

	// Body is the byte slice representation of http.Response Body
	Body []byte
}

// screenshotAPIServiceOp is the type implementing the ScreenshotAPI interface.
type screenshotAPIServiceOp struct {
	client  *Client
	baseURL *url.URL
}

var _ ScreenshotAPIService = &screenshotAPIServiceOp{}

// newRequest creates the API request with default parameters and the specified apiKey.
func (service screenshotAPIServiceOp) newRequest() (*http.Request, error) {
	req, err := service.client.NewRequest(http.MethodGet, service.baseURL, nil)
	if err != nil {
		return nil, err
	}

	query := url.Values{}
	query.Set("apiKey", service.client.apiKey)

	req.URL.RawQuery = query.Encode()

	return req, nil
}

// setOptions sets options provided as arguments.
func setOptions(v url.Values, opts ...Option) error {
	for _, opt := range opts {
		if opt == nil {
			return &ArgError{"Option", "can not be nil"}
		}
		if err := opt(v); err != nil {
			return err
		}
	}

	width, err := strconv.Atoi(v.Get("width"))
	if err != nil {
		width = 0
	}
	thumbWidth, err := strconv.Atoi(v.Get("thumbWidth"))
	if err != nil {
		thumbWidth = 0
	}
	if (width != 0 && thumbWidth > width) || (width == 0 && thumbWidth > DefaultWidth) {
		return &ArgError{"thumbWidth", "must be between 50 and width param value"}
	}

	return nil
}

// request returns intermediate API response for further actions.
func (service screenshotAPIServiceOp) request(ctx context.Context, url string, opts ...Option) (*Response, error) {
	if url == "" {
		return nil, &ArgError{"URL", "can not be empty"}
	}

	req, err := service.newRequest()
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("url", url)

	if err = setOptions(q, opts...); err != nil {
		return nil, err
	}

	req.URL.RawQuery = q.Encode()

	var b bytes.Buffer

	resp, err := service.client.Do(ctx, req, &b)
	if err != nil {
		return &Response{
			Response: resp,
			Body:     b.Bytes(),
		}, err
	}

	return &Response{
		Response: resp,
		Body:     b.Bytes(),
	}, nil
}

// apiResponse is used for parsing Screenshot API response as a model instance.
type apiResponse struct {
	ErrorMessage
}

// parse parses raw Screenshot API response.
func parse(raw []byte) (*apiResponse, error) {
	var response apiResponse

	err := json.NewDecoder(bytes.NewReader(raw)).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("cannot parse response: %w", err)
	}

	return &response, nil
}

// Get captures a screenshot to a file, or returns a parsed Screenshot API error.
func (service screenshotAPIServiceOp) Get(
	ctx context.Context,
	url string,
	filename string,
	opts ...Option,
) (err error) {
	if filename == "" {
		return &ArgError{"filename", "can not be empty"}
	}

	optsJSON := make([]Option, 0, len(opts)+2)
	optsJSON = append(optsJSON, opts...)
	optsJSON = append(optsJSON,
		OptionErrorsOutputFormat("JSON"),
		OptionImageOutputFormat("image"),
	)

	resp, err := service.request(ctx, url, optsJSON...)
	if err != nil {
		return err
	}

	screenshotAPIResp, err := parse(resp.Body)
	if err == nil && (screenshotAPIResp.Message != "" || screenshotAPIResp.Code != 0) {
		return &ErrorMessage{
			Code:    screenshotAPIResp.Code,
			Message: screenshotAPIResp.Message,
		}
	}

	if respErr := checkResponse(resp.Response); respErr != nil {
		return respErr
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)

	_, err = f.Write(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// GetRaw returns raw Screenshot API response as the Response struct with Body saved as a byte slice.
func (service screenshotAPIServiceOp) GetRaw(
	ctx context.Context,
	url string,
	opts ...Option,
) (resp *Response, err error) {
	resp, err = service.request(ctx, url, opts...)
	if err != nil {
		return resp, err
	}

	if respErr := checkResponse(resp.Response); respErr != nil {
		return resp, respErr
	}

	return resp, nil
}

// ArgError is the argument error.
type ArgError struct {
	Name    string
	Message string
}

// Error returns error message as a string.
func (a *ArgError) Error() string {
	return `invalid argument: "` + a.Name + `" ` + a.Message
}
