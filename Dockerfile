FROM golang:alpine AS build

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp

###
FROM scratch as final
COPY --from=build /app/myapp .
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Shanghai

CMD ["/myapp"]