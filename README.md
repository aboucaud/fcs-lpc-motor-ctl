fcs-lpc-motor-ctl
=================

A `html5` application using the `github.com/go-lsst/ncs` control system to control the `m702` motors.

## Installation

```sh
sh> go get github.com/go-lsst/fcs-lpc-motor-ctl
sh> fcs-lpc-motor-ctl -addr=:7777 &
sh> open 0.0.0.0:7777
```

## API

```sh

$> curl -u faux-fcs:faux-fcs clrbinetsrv.in2p3.fr:5555/api/mon
{"code":200,"infos":[{"motor":"x","online":true,"status":"ready","mode":"home","rpms":1300,"angle":0,"temps":[35,40,46,52],"histos":null,"webcam":""},{"motor":"z","online":true,"status":"ready","mode":"home","rpms":1300,"angle":0,"temps":[35,40,45,51],"histos":null,"webcam":""}]}

$> curl -u faux-fcs:faux-fcs -X POST clrbinetsrv.in2p3.fr:5555/api/cmd/req-pos --data '{"motor":"x", "value":1}'
{"code":200}

$> curl -u faux-fcs:faux-fcs -X POST clrbinetsrv.in2p3.fr:5555/api/cmd/req-angle-pos --data '{"motor":"x", "value":45}'
{"code":200}

$> curl -u faux-fcs:faux-fcs -X POST clrbinetsrv.in2p3.fr:5555/api/cmd/req-pos --data '{"motor":"z", "value":1}'
{"code":200}

$> curl -u faux-fcs:faux-fcs clrbinetsrv.in2p3.fr:5555/api/mon
{"code":200,"infos":[{"motor":"x","online":true,"status":"ready","mode":"pos","rpms":1300,"angle":13,"temps":[35,41,47,50],"histos":null,"webcam":""},{"motor":"z","online":true,"status":"ready","mode":"pos","rpms":1300,"angle":0,"temps":[36,40,46,50],"histos":null,"webcam":""}]}

$> curl -u faux-fcs:faux-fcs -X POST clrbinetsrv.in2p3.fr:5555/api/cmd/req-upload-cmds --data '{"motor":"z", "cmds":"motor z\nget 0.18.002\nmotor x\nget 0.18.002"}'
{"code":200,"script":"Pr-00.18.002: hex=[0x00 0x00 0x00 0x00] dec=[  0   0   0   0] (0)\nPr-00.18.002: hex=[0x00 0x00 0x00 0xe1] dec=[  0   0   0 225] (225)\n"}

$> curl -u faux-fcs:faux-fcs -X POST clrbinetsrv.in2p3.fr:5555/api/cmd/req-upload-script -F 'upload-file=@./test.script'
{"code":200,"script":"Pr-00.08.015: hex=[0x00 0x00 0x00 0x00] dec=[  0   0   0   0] (0)\nPr-00.08.015: hex=[0x00 0x00 0x00 0x01] dec=[  0   0   0   1] (1)\nset-x-angle-pos=20\nset-x-rpm=2000\nget-x-angle-pos=20\nget-x-rpm=2000\nset-z-angle-pos=-20\nset-z-rpm=2020\nget-z-rpm=2020\nget-z-angle-pos=-20\n"}

### high-level API

## find-home
$> curl -u faux-fcs:faux-fcs -X POST clrbinetsrv.in2p3.fr:5555/api/cmd/req-upload-cmds --data '{"motor":"z", "cmds":"z-find-home"}'
{"code":200,"script":""}

$> curl -u faux-fcs:faux-fcs clrbinetsrv.in2p3.fr:5555/api/mon
{"code":200,"infos":[{"motor":"x","online":true,"status":"ready","mode":"pos","rpms":1300,"angle":45,"temps":[35,41,45,52],"histos":null,"webcam":""},{"motor":"z","online":true,"status":"ready","mode":"home","rpms":2000,"angle":0,"temps":[35,42,46,51],"histos":null,"webcam":""}]}

## set-pos-mode
$> curl -u faux-fcs:faux-fcs -X POST clrbinetsrv.in2p3.fr:5555/api/cmd/req-upload-cmds --data '{"motor":"z", "cmds":"z-pos"}'
{"code":200,"script":""}

$> curl -u faux-fcs:faux-fcs clrbinetsrv.in2p3.fr:5555/api/mon
{"code":200,"infos":[{"motor":"x","online":true,"status":"ready","mode":"pos","rpms":1300,"angle":45,"temps":[36,41,46,52],"histos":null,"webcam":""},{"motor":"z","online":true,"status":"ready","mode":"pos","rpms":2000,"angle":0,"temps":[35,41,46,50],"histos":null,"webcam":""}]}

## get/set rpm
$> curl -u faux-fcs:faux-fcs -X POST clrbinetsrv.in2p3.fr:5555/api/cmd/req-upload-cmds --data '{"motor":"z", "cmds":"z-rpm"}'
{"code":200,"script":"get-z-rpm=1300\n"}

$> curl -u faux-fcs:faux-fcs -X POST clrbinetsrv.in2p3.fr:5555/api/cmd/req-upload-cmds --data '{"motor":"z", "cmds":"z-rpm 2000"}'
{"code":200,"script":"set-z-rpm=2000\n"}

$> curl -u faux-fcs:faux-fcs -X POST clrbinetsrv.in2p3.fr:5555/api/cmd/req-upload-cmds --data '{"motor":"z", "cmds":"z-rpm"}'
{"code":200,"script":"get-z-rpm=2000\n"}

$> curl -u faux-fcs:faux-fcs -X POST clrbinetsrv.in2p3.fr:5555/api/cmd/req-upload-cmds --data '{"motor":"z", "cmds":"z-rpm\nx-rpm"}'
{"code":200,"script":"get-z-rpm=2000\nget-x-rpm=1300\n"}

## get/set angular position

$> curl -u faux-fcs:faux-fcs -X POST clrbinetsrv.in2p3.fr:5555/api/cmd/req-upload-cmds --data '{"motor":"z", "cmds":"z-angle-pos -20"}'
{"code":200,"script":"set-z-angle-pos=-20\n"}

$> curl -u faux-fcs:faux-fcs -X POST clrbinetsrv.in2p3.fr:5555/api/cmd/req-upload-cmds --data '{"motor":"z", "cmds":"z-angle-pos"}'
{"code":200,"script":"get-z-angle-pos=-20\n"}

## sleep for some time
$> curl -u faux-fcs:faux-fcs -X POST clrbinetsrv.in2p3.fr:5555/api/cmd/req-upload-cmds --data '{"motor":"z", "cmds":"z-angle-pos +5\n sleep 20s\n z-angle-pos"}'
{"code":200,"script":"set-z-angle-pos=5\nget-z-angle-pos=5\n"}

```
