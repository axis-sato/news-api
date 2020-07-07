# News API

![Deploy API](https://github.com/c8112002/news-api/workflows/Deploy%20API/badge.svg)

## Getting started

### dev

running server
```bash
docker-compose up -d
```

generating
```bash
./bin/dev.sh generate!
./bin/dev.sh generate ./main.go
```

### production

running server
```bash
./bin/build.sh
./bin/api.sh
```
