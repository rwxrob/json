/*
Package json contains interface specifications for representing any Go
type as JSON where possible.
*/
package json

// AsJSON specifies types that can represent themselves as JSON.
type AsJSON interface {
	JSON() (string, error)  // single line, no spaces
	JSONL() (string, error) // 2-space indent and separation
}

// Stringer specifies that fmt.Stringer is fulfilled as JSON and will
// log any error if encountered. See AsJSON.
type Stringer interface {
	String() string
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
	PPrint() string
}

// Logger specifies methods for logging self as short and long JSON
// printing separately to the log if any error.
type Logger interface {
	Log() string
	PLog() string
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
