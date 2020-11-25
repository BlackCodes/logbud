package files

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/BlackCodes/logbud/flag"
)

type CloneFile struct {
	project string
	ModFile string
	dir     string
	buildDir string
	paths   map[string]string // original==>new
}
func NewCloneFile() *CloneFile{
	return &CloneFile{}
}

func (c *CloneFile) Start() error {
	if !c.FindMod() {
		return fmt.Errorf("Not found mod.go file ")
	}
	p, err := ParseMod(c.ModFile);
	if err != nil {
		return err
	}
	c.project = p

	if c.paths,err = c.ScanFile();err !=nil {
		return err
	}
	if err := c.CloneFiles(c.paths); err !=nil {
		return err
	}

	return nil
}

func (c *CloneFile) ScanFile() (map[string]string,error) {
	paths := map[string]string{}
	specialFiles := map[string]struct{}{"go.sum": {},"go.mod": {}}
	err:= filepath.Walk(c.getDir(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Err(err).Msg("scan file err")
			return err
		}
		ext := filepath.Ext(info.Name())
		_,ok := specialFiles[info.Name()];
		if ok || (ext == ".go" && !strings.HasSuffix(info.Name(),"_test"+ext)){
			paths[path] = c.newPath(path)
		}

		return nil
	})

	return paths,err
}

func (c *CloneFile) CloneFiles(data map[string]string) error {
	for old,_ := range data{
		if info,err := os.Stat(old); err !=nil {
			return err
		} else if !info.Mode().IsRegular(){
			return fmt.Errorf("%s,The file is not regular",old)
		}
	}
	for s,d := range data{
		source,err := os.Open(s)
		if err !=nil {
			return err
		}
		if _,err := os.Stat(d);os.IsNotExist(err){
			os.MkdirAll(filepath.Dir(d),os.ModePerm)
		}
		dist,err := os.Create(d)
		if err !=nil {
			source.Close()
			return err
		}
		_,err = io.Copy(dist,source);
		source.Close()
		dist.Close()
		if err != nil {
			return err
		}
	}
	return  nil
}

func (c *CloneFile) FindMod() bool {
	pwd := c.getDir()
	fileInfo, err := ioutil.ReadDir(pwd)
	if err != nil {
		return false
	}
	for _, item := range fileInfo {
		if item.Name() == "go.mod" {
			c.ModFile = fmt.Sprintf("%s/go.mod", pwd)
			return true
		}
	}
	return true
}

func (c *CloneFile) GetNewPaths() []string  {
	arr := make([]string,0,len(c.paths))
	for _,p := range c.paths{
		arr = append(arr,p)
	}
	return arr
}

func (c *CloneFile) GetDir() string  {
	return c.dir
}

func (c *CloneFile) GetBuildDir() string   {
	return c.buildDir
}

func (c *CloneFile) Destroy() error {
	return os.RemoveAll(c.buildDir)
}

func (c *CloneFile) newPath(file string ) string  {
	pathArr := strings.Split(file,"/")
	dirArr := strings.Split(c.buildDir,"/")
	dirArr = append(dirArr,pathArr[len(dirArr):]...)
	newPath := strings.Join(dirArr,"/")
	return  newPath
}

func (c *CloneFile) getDir() string  {
	if len(c.dir) == 0 {
		c.dir,_ = os.Getwd()
		c.buildDir = c.dir+flag.BuildDirSufFix
	}
	return c.dir
}