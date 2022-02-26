#!/usr/bin/sh

for i in `seq 1 6`; do ifconfig eth0:$i 10.29.1.$i/16 ; done;

for i in `seq 1 2`; do ifconfig eth0:$($i+6) 10.29.2.$i/16 ; done;
for i in `seq 1 6`;do mkdir -p storage_data/$i/objects;done;
service rabbitmq-server start

python3 /usr/bin/rabbitmqadmin declare exchange name=dataServers type=fanout
python3 /usr/bin/rabbitmqadmin declare exchange name=apiServers type=fanout
rabbitmqctl add_user test test
rabbitmqctl set_permissions -p / test ".*" ".*" ".*"