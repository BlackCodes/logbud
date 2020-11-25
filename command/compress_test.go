package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadData(t *testing.T) {
	u := NewUpx("")
	t.Log(u.osArch())

}

func TestNewBuild(t *testing.T) {
	u := NewUpx("")

	t.Run("generate file", func(t *testing.T) {
		err := u.generateBinary()
		assert.NoError(t,err)
	})

}
