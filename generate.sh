#!/bin/sh
#
# The whole Go API Client for Localizely is generated via this script
# Usage example: /bin/sh ./generate.sh 1.0.0

version=$1

if [ "$version" = "" ]; then
  echo "Version argument is not provided"
  exit
fi

openapi-generator-cli generate \
  -g go \
  -i https://api.localizely.com/api-docs \
  --git-user-id localizely \
  --git-repo-id localizely-client-go \
  --additional-properties=packageName=localizely,packageVersion=$version \
  --skip-validate-spec

echo ""
echo "Running go mod tidy..."

go mod tidy