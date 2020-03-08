
```shell
gcloud functions deploy download --source=download --entry-point=DownloadUrl --runtime=go113 --trigger-http --quiet --env-vars-file .env
gcloud alpha functions add-iam-policy-binding download --member=allUsers --role=roles/cloudfunctions.invoker
```

PROJECT_NAME=upload-tempfile-257400
curl -v -X OPTIONS -H "Host: storage.googleapis.com" -H "Access-Control-Request-Method: PUT"  -H "Origin: https://gcs.somedomain.com:8081" "https://storage.googleapis.com/upload-tempfile-257400/cors.txt"
curl -v -X POST -H "Host: storage.googleapis.com" -H "Access-Control-Request-Method: POST"  -H "Origin: https://gcs.somedomain.com:8081" -F 'files=@/Users/akiicat/Desktop/UartService.java' "https://storage.googleapis.com/upload-tempfile-257400/cors.txt"

