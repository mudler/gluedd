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
	"github.com/pkg/errors"

	"github.com/mudler/gluedd/pkg/api"
)

var _ = Describe("Generator", func() {
	Context("DefaultErrandGenerator", func() {
		It("Creates a DefaultErrand and wraps the prediction", func() {
			e := NewDefaultErrandGenerator()
			errand := e.GenerateErrand(api.Prediction{Error: errors.New("test")})
			realerrand, ok := errand.(*DefaultErrand)
			Expect(ok).To(BeTrue())
			Expect(realerrand.Prediction.Error.Error()).To(Equal("test"))
		})
	})
})
