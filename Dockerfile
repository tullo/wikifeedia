FROM golang:1.18.3-alpine3.16 as build_stage
RUN apk --no-cache add nodejs npm
WORKDIR /build/app
COPY app/package.json package.json
COPY app/package-lock.json package-lock.json
RUN npm install
WORKDIR /build
COPY . .
WORKDIR /build/app
RUN npm run build
WORKDIR /build
RUN go build

FROM alpine:3.16
WORKDIR /data
COPY --from=build_stage /build/wikifeedia /data/wikifeedia
VOLUME ["/data"]
CMD ["/data/wikifeedia"]
