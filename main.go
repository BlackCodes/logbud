package main

import (
	"fmt"
	"log"

	_ "github.com/BlackCodes/logbud/flag"

	"github.com/BlackCodes/logbud/command"
	"github.com/BlackCodes/logbud/files"
	"github.com/BlackCodes/logbud/flag"
	"github.com/BlackCodes/logbud/parse"
)

func main() {
	f := files.NewCloneFile()
	defer f.Destroy()

	if err := f.Start(); err != nil {
		log.Fatalf("clone file %v", err)
	}
	newPaths := f.GetNewPaths()
	p, err := parse.NewFileParse(f.GetDir())
	if err != nil {
		log.Fatalf("file parse %v", err)
	}
	if err := p.Start(newPaths); err != nil {
		log.Fatalf("file parse output %v", err)
	}
	b := command.NewBuild(f.GetBuildDir(), f.GetDir())

	if err := b.Build(); err != nil {
		log.Fatalf("build go err:%v", err)
		return
	}
	fmt.Println("build success...")
	if flag.Compress {
		fmt.Println("compress...")
		file, err := b.GetBinary()
		if err != nil {
			log.Fatalf("find build file err:%v", err)
			return
		}
		if err := command.NewUpx(file).Build(); err != nil {
			log.Fatalf("upx error:%v", err)
		}

	}
	if err := b.Copy(); err != nil {
		log.Fatalf("copy file error:%v", err)
	}
	fmt.Println("That's Great!!!")
}
