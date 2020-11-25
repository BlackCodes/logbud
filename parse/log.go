package parse

import (
	"fmt"
	"go/ast"
	"go/token"
	"path/filepath"
	"strings"

	"github.com/BlackCodes/logbud/flag"
)

type SetLog struct {
	fset *token.FileSet
	dir  string
}

func (s *SetLog) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	if fn, ok := n.(*ast.CallExpr); ok {
		s.SetLog(fn)
	}
	return s
}

func (s *SetLog) SetLog(fn *ast.CallExpr) {
	if fn == nil {
		return
	}
	val := func(s string) string{
		if len(s) <= 1 {
			return s
		}
		return  s[1:len(s)-1]
	}
	replaceExpr := map[string]struct{}{"Println": {}, "Printf": {}, "Infof": {}, "Info": {}, "Debugf": {}, "Debug": {}, "Warnf": {}, "Warn": {}, "Errorf": {}, "Error": {}, "Fatalf": {}, "Fatal": {}, "Msgf": {}, "Msg": {}}
	if cf, ok := fn.Fun.(*ast.SelectorExpr); ok {
		if _, ok := replaceExpr[cf.Sel.String()]; ok && len(fn.Args) > 0 {
			if str, ok := fn.Args[0].(*ast.BasicLit); ok {
				line := fmt.Sprintf(`"[%s] %s"`, s.showFile(s.fset.Position(cf.Pos()).String()),val(str.Value))
				if flag.Position == flag.PositionTail {
					line = fmt.Sprintf(`"%s [%s]"`, val(str.Value), s.showFile(s.fset.Position(cf.Pos()).String()))
				}
				str.Value = line
			}
		}
	}

}

func (s *SetLog) showFile(fp string) string {
	str := fp
	switch flag.PathMod {
	case flag.PathModRelative:
		if len(s.dir) > 0 {
			str = fp[len(s.dir):]
		}
	case flag.PathModFile:
		str = filepath.Base(fp)
	default:
	}
	return strings.Replace(str, flag.BuildDirSufFix, "", 1)
}
