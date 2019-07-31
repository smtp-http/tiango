@echo off

setlocal

echo "%0% tiango"

set CLIENT_KEY=test-w3


echo "CLIENT_KEY=%CLIENT_KEY%"

set CURRENTPATH=%~dp0

set TOOL_PATH=%CURRENTPATH%\nssm

set BIN_PATH=%CURRENTPATH%

echo "TOOL_PATH = %TOOL_PATH%"

echo "BIN_PATH = %BIN_PATH%"

reg query "HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Cryptography" /v MachineGuid

echo "install a.tiango service"

%TOOL_PATH%\nssm.exe install a.tiango  %BIN_PATH%\main.exe " -serverPort=80"

echo "set a.tiango log file"

%TOOL_PATH%\nssm.exe set a.tiango AppStdout %BIN_PATH%\log\service.log
%TOOL_PATH%\nssm.exe set a.tiango AppStderr %BIN_PATH%\log\service.log

echo "start a.tiango service"

%TOOL_PATH%\nssm.exe start a.tiango

echo finished

pause