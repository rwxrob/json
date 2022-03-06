/*
Package json contains the Printer interface and functions for marshaling
and unmarshaling that do not depend on Go reflection and do not escape
any characters unnecessarily (such as <> and & which are escaped by
default in the standard encoding/json package).
*/
package json

// Printer provides a consistent representation of any structure
// such that it an easily be read and compared as JSON whenever printed
// and test. Sadly, the default string representations for most types in
// Go are virtually unusable for consistent representations of any
// structure. And while it is true that JSON data should be supported in
// any way that is it presented, some consistent output makes for more
// consistent debugging, documentation, and testing.
//
// Printer implementations must fulfill their marshaling to any
// textual/printable representation by producing compressed,
// single-line, no-spaces, parsable, shareable,
// not-unnecessarily-escaped JSON data. When indented, long-form JSON is
// wanted utility functions and utilities (such as jq) can be used to
// expand the compressed, default JSON.
//
// All implementations must also implement the fmt.Stringer interface by
// calling JSON(), which most closely approximates the Go standard
// string marshalling. Thankfully, Go does ensure that the order of
// elements in any type will appear consistently in that same order
// during testing even though they should never be relied upon for such
// ordering other than in testing.
//
// All implementations must promise they will never escape any string
// character that is not specifically required to be escaped by the JSON
// standard as described in this PEGN specification:
//
//     String  <-- DQ (Escaped / [x20-x21] / [x23-x5B]
//                 / [x5D-x10FFFF])* DQ
//     Escaped  <- BKSLASH ("b" / "f" / "n" / "r" / "t" / "u" hex{4}
//                 / DQ / BKSLASH / SLASH)
//
// This means that binary data will always be converted to base64
// encoding and expressed as a string value.
//
// In general, implementers of Printer should not depend on
// structure tagging and reflection for unmarshaling instead
// implementing their own consistent UnmarshalJSON method. This allows
// for better error checking as the default does nothing to ensure that
// unmarshaled values are within acceptable ranges. Errors are only
// generated if the actual JSON syntax itself is incorrect.
type Printer interface {
	// MarshalJSON() ([]byte, error) // encouraged, but not required
	JSON() string   // single line, no spaces, no uneeded escapes
	JSONL() string  // 2-space indent and separation
	String() string // must return s.JSON()
	Print() string  // must call fmt.Println(s.JSON())
	PPrint() string // must call fmt.Println(s.JSONL())
	Log() string    // must call log.Print(s.JSON())
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
