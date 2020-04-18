GOCMD=go
GOBUILD=$(GOCMD) build
GOENV=$(GOCMD) env
GREP=grep
SED=sed
FIND=find
OUTDIR=out
VERSION=1.0.0
LDFLAGS=-ldflags "-X main.version=${VERSION}"
NAMEFILE=go-urlhash

build:
	rm -fr ${OUTDIR}
	mkdir -p ${OUTDIR}
	GOOS=linux GOARCH=arm GOARM=6 ${GOBUILD} ${LDFLAGS} -o ${OUTDIR}/${NAMEFILE}.rpi
	GOOS=linux ${GOBUILD}  ${LDFLAGS} -o ${OUTDIR}/${NAMEFILE}.lin
	GOOS=darwin ${GOBUILD} ${LDFLAGS} -o ${OUTDIR}/${NAMEFILE}.mac
	zip ${OUTDIR}/${NAMEFILE}.rpi.zip ${OUTDIR}/${NAMEFILE}.rpi
	zip ${OUTDIR}/${NAMEFILE}.lin.zip ${OUTDIR}/${NAMEFILE}.lin
	zip ${OUTDIR}/${NAMEFILE}.mac.zip ${OUTDIR}/${NAMEFILE}.mac
