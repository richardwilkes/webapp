#!/usr/bin/env bash
set -eo pipefail

NETFrameworkVersion=4.6.2
VisualStudioVersion=15.0
WindowsSDKVersion=10.0.17134.0/
FrameworkVersion=v4.0.30319
VCToolsVersion=14.15.26726
FrameworkDir=$SYSTEMROOT/Microsoft.NET/Framework64/
ProgramFilesx86="$PROGRAMFILES (x86)"
WindowsSdkDir="$ProgramFilesx86/Windows Kits/10/"
WindowsSdkBinPath=${WindowsSdkDir}bin/
WindowsSdkVerBinPath=$WindowsSdkBinPath$WindowsSDKVersion
VSINSTALLDIR="$ProgramFilesx86/Microsoft Visual Studio/2017/Community/"
VCINSTALLDIR=${VSINSTALLDIR}VC/
VCToolsInstallDir=${VCINSTALLDIR}Tools/MSVC/$VCToolsVersion/
export INCLUDE="${VCToolsInstallDir}include;$INCLUDE"
export LIB="${VCToolsInstallDir}lib/x64;$ProgramFilesx86/Windows Kits/NETFXSDK/$NETFrameworkVersion/lib/um/x64;${WindowsSdkDir}lib/${WindowsSDKVersion}ucrt/x64;${WindowsSdkDir}lib/${WindowsSDKVersion}um/x64;$LIB"
export LIBPATH="$FrameworkDir$FrameworkVersion;$LIBPATH"
export PATH="${VCToolsInstallDir}bin/HostX64/x64:${VSINSTALLDIR}MSBuild/$VisualStudioVersion/bin/Roslyn:${WindowsSdkVerBinPath}x64:$PATH"

WEBVIEW=Microsoft.Toolkit.Forms.UI.Controls.WebView
WEBVIEW_VERSION=5.0.0
WEBVIEW_DIR=WebView-$WEBVIEW_VERSION
if [ ! -f $WEBVIEW_DIR/$WEBVIEW.dll ]; then
    mkdir -p $WEBVIEW_DIR
    pushd $WEBVIEW_DIR > /dev/null
    if [ ! -f $WEBVIEW.$WEBVIEW_VERSION.nupkg ]; then
	curl -LO https://github.com/windows-toolkit/WindowsCommunityToolkit/releases/download/v$WEBVIEW_VERSION/$WEBVIEW.$WEBVIEW_VERSION.nupkg
    fi
    unzip -j $WEBVIEW.$WEBVIEW_VERSION.nupkg lib/net462/$WEBVIEW.dll
    popd > /dev/null
fi

rm -rf Release
mkdir -p Release

csc \
	-noconfig \
	-nowarn:1701,1702 \
	-nostdlib+ \
	-errorreport:prompt \
	-warn:4 \
	-define:TRACE \
	-highentropyva+ \
	-reference:$WEBVIEW_DIR/$WEBVIEW.dll \
	-reference:mscorlib.dll \
	-reference:System.Core.dll \
	-reference:System.Data.dll \
	-reference:System.dll \
	-reference:System.Collections.dll \
	-reference:System.Drawing.dll \
	-reference:System.Windows.Forms.dll \
	-reference:System.Xml.dll \
	-debug:pdbonly \
	-filealign:512 \
	-optimize+ \
	-out:Release/WebApp.dll \
	-subsystemversion:6.00 \
	-target:library \
	-utf8output \
	-deterministic+ \
	platform_windows.cs

cl \
	-c \
	-Zi \
	-clr \
	-nologo \
	-W3 \
	-WX- \
	-diagnostics:classic \
	-O2 \
	-DNDEBUG \
	-D_WINDLL \
	-D_UNICODE \
	-DUNICODE \
	-EHa \
	-MD \
	-GS \
	-fp:precise \
	-Zc:wchar_t \
	-Zc:forScope \
	-Zc:inline \
	-FoRelease/ \
	-FdRelease/vc141.pdb \
	-TP \
	-FUmscorlib.dll \
	-FUSystem.dll \
	-FUSystem.Windows.Forms.dll \
	-FURelease/WebApp.dll \
	-FC \
	-errorReport:queue \
	-clr:nostdlib \
	platform_windows.mcpp

link \
	-ERRORREPORT:QUEUE \
	-OUT:Release/Interop.dll \
	-INCREMENTAL:NO \
	-NOLOGO \
	-MANIFEST "-MANIFESTUAC:level='asInvoker' uiAccess='false'" \
	-manifest:embed \
	-DEBUG:FULL \
	-PDB:Release/Interop.pdb \
	-TLBID:1 \
	-DYNAMICBASE \
	-FIXED:NO \
	-NXCOMPAT \
	-MACHINE:X64 \
	-DLL \
	Release/platform_windows.obj

cp $WEBVIEW_DIR/$WEBVIEW.dll Release
