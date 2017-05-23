install: scp
	@ssh pi@raspberrypi.local PKG_CONFIG_PATH=/opt/vc/lib/pkgconfig /usr/local/go/bin/go install -x github.com/zchee/go-mmal
	${MAKE} copy

scp:
	@ssh pi@raspberrypi.local mkdir -p /home/pi/go/src/github.com/zchee/go-mmal
	@ssh pi@raspberrypi.local rm -rf /home/pi/go/src/github.com/zchee/go-mmal/*.go
	@scp -q $(shell find . -maxdepth 1 -name '*.go' -and -not -name '_*') pi@raspberrypi.local:/home/pi/go/src/github.com/zchee/go-mmal

copy:
	@rm -rf $(shell go env GOPATH)/pkg/darwin_amd64/github.com/zchee/go-mmal*
	@scp -r pi@raspberrypi.local:/home/pi/go/pkg/linux_arm/github.com/zchee/go-mmal $(shell go env GOPATH)/pkg/darwin_amd64/github.com/zchee

debug:
	@go install ./internal/cmd/mmalgen
	@mmalgen ~/src/github.com/raspberrypi/userland/interface/mmal/util/mmal_util_params.h
