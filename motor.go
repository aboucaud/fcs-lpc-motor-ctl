// Copyright ©2016 The go-lsst Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/go-lsst/ncs/drivers/m702"
)

type motor struct {
	name   string
	addr   string
	params motorParams
	histos motorHistos
	online bool // whether motors are online/connected
}

func newMotor(name, addr string) motor {
	return motor{
		name:   name,
		addr:   addr,
		params: newMotorParams(),
		histos: motorHistos{
			rows: make([]monData, 0, 128),
		},
	}
}

type motorParams struct {
	Manual     m702.Parameter
	CmdReady   m702.Parameter
	HWSafety   m702.Parameter
	Home       m702.Parameter
	Random     m702.Parameter
	RPMs       m702.Parameter
	WriteAngle m702.Parameter
	ReadAngle  m702.Parameter
	Temps      [4]m702.Parameter
}

func newMotorParams() motorParams {
	return motorParams{
		Manual:     newParameter(paramManualOverride),
		CmdReady:   newParameter(paramCmdReady),
		HWSafety:   newParameter(paramHWSafety),
		Home:       newParameter(paramHome),
		Random:     newParameter(paramRandom),
		RPMs:       newParameter(paramRPMs),
		WriteAngle: newParameter(paramWritePos),
		ReadAngle:  newParameter(paramReadPos),
		Temps: [4]m702.Parameter{
			newParameter(paramTemp0),
			newParameter(paramTemp1),
			newParameter(paramTemp2),
			newParameter(paramTemp3),
		},
	}
}

type motorHistos struct {
	rows []monData
	Temp [4][]float64
	Pos  []float64
	RPMs []float64
}

type motorInfos struct {
	Motor  string            `json:"motor"` // x,z
	Online bool              `json:"online"`
	Status string            `json:"status"` // N/A,manual,hw-safety,ready
	Mode   string            `json:"mode"`   // N/A,ready,home,random
	RPMs   int               `json:"rpms"`
	Angle  int               `json:"angle"`
	Temps  [4]float64        `json:"temps"`
	Histos map[string]string `json:"histos"`
	Webcam string            `json:"webcam"`
}
