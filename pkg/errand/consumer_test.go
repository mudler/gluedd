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

package errand_test

import (
	. "github.com/mudler/gluedd/pkg/errand"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

type DummyErrand struct {
	F string
}

func (e *DummyErrand) Apply() error {
	e.F = "fired"
	return nil
}

var _ = Describe("Consumer", func() {
	Context("Generate errands from prediction with a errandgenerator", func() {
		It("Generate consumable errand", func() {
			consumer := NewErrandConsumer()
			e := &DummyErrand{}
			c := make(chan Errand, 1)
			c <- e
			consumer.Consume(c)
			Eventually(func() string { return e.F }, 20*time.Second).Should(Equal("fired"))
		})
	})
})
