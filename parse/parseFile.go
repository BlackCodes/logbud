package parse

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/rs/zerolog/log"
)

type FileParse struct {
	dir string
	errs []error
}

func NewFileParse(dir string) (*FileParse,error) {
	if len(dir) > 0 {
		if info,err := os.Stat(dir); err != nil || !info.IsDir(){
			log.Err(err).Msg("that is not dir")
			return nil, fmt.Errorf("file path is not dir")
		}
	}
	return &FileParse{
		errs: make([]error,0,10),
		dir:dir,
	},nil

}

func (f *FileParse) Start(p []string) error{
	pathsMap := f.Dispatch(p)
	var wg sync.WaitGroup
	for _, item := range pathsMap {
		paths := item
		if len(item) > 0 {
			wg.Add(1)
			go f.run(paths, &wg)
		}
	}
	wg.Wait()
	if len(f.errs) > 0 {
		for _, item := range f.errs{
			log.Err(item).Send()
		}
		return fmt.Errorf("There are many errs ")
	}
	return nil
}

func (f *FileParse) Dispatch(files []string) map[int][]string {
	n := runtime.NumCPU()
	allocate := make(map[int][]string)
	for i, f := range files {
		m := i % n
		if s, ok := allocate[m]; ok {
			s = append(s, f)
			allocate[m] = s
		} else {
			s := make([]string, 0, len(files)/n+1)
			s = append(s, f)
			allocate[m] = s
		}
	}
	return allocate
}
func (f *FileParse) run(files []string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, p := range files {
		if filepath.Ext(p) != ".go"{
			continue
		}
		by,err := f.execute(p)
		if err != nil {
			f.errs = append(f.errs,err)
			return
		}
		if err := f.writeTo(p,by); err != nil  {
			f.errs = append(f.errs,err)
			return
		}
	}
}

func (f *FileParse) execute(file string) ([]byte, error) {
	_, err := os.Stat(file)
	if err != nil {
		log.Err(err).Msg("open file error")
		return nil, err
	}
	fset := token.NewFileSet()
	fNode, err1 := parser.ParseFile(fset, file, nil, parser.AllErrors)
	if err1 != nil {
		log.Err(err1).Send()
		return nil, err
	}
	vl := &SetLog{fset: fset,dir: f.dir}
	ast.Walk(vl, fNode)
	buf := &bytes.Buffer{}
	if err := format.Node(buf, fset, fNode); err != nil {
		log.Err(err).Send()
		return nil, err
	}
	return buf.Bytes(), nil
}

func (f *FileParse) writeTo(path string,data []byte) error  {
	fh,err := os.Create(path)
	if err !=nil {
		return err
	}
	_,err = fh.Write(data)
	if err == nil {
		err = fh.Sync()
	}
	fh.Close()
	return err
}