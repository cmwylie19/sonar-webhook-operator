# Unit Test
```
go test -covermode=atomic -coverprofile=coverage.out -v ./... 
```
# TLS Config
```
ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
insecure_skip_verify: true
bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token
```

Quick Build
```
docker build -t docker.io/cmwylie19/sonar-webhook .; docker push docker.io/cmwylie19/sonar-webhook
```

## Test
```
curl https://freshlist.us/webhook/healthz

curl http://localhost:8080/webhook/x/results


curl -X POST -d '
{
    "serverUrl": "http://localhost:9000",
    "taskId": "AVh21JS2JepAEhwQ-b3u",
    "status": "SUCCESS",
    "analysedAt": "2016-11-18T10:46:28+0100",
    "revision": "c739069ec7105e01303e8b3065a81141aad9f129",
    "project": {
        "key": "myproject",
        "name": "My Project",
        "url": "https://mycompany.com/sonarqube/dashboard?id=myproject"
    },
    "properties": {
        "buildNumber":"69"
    },
    "qualityGate": {
        "conditions": [
            {
                "errorThreshold": "1",
                "metric": "new_security_rating",
                "onLeakPeriod": true,
                "operator": "GREATER_THAN",
                "status": "OK",
                "value": "1"
            },
            {
                "errorThreshold": "1",
                "metric": "new_reliability_rating",
                "onLeakPeriod": true,
                "operator": "GREATER_THAN",
                "status": "OK",
                "value": "1"
            },
            {
                "errorThreshold": "1",
                "metric": "new_maintainability_rating",
                "onLeakPeriod": true,
                "operator": "GREATER_THAN",
                "status": "OK",
                "value": "1"
            },
            {
                "errorThreshold": "80",
                "metric": "new_coverage",
                "onLeakPeriod": true,
                "operator": "LESS_THAN",
                "status": "NO_VALUE"
            }
        ],
        "name": "SonarQube way",
        "status": "OK"
    }
}' http://localhost:8080/webhook/x/post


# }' https://freshlist.us/webhook/store

```
curl -X POST -H "X-Sonar-Webhook-HMAC-SHA256:2dce505d96a53c5768052ee90f3df2055657518dad489160df9913f66042e160" -d '
{
    "serverUrl": "http://localhost:9000",
    "taskId": "AVh21JS2JepAEhwQ-b3u",
    "status": "SUCCESS",
    "analysedAt": "2016-11-18T10:46:28+0100",
    "revision": "c739069ec7105e01303e8b3065a81141aad9f129",
    "project": {
        "key": "myproject",
        "name": "My Project",
        "url": "https://mycompany.com/sonarqube/dashboard?id=myproject"
    },
    "properties": {
        "buildNumber":"69"
    },
    "qualityGate": {
        "conditions": [
            {
                "errorThreshold": "1",
                "metric": "new_security_rating",
                "onLeakPeriod": true,
                "operator": "GREATER_THAN",
                "status": "OK",
                "value": "1"
            },
            {
                "errorThreshold": "1",
                "metric": "new_reliability_rating",
                "onLeakPeriod": true,
                "operator": "GREATER_THAN",
                "status": "OK",
                "value": "1"
            },
            {
                "errorThreshold": "1",
                "metric": "new_maintainability_rating",
                "onLeakPeriod": true,
                "operator": "GREATER_THAN",
                "status": "OK",
                "value": "1"
            },
            {
                "errorThreshold": "80",
                "metric": "new_coverage",
                "onLeakPeriod": true,
                "operator": "LESS_THAN",
                "status": "NO_VALUE"
            }
        ],
        "name": "SonarQube way",
        "status": "OK"
    }
}' http://localhost:8080/webhook/store
```

```
curl -X POST -H "X-Sonar-Webhook-HMAC-SHA256:2dce505d96a53c5768052ee90f3df2055657518dad489160df9913f66042e160" -d '

https://gist.github.com/maoueh/624f108ee2f3e6ca0b496d6c2f75bcd7

https://www.mongodb.com/community/forums/t/finding-data-between-two-dates-by-using-a-query-in-mongodb-charts/102506/3

https://www.alexedwards.net/blog/making-and-using-middleware
https://betterprogramming.pub/hands-on-with-jwt-in-golang-8c986d1bb4c0
https://github.com/hamzawix/jwt-auth-go

https://betterprogramming.pub/hands-on-with-jwt-in-golang-8c986d1bb4c0 # Has Testing

echo -n "value" | openssl dgst -sha256 -hmac "secret"


cover () { 
    t="/tmp/go-cover.$$.tmp"
    go test -coverprofile=$t $@ && go tool cover -html=$t && unlink $t
}