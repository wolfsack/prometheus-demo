# get base image from docker hub
FROM prom/prometheus:v2.30.3

# copy the prometheus config file into into to the image folder
# where it will be read on start
ADD prometheus.yml /etc/prometheus/

# prometheus runs on port 9090 by default inside the container
EXPOSE 9090