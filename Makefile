#/bin/bash
# This is how we want to name the binary output
OUTPUT=savvy_life

# These are the values we want to pass for Version and BuildTime
GitTag=v1.0
BuildTime=`date +%Y%m%d%H%M%S`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=${GitTag} -X main.BuildTime=${BuildTime}"

debug:
	go install ${LDFLAGS} ${OUTPUT}
	mv ${GOPATH}/bin/savvy_life ./bin/

release:
	rm -f ./bin/savvy_life
	go install ${LDFLAGS} ${OUTPUT}
	mv ${GOPATH}/bin/savvy_life ./bin/
	cd .. && \
tar -zcvf savvy_life_release.tar.gz \
--exclude savvy_life/.git \
--exclude savvy_life/.idea \
--exclude savvy_life/build \
--exclude savvy_life/common \
--exclude savvy_life/extern \
--exclude savvy_life/middlewares \
--exclude savvy_life/proto \
--exclude savvy_life/server \
--exclude savvy_life/timer \
--exclude savvy_life/go.mod \
--exclude savvy_life/go.sum \
--exclude savvy_life/savvy_life.go \
--exclude savvy_life/Makefile \
savvy_life && \
mv -f savvy_life_release.tar.gz savvy_life/build/savvy_life_release_${GitTag}_${BuildTime}.tar.gz && \
rm -f savvy_life/bin/savvy_life

clean:
	rm -f ./bin/savvy_life
	rm -f ./build/savvy_life_*
