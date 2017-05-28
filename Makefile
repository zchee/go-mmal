install: scp
	@ssh pi@raspberrypi.local PKG_CONFIG_PATH=/opt/vc/lib/pkgconfig /usr/local/go/bin/go install -x github.com/zchee/go-mmal/cmd/...

scp:
	@ssh pi@raspberrypi.local rm -rf /home/pi/go/src/github.com/zchee/go-mmal/cmd
	@scp -rq $(CURDIR)/cmd pi@raspberrypi.local:/home/pi/go/src/github.com/zchee/go-mmal
	@ssh pi@raspberrypi.local /usr/local/go/bin/go get -d github.com/zchee/go-mmal/cmd/...

run/preview: install
	@ssh pi@raspberrypi.local GODEBUG=cgocheck=0 /home/pi/go/bin/preview

run/list-encoding: install
	@ssh pi@raspberrypi.local GODEBUG=cgocheck=0 /home/pi/go/bin/list-encoding vc.ril.camera

install/lib: scp/lib
	@ssh pi@raspberrypi.local PKG_CONFIG_PATH=/opt/vc/lib/pkgconfig /usr/local/go/bin/go install -x github.com/zchee/go-mmal github.com/zchee/go-mmal/bcmhost
	${MAKE} copy

scp/lib:
	@ssh pi@raspberrypi.local mkdir -p /home/pi/go/src/github.com/zchee/go-mmal
	@ssh pi@raspberrypi.local rm -rf /home/pi/go/src/github.com/zchee/go-mmal/*.go
	@scp -q $(shell find . -maxdepth 1 -name '*.go' -and -not -name '_*') pi@raspberrypi.local:/home/pi/go/src/github.com/zchee/go-mmal
	@scp -rq ./bcmhost pi@raspberrypi.local:/home/pi/go/src/github.com/zchee/go-mmal

copy:
	@rm -rf $(shell go env GOPATH)/pkg/darwin_amd64/github.com/zchee/go-mmal*
	@scp -r pi@raspberrypi.local:/home/pi/go/pkg/linux_arm/github.com/zchee/go-mmal.a $(shell go env GOPATH)/pkg/darwin_amd64/github.com/zchee

debug:
	@go install ./internal/cmd/mmalgen
	@mmalgen ~/src/github.com/raspberrypi/userland/interface/mmal/util/mmal_util_params.h

todo:
	@pt TODO --after=1 --ignore Makefile
