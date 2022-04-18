package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var ChasCH = time.Date(2022, time.April, 18, 12, 0, 0, 0, time.Local)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: likscan SOURCE TARGET\n")
		return
	}
	fmt.Printf("LIkScan started\n")
	src := strings.TrimSuffix(os.Args[1], "/")
	trg := strings.TrimSuffix(os.Args[2], "/")
	scanDirectory(src, trg)
	fmt.Printf("\nLIkScan done\n")
}

func scanDirectory(src string, trg string) {
	dirs, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}
	ms := len(dirs)
	dirt, err := ioutil.ReadDir(trg)
	if err != nil {
		return
	}
	mt := len(dirt)
	ns, nt := 0, 0
	for ns < ms && nt < mt {
		elms := dirs[ns]
		elmt := dirt[nt]
		names := elms.Name()
		namet := elmt.Name()
		if names < namet {
			ns++
		} else if names > namet {
			nt++
		} else if elms.IsDir() != elmt.IsDir() {
			ns++
			nt++
		} else {
			target := trg + "/" + namet
			if elmt.IsDir() {
				scanDirectory(src+"/"+names, target)
			}
			mods := elms.ModTime()
			modt := elmt.ModTime()
			if mods.Unix() != modt.Unix() {
				if modt.Before(ChasCH) {
					fmt.Printf("%s\n", target)
					os.Chtimes(target, elms.ModTime(), elms.ModTime())
				}
			}
			ns++
			nt++
		}
	}
}
