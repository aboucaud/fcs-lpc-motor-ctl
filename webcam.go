// Copyright ©2016 The go-lsst Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func (srv *server) handleWebcam(w http.ResponseWriter, r *http.Request) {
	img, err := http.Get("http://195.221.117.245:80/axis-cgi/jpg/image.cgi")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer img.Body.Close()
	buf, err := ioutil.ReadAll(img.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buf)))
	_, err = w.Write(buf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (srv *server) fetchWebcamImage() string {
	url := "http://" + srv.Addr + "/webcam"
	//url := "http://195.221.117.245/axis-cgi/jpg/image.cgi"
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error fetching webcam image: %v\n", err)
		return ""
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error reading webcam image: %v\n", err)
		return ""
	}

	return base64.StdEncoding.EncodeToString(buf)
}