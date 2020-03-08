
```shell
gcloud functions deploy download --source=download --entry-point=DownloadUrl --runtime=go113 --trigger-http --quiet --env-vars-file .env
gcloud alpha functions add-iam-policy-binding download --member=allUsers --role=roles/cloudfunctions.invoker
```

