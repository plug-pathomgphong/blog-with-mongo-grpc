FROM envoyproxy/envoy-dev:9105f45c7fb872d1db2bf8a9bc908368effe77cd
COPY ./envoy.yaml /etc/envoy/envoy.yaml
RUN chmod go+r /etc/envoy/envoy.yaml



# FROM envoyproxy/envoy-dev:9105f45c7fb872d1db2bf8a9bc908368effe77cd
# COPY ./envoy.yaml /etc/envoy/envoy.yaml

# CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml

