# Ajax Temperature Monitor

The part of the project for monitoring temperature in the apartment with Ajax app + Apple Shortcuts + Go + InfluxDB + Grafana.
More on https://morskov.com/blog/2024/01/14/ajax-temperature-monitoring

## Installation

1. Clone

2. Setup InfluxDB:

```
docker-compose up -d influxdb
docker-compose exec influxdb influx setup
```

4. Get auth token for InfluxDB:

```
docker-compose exec influxdb influx auth create \
  --org org \
  --all-access
```

5. Put InfluxDB config params in .env file

6. Create the app `docker-compose up -d app`

## Deploy

You can fetch changes on the server with `./deploy.sh`. It's just run several commands by SSH.
