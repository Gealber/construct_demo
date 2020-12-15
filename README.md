# Construct Demo

This is a simple demo of an API for a constructor application

# How to run?

> Common for all plattform
* Make sure you're on the `script` folder and change in depend on your plattform.
* Make sure you have an installation of Redis and MongoDB.
* Run the redis server from the folder `script/config`, in this way `redis-server redis.conf`.
* Make sure the MongoDB is working properly.

> Ubuntu or Linux based OS, only tested on Ubuntu.

```bash
./run.sh
```

> Windows

Not working properly, yet.

> Using Docker and GNU Makefile

1. First we'll need to build the image of the server

```bash
make
```

Just that, 

2. Second after the creation of the server image we'll need to run the containers

```bash
make up
```

This will download and setup MongoDB, Redis and the server. After this you will be able to
access the `API` on port `3000`, take a look at the documentation of the `API`.

3. When you're done and want to tear down the docker containers, you can achive that with

```bash
make down
```

**NOTE**: In case you're using `mingw32-make` you should run `mingw32-make` instead of `make`.
