# weather/api

> [!NOTE]
> Deployment is maintained externally at [fruzitent/infra](https://git.fruzit.pp.ua/fruzitent/infra/compare/main...ses/5.0)

## Quick Start

Requirements:

- [atlas](https://atlasgo.io/getting-started#installation)

```shell
gum input --password | install -D "/dev/stdin" "./weatherapi-token"
weather-api \
  --smtp.from "weather@example.org" \
  --smtp.host "mail.example.org" \
  --smtp.password "file://smtp-password" \
  --smtp.port 465 \
  --smtp.username "smtp-username" \
  --weatherApi.secret "file://weatherapi-secret" \
  daemon
```
