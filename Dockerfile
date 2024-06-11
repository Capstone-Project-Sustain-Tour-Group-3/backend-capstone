# syntax=docker/dockerfile:1

FROM golang:1.22.2 AS build

WORKDIR /

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /tourease .

FROM gcr.io/distroless/base-debian11 AS build-release

WORKDIR /

COPY --from=build /tourease .
COPY --from=build /static/ /static/
COPY --from=build /docs/ /docs/
COPY --from=build /helpers/email_template.html /helpers/email_template.html
COPY --from=build /storages/ /storages/

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/tourease", "--seed" ]