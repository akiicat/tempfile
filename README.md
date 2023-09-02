
# Short Url on Google Cloud Function

```shell
pip install google-cloud-firestore
```

## GCP

Init GCP

```sh
export GOOGLE_CLOUD_PROJECT=<project_id>
export BUCKET=<bucket_name>

gcloud auth login
gcloud config set project PROJECT_ID

# enable service (for new project and only first time)
gcloud services enable run.googleapis.com

# setup key for debugging
export GOOGLE_APPLICATION_CREDENTIALS="path/to/xxxxxx-xxxxxx-xxxxxxxxxxxx.json"
```

Create firestore

```shell
gcloud alpha firestore databases create \
--database=note \
--location=asia-east1 \
--type=firestore-native
```

## Setup Storage cors configuration

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
gsutil cors set cors.json gs://${BUCKET}
gsutil cors get gs://${BUCKET}
```

## Routes

Edit **firebase.json** 

```shell
firebase deploy --only hosting
```

## Debug

```shell
python -m venv venv
source venv/bin/activate
pip install functions-framework
pip install -r function/note/requirements.txt
```

```shell
./debug.sh
```

## Deploy

```shell
./deploy.sh
```