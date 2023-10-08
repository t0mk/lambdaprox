#!/bin/bash

# example of deployment

RT=go121

while read REG; do
    gcloud functions deploy lambdaprox \
        --runtime $RT --trigger-http \
        --allow-unauthenticated  --region $REG \
        --memory=128Mi --gen2 --source=. \
        --entry-point LambdaProx
done < regs.gc
