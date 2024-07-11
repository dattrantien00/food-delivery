# FROM alpine
# RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
# RUN update-ca-certificates

# WORKDIR /app/
# ADD ./app /app/
# # ADD ./zoneinfo.zip /usr/local/go/lib/time/
# ADD ./demo.html /app/
# RUN ls -la /app/ && file /app/app
# ENTRYPOINT ["./app"]
FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates

# Cài đặt thêm các gói cần thiết cho kiểm tra
RUN apk add file

WORKDIR /app/
ADD ./app /app/
# ADD ./zoneinfo.zip /usr/local/go/lib/time/
ADD ./demo.html /app/

# Thêm quyền thực thi và kiểm tra file /app/app
RUN ls -la /app/ && file /app/app

ENTRYPOINT ["./app"]