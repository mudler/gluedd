// Copyright © 2019 Ettore Di Giacinto <mudler@gentoo.org>
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 2 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program; if not, see <http://www.gnu.org/licenses/>.

package api

import (
	"fmt"
	"strconv"

	"github.com/jolibrain/godd"
)

type Prediction struct {
	godd.PredictResult
	Error error
}

func (p Prediction) Explain() {
	fmt.Println("")

	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	fmt.Println("Predictions:", len(p.Body.Predictions))
	for _, i := range p.Body.Predictions {
		fmt.Println("URI", i.URI)
		for _, c := range i.Classes {
			fmt.Println("Category:", c.Cat+" [prob "+strconv.FormatFloat(c.Prob, 'f', 6, 64)+"]")
		}
	}
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
}
