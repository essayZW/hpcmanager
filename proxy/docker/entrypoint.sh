#!/bin/bash

sed -i s/host:\ 127.0.0.1/host:\ $MYSQL_HOST/g /config-template.yaml
sed -i s/port:\ 3306/port:\ $MYSQL_PORT/g /config-template.yaml
sed -i s/database:\ hpcmanager/database:\ $MYSQL_DATABASE/g /config-template.yaml
sed -i s/username:\ mysqlroot/username:\ $MYSQL_USERNAME/g /config-template.yaml
sed -i s/password:\ mysqlpass/password:\ $MYSQL_PASSWORD/g /config-template.yaml
sed -i s/address:\ etcdaddress/address:\ $ETCD_ADDRESS/g /config-template.yaml

sed -i s/gatewayaddress/$GATEWAY_ADDRESS/g /etc/nginx/nginx.conf



mkdir -p /root/.config/hpcmanager
cp /config-template.yaml /root/.config/hpcmanager/config-$HPCMANAGER_ENV.yaml

# run nginx
nginx -g 'daemon on;'

# run app
/main
