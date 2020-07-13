GC=go
MAIN_GO_FILE=callback_tool.go
RELEASE_DIR=release
BIN_NAME=callback_tool

build-unix :
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GC} build -ldflags -w -o ${RELEASE_DIR}/${BIN_NAME}_unix ${MAIN_GO_FILE}

build-mac :
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 ${GC} build -ldflags -w -o ${RELEASE_DIR}/${BIN_NAME}_mac ${MAIN_GO_FILE}

build-win :
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 ${GC} build -ldflags -w -o ${RELEASE_DIR}/${BIN_NAME}_win ${MAIN_GO_FILE}

build :
	make build-unix
	make build-mac
	make build-win

release :
	make build
	upx ${RELEASE_DIR}/${BIN_NAME}_*
	tar -czvf ${RELEASE_DIR}/${BIN_NAME}_mac.tar.gz -C ${RELEASE_DIR}/ ${BIN_NAME}_mac
	tar -czvf ${RELEASE_DIR}/${BIN_NAME}_unix.tar.gz -C ${RELEASE_DIR}/ ${BIN_NAME}_unix
	tar -czvf ${RELEASE_DIR}/${BIN_NAME}_win.tar.gz -C ${RELEASE_DIR}/ ${BIN_NAME}_win

clean :
	@if [ -d ${RELEASE_DIR} ] ; then rm -rf ${RELEASE_DIR}; fi