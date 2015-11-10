# Go + Docker
An example built for the Medellin DevOps meetup.

### Dependencies
- [Docker](http://docs.docker.com/)
- [Docker Machine](http://docs.docker.com/machine/install-machine/) *(If you're not on Linux)*

### Run the example
 ~~~sh
 $ docker build -t server
 $ docker run --publish 6060:8080 --rm server
 ~~~
Go to http://localhost:6060/Johnny, and say hello!

**Note:** If you're not on Linux, make sure you installed Docker Machine.
Then run
```sh
$ docker-machine <name of your VM> ip
```
To get your virtual machine's IP. Replace "localhost" with that IP in the above URL,
and you should be ready to go!
