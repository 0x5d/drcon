# Service discovery with DRCoN
### Docker + Consul + Registrator + Consul Template + Nginx
An example built for the Medellin DevOps meetup, based on [this
article](http://www.maori.geek.nz/scalable_architecture_dr_con_docker_registrator_consul_nginx/)
.

## Dependencies
- [Docker](http://docs.docker.com/)
- [Docker Machine](http://docs.docker.com/machine/install-machine/) *(If you're not on Linux)*
- [DR-CoN Container](https://github.com/grahamjenson/DR-CoN)

## Run the example

### Run Consul:
```sh
$ docker run -it -h node -p 8500:8500 -p 8600:53/udp progrium/consul -server -bootstrap -advertise
$(docker-machine ip <machine name>) -log-level debug
 ```

### Run Registrator:
```sh
$ docker run -it -v /var/run/docker.sock:/tmp/docker.sock -h $(docker-machine ip <name of your VM>)
gliderlabs/registrator consul://$(docker-machine ip <machine name>):8500
```

### Run DR-CoN
If you haven't pull the DR-CoN container Dockerfile.
```sh
$ git clone https://github.com/grahamjenson/DR-CoN.git
```
Build it:
```sh
docker run -it -e "CONSUL=$(docker-machine env <machine name>):8500" -e "SERVICE=server" -p 80:80
dr-con
```

### Run the Go server:
```sh
$ docker build -t server
$ docker run --publish 6060:8080 --rm server
```
Go to [http://localhost:6060/Johnny](http://localhost/johnny), and say hello!

**Note:** If you're not on Linux, make sure you installed Docker Machine.
Then run
```sh
$ docker-machine ip <name of your VM>
```
To get your virtual machine's IP. Replace "localhost" with that IP in the above URL,
and you should be ready to go!
Notice that if you reload the page multiple times, the displayed IP changes. That's Consul
Template doing its magic!

You can also check how Registrator detects when a new service is run and registers it on the
Consul server. Go to
[http://localhost:8500](http://localhost:8500) and you should see it on Consul's UI.
*Replace "localhost" for your VM's IP if you're not on linux.*
