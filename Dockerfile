ARG GLOBAL_WARMING
FROM ${GLOBAL_WARMING}/golang:1.18.1-alpine3.15 as dev
WORKDIR /main/
COPY go.* ./
ENV CGO_ENABLED=0
RUN go env -w GOPRIVATE=github.com/sayan/GlobalWarming

FROM dev as build
ARG GITHUB_TOKEN

RUN apk add git
RUN git config --global url."https://golang:${GITHUB_TOKEN}@github.com".insteadOf "https://github.com"
RUN apk add --no-cache make
COPY . ./
RUN make build

FROM ${GLOBAL_WARMING}/alpine:3.14.0 as release
WORKDIR /app/
COPY --from=build /main/dist/ /app/

CMD ["./GlobalWarming"]
