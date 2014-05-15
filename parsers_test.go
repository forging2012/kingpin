package kingpin

import (
	"net"
	"net/url"

	"github.com/stretchrcom/testify/assert"

	"testing"
)

func TestParseStrings(t *testing.T) {
	p := parserMixin{}
	v := p.Strings()
	p.value.Set("a")
	p.value.Set("b")
	assert.Equal(t, []string{"a", "b"}, *v)
}

func TestParseStringMap(t *testing.T) {
	p := parserMixin{}
	v := p.StringMap()
	p.value.Set("a=b")
	p.value.Set("b=c")
	assert.Equal(t, map[string]string{"a": "b", "b": "c"}, *v)
}

func TestParseIP(t *testing.T) {
	p := parserMixin{}
	v := p.IP()
	p.value.Set("10.1.1.2")
	ip := net.ParseIP("10.1.1.2")
	assert.Equal(t, ip, *v)
}

func TestParseURL(t *testing.T) {
	p := parserMixin{}
	v := p.URL()
	p.value.Set("http://w3.org")
	u, err := url.Parse("http://w3.org")
	assert.NoError(t, err)
	assert.Equal(t, *u, **v)
}

func TestParseExistingFile(t *testing.T) {
	p := parserMixin{}
	v := p.ExistingFile()
	err := p.value.Set("/etc/hosts")
	assert.NoError(t, err)
	assert.Equal(t, "/etc/hosts", *v)
	err = p.value.Set("/etc/hostsDEFINITELYMISSING")
	assert.Error(t, err)
}
