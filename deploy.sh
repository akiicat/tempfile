#!/bin/bash

: ${GOOGLE_CLOUD_PROJECT?: GOOGLE_CLOUD_PROJECT project id is not provided}
: ${BUCKET?: BUCKET is not provided}

pushd function/tempfile

gcloud functions deploy tempfile \
--gen2 \
--runtime=python311 \
--region=asia-east1 \
--source=. \
--entry-point=entrypoint \
--trigger-http \
--allow-unauthenticated \
--set-env-vars GCP_PROJECT=$GOOGLE_CLOUD_PROJECT
--set-env-vars GCP_PROJECT=$BUCKET

popd

