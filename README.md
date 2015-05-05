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

TODO

## Developing

Looking to play around with code? This section outlines how to get the dev environment up and running. 

TODO

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
| **Description** | Used to register a _click_ on a specific object. The `object_id` must be valid object identifiers as return in the _Get Scene Info_ call. This call must include a Http header named `X-Clikr-User` that is a unique identifier representing the user who made this call. For this demo the client just picks an identifier on boot to use, it then uses this to tie information together between calls.|
| **Returns** | An array containing the suggested objects that the user may enjoy based on the current bonds between objects. |

# Concessions

The following lists the concessions made on this tech demo:

1. Multiple cliks on an object will register multiple points of interest. This makes testing easier, however, would likely not be somethingy ou wanted in reality. 
1. The build/watch cycle is brittle and the gulp subsystem will exit on errors.
1. CSS not written to be responsive. Once window width is reduced below a certain amount things will fall apart.
1. Docker container `test` is built from the `api` container, as such, any changes in `api` container will cause that container to be rebuilt - sometimes unecessarily.
1. Images should probably be minified in gulp build, instead just used image optim manually.
1. Scene scrubber should show a preview of scene's as you scrub through them instead of loading them right away (see youtube).
1. Images should be pre-loaded on boot. 
1. In reality you probably wouldn't want multiple clicks to be registered over and over, but for testing it's useful to see how the object bond is increased with other objects as you do.