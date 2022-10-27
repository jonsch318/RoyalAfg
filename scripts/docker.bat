@echo off

if exist ".\services\%1\Dockerfile" (
    echo "Dockerfile found at .\services\%1\Dockerfile"
    COPY .\services\%1\Dockerfile .\Dockerfile
) else (
    echo "No Dockerfile found at .\services\%1\Dockerfile"

    COPY .\universal_Dockerfile .\Dockerfile
)

echo "Starting docker build of %1"


set RESTVAR=
shift
:loop1
if "%1"=="" goto after_loop
set RESTVAR=%RESTVAR% %1
shift
goto loop1

:after_loop



echo "Calling docker build%RESTVAR% ."
docker build %RESTVAR% .

echo "Finished docker build removing temporary Dockerfile"

echo "Finished"