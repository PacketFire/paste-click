#!/bin/bash

kubectl --server="$GKE_APISERVER" \
        --namespace="$GKE_NAMESPACE" \
        --token="$GKE_BEARER_TOKEN" \
        --insecure-skip-tls-verify=true \
        set image deployment/paste-click paste-click=gcr.io/$GCP_PROJECT_ID/paste-click:$TRAVIS_COMMIT
        
kubectl --server="$GKE_APISERVER" \
        --namespace="$GKE_NAMESPACE" \
        --token="$GKE_BEARER_TOKEN" \
        --insecure-skip-tls-verify=true \
        set image deployment/nginx nginx=gcr.io/$GCP_PROJECT_ID/openresty-paste-click:$TRAVIS_COMMIT
