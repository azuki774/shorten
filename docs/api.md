## Redirector

### GET /
- HealthCheck

### GET /{short_url}
- 301 StatusFound

## Register

### GET /
- HealthCheck

### POST /regist
- JSON
- request
    - JSON : `{"target_url" : "<target_url>"}`
    - Optional field: `{"expired_in" : <int (sec)>}`
- response
    - JSON : `{"short_url":"<short_url>","target_url":"<target_url>","expired_at":"2038-01-01T00:00:00+09:00"}`
