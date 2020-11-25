package parse

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/BlackCodes/logbud/flag"
)

func TestParsFile(t *testing.T) {


		var crateFile = func(f string) {
			dst,err := os.Create(f)
			assert.NoError(t,err)
			defer dst.Close()
			file := "../example/example.go"
			source ,err := os.Open(file)
			defer source.Close()
			assert.NoError(t,err)
			_,err = io.Copy(dst,source)
			assert.NoError(t,err)
		}
		t.Run("Position Tail", func(t *testing.T) {
			newFile := "../example/example_tail.go"
			crateFile(newFile)
			flag.Position = flag.PositionTail
			fp,err := NewFileParse("../example")
			assert.NoError(t,err)
			err = fp.Start([]string{newFile})
			assert.NoError(t,err)
		})

	t.Run("Position Head", func(t *testing.T) {
		newFile := "../example/example_head.go"
		crateFile(newFile)
		flag.Position = flag.PositionHead
		fp,err := NewFileParse("../example")
		assert.NoError(t,err)
		err = fp.Start([]string{newFile})
		assert.NoError(t,err)
	})

	t.Run("PathMod Relative", func(t *testing.T) {
		newFile := "../example/example_pathmod_relative.go"
		crateFile(newFile)
		flag.PathMod = flag.PathModRelative
		fp,err := NewFileParse("../example")
		assert.NoError(t,err)
		err = fp.Start([]string{newFile})
		assert.NoError(t,err)
	})

	t.Run("PathMod File", func(t *testing.T) {
		newFile := "../example/example_pathmod_file.go"
		crateFile(newFile)
		flag.PathMod = flag.PathModFile
		fp,err := NewFileParse("../example")
		assert.NoError(t,err)
		err = fp.Start([]string{newFile})
		assert.NoError(t,err)
	})


}
