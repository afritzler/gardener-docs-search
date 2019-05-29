# gardener-docs-search

This project allows to search the Gardener documentation and integrate it into a bot framework. The overall flow looks like the following

![architecture](images/architecture.png)

## Run locally using Docker

To startup a local instance of the service run

```shell
docker run -p 8080:8080 afritzler/gardener-docs-search
```

The service should response with a `200`OK response on http://localhost:8080.

You can run an example search query via

```shell
curl -X POST -H "Content-Type: application/json" -d @example/request.json http://localhost:8080/search
```

## Deploying Cloud Function

The search backend can be deployed to GCP as a Cloud Function via

```shell
gcloud functions deploy gardener-search --entry-point Search --runtime go111 --trigger-http --memory 128
```