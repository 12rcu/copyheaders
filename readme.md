# Copy header middleware

Copies the value of a header and pastes it into another header.
This can solve for example an issue with the oauth2 proxy as this project will provide the id token in the 
Authorization header which some applications can't handle and need the X-Auth-Request-Access-Token header value.

## Installation

Add this to the values.yaml of traefik

```yaml
experimental:
  plugins:
    copyheaders:
      moduleName: github.com/12rcu/copyheaders
      version: v0.1.4
```

## Usage

Create the middleware:

```yaml
apiVersion: traefik.us/v1alpha1
kind: Middleware
metadata:
  name: rewrite-auth-header
spec:
  plugin:
    copyheaders:
      headers:
        - from: "X-Auth-Request-Access-Token"
          to: "Authorization"
          prefix: "Bearer "
```

Use it in an ingress file via annotations:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    traefik.ingress.kubernetes.io/router.middlewares: |
      namespace-rewrite-auth-header@kubernetescrd
```