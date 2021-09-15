# rpc-ops

## 运维

```shell
ops --variable .cfg/dev.yaml -a run --env dev --task image
ops --variable .cfg/dev.yaml -a run --env dev --task helm
ops --variable .cfg/dev.yaml -a run --env dev --task helm --cmd=install
ops --variable .cfg/dev.yaml -a run --env dev --task helm --cmd=upgrade

curl -X POST -H "Origin:hello" 127.0.0.1/v1/repository \
  -d '{
  "username": "hatlonely",
  "password": "xx",
  "endpoint": "github.com",
  "team": "hatlonely",
  "name": "ops"
}'


curl -H "Origin:hello" 127.0.0.1/v1/repository/d69658805d8340d38ece83752ff27288

curl -X POST  -H "Origin:hello" 127.0.0.1/v1/describeRepository \
  -d '{
  "id": "d69658805d8340d38ece83752ff27288"
}'
```