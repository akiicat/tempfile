
![](https://github.com/akiicat/tempfile/workflows/Deploy%20to%20Google%20Cloud/badge.svg)

# Short Url on Google Cloud Function

## Create Google Cloud Storage

```sh
export ProjectID=<project_id>
export BucketName=<bucket_name>

# Set Project ID
gcloud config set project ${ProjectID}

# Create Bucket
gsutil mb -b on -c Standard -l asia gs://${BucketName}
```

## enable lifecycle management

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

## setup cors configuration

Edit the **gcs_cors.json** to update the origin url.

```json
# gcs_cors.json
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
gsutil cors set gcs_cors.json gs://${BucketName}
gsutil cors get gs://${BucketName}
```

## Signed Url Key

```shell
# Create Service Account
gcloud iam service-accounts create "signed-url" --display-name "signed-url"

# Create Key
gcloud iam service-accounts keys create signed-url-key.json --iam-account signed-url@${ProjectID}.iam.gserviceaccount.com
```

## Install Require Libaray

```shell
pip install -r requirements.txt -t lib
```

## Deploy

```shell
gcloud app deploy app.yaml --quiet --stop-previous-version
```


