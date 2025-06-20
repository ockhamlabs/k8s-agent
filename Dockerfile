FROM gcr.io/distroless/static-debian12:nonroot
ARG TARGETARCH
COPY bin/helios-agent-$TARGETARCH /usr/local/bin/helios-agent
CMD ["helios-agent"]
