FROM alpine

COPY build/main /main
COPY docker/config-template.yaml /config-template.yaml
COPY docker/entrypoint.sh /entrypoint.sh


ENV MYSQL_HOST=127.0.0.1
ENV MYSQL_PORT=3306
ENV MYSQL_DATABASE=hpcmanager
ENV MYSQL_USERNAME=root
ENV MYSQL_PASSWORD=password
ENV ETCD_ADDRESS=172.17.0.1

ENV PORT=80


ENV HPCMANAGER_ENV=production

CMD ["sh", "/entrypoint.sh"]
