@echo off

set NETFrameworkVersion=4.6.2
set VisualStudioVersion=15.0
set WindowsSDKVersion=10.0.17134.0\
set FrameworkVersion=v4.0.30319
set VCToolsVersion=14.15.26726
set FrameworkDir=%SystemRoot%\Microsoft.NET\Framework64\
set WindowsSdkDir=%ProgramFiles(x86)%\Windows Kits\10\
set WindowsSdkBinPath=%WindowsSdkDir%bin\
set WindowsSdkVerBinPath=%WindowsSdkBinPath%%WindowsSDKVersion%
set VSINSTALLDIR=%ProgramFiles(x86)%\Microsoft Visual Studio\2017\Community\
set VCINSTALLDIR=%VSINSTALLDIR%VC\
set VCToolsInstallDir=%VCINSTALLDIR%Tools\MSVC\%VCToolsVersion%\
set INCLUDE=%VCToolsInstallDir%include;%INCLUDE%
set LIB=%VCToolsInstallDir%lib\x64;%ProgramFiles(x86)%\Windows Kits\NETFXSDK\%NETFrameworkVersion%\lib\um\x64;%WindowsSdkDir%lib\%WindowsSDKVersion%ucrt\x64;%WindowsSdkDir%lib\%WindowsSDKVersion%um\x64;%LIB%
set LIBPATH=%FrameworkDir%%FrameworkVersion%;%LIBPATH%
set Path=%VCToolsInstallDir%bin\HostX64\x64;%VSINSTALLDIR%MSBuild\%VisualStudioVersion%\bin\Roslyn;%WindowsSdkVerBinPath%x64;%Path%
set WEBVIEW_DIR=%USERPROFILE%\.nuget\packages\microsoft.toolkit.forms.ui.controls.webview\5.0.0\lib\net462

rmdir /S /Q Release
mkdir Release

csc -noconfig -nowarn:1701,1702 -nostdlib+ -errorreport:prompt -warn:4 -define:TRACE -highentropyva+ -reference:%WEBVIEW_DIR%\Microsoft.Toolkit.Forms.UI.Controls.WebView.dll -reference:mscorlib.dll -reference:System.Core.dll -reference:System.Data.dll -reference:System.dll -reference:System.Collections.dll -reference:System.Drawing.dll -reference:System.Windows.Forms.dll -reference:System.Xml.dll -debug:pdbonly -filealign:512 -optimize+ -out:Release\WebApp.dll -subsystemversion:6.00 -target:library -utf8output -deterministic+ platform_windows.cs %USERPROFILE%\AppData\Local\Temp\.NETFramework,Version=v%NETFrameworkVersion%.AssemblyAttributes.cs
copy %WEBVIEW_DIR%\Microsoft.Toolkit.Forms.UI.Controls.WebView.dll Release\Microsoft.Toolkit.Forms.UI.Controls.WebView.dll
copy %WEBVIEW_DIR%\Microsoft.Toolkit.Forms.UI.Controls.WebView.pdb Release\Microsoft.Toolkit.Forms.UI.Controls.WebView.pdb
copy %WEBVIEW_DIR%\Microsoft.Toolkit.Forms.UI.Controls.WebView.xml Release\Microsoft.Toolkit.Forms.UI.Controls.WebView.xml

CL -c -Zi -clr -nologo -W3 -WX- -diagnostics:classic -O2 -D NDEBUG -D _WINDLL -D _UNICODE -D UNICODE -EHa -MD -GS -fp:precise -Zc:wchar_t -Zc:forScope -Zc:inline -FoRelease\ -FdRelease\vc141.pdb -TP -FUmscorlib.dll -FUSystem.dll -FUSystem.Windows.Forms.dll -FURelease\WebApp.dll -FC -errorReport:queue -clr:nostdlib platform_windows.mcpp
CL -c -Zi -clr -nologo -W3 -WX- -diagnostics:classic -O2 -D NDEBUG -D _WINDLL -D _UNICODE -D UNICODE -EHa -MD -GS -fp:precise -Zc:wchar_t -Zc:forScope -Zc:inline -FpRelease\Interop.pch -FaRelease\ -FoRelease\ -FdRelease\vc141.pdb -FC -errorReport:queue %USERPROFILE%\AppData\Local\Temp\.NETFramework,Version=v%NETFrameworkVersion%.AssemblyAttributes.cpp
link -ERRORREPORT:QUEUE -OUT:Release\Interop.dll -INCREMENTAL:NO -NOLOGO -MANIFEST -MANIFESTUAC:"level='asInvoker' uiAccess='false'" -manifest:embed -DEBUG:FULL -PDB:Release\Interop.pdb -TLBID:1 -DYNAMICBASE -FIXED:NO -NXCOMPAT -MACHINE:X64 -DLL Release\platform_windows.obj

pause
