package screenshotapi

import (
	"fmt"
	"strings"
)

// Cookies is a help wrapper on map.
type Cookies map[string]string

// toString converts Cookies map to string in the following format: name1=value1;name2=value2.
func (c Cookies) toString() string {
	str := make([]string, 0, len(c))
	for key, value := range c {
		str = append(str, key+"="+value)
	}
	return strings.Join(str, ";")
}

// ErrorMessage is the error message.
type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"messages"`
}

// Error returns error message as a string.
func (e *ErrorMessage) Error() string {
	return fmt.Sprintf("API error: [%d] %s", e.Code, e.Message)
}
