package version_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/VKCOM/php-parser/pkg/version"
)

func Test(t *testing.T) {
	ver, err := version.New("7.4")
	assert.NilError(t, err)

	assert.Equal(t, *ver, version.Version{
		Major: 7,
		Minor: 4,
	})
}

func TestLeadingZero(t *testing.T) {
	ver, err := version.New("07.04")
	assert.NilError(t, err)

	assert.Equal(t, *ver, version.Version{
		Major: 7,
		Minor: 4,
	})
}

func TestInRange(t *testing.T) {
	s, err := version.New("7.0")
	assert.NilError(t, err)

	e, err := version.New("7.4")
	assert.NilError(t, err)

	ver, err := version.New("7.0")
	assert.NilError(t, err)
	assert.Assert(t, ver.InRange(s, e))

	ver, err = version.New("7.2")
	assert.NilError(t, err)
	assert.Assert(t, ver.InRange(s, e))

	ver, err = version.New("7.4")
	assert.NilError(t, err)
	assert.Assert(t, ver.InRange(s, e))
}

func TestInRangePHP8(t *testing.T) {
	s, err := version.New("8.0")
	assert.NilError(t, err)

	e, err := version.New("8.1")
	assert.NilError(t, err)

	ver, err := version.New("8.0")
	assert.NilError(t, err)
	assert.Assert(t, ver.InRange(s, e))

	ver, err = version.New("8.1")
	assert.NilError(t, err)
	assert.Assert(t, ver.InRange(s, e))
}
