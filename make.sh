#!/usr/bin/env bash
#---------------------
# Golang manager for make
# https://github.com/verystar/gomake
#---------------------

APP="app"
SHELL_PATH=$(cd `dirname $0`; pwd)
VERSION="1.0.1"
APP_VERSION=`date +%Y%m%d%H%M%S`

fail() {
	echo -e "\033[31m[Error]$1\033[0m"
	exit 1
}

info() {
	echo -e "\033[34m$1\033[0m"
}

info "[App]"${APP}
info "[Path]"${SHELL_PATH}

if [ "${SHELL_PATH}" = "" ]; then
    fail "SHELL_PATH empty!"
fi

gop(){
    currpath=${SHELL_PATH}/
    gopath=${currpath%src/*}

    if [[ ${gopath} = "" ]];then
        echo "path not found src"
        exit
    else
        export GOPATH=${currpath%src/*}
    fi
}

build(){
    if [ ! -d ${SHELL_PATH}/bin ]; then
        mkdir ${SHELL_PATH}/bin/
    fi

    if [ "${GOPATH}" = "" ]; then
        fail "GOPATH empty!"
    fi
    info "[GOPATH]"${GOPATH}
    go build -ldflags "-X app/server.VERSION=${APP_VERSION}" -o ${SHELL_PATH}/bin/${APP}-${APP_VERSION} main.go

    changeVersion ${APP_VERSION}
}

changeVersion(){
    if [ "$1" != "" ]; then
        if [ -f ${SHELL_PATH}/bin/${APP}-${APP_VERSION} ]; then
            rm -f ${SHELL_PATH}/${APP}
            ln -s ${SHELL_PATH}/bin/${APP}-${APP_VERSION} ${SHELL_PATH}/${APP}
            ls -t ${SHELL_PATH}/bin/ | awk '{if(NR>11){print "./bin/"$0}}' | xargs rm -f
        else
        fail "Not found bin file:./bin/${APP}-${APP_VERSION}"
        fi
    else
       fail "Must params version!"
    fi
}

back(){
    currversion=$(ls -ld ${APP}|awk '{print $NF}')
    prefile=""
    echo $1
    if [ "$1" != "" ]; then
        APP_VERSION=$1
    else
        for file in `ls ${SHELL_PATH}/bin/${APP}-* | sort`
        do
            if [ "${file}" == "${currversion}" ]; then
                if [ "${prefile}" == "" ]; then
                    fail "Not prev version"
                else
                    APP_VERSION=${prefile#*${APP}-}
                    break
                fi
            fi
            prefile=${file}
        done
    fi

    changeVersion ${APP_VERSION}
}


#upgrade
initDownloadTool() {
	if type "curl" > /dev/null; then
		DOWNLOAD_TOOL="curl"
	elif type "wget" > /dev/null; then
		DOWNLOAD_TOOL="wget"
	else
		fail "You need curl or wget as download tool. Please install it first before continue"
	fi
	echo "Using $DOWNLOAD_TOOL as download tool, install gomake"
}

getFile() {
	local url="$1"
	local filePath="$2"
	if [ "$DOWNLOAD_TOOL" = "curl" ]; then
		curl -L "$url" -o "$filePath"
	elif [ "$DOWNLOAD_TOOL" = "wget" ]; then
		wget -O "$filePath" "$url"
	fi
}

install() {
    DOWNLOAD_MAKESHELL_URL=https://raw.githubusercontent.com/verystar/gomake/master/make.sh
    DOWNLOAD_MAKEFILE_URL=https://raw.githubusercontent.com/verystar/gomake/master/Makefile
    TOPATH=./
    getFile "${DOWNLOAD_MAKESHELL_URL}" "${TOPATH}"make.sh
    getFile "${DOWNLOAD_MAKEFILE_URL}" "${TOPATH}"Makefile
}

upgradeGoMake() {
    initDownloadTool
    install
}

#set gopath
gop

if [ "$1" = "back" ];then
    back $2
    info "[Current Version]"$(ls -ld ${APP}|awk '{print $NF}')
elif [ "$1" = "version" ];then
    info "[Current Version]"$(ls -ld ${APP}|awk '{print $NF}')
elif [ "$1" = "list" ]; then
    ls -l ${SHELL_PATH}/bin/
elif [ "$1" = "publish" ]; then
    cd ../../
    /bin/sh rsync.sh
elif [ "$1" = "test" ]; then
     goconvey -workDir=${SHELL_PATH} -excludedDirs="vendor,storage,examples,bin"
elif [ "$1" = "upgrade" ]; then
    upgradeGoMake
else
    build
fi