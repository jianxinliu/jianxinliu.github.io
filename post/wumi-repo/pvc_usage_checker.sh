#!/bin/bash

usage=$(kubectl -n loki-stack exec -it pod/loki-0 -- df -h -t ext4 | grep vdc | awk '{print $5}')
p=$(echo ${usage} | grep -E -o "[0-9]+")
wechat_bot_addr="https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=b7c865fe-2c12-40fb-9570-4dc9d50c1402"
echo ${p}
if [ ${p} -ge 80 ]; then
        echo send
        #curl ${wechat_bot_addr} -H 'Content-Type: application/json' -d '{"msgtype": "markdown", "markdown": {"content": "hk k8s loki pv usage: <font color=\"warning\">'${usage}'</font>"}}'
fi
