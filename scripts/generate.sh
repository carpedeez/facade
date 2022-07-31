#!/usr/bin/bash
(cd api; widdershins openapi-v1.yaml -o README.md --omitHeader -c > /dev/null)
(cd api; oapi-codegen --config oapi-config.yaml openapi-v1.yaml > ../facade/wrapper.go)