# AZ Go Starter

## Barebones Go web-app starter

### Features

- JWT authentication via [appleboy-jwt](https://github.com/appleboy/gin-jwt)
- MongoDB integration via [mgo](https://labix.org/mgo)
- API route versioning
- Govendor for dependency management

### Getting Started

- Set the GOPATH and GOBIN environment vars to the current directory
- It is recommended you set up a bash alias for this `alias gp='GOPATH=`pwd`; export GOPATH; GOBIN=$GOPATH/bin; export GOBIN'`
- If you don't want to set up a bash alias do `export GOPATH=`pwd && export GOBIN=$GOPATH/bin`

Run `./SETUP.sh`

After that run `make`

Then before starting the server copy over the sample configuration file to the `bin/` folder

Make sure you have MongoDB installed and `mongod` is running [Instructions to install Mongo](https://docs.mongodb.com/manual/installation/)

Start the server by doing `./bin/server`

To start the server from a different port modify the config.json file in the bin directory
