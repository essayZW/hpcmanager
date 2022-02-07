#!/bin/sh

sed -i s/host:\ 127.0.0.1/host:\ $MYSQL_HOST/g /config-template.yaml
sed -i s/port:\ 3306/port:\ $MYSQL_PORT/g /config-template.yaml
sed -i s/database:\ hpcmanager/database:\ $MYSQL_DATABASE/g /config-template.yaml
sed -i s/username:\ mysqlroot/username:\ $MYSQL_USERNAME/g /config-template.yaml
sed -i s/password:\ mysqlpass/password:\ $MYSQL_PASSWORD/g /config-template.yaml



mkdir -p /root/.config/hpcmanager
cp /config-template.yaml /root/.config/hpcmanager/config-$HPCMANAGER_ENV.yaml

# run app
if [ "$HPCMANAGER_ENV" == "production" ]
then
    /main -debug=false -port=$PORT
else
    /main -debug=true -port=$PORT
fi
