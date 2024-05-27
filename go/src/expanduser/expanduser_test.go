package expanduser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpandUser(t *testing.T) {
	v, err := ExpandUser("~")
	assert.Equal(t, v, "/home/nikita")
	assert.Nil(t, err)
	v, err = ExpandUser("~gnunet")
	assert.Equal(t, v, "/var/chroot/gnunet")
	assert.Nil(t, err)
	v, err = ExpandUser("~/~/foo/bar")
	assert.Equal(t, v, "/home/nikita/~/foo/bar")
	assert.Nil(t, err)
}
