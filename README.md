# News API

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
docker build -t news-api .
docker run -p 3001:3001 news-api
```
