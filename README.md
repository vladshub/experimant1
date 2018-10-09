# Experiment 1

## Secrets
copy .secrets.example to .secrets and fill it out

## Run
```
docker-compose up -d
docker-compose logs -f
```
Sometimes you will see that api and topicextractor had shut down because of missing kafka broker just run `docker-compose up -d` again.

open [http://127.0.0.1:9999](http://127.0.0.1:9999)

Adding some stuff
