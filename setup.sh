#!/usr/bin/env bash
set -eo pipefail
set -x

CEF_VERSION=3.3538.1849.g458cc98

# Setup PLATFORM
case $(uname -s) in
    Darwin*)  PLATFORM=macosx64 ;;
    Linux*)   PLATFORM=linux64 ;;
    MINGW64*) PLATFORM=windows64 ;;
    *)        echo "Unsupported OS: $(uname -s)"; false ;;
esac

if [ -e cef/include/cef_version.h ]; then
    EXISTING=`grep "#define CEF_VERSION " cef/include/cef_version.h | cut -f 2 -d '"'`
fi

if [ $CEF_VERSION != "$EXISTING" ]; then
    /bin/rm -rf cef
    /bin/rm -f cef_binary_${CEF_VERSION}_${PLATFORM}_minimal.tar*
    curl -LO http://opensource.spotify.com/cefbuilds/cef_binary_${CEF_VERSION}_${PLATFORM}_minimal.tar.bz2
    # Have to break this into two steps because Windows fails to do them in one...
    bunzip2 cef_binary_${CEF_VERSION}_${PLATFORM}_minimal.tar.bz2
    tar xf cef_binary_${CEF_VERSION}_${PLATFORM}_minimal.tar
    mv cef_binary_${CEF_VERSION}_${PLATFORM}_minimal cef
    /bin/rm -f cef_binary_${CEF_VERSION}_${PLATFORM}_minimal.tar
fi
