#!/bin/bash

ssh apps@ajax-monitor "cd /opt/docker/ajax-monitor; git pull; docker-compose up -d --build app"

