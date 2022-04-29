FROM alpine

RUN apk update && apk add tzdata

COPY build/main /main
COPY docker/config-template.yaml /config-template.yaml
COPY docker/entrypoint.sh /entrypoint.sh


ENV MYSQL_HOST=127.0.0.1
ENV MYSQL_PORT=3306
ENV MYSQL_DATABASE=hpcmanager
ENV MYSQL_USERNAME=root
ENV MYSQL_PASSWORD=password
ENV ETCD_ADDRESS=172.17.0.1
ENV REDIS_ADDRESS=172.17.0.1
ENV RABBITMQ_ADDRESS=172.17.0.6:5672


# 由于hpc服务部署在docker中只是为了测试,所以设置为dev环境
ENV HPCMANAGER_ENV=dev

CMD ["sh", "/entrypoint.sh"]
