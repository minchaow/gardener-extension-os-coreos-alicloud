#!/bin/bash
mkdir -p '/'
cat << EOF | base64 -d > '/foo'
YmFy
EOF
chmod '0600' '/foo'

mkdir -p '/etc/systemd/system/containerd.service'
cat << EOF | base64 -d > '/etc/systemd/system/containerd.service/11-exec-start-config.conf'
W1NlcnZpY2VdCkV4ZWNTdGFydD0KRXhlY1N0YXJ0PS9ydW4vdG9yY3gvYmluL2NvbnRhaW5lcmQgLS1jb25maWcgL2V0Yy9jb250YWluZXJkL2NvbmZpZy50b21sClJlc3RhcnQ9b24tZmFpbHVyZQ==
EOF

mkdir -p '/etc/systemd/system/docker.service'
cat << EOF | base64 -d > '/etc/systemd/system/docker.service/11-exec-start-config.conf'
W1NlcnZpY2VdCkV4ZWNTdGFydD0KRXhlY1N0YXJ0PS9ydW4vdG9yY3gvYmluL2RvY2tlcmQgLS1ob3N0PWZkOi8vIC0tY29udGFpbmVyZD0vcnVuL2NvbnRhaW5lcmQvY29udGFpbmVyZC5zb2NrIC0tc2VsaW51eC1lbmFibGVkPXRydWU=
EOF
