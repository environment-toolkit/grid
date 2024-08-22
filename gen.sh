#!/bin/bash

oapi-codegen -config ./pkg/identity/oapi-config.yaml http://localhost:5200/docs/openapi.json
oapi-codegen -config ./pkg/bond/oapi-config.yaml http://localhost:5600/docs/openapi.json
