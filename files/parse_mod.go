package files

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// ParseMod parse mod file obtain project name
func  ParseMod(modFile string ) (string,error) {
	if len(modFile) == 0 {
		return "",fmt.Errorf("not fond modFile file")
	}
	fh,err := os.OpenFile(modFile,os.O_RDONLY,0755)
	if err != nil {
		return "",err
	}
	defer fh.Close()
	if info,err:= fh.Stat(); info.Size() == 0 {
		return "",err
	}
	buf := bufio.NewReader(fh)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return "",err
			}
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line,"module"){
			reg,_ := regexp.Compile(`^module\s+(\w+)`)
			matchStr := reg.FindStringSubmatch(line)
			if len(matchStr)> 1{
				arr := strings.Split(matchStr[1],"/")
				if len(arr)>0{
					return arr[len(arr)-1],nil
				}
			}
		}
	}
	return "",fmt.Errorf("parse modFile file failed")
}