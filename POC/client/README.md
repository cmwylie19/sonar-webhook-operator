# Test Endpoints
- `/webhook/healthz`
- `/webhook/store`
- `/webhooker/post`
- `/webhook/post`
- `/webhook/views`

```
# Healthz
curl http://localhost:8080/webhook/healthz

# Sonarendpoint
curl -d '{"status":"Success"}' http://localhost:9001/

# GetPost
curl -d '{"status":"Success"}' http://localhost:9001/post

# GetPostTest
curl -d '{"status":"Success"}' http://localhost:8080/webhook/post
```