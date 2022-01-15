# How to setup Letsencrypt by cert-manager with Kubernetes microk8s using default Ingress?
[stackoverflow](https://stackoverflow.com/questions/67430592/how-to-setup-letsencrypt-with-kubernetes-microk8s-using-default-ingress)
[cert-manager](https://cert-manager.io/docs/tutorials/acme/nginx-ingress/)


**Issuer** is specific to a single namespace in Kubernetes, and a \
**ClusterIssuer** is meant to be a cluster-wide definition for the same purpose.


Cert-manager uses two different custom resources, also known as CRD’s, to configure and control how it operates, as well as share status of its operation.
- letsencrypt-staging
- letsencrypt-prod *Production issuer has very strict rate limits!*

```sh
# if you in microk8s as user use
sudo microk8s kubectl
# if you in microk8s as root use
microk8s kubectl
# if you in local in remote microk8s use
kubectl
```


### 0 Creates deploys and services

kubectl apply -f 0webserver-depl-svc.yaml

Tip! write 
```sh
kubectl apply -f 0
```
and push tab

### 1 Creates simple Ingress and check domains to deploys on 80 ports

kubectl apply -f 1ingress-routes.yaml

### 2 Install jetstack/cert-manager (latest version respectevly)

kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.3.1/cert-manager.yaml

### 3 Add certificate issuer config

kubectl apply -f 3letsencrypt-clusterissuer.yaml

An Issuer is the definition for where cert-manager will get request TLS certificates. An Issuer is specific to a single namespace in Kubernetes, and a ClusterIssuer is meant to be a cluster-wide definition for the same purpose.

### 4 Run letsencrypt-staging

!!! strong with \
cert-manager.io/cluster-issuer: "letsencrypt-staging"

apply -f 45ingress-routes_upd.yaml


Check domains to deploys (it redirect on https port), it should work with unsecure caution.

### 5 Run letsencrypt-prod

checks due to about 5 min
```sh
kubectl get certificate
```
Your secretName should be Ready=True. If not check all steps before.

Warning! You can try a fiew times to run prod per day.

Change to \
cert-manager.io/cluster-issuer: "letsencrypt-prod"

kubectl apply -f 45ingress-routes_upd.yaml

### 6 Profit!

checks due to about 5 min
```sh
kubectl get certificate
```
Your secretName should be Ready=True. 

And check 
```sh
kubectl describe certificate name-from-get-sertificate

#right anwser:

  Issuer Ref:
    ...
    Name:       letsencrypt-prod
```
Check domains to deploys (it redirect on https port), it should work right (without unsecure caution).

If not pass check all steps before.

### 7 Cleaning (optional)

kubectl delete -f 45ingress-routes_upd.yaml
kubectl delete -f 0webserver-depl-svc.yaml