#!/bin/bash

: ${GOOGLE_CLOUD_PROJECT?: GOOGLE_CLOUD_PROJECT project id is not provided}
: ${BUCKET?: BUCKET is not provided}

cd function/tempfile
GCP_PROJECT="$GOOGLE_CLOUD_PROJECT" BUCKET="$BUCKET" functions-framework --target entrypoint --debug
