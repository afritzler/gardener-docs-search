# gardener-docs-search

## Deploying Cloud Function

```bash
gcloud functions deploy gardener-search --entry-point Search --runtime go111 --trigger-http --memory 128
```