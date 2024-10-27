#!/bin/bash

# 查看 secret 使用证书的 serial, 可以与更新后的证书文件进行对比
# openssl x509 -noout -serial -in <(kubectl -n ${namespace} get secret/${secret_name} -o jsonpath='{.data.tls\.crt}' | base64 -d)

namespace=sdk-h5

do_update() {
    local domain=$1
    local secret_name=$2
    local cert_dir=/home/ubuntu/ingress/ssl
    local cert_file=${cert_dir}/fullchain.cer
    local key_file=${cert_dir}/${domain}.key

    if [ "$(openssl x509 -noout -serial -in ${cert_file})" != "$(openssl x509 -noout -serial -in <(kubectl -n ${namespace} get secret/${secret_name} -o jsonpath='{.data.tls\.crt}' | base64 -d))" ]; then
        kubectl create secret tls ${secret_name} -n ${namespace} --cert=${cert_file} --key=${key_file} --dry-run=client -o yaml | kubectl apply -f -
        echo secret ${domain} renew
    else
        echo no need renew ${domain} secret
    fi
}

do_update zrqsmcx.top zrqsmcx.top