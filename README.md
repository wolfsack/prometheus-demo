# Prometheus-Demo

## Try it yourself
    clone repository with:
    
    git clone https://github.com/wolfsack/prometheus-demo.git

## How to run it?
inside the local repository directory:

    with docker and docker-compose installed run:
    
        docker-compose up
        
        
    If you don't have permission to create new files on the host system prometheus won't be able to start.
    In the docker-compose.yaml is a bind-mount for the directory "prom" defined. Prometheus needs permission to write files there.

## Endpoints
----
Prometheus Dashboard

    localhost:9090      
----
Prometheus Self-Monitoring

    localhost:9090/metrics
---- 
Node-Exporter

    localhost:9100/metrics
---- 
Go-Demo

    localhost:8090/metrics  
----



    
