# Prometheus-Demo

### Try it yourself
    clone repository with:
    
    git clone https://github.com/wolfsack/prometheus-demo.git

### How to run it?
inside the local repository directory:

    with docker and docker-compose installed run:
    
        docker-compose up

----

    or from a custom image run:

        docker build -t prom-demo .

        docker run --rm -d -p 9090:9090 --name prom_demo -v $pwd/prom:/prometheus prom-demo --config.file=/etc/prometheus/prometheus.yml 


    - build docker image from docker file
      - with tag/name prom-demo
      - Dockerfile in this directory
    - run the container with
      - external port 9090 mapped to internal port 9090
      - container name prom_demo
      - mount /prom to /prometheus prom-demo
      - select image (prom-demo)
      - run with command that specifies the config file (prometheus command)
----
Prometheus dashboard can be accessed on 

    host_ip:9090 
    localhost:9090
        
----
Prometheus self-monitoring

    hostip:9090/metrics
    localhost:9090/metrics

---- 
    