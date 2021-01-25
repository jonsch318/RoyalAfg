<h1 align="center">RoyalAfg - Online Casino</h1>
<p align="center">
	RoyalAfg is a online casino developed using react <a href="https://nextjs.org/">Nextjs</a> and <a href="https://golang.org/">Go</a>
  <i>It was designed in an microservice architecture
    <br>for a <b>special learning achievement.</b></i>
  <br>
</p>


## Documentation

To run individual Services review the build and run guide of your choosing.

 - [Motivations](#motivations)
 - [Getting Started](#Installation)
 - [Architecture](#architecture)

For a complete deployment see this.
 - [Deploy](#Deploy)
 
 ## Motivations
 While this is a project for a special learning achievement, it also serves as a full example of a Kubernetes and microservice oriented application.

A online casino is a perfect example for a microservice architecture, because many services need to communicate with each other, which is the main problem of this pattern.

## Installation
### Bazel
This project is build with [Bazel](https://bazel.build/) tools, which create a sandbox to compile the source code.
The authentication service would be build using the following output

	#Builds the authentication service
	bazel build //services/auth:auth
	
	#Run the authentication service (does not 
	require the previous step)
	bazel run //services/auth:auth
	
if you want to specify the configuration used by the service you can set the `--config` Flag like this
	
	bazel run //services/auth:auth --config=./pathToYourConfig
	
Bazel is able to build directly to containers from the source code.
This can be done using the 

	bazel build //services/auth:image --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64

Notice that the `--platform` is important to tell Bazel that the go code should be compiled for the Linux runtime, since the Container is run in Linux.

### Standard Go Compiler
Apart from the  [Bazel](#Bazel) build tool every service can be build using the standard go tools
	
	#Build authentication service
	go build ./services/auth/main.go
	
	#Run the service
	go run ./serivces/auth/main.go --config=./pathToConfig
	
### Build the docker image
The Docker file in the root directory can accommodate every go service of the system.
To specify a service to be build to an Docker image use the following command

	docker build -t royalafg_auth --build-arg service=./services/auth/main.go .

This will create the Docker image with the tag of `royalafg_auth`. Run `docker run royalafg_auth` to run the application

## Deploy
I publish each service as a docker container in my [docker hub page](https://hub.docker.com/u/johnnys318). This is used by the deployment scripts aswell.


To deploy the application on linux simply run
	
	make deploy

which should deploy each service configured with default configuration. **Keep in mind this also includes encryption keys and other private keys. Only use it for testing purposes**

This script has the follwing requirements:

 - access to a terminal
 - docker installed
 - [minikube](https://minikube.sigs.k8s.io/docs/) installed. (On Linux the `docker` or `none` options are available to prevent the extra virutal machine)
 - [kubectl](https://kubernetes.io/de/docs/tasks/tools/install-kubectl/) installed (comes with minikube)
 - [Helm](https://helm.sh/) installed

It will fire of a new minikube instance, which will configure kubectl and helm and the deploy the deployments described in `/deployments` in the correct order. In addition to that it will add the required helm repositories and install them. This includes [Agones](https://agones.dev/site/) .

## Architecture
Following services are integrated into the system:

 - [Auth](https://github.com/JohnnyS318/RoyalAfg/tree/master/services/auth)
 - [User](https://github.com/JohnnyS318/RoyalAfg/tree/master/services/user)
 - [Bank](https://github.com/JohnnyS318/RoyalAfg/tree/master/services/bank)
 - [Poker-Matchmaker](https://github.com/JohnnyS318/RoyalAfg/tree/master/services/poker-matchmaker)
 - [Poker](https://github.com/JohnnyS318/RoyalAfg/tree/master/services/poker)
 - [Search](https://github.com/JohnnyS318/RoyalAfg/tree/master/services/search)
 - [Docs](https://github.com/JohnnyS318/RoyalAfg/tree/master/services/docs)
 - [Web](https://github.com/JohnnyS318/RoyalAfg/tree/master/services/web)

//TODO system overview.
