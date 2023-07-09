[![Docker build](https://img.shields.io/github/actions/workflow/status/gavrylenkoIvan/gonotes/docker-image.yml?branch=master&label=docker%20build&logo=github)](https://github.com/gavrylenkoIvan/gonotes/actions?query=workflow)
[![Build Status](https://img.shields.io/github/actions/workflow/status/gavrylenkoIvan/gonotes/go.yml?branch=master&label=build&logo=github)](https://github.com/gavrylenkoIvan/gonotes/actions?query=workflow)
[![Go Report Card](https://goreportcard.com/badge/github.com/gavrylenkoIvan/gonotes)](https://goreportcard.com/report/github.com/gavrylenkoIvan/gonotes)
![Scrutinizer Code Quality](https://img.shields.io/scrutinizer/quality/g/gavrylenkoIvan/gonotes/master)
![GitHub](https://img.shields.io/github/license/gavrylenkoIvan/gonotes)
![GitHub last commit](https://img.shields.io/github/last-commit/gavrylenkoIvan/gonotes)

<p align="center">
  <img src="https://github.com/gavrylenkoIvan/gonotes/blob/master/images/logo.png" />
</p>

## Note taking app written with golang an vue 3

## Installation
### You can simply install package:
```sh
$ docker pull ghcr.io/gavrylenkoivan/gonotes:master
```
### Or build it locally:
* #### First of all, you will need [docker](https://www.docker.com) dowloaded and started on your pc.
* #### Dowload this repo and open it.
* #### Set environment vars in .env file (those are really informative so I think you won't have any troubles), change enironment in docker-compose file if needed
* #### To start container and init database, run: 
  ```sh
  $ make compose
  ```
* #### Enjoy!

---

<div style="margin-top: 100px;">
  
  ### Screenshots:

  ![alt text](https://github.com/gavrylenkoIvan/gonotes/blob/master/images/main-page.png)

  ![alt text](https://github.com/gavrylenkoIvan/gonotes/blob/master/images/add-note.png)

  ![alt text](https://github.com/gavrylenkoIvan/gonotes/blob/master/images/notes.png)
  
</div>
