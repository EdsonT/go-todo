FROM envoyproxy/envoy-dev
RUN apt-get update
COPY envoy.yml /etc/envoy.yml
CMD /usr/local/bin/envoy -c /etc/envoy.yml