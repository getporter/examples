FROM ghcr.io/getporter/examples/images/whalesay

COPY whalesayd /usr/local/bin/whalesayd
RUN ls /usr/local/bin/
EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/whalesayd"]