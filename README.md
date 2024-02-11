# if-i-am-gone-web-backend

Backend part of the **if-i-am-gone** project

## Checklist before start:

Config files path (relative to the root of the project - this readme):

- /config-dev.yaml
- /config-prod.yaml

**DEV** config file:

```yaml
api: # config parameters for api server
  hostname: 127.0.0.1
  port: 8080

postgres: # config parameters for postgres database
  host: ifidie-dev-pg # dev name
  port: 5432
  db_name: IFIDIE
  user: xxx # confidential
  password: xxx # confidential
```

**PROD** config file:

```yaml
api: # config parameters for api server
  hostname: 127.0.0.1
  port: 8080

postgres: # config parameters for postgres database
  host: ifidie-prod-pg # prod name
  port: 5432
  db_name: IFIDIE
  user: xxx # confidential
  password: xxx # confidential
```
