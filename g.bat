@ ECHO OFF
set STARTTIME=%TIME%
echo          ---------------
echo Start    : %STARTTIME% %0
echo          ---------------

: -------------------------------------------------

SET WORK_PATH=%CD%
SET cpss_tool=%WORK_PATH%\tools
SET TOOL_COMMON=%cpss_tool%\common
set SCRIPT_DIR=%WORK_PATH%\tools\admin
SET WORK_DRIVE=%WORK_PATH:~0,2%

echo go run main_1.go
     go run main_1.go

:: %WORK_DRIVE%
:: cd %WORK_PATH%

CALL %TOOL_COMMON%\duration.bat %STARTTIME% %0
