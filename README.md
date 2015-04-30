# Clikr

A really simple test bed to play with some tech. A simple one page app with a backend to show how one might implement a system to support TV metadata. The app uses the following tech:

* **Go:**  There is a golang based server that uses some sample in-memory data which is then retrieved from the javascript based front end. 
* **React:** The front infrastructure is based on facebook's react.
* **Flux:** We utilize facebook's flux pattern within react to model the information flow within the front-end.
* **Goconvey:** We use goconvey to automate the running of our tests, the goconvey server is run in it's own container.
* **Docker:** We utilize docker (along with docker-compose) to bootstrap various containers both for development as well as the hypothetical production environent.
* **Gulp:** We use gulp (with a suite of plugins like browser-sync, browserify, coffee-script, etc) to build our front-end layer.

## Project Layout

* **/api** contains all the go code used to run the api server.
* **/bin** some simple shell scripts (mostly around docker) to do some simple stuff like re-build all docker instances (_build_), run a bash shell (_shell_), startup all the docker containers via docker-compose (_dev_).
* **/provision** contains a script used by vagrant to initialize an environment you can develop within.
* **/test** contains the code (docker based) needed to build the goconvey based test server.
* **/web** contains the code needed to build and run the front end. builds a nginx docker container which servers all our html/css/js and also acts as a reverse proxy to our api server.

## Running

This section will outline how to quickly get the servers up and running using docker. If you just want to run the app and not worry about developing this is your fastest way to get things up and running.

TODO

## Developing

Looking to play around with code? This section outlines how to get the dev environment up and running. 

TODO

# Architecture

This section outlines some high level decisions made about how things run, are structured, etc. It's not meant to be exaustive, but rather just a quick reference to understand things at a high level.
