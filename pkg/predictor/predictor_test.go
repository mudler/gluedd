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

package predictor_test

import (
	"github.com/mudler/gluedd/pkg/api"
	"github.com/mudler/gluedd/pkg/errand"
	. "github.com/mudler/gluedd/pkg/predictor"
	"github.com/mudler/gluedd/pkg/resource"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
)

var ok string

type DummyErrand struct {
	Prediction api.Prediction
}

func (e *DummyErrand) Apply() error {
	ok = e.Prediction.Error.Error()
	return nil
}

func (e *DummyErrand) Generate(api.Detector) *api.Prediction {
	return nil
}

type DummyGenerator struct{}
type DeepDetect struct{ s string }

func NewDummyGenerator() errand.ErrandGenerator {
	return &DummyGenerator{}
}
func (l *DummyGenerator) GenerateErrand(p api.Prediction) errand.Errand {
	return &DummyErrand{Prediction: p}
}

func NewDummy() resource.Resource {
	return &DummyResource{}
}

// fs listener
type DummyResource struct {
}

func (l *DummyResource) Listen() chan string {

	files := make(chan string, 1)

	files <- ""
	return files
}

func NewDeepDetect(server string) api.Detector {
	return &DeepDetect{s: server}
}

func (d *DeepDetect) Detect(photo string) api.Prediction {
	return api.Prediction{Error: errors.New(d.s)}
}

func (d *DeepDetect) DetectService(photo, s string) api.Prediction {
	return api.Prediction{Error: errors.New(d.s)}
}

func (d *DeepDetect) WithService(s string) api.Detector {
	return d
}

type DefaultErrandConsumer struct{}

func NewErrandConsumer() errand.ErrandConsumer {
	return &DefaultErrandConsumer{}
}

func (p *DefaultErrandConsumer) Consume(e chan errand.Errand) {
	j := <-e
	j.Apply()
}

var _ = Describe("Predictor", func() {
	Context("Generate errands from prediction with a errandgenerator", func() {
		It("Generate consumable errand", func() {
			errandgen := NewDummyGenerator()
			predictor := NewPredictor(NewDeepDetect("fired"), NewDummy(), errandgen)
			consumer := NewErrandConsumer()
			consumer.Consume(predictor.Generate())
			Expect(ok).To(Equal("fired"))
		})
	})
})
