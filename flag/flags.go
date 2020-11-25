package flag

import (
	"flag"
	"fmt"
	"os"
	"strings"
)
const(
	BuildDirSufFix = "_build"
	PathModFull     = "full"
	PathModRelative = "relative"
	PathModFile     = "file"
	PositionHead = "head"
	PositionTail = "tail"
	IsNoCompress = false
)

var (
	PathMod string // PathMod:full、relative、file. full path is absolute path ('/aa/bb/xx.go') relative is from project begin,file path is only file name
	BuildArgs string // BuildArgs: exclude 'go build' keywords, and use go build command arguments
	Position string // Position: where is the position inert it, head or tail, default head
	Compress bool // Position: compress go binary file. ture or false
)

func init() {
	flag.StringVar(&PathMod,"pmod",lookupStrEnv("PATH_MOD",PathModRelative),"Env $PATH_MOD,there are three options:full、relative、file. full path is absolute path ('/aa/bb/xx.go') relative is from project begin,file path is only file name")
	flag.StringVar(&BuildArgs,"bargs",lookupStrEnv("BUILD_ARGS",""),`Env $BUILD_ARGS,build go source,exclude 'go build' keywords, and use go build command arguments，default empty,example:-bagrs="-o helloGood"`)
	flag.StringVar(&Position,"pos",lookupStrEnv("POSITION",PositionHead),"Env $POSITION,where is the position inert it, head or tail")
	flag.BoolVar(&Compress,"cp",lookupBoolEnv("COMPRESS",IsNoCompress),"Env $COMPRESS,it can be compress go binary file. ture or false")
	flag.CommandLine.Usage= func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	allowPathMod := map[string]struct{}{PathModFull: {}, PathModRelative:{}, PathModFile:{}}
	if _,ok := allowPathMod[PathMod]; !ok {
		PathMod = PathModFull
	}

	if Position != PositionHead && Position != PositionTail{
		Position = PositionHead
	}

	// check shell
	if _,err := os.Stat("/bin/bash"); os.IsExist(err){
		fmt.Println("No exist /bin/bash,build failed")
		os.Exit(1)
	}
}

func lookupStrEnv(key string,defaultVal string )  string {
	if s,ok := os.LookupEnv(key);ok && len(s)> 0 {
		return s
	}
	return defaultVal
}
func lookupBoolEnv(key string,defaultVal bool )  bool {
	if s,ok := os.LookupEnv(key);ok && len(s)> 0 {
		if strings.ToLower(s) == "true"{
			return true
		}
		return false
	}
	return defaultVal
}
