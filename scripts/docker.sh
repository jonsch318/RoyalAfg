#!/bin/sh


echo "Starting docker build of $1"

if test -f ./services/$1/Dockerfile; then 
    echo "Dockerfile found at ./services/$1/Dockerfile"
    cp ./services/$1/Dockerfile ./Dockerfile
else
    echo "No Dockerfile found at ./services/$1/Dockerfile"

    cp ./universal_Dockerfile ./Dockerfile
fi

echo "Calling docker build . ${@:2}"
docker build ${@: 2} .

echo "Finished docker build removing temporary Dockerfile"

rm ./Dockerfile

echo "Finished"