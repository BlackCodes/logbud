package command

import (
	"encoding/base64"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/rs/zerolog/log"
)

var NoSupport = errors.New("Current Os not support compress ")

type Upx struct {
	binPath      string
	compressPath string
}

func NewUpx(cPath string) *Upx {
	return &Upx{compressPath: cPath}
}

func (u *Upx) Build() error {
	if err := u.generateBinary(); err != nil {
		log.Err(err).Msg("generate build upx binary")
		if err == NoSupport {
			return nil
		}
		return err
	}
	c := exec.Command(u.binPath, u.compressPath)
	_, err := c.Output()
	if err != nil {
		return err
	}
	return u.destroy()
}

func (u *Upx) destroy() error{
	return os.RemoveAll(filepath.Dir(u.binPath))
}
func (u *Upx) generateBinary() error {
	f := func(s string) error {
		dir, _ := os.Getwd()
		dir = dir + "/bin"
		if _, err := os.Stat(dir); err != nil {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
		}
		p := dir + "/compress_build"
		fh, err := os.Create(p)
		if err != nil {
			return err
		}
		defer fh.Close()
		b, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			return err
		}
		if _, err := fh.Write(b); err != nil {
			return err
		}
		os.Chmod(p, 0777)
		u.binPath = p
		return nil
	}
	switch u.osArch() {
	case "darwin":
		return f(darwin)
	default:
		return NoSupport
	}
	return nil
}

func (u *Upx) osArch() string {
	return runtime.GOOS
}

//func (u *Upx )ReadData() {
//	fh, err := os.Open("")
//	defer fh.Close()
//	fmt.Println(err)
//	by, err := ioutil.ReadAll(fh)
//	s := base64.StdEncoding.EncodeToString(by)
//	str := `package command
//var darwin string ="%s"`
//	b := fmt.Sprintf(str, s)
//	nf, err := os.Create("darwin.go")
//	fmt.Println(err)
//	nf.WriteString(b)
//	nf.Close()
//}
