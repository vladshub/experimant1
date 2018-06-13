# Experiment 1

## Secrets
copy .secrets.example to .secrets and fill it out

## Run
```
docker-compose up -d
sleep 50
docker-compose up -d
```
yes twice it takes kafka a while to be ready (about 50 seconds)

open [http://127.0.0.1:9999](http://127.0.0.1:9999)
