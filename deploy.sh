#!/bin/bash
gcloud functions deploy ProxyLineNoti --runtime go116 \
--trigger-http \
--allow-unauthenticated \
--service-account=$GCLOUD_EMAIL \
--region=asia-east2 \
--env-vars-file config/config.yaml \
--memory 128 \
--timeout 540