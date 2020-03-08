
```shell
gcloud functions deploy upload --source=upload --entry-point=UploadUrl --runtime=go113 --trigger-http --quiet --env-vars-file .env
gcloud alpha functions add-iam-policy-binding upload --member=allUsers --role=roles/cloudfunctions.invoker
```

