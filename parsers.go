package kingpin

import (
	"net"
	"net/url"
	"os"
	"time"
)

type Settings interface {
	SetValue(value Value)
}

type FlagSettings interface {
	SetIsBoolean()
}

type parserMixin struct {
	value    Value
	required bool
}

// String sets the parser to a string parser.
func (p *parserMixin) String() (target *string) {
	target = new(string)
	p.SetValue(newStringValue("", target))
	return
}

// Strings appends multiple occurrences to a string slice.
func (p *parserMixin) Strings() (target *[]string) {
	target = new([]string)
	p.SetValue(newStringsValue(target))
	return
}

// StringMap provides key=value parsing into a map.
func (p *parserMixin) StringMap() (target *map[string]string) {
	target = &map[string]string{}
	p.SetValue(newStringMapValue(target))
	return
}

// Bool sets the parser to a boolean parser. Supports --no-<X> to disable the flag.
func (p *parserMixin) Bool() (target *bool) {
	target = new(bool)
	p.SetValue(newBoolValue(false, target))
	return
}

// Int sets the parser to an int parser.
func (p *parserMixin) Int() (target *int) {
	target = new(int)
	p.SetValue(newIntValue(0, target))
	return
}

// Int64 parses an int64
func (p *parserMixin) Int64() (target *int64) {
	target = new(int64)
	p.SetValue(newInt64Value(0, target))
	return
}

// Uint64 parses a uint64
func (p *parserMixin) Uint64() (target *uint64) {
	target = new(uint64)
	p.SetValue(newUint64Value(0, target))
	return
}

// Float sets the parser to a float64 parser.
func (p *parserMixin) Float() (target *float64) {
	target = new(float64)
	p.SetValue(newFloat64Value(0, target))
	return
}

// Duration sets the parser to a time.Duration parser.
func (p *parserMixin) Duration() (target *time.Duration) {
	target = new(time.Duration)
	p.SetValue(newDurationValue(time.Duration(0), target))
	return
}

// IP sets the parser to a net.IP parser.
func (p *parserMixin) IP() (target *net.IP) {
	target = new(net.IP)
	p.SetValue(newIPValue(target))
	return
}

// ExistingFile sets the parser to one that requires and returns an existing file.
func (p *parserMixin) ExistingFile() (target *string) {
	target = new(string)
	p.SetValue(newFileStatValue(target, func(s os.FileInfo) bool { return !s.IsDir() }))
	return
}

// ExistingDir sets the parser to one that requires and returns an existing directory.
func (p *parserMixin) ExistingDir() (target *string) {
	target = new(string)
	p.SetValue(newFileStatValue(target, func(s os.FileInfo) bool { return s.IsDir() }))
	return
}

// File sets the parser to one that requires and opens a valid os.File.
func (p *parserMixin) File() (target **os.File) {
	target = new(*os.File)
	p.SetValue(newFileValue(target))
	return
}

// URL provides a valid, parsed url.URL.
func (p *parserMixin) URL() (target **url.URL) {
	target = new(*url.URL)
	p.SetValue(newURLValue(target))
	return
}

func (p *parserMixin) SetValue(value Value) {
	p.value = value
}
