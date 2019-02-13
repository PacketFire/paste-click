#!/bin/bash

kubectl --server="$GKE_APISERVER" \
        --token="$GKE_BEARER_TOKEN" \
        --insecure-skip-tls-verify=false \
        set image deployment/paste-click nginx=gcr.io/$GCP_PROJECT_ID/paste-click:latest
        
kubectl --server="$GKE_APISERVER" \
        --token="$GKE_BEARER_TOKEN" \
        --insecure-skip-tls-verify=false \
        set image deployment/paste-click nginx=gcr.io/$GCP_PROJECT_ID/paste-click

kubectl --server="$GKE_APISERVER" \
        --token="$GKE_BEARER_TOKEN" \
        --insecure-skip-tls-verify=false \
        set image deployment/nginx nginx=gcr.io/$GCP_PROJECT_ID/openresty-paste-click:latest
        
kubectl --server="$GKE_APISERVER" \
        --token="$GKE_BEARER_TOKEN" \
        --insecure-skip-tls-verify=false \
        set image deployment/nginx nginx=gcr.io/$GCP_PROJECT_ID/openresty-paste-click
