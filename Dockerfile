FROM golang:1.27 AS build-env
WORKDIR /banks

COPY ./go.mod /banks/
COPY ./go.sum /banks/

RUN go mod download
RUN go mod verify

COPY . .

RUN CGO_ENABLED=0
RUN go install .

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin /
CMD ["/banks"]
