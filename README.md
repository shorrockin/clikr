# Overview

A really simple test bed to play with some tech. 

A simple one page app with an api backend to show how one might implement a system to support TV metadata. The app presents the users with a simple interface to scrub through some screenshots of a TV show. As they do so they can express interest in objects on screen by clicking on them. As they do this bonds are created between items. These bonds, in theory, represent shared interests - and as such - can be used to show or predict what other items one might be interested in based on the bonds and relationshipts created by other users.

This of course is all theoretical, and only really relevant at scale. This example could be expanded to show relationships between the metadata about objects, that however, is outside the scope of his demo.

The app uses the following tech:

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

# Getting Started

## Running

This section will outline how to quickly get the servers up and running using docker. If you just want to run the app and not worry about developing this is your fastest way to get things up and running.

Because this utilizes docker, that will be a pre-requisite for this section. To install docker please check out their [installation instructions](https://docs.docker.com/installation/).

Once you installed docker, and confirmed that it's functional (if using OSX, I recommend using Boot2Docker), you can do the following to download and run the clikr api image:

`> docker run --rm --name api shorrockin/clikr_api`

This will download everything needed (os, binaries, code - so it may take a bit the first time) needed to run the clikr golang based api server. It will run it within this container and bind to port 3000 inside the container. We do not need to expose the api port to our host machine, since we're going to link our containers together so they can talk directly.

The `--rm` parameter indicates that we want this container removed once it finishes. This does not mean we'll need to re-download the image, only that the container instance itself is transient.

We name this container `api` using the `--name` parameter. We'll then reference this name when we run our web server:

`> docker run --rm --name web --link api:api -p 8080:80 shorrockin/clikr_web`

Similar to before this will download everything we need to get the server up and running. This also does 2 other new things. 

The `--link` command makes our api server available within our web server. Our nginx server then acts as a reverse proxy to our api server sending it requests prefixed with `/api`. Within our container `api` acts as a named host (an entry in our `/etc/hosts` file) that we use in our nginx config to forward to. 

The `-p` command exposes our nginx port to our host machine. In this example we've bound it to `8080`, but you're welcome to bind to whatever you wish. 

You should now be able to visit the demo at [http://localhost:8080](http://localhost:8080) (modifying the port if you didn't use 8080).

**NOTE** If you're using Boot2Docker, your server **will not** be available at `localhost`. You'll need to access it by the IP of your Boot2Docker virtual instance (returned by `> boot2docker shellinit`).

## Developing

Looking to play around with code? This section outlines how to get the dev environment up and running. 

The dev setup for this project makes liberal use of _doing stuff on save_, this involves running our test suite, compiling our go code, restarting servers, compiling our less to css, etc. All this happens automatically on save. With that in mind, Boot2Docker (at the time of writing) is not up to the task, as a mounted NFS drive will not trigger the file system events necessary to detect file changes.

For development we utilize [Vagrant](https://www.vagrantup.com/) to build a small Ubuntu server that we run docker within. Instead of using a mounted NFS drive we use `vagrant rsync-auto` to copy our changes to our vagrant box. This section assumes some familiarity with Vagrant, and as such will not walk you through everything.

To get started, build our Vagrant instance with:

`> vagrant up`

Once you have your vagrant instance running, you you can copy the changes from your host box to the vagrant box with `vagrant rsync-auto`.

Now you're ready to start developing. `ssh` to your vagrant box and `cd` to the `clikr` directory. From here you can build all the docker containers we use in dev. We use [docker compose](https://docs.docker.com/compose/) in development to orchestrate building the dev containers, and setting up the shared resources between them. In dev mode we build the following containers:

* **gulp**: [gulp](http://gulpjs.com/) based build system to automatically watch for changes and build our coffeescript and less files into javascript and css. shares a file volume with the _web_ container so it can serve out the most up to date files.
* **web**: our nginx web server which serves our html, css and js. Also acts as a remote proxy for our _api_ server for any request starting with `/api`.
* **api**: our go based api server. Responds to api requests, more details on the api routes available is documented below.
* **test**: our automatic test server which uses [goconvey](http://goconvey.co/) to automatically run and report on the status of tests to our browser. shares a file system with our _api_ server. 

To build all these containers you can use the `./bin/build` script, which is simply a shortcut for:

`> docker-compose build`

Then start them all up together the `./bin/dev` script, which does:

`> docker-compose up`

Once you have all the containers running, you should be able to access (from your host box) the web server on port `8080`, the go-convey test results on port `9090` and the api server on `3000`. These port mappings are defined in the main `Vagrantfile`. 


# Architecture

This section outlines some high level decisions made about how things run, are structured, etc. It's not meant to be exaustive, but rather just a quick reference to understand things at a high level.

## Api Server

The api server (found in _/api_) is a simple go based server responding to a variety of requests and serving back the data over JSON. The following outlines the URLs currently supported by this sample app:


| **Name** | Get Episode Info |
|------------:|:-----|
| **Method** | GET |
| **URL** | `/series/[series_slug]/season/[season_num]/episode/[episode_num]` |
| **Description** | Retrieves the metadata about a given episode. Hypothetical at this point, server coded such that this will always return the same information regardless of what parameters are passed in. For this demo, this will be called on page load. |
| **Returns** | A json structure containing the episode, all the scenes, all the interaction points within those scenes, and for each interaction point the object that can be interacted with. |

| **Name** | Register Object Click |
|------------:|:-----|
| **Method** | PUT |
| **URL** | `/object/[object_id]/click` |
| **Description** | Used to register a _click_ on a specific object. The `object_id` must be valid object identifiers as return in the _Get Scene Info_ call. This call must include a Http header named `X-Clikr-User` that is a unique identifier representing the user who made this call. For this demo the client just picks an identifier on boot to use, it then uses this to tie information together between calls. The data about object clicks is stored in memory, and will be reset when the server reboots.|
| **Returns** | An array containing the suggested objects that the user may enjoy based on the current bonds between objects. |

# Concessions & Errata

The following lists the concessions made on this tech demo:

1. Multiple cliks on an object will register multiple points of interest. This makes testing easier, however, would likely not be somethingy ou wanted in reality. 
1. The build/watch cycle is brittle and the gulp subsystem will exit on errors.
1. CSS not written to be responsive. Once window width is reduced below a certain amount things will fall apart.
1. Docker container `test` is built from the `api` container, as such, any changes in `api` container will cause that container to be rebuilt - sometimes unecessarily.
1. Images should probably be minified in gulp build, instead just used image optim manually.
1. Scene scrubber should show a preview of scene's as you scrub through them instead of loading them right away (see youtube).
1. Images should be pre-loaded on boot. 
1. In reality you probably wouldn't want multiple clicks to be registered over and over, but for testing it's useful to see how the object bond is increased with other objects as you do.
1. Tons of other UI improvements could be made. But all tech demos need to stop at some point. 