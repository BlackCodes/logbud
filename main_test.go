package main

import (
	"testing"

	"github.com/BlackCodes/logbud/files"
	"github.com/stretchr/testify/assert"
)

func TestCloneStart(t *testing.T) {
	c := files.NewCloneFile()

	t.Run("parse mod", func(t *testing.T) {
		b := c.FindMod()
		assert.Equal(t,true,b,"find mod file")

		p,err := files.ParseMod(c.ModFile)
		assert.NoError(t,err)
		t.Logf("according to the mod get projectName:%s",p)
	})

	t.Run("clone start", func(t *testing.T) {
		err := c.Start()
		assert.NoError(t,err)
	})

	t.Run("destroy", func(t *testing.T) {
		err := c.Destroy()
		assert.NoError(t,err )
	})
}