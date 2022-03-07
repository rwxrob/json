/*
Package json contains interface specifications for representing any Go
type as JSON where possible.
*/
package json

// AsJSON specifies types that can represent themselves as JSON both in
// a single-line form with no spaces and an indented ("pretty") form
// with consistent 2-space indentation and separation.
type AsJSON interface {
	JSON() (string, error)  // single line, no spaces
	JSONI() (string, error) // 2-space indent and separation
}

// Stringer specifies that rwxrob/to.Stringer is fulfilled as JSON and
// will log any error if encountered. See AsJSON.
type Stringer interface {
	String() string
}

// StringerLong specifies that rwxrob/to.Stringer is fulfilled as JSONI
// and will log any error if encountered. See AsJSON.
type StringerLong interface {
	StringLong() string
}

// Printer specifies methods for printing self as JSON and will log any
// error if encountered. Printer provides a consistent representation of
// any structure such that it an easily be read and compared as JSON
// whenever printed and test. Sadly, the default string representations
// for most types in Go are virtually unusable for consistent
// representations of any structure. And while it is true that JSON data
// should be supported in any way that is it presented, some consistent
// output makes for more consistent debugging, documentation, and
// testing.
type Printer interface {
	Print() string
	PrintLong() string
}

// Logger specifies methods for logging self as short and long JSON
// printing separately to the log if any error.
type Logger interface {
	Log() string
	LogLong() string
}

// Escape only escapes those runes that require it according to the JSON
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
