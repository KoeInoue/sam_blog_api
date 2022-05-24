.PHONY: build

build:
	sam build

start:
	sam local start-api --env-vars test/env.json

create-table:
	aws dynamodb create-table --cli-input-json file://test/post_table.json --endpoint-url http://0.0.0.0:8000
add-data:
	aws dynamodb batch-write-item --request-items file://test/post_table_data.json --endpoint-url http://0.0.0.0:8000