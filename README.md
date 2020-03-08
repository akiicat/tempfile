
![](https://github.com/akiicat/tempfile/workflows/Deploy%20to%20Google%20Cloud/badge.svg)

# Short Url on Google Cloud Function

- [Setup Signed Url Service](sign)


## Create Google Cloud Storage

```sh
export ProjectID=<project_id>
export BucketName=<bucket_name>

# Set Project ID
gcloud config set project ${ProjectID}

# Create Bucket
gsutil mb -b on -c Standard -p ${ProjectID} -l asia gs://${BucketName}
```

### Enable lifecycle management

Lifecycle default 2 day

```json
# lifecycle.json
{
  "lifecycle": {
    "rule": [
      {
        "action": {"type": "Delete"},
        "condition": {
          "age": 2
        }
      }
    ]
  }
}
```

```shell
gsutil lifecycle set lifecycle.json gs://${BucketName}
gsutil lifecycle get gs://${BucketName}
```

### Setup Storage cors configuration

Edit the **cors.json** to update the origin url.

```json
# cors.json
[
    {
      "origin": ["https://<example.com>", "http://localhost:3000"],
      "responseHeader": ["Content-Type"],
      "method": ["GET", "HEAD", "DELETE", "PUT", "POST"],
      "maxAgeSeconds": 3600
    }
]
```

```shell
gsutil cors set cors.json gs://${BucketName}
gsutil cors get gs://${BucketName}
```

## Setup Datastore Index

```shell
gcloud datastore create-indexes db/index.yaml
```

## Routes

Edit **firebase.json** 

```shell
firebase deploy --only hosting
```

## Signed Url Key

```shell
# Create Service Account
gcloud iam service-accounts create "signed-url" --display-name "signed-url"

# Grant Service Account with storage object admin
gcloud projects add-iam-policy-binding ${ProjectID} \
  --member serviceAccount:signed-url@${ProjectID}.iam.gserviceaccount.com \
  --role roles/storage.objectAdmin

# Create Key
gcloud iam service-accounts keys create signed-url-key.json --iam-account signed-url@${ProjectID}.iam.gserviceaccount.com
```

## Deploy

```shell
gcloud functions deploy home --source=home --entry-point=HomePage --runtime=go113 --trigger-http --quiet --env-vars-file .env
gcloud alpha functions add-iam-policy-binding home --member=allUsers --role=roles/cloudfunctions.invoker

gcloud functions deploy upload --source=upload --entry-point=UploadUrl --runtime=go113 --trigger-http --quiet --env-vars-file .env
gcloud alpha functions add-iam-policy-binding upload --member=allUsers --role=roles/cloudfunctions.invoker

gcloud functions deploy download --source=download --entry-point=DownloadUrl --runtime=go113 --trigger-http --quiet --env-vars-file .env
gcloud alpha functions add-iam-policy-binding download --member=allUsers --role=roles/cloudfunctions.invoker
```

## Deploy Key (Optional)

```shell
# Create Service Account
gcloud iam service-accounts create "deploy-app-engine" --display-name "deploy-app-engine"

# Grant Service Account with appengine admin, storage admin, and datastore index admin roles
gcloud projects add-iam-policy-binding ${ProjectID} \
  --member serviceAccount:deploy-app-engine@${ProjectID}.iam.gserviceaccount.com \
  --role roles/appengine.appAdmin

gcloud projects add-iam-policy-binding ${ProjectID} \
  --member serviceAccount:deploy-app-engine@${ProjectID}.iam.gserviceaccount.com \
  --role roles/storage.admin

gcloud projects add-iam-policy-binding ${ProjectID} \
  --member serviceAccount:deploy-app-engine@${ProjectID}.iam.gserviceaccount.com \
  --role roles/datastore.indexAdmin

# Create Key
gcloud iam service-accounts keys create deploy-key.json --iam-account deploy-app-engine@${ProjectID}.iam.gserviceaccount.com
```

