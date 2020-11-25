package command

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/BlackCodes/logbud/flag"
	"github.com/rs/zerolog/log"
)

type Build struct {
	projectDir string
	buildDir   string
}

func NewBuild(buildDir, projectDir string) *Build {
	return &Build{projectDir: projectDir, buildDir: buildDir}
}

func (b *Build) Build() error {
	build := "go build "
	if len(flag.BuildArgs) > 0 {
		build += flag.BuildArgs
	}

	if len(b.projectDir) > 0 {
		os.Chdir(b.buildDir)
	}
	if err := b.clearBinary(); err != nil {
		log.Err(err).Msg("clear binary error")
		return err
	}
	fmt.Println(build)
	cmd := exec.Command("/bin/bash", "-c", build)
	_, err := cmd.Output()
	if err != nil {
		log.Err(err).Msg("build failed")
		return err
	}
	return nil
}

func (b *Build) GetBinary() (string, error) {
	bf, err := b.findBinary()
	if err != nil {
		return "", err
	}
	if len(bf) > 0 {
		return bf[0], nil
	}
	return "", fmt.Errorf("nof found binary")
}

func (b *Build) Copy() error {
	bf, err := b.findBinary()
	if err != nil {
		return err
	}

	for _, s := range bf {
		newPath := fmt.Sprintf("%s/%s", b.projectDir, filepath.Base(s))
		if err := os.Rename(s, newPath); err != nil {
			return err
		}
	}
	return nil
}

func (b *Build) clearBinary() error {
	files, err := b.findBinary()
	if err != nil {
		return err
	}
	for _, item := range files {
		if err := os.Remove(item); err != nil {
			log.Err(err).Msg("remove exist binary file error")
			return err
		}
	}
	return nil
}
func (b *Build) findBinary() ([]string, error) {
	files, err := filepath.Glob(b.buildDir + "/*")
	binaries := make([]string, 0, 5)
	if err != nil {
		return binaries, err
	}

	for _, f := range files {
		fh, err := os.Open(f)
		if err != nil {
			log.Err(err).Msg("open file error")
			continue
		}
		info, _ := fh.Stat()
		if info.IsDir() {
			continue
		}
		mType, err := b.mimeType(fh)
		if err != nil {
			log.Err(err).Str("file", f).Msg("get file mime type error")
			continue
		}
		if strings.HasPrefix(mType, "application") {
			binaries = append(binaries, f)
		}
	}
	return binaries, nil
}

func (b *Build) mimeType(f *os.File) (string, error) {
	buf := make([]byte, 512)
	if _, err := f.Read(buf); err != nil {
		return "", err
	}
	mtype := http.DetectContentType(buf)
	return mtype, nil
}
