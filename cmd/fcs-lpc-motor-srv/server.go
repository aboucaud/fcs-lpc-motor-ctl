// Copyright ©2019 The go-lsst Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-lsst/fcs-lpc-motor-ctl/bench"
	"github.com/go-lsst/ncs/drivers/m702"
)

var (
	errInvalidHTTPMethod = errors.New("invalid HTTP method")
)

type server struct {
	motors []motor
}

func newServer() *server {
	return &server{
		motors: []motor{
			newMotor("x", "134.158.155.16:5021"),
			newMotor("z", "134.158.155.16:5023"),
		},
	}
}

func (srv *server) apiMonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		srv.apiError(w, errInvalidHTTPMethod, http.StatusMethodNotAllowed)
		return
	}

	var infos [2]bench.MotorInfos
	for i := range srv.motors {
		m := &srv.motors[i]
		info, err := m.infos(motorTimeout)
		if err != nil {
			srv.apiError(w, err, http.StatusServiceUnavailable)
			return
		}
		infos[i] = info
	}
	var resp = struct {
		Err   string              `json:"error,omitempty"`
		Code  int                 `json:"code"`
		Infos [2]bench.MotorInfos `json:"infos"`
	}{
		Err:   "",
		Code:  http.StatusOK,
		Infos: infos,
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(resp)
	if err != nil {
		srv.apiError(w, fmt.Errorf("could not encode monitoring infos to JSON: %v", err), http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(w, &buf)
	if err != nil {
		srv.apiError(w, fmt.Errorf("error writing JSON response: %v", err), http.StatusInternalServerError)
		return
	}
}

func (srv *server) apiCmdReqReadyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		srv.apiError(w, errInvalidHTTPMethod, http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var req cmdRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		srv.apiError(w, fmt.Errorf("error decoding JSON request: %v", err), http.StatusBadRequest)
		return
	}
	req.tstamp = time.Now().UTC()
	req.Type = "ctl"

	m, ok := srv.apiCheck(req, w, r)
	if !ok {
		return
	}

	p := newParameter(bench.ParamCmdReady)
	codec.PutUint32(p.Data[:], uint32(req.Value))
	err = m.Motor().WriteParam(p)
	if err != nil {
		srv.apiError(w, fmt.Errorf("error writing parameter %v to motor-%v: %v", p, m.name, err), http.StatusInternalServerError)
		return
	}
	srv.apiOK(w, http.StatusOK)
}

func (srv *server) apiCmdReqFindHomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		srv.apiError(w, errInvalidHTTPMethod, http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var req cmdRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		srv.apiError(w, fmt.Errorf("error decoding JSON request: %v", err), http.StatusBadRequest)
		return
	}
	req.tstamp = time.Now().UTC()
	req.Type = "ctl"

	m, ok := srv.apiCheck(req, w, r)
	if !ok {
		return
	}

	params := append([]m702.Parameter{},
		newParameter(bench.ParamCmdReady),
		newParameter(bench.ParamModePos),
		newParameter(bench.ParamHome),
		newParameter(bench.ParamCmdReady),
	)

	codec.PutUint32(params[0].Data[:], 0)
	codec.PutUint32(params[1].Data[:], 0)
	codec.PutUint32(params[2].Data[:], 1)
	codec.PutUint32(params[3].Data[:], 1)

	for _, p := range params {
		err = m.Motor().WriteParam(p)
		if err != nil {
			srv.apiError(w, fmt.Errorf("error writing parameter %v to motor-%v: %v", p, m.name, err), http.StatusInternalServerError)
			return
		}
	}

	srv.apiOK(w, http.StatusOK)
}

func (srv *server) apiCmdReqPosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		srv.apiError(w, errInvalidHTTPMethod, http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var req cmdRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		srv.apiError(w, fmt.Errorf("error decoding JSON request: %v", err), http.StatusBadRequest)
		return
	}
	req.tstamp = time.Now().UTC()
	req.Type = "ctl"

	m, ok := srv.apiCheck(req, w, r)
	if !ok {
		return
	}

	params := append([]m702.Parameter{},
		newParameter(bench.ParamCmdReady),
		newParameter(bench.ParamModePos),
		newParameter(bench.ParamHome),
		newParameter(bench.ParamCmdReady),
	)

	codec.PutUint32(params[0].Data[:], 0)
	codec.PutUint32(params[1].Data[:], 1)
	codec.PutUint32(params[2].Data[:], 0)
	codec.PutUint32(params[3].Data[:], 1)

	for _, p := range params {
		err = m.Motor().WriteParam(p)
		if err != nil {
			srv.apiError(w, fmt.Errorf("error writing parameter %v to motor-%v: %v", p, m.name, err), http.StatusInternalServerError)
			return
		}
	}

	srv.apiOK(w, http.StatusOK)
}

func (srv *server) apiCmdReqRPMHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		srv.apiError(w, errInvalidHTTPMethod, http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var req cmdRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		srv.apiError(w, fmt.Errorf("error decoding JSON request: %v", err), http.StatusBadRequest)
		return
	}
	req.tstamp = time.Now().UTC()
	req.Type = "ctl"

	if req.Value > 3000 {
		srv.apiError(w, fmt.Errorf("invalid RPM value (%v > 3000)", req.Value), http.StatusBadRequest)
		return
	}
	if req.Value < 0 {
		srv.apiError(w, fmt.Errorf("invalid RPM value (%v < 0)", req.Value), http.StatusBadRequest)
		return
	}

	m, ok := srv.apiCheck(req, w, r)
	if !ok {
		return
	}

	p := newParameter(bench.ParamRPMs)
	codec.PutUint32(p.Data[:], uint32(req.Value))
	err = m.Motor().WriteParam(p)
	if err != nil {
		srv.apiError(w, fmt.Errorf("error writing parameter %v to motor-%v: %v", p, m.name, err), http.StatusInternalServerError)
		return
	}

	srv.apiOK(w, http.StatusOK)
}

func (srv *server) apiCmdReqAnglePosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		srv.apiError(w, errInvalidHTTPMethod, http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var req cmdRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		srv.apiError(w, fmt.Errorf("error decoding JSON request: %v", err), http.StatusBadRequest)
		return
	}
	req.tstamp = time.Now().UTC()
	req.Type = "ctl"
	if req.Value > +90 {
		srv.apiError(w, fmt.Errorf("invalid angle position (%v > +90.0)", req.Value), http.StatusBadRequest)
		return
	}
	if req.Value < -90 {
		srv.apiError(w, fmt.Errorf("invalid angle position (%v < -90.0)", req.Value), http.StatusBadRequest)
		return
	}
	req.Value *= 10

	m, ok := srv.apiCheck(req, w, r)
	if !ok {
		return
	}

	p := newParameter(bench.ParamWritePos)
	codec.PutUint32(p.Data[:], uint32(req.Value))
	err = m.Motor().WriteParam(p)
	if err != nil {
		srv.apiError(w, fmt.Errorf("error writing parameter %v to motor-%v: %v", p, m.name, err), http.StatusInternalServerError)
		return
	}

	srv.apiOK(w, http.StatusOK)
}

func (srv *server) apiCmdReqUploadCmdsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		srv.apiError(w, errInvalidHTTPMethod, http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var req cmdRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		srv.apiError(w, fmt.Errorf("error decoding JSON request: %v", err), http.StatusBadRequest)
		return
	}
	req.tstamp = time.Now().UTC()
	req.Type = "ctl"

	m, ok := srv.apiCheck(req, w, r)
	if !ok {
		return
	}

	buf := new(bytes.Buffer)
	script := newScripter(srv, m.Motor())
	cmds := bytes.NewReader([]byte(req.Cmds))
	err = script.run(m.Motor(), cmds, buf)
	if err != nil {
		srv.apiError(w, fmt.Errorf("error running script: %v", err), http.StatusInternalServerError)
		return
	}

	var reply = struct {
		Err    string `json:"error,omitempty"`
		Code   int    `json:"code"`
		Script string `json:"script"`
	}{
		Err:    "",
		Code:   http.StatusOK,
		Script: string(buf.Bytes()),
	}

	var o = new(bytes.Buffer)
	err = json.NewEncoder(o).Encode(reply)
	if err != nil {
		srv.apiError(w, fmt.Errorf("error encoding JSON reply: %v", err), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(w, o)
	if err != nil {
		srv.apiError(w, fmt.Errorf("error sending JSON reply: %v", err), http.StatusInternalServerError)
		return
	}
}

func (srv *server) apiCmdReqUploadScriptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		srv.apiError(w, errInvalidHTTPMethod, http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	r.ParseMultipartForm(500 << 20)
	f, handler, err := r.FormFile("upload-file")
	if err != nil {
		srv.apiError(w, fmt.Errorf("error parsing form-file %q: %v", handler.Filename, err), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	cmd := new(bytes.Buffer)
	_, err = io.Copy(cmd, f)
	if err != nil {
		srv.apiError(w, fmt.Errorf("error reading form-file %q: %v", handler.Filename, err), http.StatusInternalServerError)
		return
	}

	req := cmdRequest{
		tstamp: time.Now().UTC(),
		Motor:  "x",
		Cmds:   string(cmd.Bytes()),
		Type:   "ctl",
	}

	m, ok := srv.apiCheck(req, w, r)
	if !ok {
		return
	}

	buf := new(bytes.Buffer)
	script := newScripter(srv, m.Motor())
	cmds := bytes.NewReader([]byte(req.Cmds))
	err = script.run(m.Motor(), cmds, buf)
	if err != nil {
		srv.apiError(w, fmt.Errorf("error running script: %v", err), http.StatusInternalServerError)
		return
	}

	var reply = struct {
		Err    string `json:"error,omitempty"`
		Code   int    `json:"code"`
		Script string `json:"script"`
	}{
		Err:    "",
		Code:   http.StatusOK,
		Script: string(buf.Bytes()),
	}

	var o = new(bytes.Buffer)
	err = json.NewEncoder(o).Encode(reply)
	if err != nil {
		srv.apiError(w, fmt.Errorf("error encoding JSON reply: %v", err), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(w, o)
	if err != nil {
		srv.apiError(w, fmt.Errorf("error sending JSON reply: %v", err), http.StatusInternalServerError)
		return
	}
}

func (srv *server) apiError(w http.ResponseWriter, err error, code int) {
	http.Error(w, fmt.Sprintf(`{"error":%q, "code":%v}`, err.Error(), code), code)
}

func (srv *server) apiOK(w http.ResponseWriter, code int) {
	http.Error(w, fmt.Sprintf(`{"code":%v}`, code), code)
}

func (srv *server) apiCheck(req cmdRequest, w http.ResponseWriter, r *http.Request) (*motor, bool) {
	m, err := srv.getMotor(req.Motor)
	if err != nil {
		srv.apiError(w, err, http.StatusBadRequest)
		return nil, false
	}

	if online, err := m.isOnline(motorTimeout); err != nil || !online {
		if err != nil {
			srv.apiError(w, err, http.StatusServiceUnavailable)
			return nil, false
		}
		srv.apiError(w, bench.ErrMotorOffline, http.StatusServiceUnavailable)
		return nil, false
	}

	if m.isHWLocked() {
		srv.apiError(w, bench.ErrMotorHWLock, http.StatusServiceUnavailable)
		return nil, false
	}

	if m.isManual() {
		srv.apiError(w, bench.ErrMotorManual, http.StatusServiceUnavailable)
		return nil, false
	}

	return m, true
}

func (srv *server) apiAuthenticated(h func(w http.ResponseWriter, r *http.Request), needACL bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// FIXME(sbinet): check user+passord
		w.Header().Set("Content-Type", "application/json")
		h(w, r)
		return
	}
}
