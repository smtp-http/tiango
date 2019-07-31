@echo off

setlocal

echo "%0% tiango"

set CURRENTPATH=%~dp0

set TOOL_PATH=%CURRENTPATH%\nssm

set BIN_PATH=%CURRENTPATH%\bin

echo "TOOL_PATH = %TOOL_PATH%"

echo "BIN_PATH = %BIN_PATH%"

echo "stop a.tiango service"

%TOOL_PATH%\nssm.exe stop  a.tiango

echo "remove tiango service"

%TOOL_PATH%\nssm.exe remove a.tiango confirm

echo finished

pause