# AdobeSign Integration APIs

## Digital signature service via AdobeSign

## Step 1: Install gcloud SDK using [Google Cloud Quick Starts](https://cloud.google.com/sdk/docs/quickstarts)

## Step 2: clone "octupi/aws-ses" repo

## Step 3: Make sure you are inside "octupi/aws-ses" folder

## Step 4: Run "gcloud init"

## Step 5: Run following commands:

#Platform:

```
gcloud config set run/platform gke
```

#Cluster:

```
gcloud config set run/cluster cloudrun && gcloud config set run/cluster_location us-central1-a
```

#Zone:

```
gcloud config set compute/zone us-west1-a
```

#Build:

```
gcloud builds submit --tag gcr.io/octupi/adobesign-integration
```

#Deploy:

```
gcloud beta run deploy --project octupi --image gcr.io/octupi/adobesign-integration:latest --platform gke --connectivity external
```
