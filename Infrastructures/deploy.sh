#!/bin/bash

SERVICE_NAME="api-go"
PROJECT_PATH="./public/api-go.homelab.test"

echo "Pulling latest code..."
cd $PROJECT_PATH
git pull origin main
cd ../../

echo "Rebuilding and restarting $SERVICE_NAME..."
docker compose up -d --build --force-recreate $SERVICE_NAME

echo "Deployment complete!"
docker compose ps | grep $SERVICE_NAME
echo ""
echo "Latest logs:"
docker compose logs --tail 50 $SERVICE_NAME