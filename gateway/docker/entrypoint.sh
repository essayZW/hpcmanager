#!/bin/sh

sed -i s/address:\ etcdaddress/address:\ $ETCD_ADDRESS/g /config-template.yaml
sed -i s/address:\ 172.17.0.4:6379/address:\ $REDIS_ADDRESS/g /config-template.yaml

mkdir -p /root/.config/hpcmanager
cp /config-template.yaml /root/.config/hpcmanager/config-$HPCMANAGER_ENV.yaml

sleep 20

# run app
if [ "$HPCMANAGER_ENV" == "production" ]
then
    /main -debug=false -port=$PORT
else
    /main -debug=true -port=$PORT
fi
