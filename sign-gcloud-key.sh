#!/bin/bash

ProjectID=

# deploy google cloud function
# Create Service Account
gcloud iam service-accounts create "deploy-func" --display-name "deploy-func"

# Grant Service Account with appengine admin, storage admin, and datastore index admin roles
gcloud projects add-iam-policy-binding ${ProjectID} \
  --member serviceAccount:deploy-func@${ProjectID}.iam.gserviceaccount.com \
  --role roles/iam.serviceAccountUser
gcloud projects add-iam-policy-binding ${ProjectID} \
  --member serviceAccount:deploy-func@${ProjectID}.iam.gserviceaccount.com \
  --role roles/cloudfunctions.admin
gcloud projects add-iam-policy-binding ${ProjectID} \
  --member serviceAccount:deploy-func@${ProjectID}.iam.gserviceaccount.com \
  --role roles/datastore.indexAdmin
gcloud projects add-iam-policy-binding ${ProjectID} \
  --member serviceAccount:deploy-func@${ProjectID}.iam.gserviceaccount.com \
  --role roles/storage.admin

# Create Key 
gcloud iam service-accounts keys create deploy-key.json --iam-account deploy-func@${ProjectID}.iam.gserviceaccount.com

