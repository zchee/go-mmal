install: scp
	@ssh pi@raspberrypi.local PKG_CONFIG_PATH=/opt/vc/lib/pkgconfig /usr/local/go/bin/go install -x github.com/zchee/go-mmal
	${MAKE} copy

scp:
	@ssh pi@raspberrypi.local rm -rf /home/pi/go/src/github.com/zchee/go-mmal/*.go
	@scp -q $(shell find . -maxdepth 1 -name '*.go' -and -not -name '_*') pi@raspberrypi.local:/home/pi/go/src/github.com/zchee/go-mmal

copy:
	@scp pi@raspberrypi.local:/home/pi/go/pkg/linux_arm/github.com/zchee/go-mmal.a $(shell go env GOPATH)/pkg/darwin_amd64/github.com/zchee
