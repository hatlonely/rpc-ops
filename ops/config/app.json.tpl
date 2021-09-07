{
  "grpcGateway": {
    "httpPort": 80,
    "grpcPort": 6080,
    "exitTimeout": "20s",
    "validators": [
      "Default"
    ],
    "usePascalNameLogKey": false,
    "usePascalNameErrKey": false,
    "marshalUseProtoNames": true,
    "marshalEmitUnpopulated": false,
    "unmarshalDiscardUnknown": true,
    "enableTrace": false,
    "enableMetric": false,
    "enablePprof": false,
    "jaeger": {
      "serviceName": "rpc-ancient",
      "sampler": {
        "type": "const",
        "param": 1
      },
      "reporter": {
        "logSpans": false
      }
    }
  },
  "service": {
    "ops": {
      "mongo": {
        "retry": {
          "attempts": 1,
          "delay": "1s",
          "maxDelay": "3m",
          "lastErrorOnly": true,
          "delayType": "BackOff"
        },
        "wrapper": {
          "name": "mongo",
          "enableTrace": false,
          "enableMetric": false
        },
        "mongo": {
          "uri": "mongodb://localhost:27017",
          "connectTimeout": "3s",
          "pingTimeout": "2s"
        }
      },
      "database": "ops",
      "repositoryCollection": "repository",
      "timeout": "10s",
      "jobExpiration": "72h"
    }
  },
  "logger": {
    "grpc": {
      "level": "Info",
      "writers": [{
        "type": "RotateFile",
        "options": {
          "filename": "log/app.rpc",
          "maxAge": "24h",
          "formatter": {
            "type": "Json",
            "options": {
              "flatMap": true,
              "pascalNameKey": true
            }
          }
        }
      }]
    },
    "info": {
      "level": "Info",
      "writers": [{
        "type": "RotateFile",
        "options": {
          "filename": "log/app.log",
          "maxAge": "24h",
          "formatter": {
            "type": "Json",
            "options": {
              "pascalNameKey": true
            }
          }
        }
      }]
    }
  }
}