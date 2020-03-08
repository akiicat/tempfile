

```shell
gcloud functions deploy home --source=home --entry-point=HomePage --runtime=go113 --trigger-http --quiet --env-vars-file .env
gcloud alpha functions add-iam-policy-binding home --member=allUsers --role=roles/cloudfunctions.invoker
```
