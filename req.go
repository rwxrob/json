package json

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// TimeOut is a package global timeout for any of the high-level https
// query functions in this package. The default value is 60 seconds.
var TimeOut int = 60

// Client provides a way to change the default HTTP client for
// any further package HTTP request function calls. By default, it is
// set to http.DefaultClient. This is particularly useful when creating
// mockups and other testing.
var Client = http.DefaultClient

// Req passes the requested method with the given URL, query string, and
// input data values to the HTTP Client and JSON unmarshals the response
// into the data struct passed by pointer (data, which may already
// contain populated data fields). The data (any) argument must be
// a pointer to something that can be unmarshaled as JSON.  Request also
// observes the package global json.TimeOut. Status codes not in th 200s
// range will return an error with the status message. The
// http.DefaultClient is used by default but can be changed by setting
// json.Client. Note that passing the query string as url.Values (3rd
// argument) will automatically add a question mark (?) followed by the
// URL encoded values to the end of the URL which may present a problem
// if the URL already has a query string. Encouraging the use of
// url.Values for passing the query string serves as a reminder that all
// query strings should be URL encoded (as is often forgotten). Req does
// this automatically.
func Req(method, url string, qs, body url.Values, data any) error {
	var err error
	var bodyreader io.Reader
	var bodylength string

	url = url + "?" + qs.Encode()

	if body != nil {
		encoded := body.Encode()
		bodyreader = strings.NewReader(encoded)
		bodylength = strconv.Itoa(len(encoded))
	}

	req, err := http.NewRequest(method, url, bodyreader)
	if body != nil {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", bodylength)
	}
	if err != nil {
		return err
	}

	dur := time.Duration(time.Second * time.Duration(TimeOut))
	ctx, cancel := context.WithTimeout(context.Background(), dur)
	defer cancel()
	req = req.WithContext(ctx)

	res, err := Client.Do(req)
	if err != nil {
		return err
	}

	if !(200 <= res.StatusCode && res.StatusCode < 300) {
		return fmt.Errorf(res.Status)
	}

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, data)
}
