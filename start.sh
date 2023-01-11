#!/bin/bash
docker stop $(docker ps -qa) && docker system prune -af --volumes
docker compose up