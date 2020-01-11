FROM golang:1.13-buster AS build-stage
WORKDIR /app
COPY . .
RUN make build

FROM debian:buster-slim
WORKDIR /root/
COPY --from=build-stage app/baelfire .
CMD [ "./baelfire" ]
