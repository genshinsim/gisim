FROM alpine:3.16.3 as backend
WORKDIR /
COPY backend/artifacts/preview /preview
RUN ls -la
ENTRYPOINT ["/preview"]