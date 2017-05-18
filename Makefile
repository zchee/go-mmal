build: scp
	@ssh pi@192.168.1.21 PKG_CONFIG_PATH=/opt/vc/lib/pkgconfig /usr/local/go/bin/go build -x github.com/zchee/go-mmal

scp:
	@ssh pi@192.168.1.21 rm -rf /home/pi/go/src/github.com/zchee/go-mmal/*.go
	@scp -q $(shell find . -maxdepth 1 -name '*.go' -and -not -name '_*') pi@$(PI_SSH_IP):/home/pi/go/src/github.com/zchee/go-mmal
