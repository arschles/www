---
title: "I Can Haz Services?"
revealOptions:
    transition: 'fade'
---

# I Can Haz Services?

### Kubeapps, OSB, Service Catalog, Oh my!

---

## In The Beginning...

There was Service Catalog and OSB

And there was Monocular

---

## And Then

Bitnami and Microsoft got together

![rad!](images/2018-02-21-helm-summit/rad.png)

---

## The Result!

---

![the result!](images/2018-02-21-helm-summit/the-result.png)

---

![monocular screenshot](images/2018-02-21-helm-summit/monocular.png)

---

## The Mighty Service Catalog
### An Illustraed Guide

---

## The Why

![homer](images/2018-02-21-helm-summit/homer.jpg)

---

## Terminology

- OSB Broker
- Service Catalog
- Class
- Provision
- Bind

---

## Kubeapps

<img src="images/2018-02-21-helm-summit/kubeapps-logo.jpg" width="100" />

The Easiest Way to Deploy Applications to Your Kubernetes Cluster

---

##### Web-based Kubernetes App Community

Browse, Rate and Review the Kubernetes Community Charts
[hub.kubeapps.com](https://hub.kubeapps.com)

<img src="images/2018-02-21-helm-summit/kubeapps-hub.jpg" height="450" />

---

Kubeapps Combines our Kubenetes application packaging expertise

<img src="images/2018-02-21-helm-summit/kubeapps-hub-and-spoke.png" height="400" />

--- 

## No More Empty Clusters!

- Complete application delivery environment
- App-focused dashboard UI
- Simple browse and click deployment of apps
- Deploy Helm Charts and Kubeless Functions
- ... and now provision Service Instances and Bindings!

---

![kubeapps up screenshot](images/2018-02-21-helm-summit/kubeapps-up-screenshot.jpg)

---

![kubeapps screenshot](images/2018-02-21-helm-summit/kubeapps-screenshot.jpg)

---

## Let's See Some **Wordpress!**

---

## A Fountain of Ideas

- Binding to apps is less sad
- `kubeapps up --cloud=$YOUR_CLOUD`
    - Installs service-catalog + `$YOUR_BROKER`
- Standard way of consuming secrets in upstream charts
    - PoC: [github.com/azure/helm-charts](https://github.com/azure/helm-charts)

---

## Thanks To

- Adnan Abdulhussein (@prydonius)
- Ara Pulido (@arapulido)
- Rita Zhang (@ritazzhang)
- Sertac Ozercan (@sozercan)
- Evan Louie (@evanlouie)
- Angel M Miguel (@_angelmm)
- Angus Lees
- Miguel Martinez (@migmartri)
- *Plenty More Folks*

[github.com/kubeapps/kubeapps/graphs/contributors](https://github.com/kubeapps/kubeapps/graphs/contributors)

