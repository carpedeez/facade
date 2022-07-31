#!/usr/bin/bash
(cd api; widdershins openapi-v1.yaml -o README.md --omitHeader -c > /dev/null)
(cd api; goapi-gen --config goapi-config.yaml openapi-v1.yaml > ../facade/wrapper.gen.go)