

```shell
gcloud functions deploy home --source=home --entry-point=HomePage --runtime=go113 --trigger-http --quiet --env-vars-file .env
```

https://cloud.google.com/sdk/gcloud/reference/alpha/functions/add-iam-policy-binding
https://stackoverflow.com/a/57193899

```shell
gcloud alpha functions add-iam-policy-binding home --member=allUsers --role=roles/cloudfunctions.invoker
```
