/*
Package json contains interface specifications for representing any Go
type as JSON where possible. Using the goprintasjson tool allows for quick code generation of scaffolding to make any Go type easily used as JSON.
*/
package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/rwxrob/to"
	yq "github.com/rwxrob/yq/pkg"
)

// AsJSON specifies a type that must support marshaling using the
// rwxrob/json package with its defaults for marshaling and unmarshaling
// which do not have unnecessary escaping.
//
// String is from fmt.Stringer, but fulfilling this interface in this
// package promises to render the string specifically using rwxrob/json
// default output marshaling --- especially when it comes to consistent
// indentation, wrapping, and escaping. While JSON is a flexible format,
// consistency ensures the most efficient and sustainable creation of
// tests and other systems that require such consistency, whether or not
// dependency on such consistency is a "good idea".
//
// Printer specifies methods for printing self as JSON and will log any
// error if encountered. Printer provides a consistent representation of
// any structure such that it an easily be read and compared as JSON
// whenever printed and test. Sadly, the default string representations
// for most types in Go are virtually unusable for consistent
// representations of any structure. And while it is true that JSON data
// should be supported in any way that is it presented, some consistent
// output makes for more consistent debugging, documentation, and
// testing.
//
// AsJSON implementations must Print and Log the output of String from
// the same interface.
//
// MarshalJSON and UnmarshalJSON must be explicitly defined and use the
// rwxrob/json package to avoid confusion. Use of the helper json.This
// struct may facilitate this for existing types that do not wish to
// implement the full interface.
type AsJSON interface {
	JSON() ([]byte, error)
	String() string
	Print()
	Log() string
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(buf []byte) error
}

// specification (unlike the encoding/json standard which defaults to
// escaping many other characters as well unnecessarily).
func Escape(in string) string {
	out := ``
	for _, r := range in {
		switch r {
		case '\t':
			out += `\t`
		case '\b':
			out += `\b`
		case '\f':
			out += `\f`
		case '\n':
			out += `\n`
		case '\r':
			out += `\r`
		case '\\':
			out += `\\`
		case '"':
			out += `\"`
		default:
			out += string(r)
		}
	}
	return out
}

// Marshal mimics json.Marshal from the encoding/json package without
// the broken, unnecessary HTML escapes and extraneous newline that the
// json.Encoder adds. Call this from your own MarshalJSON methods to get
// JSON rendering that is more readable and compliant with the JSON
// specification (unless you are using the extremely rare case of
// dumping that into HTML, for some reason). Note that this cannot be
// called from any structs MarshalJSON method on itself because it will
// cause infinite functional recursion. Write a proper MarshalJSON
// method or create a dummy struct and call json.Marshal on that
// instead.
func Marshal(v any) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(v)
	return []byte(strings.TrimSpace(buf.String())), err
}

// MarshalIndent mimics json.Marshal from the encoding/json package but
// without the escapes, etc. See Marshal.
func MarshalIndent(v any, a, b string) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent(a, b)
	err := enc.Encode(v)
	return []byte(strings.TrimSpace(buf.String())), err
}

// Unmarshal mimics json.Unmarshal from the encoding/json package.
func Unmarshal(buf []byte, v any) error {
	return json.Unmarshal(buf, v)
}

// This encapsulates anything with the AsJSON interface from this package
// by simply assigning a new variable with that item as the only value
// in the structure:
//
//     something := []string{"some","thing"}
//     jsonified := json.This{something}
//     jsonified.Print()
//
type This struct{ This any }

// UnmarshalJSON implements AsJSON
func (s *This) UnmarshalJSON(buf []byte) error {
	return json.Unmarshal(buf, &s.This)
}

// JSON implements AsJSON.
func (s This) JSON() ([]byte, error) { return json.Marshal(s.This) }

// String implements AsJSON and logs any error.
func (s This) String() string {
	byt, err := s.JSON()
	if err != nil {
		log.Print(err)
	}
	return string(byt)
}

// Print implements AsJSON printing with fmt.Println (adding a line
// return).
func (s This) Print() { fmt.Println(s.String()) }

// Log implements AsJSON.
func (s This) Log() { log.Print(s.String()) }

// Query provides YAML/JSON query responses.
func (s This) Query(q string) (string, error) {
	return yq.EvaluateToString(to.String(s.This), q)
}

// QueryPrint prints YAML/JSON query responses.
func (s This) QueryPrint(q string) error {
	return yq.Evaluate(to.String(s.This), q)
}
