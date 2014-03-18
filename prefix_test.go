package prefixes

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var prefixTests = []struct {
	number string
	name   string
}{
	{"46485562003", "Sweden"},
	{"61296384929", "Australia"},
	{"74956034085", "Russia"},
	{"14158575430", "United States"},
	{"14169323338", "Canada"},
	{"12504798004", "Canada"},
	{"6493662929", "New Zealand"},
	{"420737471423", "Czech Republic"},
	{"390669812345", "Vatican City"},
	{"3799812345", "Vatican City"},
}

var invalidPrefixes = []string{
	"",
}

func TestPrefixes(t *testing.T) {

	for i, test := range prefixTests {
		msg := fmt.Sprintf("Test %d", i)
		c, err := Lookup(test.number)
		assert.NoError(t, err, msg)
		assert.Equal(t, test.name, c.Name, msg)
	}

}

func TestInvalidPrefixes(t *testing.T) {
	for i, prefix := range invalidPrefixes {
		msg := fmt.Sprintf("Test %d", i)
		_, err := Lookup(prefix)
		assert.Error(t, err, msg)
	}
}
