FROM alpine:3.16.3 as backend
WORKDIR /
COPY binary/jadechamber /jadechamber
RUN ls -la
ENTRYPOINT ["/jadechamber"]