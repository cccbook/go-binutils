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

package common

import (
	"debug/elf"
)

// util.go: Common utilities and the Util interface

func Init(fileName string) (*elf.File, error) {
	return elf.Open(fileName)
}

func Output(content string) {
	return
}

type Util interface {
	DefineFlags() map[string]interface{}
	Run(args map[string]interface{}) (string, error)
}
