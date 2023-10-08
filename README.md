# LambdaProx - Lightweight Serverless Proxy

[![Go Report Card](https://goreportcard.com/badge/github.com/t0mk/lambdaprox)](https://goreportcard.com/report/github.com/t0mk/lambdaprox)

**LambdaProx** is a minimalistic serverless proxy service designed to work with various cloud providers such as Google Cloud Functions, AWS Lambda, and Azure Functions. It allows you to fetch data from an external URL by making a GET request to it. The goal is to make lightweight and cost-effective solution to utilize multiple source IPs across different cloud regions.

It's a http service which will fetch URL that you pass in `url` GET param, like 

```bash
curl https://your-function-url.com?url=https://example.com
```

(don't forget to url encode the url...)

Only GCP works now.

If you deploy this to all regions, considering free tiers, inter-cloud round-roubin usage, it might be a cheap datacenter proxy service. The cost-effectiveness depends on your use case (how often you call, how much you download). Consult pricing of the functions service:
- https://cloud.google.com/functions/pricing

This approach works long as you only fetch single urls. In other words, it won't work as a proxy for all HTTPS(TLS) connections in a browser.

## Features

- **Proxy Service**: Fetch data from an external URL via a serverless function.
- **Multi-Cloud Support**: Designed to work with Google Cloud Functions, AWS Lambda, Azure Functions, and more.
- **Thin and Cost-Efficient**: Minimal code and resource usage for cost-effective operation.

## How to Use

### Deployment

You can deploy **LambdaProx** to various cloud providers, including Google Cloud, AWS, and Azure, by adapting the deployment script to the respective cloud's deployment mechanism. Below are examples of how you can deploy this service to different cloud providers:

#### Google Cloud Functions (GCP)

Use the provided deployment script for Google Cloud Functions as a reference. Customize it to match your project and preferences.

```bash
#!/bin/bash

# Example deployment script for Google Cloud Functions

RT=go121

while read REG; do
    gcloud functions deploy lambdaprox \
        --runtime $RT --trigger-http \
        --allow-unauthenticated  --region $REG \
        --memory=128Mi --gen2 --source=. \
        --entry-point LambdaProx
done < regs.gc
```

#### AWS Lambda (AWS)

For AWS Lambda, you can use AWS CLI or Serverless Framework to deploy the function. Make sure to configure the necessary IAM roles and permissions.

#### Azure Functions (Azure)

For Azure Functions, use the Azure CLI or Azure Portal to create and deploy the function. Ensure you have the required Azure resources and permissions set up.

### Accessing the Proxy

Once deployed, you can access the proxy service using the following `curl` command:

```bash
curl https://your-function-url.com?url=https://example.com
```

Replace `your-function-url.com` with the URL of your deployed function and change the `url` parameter to the target URL you want to fetch data from.

## Dependencies

- [Functions Framework for Go](https://pkg.go.dev/github.com/GoogleCloudPlatform/functions-framework-go/functions): A framework for writing portable functions for various serverless platforms.

## Contributing

If you find any issues, have questions, or want to contribute to this project, please feel free to create an issue or submit a pull request on the [GitHub repository](https://github.com/t0mk/lambdaprox).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

**Note:** Please ensure you have the necessary permissions and credentials set up in your chosen cloud provider's environment before deploying the function.

Thank you for using **LambdaProx**!
