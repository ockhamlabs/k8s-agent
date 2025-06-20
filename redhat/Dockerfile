FROM redhat/ubi9-minimal:9.1
ARG TARGETARCH=amd64
COPY bin/helios-agent-$TARGETARCH /usr/local/bin/helios-agent
COPY LICENCE /licenses/LICENCE
USER 1001010000
CMD ["helios-agent"]
