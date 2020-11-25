package files

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloneFile_CloneFile(t *testing.T) {

	c := NewCloneFile()

	paths,err := c.ScanFile()
	assert.NoError(t, err)
	t.Run("clone files", func(t *testing.T) {
		err = c.CloneFiles(paths)
		assert.NoError(t, err)
	})

	t.Run("destroy", func(t *testing.T) {
		err := c.Destroy()
		assert.NoError(t, err)
	})
}
