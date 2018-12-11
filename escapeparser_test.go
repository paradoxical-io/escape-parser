package escaper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser(t *testing.T) {
	parsed := Parse("foo bar biz")

	assert.Equal(t, []string{"foo", "bar", "biz"}, parsed)
}

func TestParserEscaped(t *testing.T) {
	parsed := Parse("foo \"bar\" biz")

	assert.Equal(t, []string{"foo", "bar", "biz"}, parsed)
}

func TestParserEscapedWithSpaces(t *testing.T) {
	parsed := Parse("foo \"test me!\" biz")

	assert.Equal(t, []string{"foo", "test me!", "biz"}, parsed)
}

func TestParserEscapedWithEndingEscapes(t *testing.T) {
	parsed := Parse("foo \"test me! biz\"")

	assert.Equal(t, []string{"foo", "test me! biz"}, parsed)
}

func TestParserEscapedWithQuotesInEscaped(t *testing.T) {
	parsed := Parse("foo \"test \\\"me!\\\" biz\"")

	assert.Equal(t, []string{"foo", "test \\\"me!\\\" biz"}, parsed)
}

func TestParserEscapedWithQuotesInEscapedMultipleSpaces(t *testing.T) {
	parsed := Parse("foo    \"test   \\\"me!\\\" biz\"   ")

	assert.Equal(t, []string{"foo", "test   \\\"me!\\\" biz"}, parsed)
}
