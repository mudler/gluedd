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
	"github.com/jolibrain/godd"
)

type DeepDetect struct {
	Server  string
	service string
}

func NewDeepDetect(server string) Detector {
	return &DeepDetect{Server: server}
}

func (d *DeepDetect) Detect(photo string) Prediction {

	// Create predict structure for request parameters
	var predict godd.PredictRequest

	if len(d.service) > 0 {
		predict.Service = d.service
	} else {
		predict.Service = "detection_600"
	}

	predict.Data = append(predict.Data, photo)
	predict.Parameters.Output.Bbox = true
	predict.Parameters.Output.ConfidenceThreshold = 0.1

	predictResult, err := godd.Predict(d.Server, &predict)
	if err != nil {
		return Prediction{Error: err}
	}

	return Prediction{PredictResult: predictResult}
}

func (d *DeepDetect) WithService(s string) Detector {
	d.service = s
	return d
}
