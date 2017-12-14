//
// Copyright (C) 2017  Alan (Quey-Liang) Kao  alankao@andestech.com
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package main

// main.go: The unified entry point

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/NonerKao/go-binutils/common"
	"github.com/NonerKao/go-binutils/readelf"
)

func main() {

	util, err := route()
	if err != nil {
		fmt.Println(err.Error())
		printUsage()
		return
	}

	args := util.DefineFlags()
	flag.Usage = printUsage
	flag.Parse()

	raw, err2 := util.Run(args)
	if err2 != nil {
		return
	}

	common.Output(raw)

}

func route() (common.Util, error) {

	switch {
	case strings.HasSuffix(os.Args[0], "readelf"):
		util, err := readelf.Init(os.Args[len(os.Args)-1])
		if err != nil {
			return nil, err
		}
		return util, nil
	default:
		return nil, errors.New("No such usage!")
	}
}

func printUsage() {
	fmt.Printf("Usage of go-binutils: %s\n", os.Args[0])
	fmt.Printf("    example7 file1 file2 ...\n")
	flag.PrintDefaults()
}
