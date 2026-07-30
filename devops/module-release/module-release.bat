@echo off
setlocal EnableExtensions DisableDelayedExpansion

set "SCRIPT_DIR=%~dp0"
for %%I in ("%SCRIPT_DIR%..\..") do set "REPO_ROOT=%%~fI"
set "MANIFEST=%SCRIPT_DIR%modules.tsv"
if not defined REMOTE set "REMOTE=origin"

goto :main

:usage
echo Usage:
echo   module-release.bat list
echo   module-release.bat check ^<module^> ^<version^>
echo   module-release.bat tag ^<module^> ^<version^>
echo   module-release.bat push ^<module^> ^<version^>
echo.
echo Examples:
echo   module-release.bat check kit v3.0.0
echo   module-release.bat tag kit v3.0.0
echo   module-release.bat push kit v3.0.0
exit /b 0

:die
>&2 echo error: %~1
exit /b 1

:load_module
set "REQUESTED_MODULE=%~1"
set "MODULE_NAME="
set "MODULE_DIR="
set "MODULE_PATH="
set "TAG_PREFIX="

if not exist "%MANIFEST%" (
  call :die "missing %MANIFEST%"
  exit /b 1
)

for /f "usebackq tokens=1-4 eol=#" %%A in ("%MANIFEST%") do (
  if "%%A"=="%REQUESTED_MODULE%" (
    set "MODULE_NAME=%%A"
    set "MODULE_DIR=%%B"
    set "MODULE_PATH=%%C"
    set "TAG_PREFIX=%%D"
    goto :load_module_found
  )
)

call :die "unknown module '%REQUESTED_MODULE%'; run '%~nx0 list'"
exit /b 1

:load_module_found
exit /b 0

:validate_version
set "VERSION=%~1"
set "MODULE_RELEASE_VERSION=%VERSION%"
%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe -NoLogo -NoProfile -NonInteractive -Command "$value = $env:MODULE_RELEASE_VERSION; if ($value -match '^v3\.[0-9]+\.[0-9]+([+-][0-9A-Za-z.-]+)?$') { exit 0 }; exit 1"
if errorlevel 1 (
  call :die "version must be a v3 semantic version, got '%VERSION%'"
  exit /b 1
)
exit /b 0

:set_release_tag
if "%TAG_PREFIX%"=="-" (
  set "RELEASE_TAG=%VERSION%"
) else (
  set "RELEASE_TAG=%TAG_PREFIX%/%VERSION%"
)
exit /b 0

:prepare_release
call :load_module "%~1"
if errorlevel 1 exit /b 1
call :validate_version "%~2"
if errorlevel 1 exit /b 1
call :set_release_tag
exit /b 0

:check_release
if not exist "%REPO_ROOT%\%MODULE_DIR%\go.mod" (
  call :die "missing %MODULE_DIR%/go.mod"
  exit /b 1
)

set "DECLARED_MODULE="
for /f "usebackq tokens=1,*" %%A in ("%REPO_ROOT%\%MODULE_DIR%\go.mod") do (
  if "%%A"=="module" if not defined DECLARED_MODULE set "DECLARED_MODULE=%%B"
)
if not "%DECLARED_MODULE%"=="%MODULE_PATH%" (
  call :die "%MODULE_DIR%/go.mod declares '%DECLARED_MODULE%', expected '%MODULE_PATH%'"
  exit /b 1
)

git -C "%REPO_ROOT%" status --porcelain >nul 2>&1
if errorlevel 1 (
  call :die "cannot inspect the Git working tree"
  exit /b 1
)
set "DIRTY_WORKTREE="
for /f "delims=" %%A in ('git -C "%REPO_ROOT%" status --porcelain') do set "DIRTY_WORKTREE=1"
if defined DIRTY_WORKTREE (
  call :die "working tree is not clean; commit and push the release changes first"
  exit /b 1
)

git -C "%REPO_ROOT%" rev-parse --verify --quiet "refs/tags/%RELEASE_TAG%" >nul 2>&1
if not errorlevel 1 (
  call :die "tag '%RELEASE_TAG%' already exists locally"
  exit /b 1
)

echo module:  %MODULE_NAME%
echo path:    %MODULE_PATH%
echo version: %VERSION%
echo tag:     %RELEASE_TAG%
set /p "=commit:  " <nul
git -C "%REPO_ROOT%" rev-parse HEAD
exit /b %ERRORLEVEL%

:print_module
set "PRINT_NAME=%~1"
set "PRINT_PATH=%~2"
set "PRINT_PREFIX=%~3"
set "LATEST_TAG="
if "%PRINT_PREFIX%"=="-" (
  set "TAG_PATTERN=v3.*"
) else (
  set "TAG_PATTERN=%PRINT_PREFIX%/v3.*"
)
for /f "delims=" %%T in ('git -C "%REPO_ROOT%" tag --list "%TAG_PATTERN%" --sort=-version:refname') do if not defined LATEST_TAG set "LATEST_TAG=%%T"
if not defined LATEST_TAG set "LATEST_TAG=-"
echo %PRINT_NAME%  %PRINT_PATH%  %LATEST_TAG%
exit /b 0

:list_modules
if not exist "%MANIFEST%" (
  call :die "missing %MANIFEST%"
  exit /b 1
)
git -C "%REPO_ROOT%" rev-parse --is-inside-work-tree >nul 2>&1
if errorlevel 1 (
  call :die "cannot access the Git repository at %REPO_ROOT%"
  exit /b 1
)
echo MODULE  MODULE PATH  LATEST TAG
for /f "usebackq tokens=1-4 eol=#" %%A in ("%MANIFEST%") do call :print_module "%%A" "%%C" "%%D"
exit /b 0

:main
set "COMMAND=%~1"
if /i "%COMMAND%"=="list" goto :command_list
if /i "%COMMAND%"=="check" goto :command_check
if /i "%COMMAND%"=="tag" goto :command_tag
if /i "%COMMAND%"=="push" goto :command_push
if /i "%COMMAND%"=="help" goto :command_help
if /i "%COMMAND%"=="-h" goto :command_help
if /i "%COMMAND%"=="--help" goto :command_help
call :usage
exit /b 1

:command_help
call :usage
exit /b 0

:command_list
if not "%~2"=="" (
  call :die "list does not accept arguments"
  exit /b 1
)
call :list_modules
exit /b %ERRORLEVEL%

:command_check
if "%~2"=="" goto :invalid_release_args
if "%~3"=="" goto :invalid_release_args
if not "%~4"=="" goto :invalid_release_args
call :prepare_release "%~2" "%~3"
if errorlevel 1 exit /b 1
call :check_release
exit /b %ERRORLEVEL%

:command_tag
if "%~2"=="" goto :invalid_release_args
if "%~3"=="" goto :invalid_release_args
if not "%~4"=="" goto :invalid_release_args
call :prepare_release "%~2" "%~3"
if errorlevel 1 exit /b 1
call :check_release
if errorlevel 1 exit /b 1
git -C "%REPO_ROOT%" tag -a "%RELEASE_TAG%" -m "release %MODULE_PATH% %VERSION%"
if errorlevel 1 exit /b 1
echo created tag %RELEASE_TAG%
exit /b 0

:command_push
if "%~2"=="" goto :invalid_release_args
if "%~3"=="" goto :invalid_release_args
if not "%~4"=="" goto :invalid_release_args
call :prepare_release "%~2" "%~3"
if errorlevel 1 exit /b 1
git -C "%REPO_ROOT%" rev-parse --verify --quiet "refs/tags/%RELEASE_TAG%" >nul 2>&1
if errorlevel 1 (
  call :die "tag '%RELEASE_TAG%' does not exist locally"
  exit /b 1
)
git -C "%REPO_ROOT%" push "%REMOTE%" "refs/tags/%RELEASE_TAG%"
exit /b %ERRORLEVEL%

:invalid_release_args
call :usage
exit /b 1
