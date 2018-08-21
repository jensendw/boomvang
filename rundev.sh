#!/usr/bin/env bash
export BOOMVANG_RABBITMQ_URL=https://0fc599c9-5f1b-48bb-82a3-5d3057bce11c.mock.pstmn.io
export BOOMVANG_TEST=foobar
export BOOMVANG_RABBITMQ_PASSWORD=rabbitmqpassword
export BOOMVANG_RABBITMQ_USERNAME=rabbitmqusername
export BOOMVANG_DATADOG_API_KEY=xxxxx
export BOOMVANG_DATADOG_APP_KEY=yyyy
export BOOMVANG_DATADOG_API_ENDPOINT=https://0fc599c9-5f1b-48bb-82a3-5d3057bce11c.mock.pstmn.io
export BOOMVANG_COLLECTION_INTERVAL=5

go run main.go scheduler.go
