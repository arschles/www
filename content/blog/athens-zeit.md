---
author: "Aaron Schlesinger"
date: 2018-10-24T12:31:07-07:00
title: 'Athens Running on Zeit'

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
meta_img: "/images/athens-zeit/yessss.jpg"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

I've been a long time admirer-er of [Zeit](https://zeit.co). Besides the technical part (which is really good), their platform gives people really easy, understandable and (most importantly) trustworthy tools to get a real app into production. I've wanted something like this to exist for Kubernetes for a while (I even [wrote](https://gist.github.com/arschles/7a91b621c9b259bb17c371bb4a2a8773) about that, but I digress.)

The platform was for just Node applications for a while, but containers are a thing y'all, and Zeit lets you deploy them to its platform now (they announced it [mid August](https://zeit.co/blog/serverless-docker).) It does all the same stuff as it does for Node, but with Docker:

1. Write a `Dockerfile`
1. Run `now`
1. Profit

There's a ton going on there in step 2, but the basics:

- Push the `Dockerfile` and context up to Zeit
- Zeit does the `docker build` on their infrastructure
- Zeit pushes the image to their registry (I assume they run an internal registry but no idea really)
- Zeit pulls the image into their serving infrastructure, configures it, does routing, ~~checks if you paid and gives you all the amazing hardware if you have and otherwise runs you on a RaspberryPi~~ and runs your image on their servers

I mean, I've blogged about [ACR Builds](https://arschles.com/blog/az-acr-build-or-docker-build/) before, but this is next level. One command gives you everything.

<img src="/images/athens-zeit/yessss.jpg" />

# Athens-ing it Up

There's this [rad project](https://docs.gomods.io) I know about. Something about Go and dependencies IDK but it has a `Dockerfile` and you can do cool stuff with it. So why not see if Zeit can handle this beast??

I had to tweak the Dockerfile a little bit because the CLI (`now`) doesn't let you specify custom `Dockerfile` locations (I don't think???) but that was pretty much it (check out the [PR](https://github.com/arschles/athens/pull/2).) Now I can do this:

```console
➜  athens git:(zeit-now) now --docker
> Deploying ~/github/athens under arschles
> Your deployment's code and logs will be publicly accessible because you are subscribed to the OSS plan.

> NOTE: You can use `now --public` or upgrade your plan (https://zeit.co/account/plan) to skip this prompt
> https://athens-ukiespgiyz.now.sh [in clipboard] (sfo1) [5s]
> Building…
> Sending build context to Docker daemon  37.83MB
> Step 1/16 : FROM golang:1.11-alpine AS builder
>  ---> 95ec94706ff6
> Step 2/16 : RUN mkdir /proj
>  ---> Using cache
>  ---> 7690d95ac3b4
> Step 3/16 : WORKDIR /proj
>  ---> Using cache
>  ---> 89a1e2d291ec
> Step 4/16 : COPY . .
>  ---> 6fdf77ac18ec
> Step 5/16 : ENV GO111MODULE=on
>  ---> Running in 61601842b526
> Removing intermediate container 61601842b526
>  ---> 75be7b20b225
> Step 6/16 : ENV GOPROXY=https://microsoftgoproxy.azurewebsites.net
>  ---> Running in 8e92f2e31291
> Removing intermediate container 8e92f2e31291
>  ---> 8f7e12c7fef4
> Step 7/16 : RUN GO111MODULE=on CGO_ENABLED=0 go build -o /bin/athens-proxy ./cmd/proxy
>  ---> Running in b25589823d39
> go: finding github.com/fatih/color v1.7.0
> go: finding github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
> go: finding github.com/minio/minio-go v6.0.5+incompatible
> go: finding github.com/gobuffalo/mw-csrf v0.0.0-20180802151833-446ff26e108b
> go: finding github.com/tinylib/msgp v1.0.2
[SNIP]
> go: downloading github.com/markbates/goth v1.46.0
> go: downloading github.com/rs/cors v1.5.0
> go: downloading github.com/gobuffalo/buffalo v0.13.1
> go: downloading github.com/unrolled/secure v0.0.0-20181005190816-ff9db2ff917f
> go: downloading github.com/mitchellh/go-homedir v1.0.0
[SNIP]
> Removing intermediate container b25589823d39
>  ---> d61630e38cd9
> Step 8/16 : FROM alpine
>  ---> 196d12cf6ab1
> Step 9/16 : ENV GO111MODULE=on
>  ---> Using cache
>  ---> cd147951c9a7
> Step 10/16 : COPY --from=builder /bin/athens-proxy /bin/athens-proxy
>  ---> Using cache
>  ---> a462fc7bd78b
> Step 11/16 : COPY --from=builder /proj/config.dev.toml /config/config.toml
>  ---> 3c99d68507fe
> Step 12/16 : COPY --from=builder /usr/local/go/bin/go /bin/go
>  ---> bcac2f09281a
> Step 13/16 : RUN apk update &&      apk add --no-cache bzr git mercurial openssh-client subversion procps fossil &&     mkdir -p /usr/local/go
>  ---> Running in 49fd4fa48969
> fetch http://dl-cdn.alpinelinux.org/alpine/v3.8/main/x86_64/APKINDEX.tar.gz
> fetch http://dl-cdn.alpinelinux.org/alpine/v3.8/community/x86_64/APKINDEX.tar.gz
> v3.8.1-38-g898a0bb28a [http://dl-cdn.alpinelinux.org/alpine/v3.8/main]
> v3.8.1-35-ga062ffc9e8 [http://dl-cdn.alpinelinux.org/alpine/v3.8/community]
> OK: 9539 distinct packages available
> fetch http://dl-cdn.alpinelinux.org/alpine/v3.8/main/x86_64/APKINDEX.tar.gz
> fetch http://dl-cdn.alpinelinux.org/alpine/v3.8/community/x86_64/APKINDEX.tar.gz
> (1/33) Installing libbz2 (1.0.6-r6)
> (2/33) Installing expat (2.2.5-r0)
> (3/33) Installing libffi (3.2.1-r4)
> (4/33) Installing gdbm (1.13-r1)
> (5/33) Installing ncurses-terminfo-base (6.1_p20180818-r1)
[SNIP]
> Executing subversion-1.10.0-r0.pre-install
> Executing busybox-1.28.4-r1.trigger
> Executing ca-certificates-20171114-r3.trigger
> OK: 123 MiB in 46 packages
> Removing intermediate container 49fd4fa48969
>  ---> e8a15877e355
> Step 14/16 : ENV GO_ENV=production
>  ---> Running in ac91ac08096a
> Removing intermediate container ac91ac08096a
>  ---> 096f14de684e
> Step 15/16 : EXPOSE 3000
>  ---> Running in eb60ce775db6
> Removing intermediate container eb60ce775db6
>  ---> 0075a1cdb3c5
> Step 16/16 : CMD ["athens-proxy", "-config_file=/config/config.toml"]
>  ---> Running in 33e543286967
> Removing intermediate container 33e543286967
>  ---> 211b19e9d4a4
> Successfully built 211b19e9d4a4
> Successfully tagged build:RWio8EYjdNbuK5HXd31rKHAp_1540411816
> ▲ Assembling image
> ▲ Storing image (64.0M)
> Build completed
> Verifying instantiation in sfo1
> [0] buffalo: Unless you set SESSION_SECRET env variable, your session storage is not protected!
> [0] time="2018-10-24T20:12:29Z" level=info msg="Exporter not specified. Traces won't be exported"
> [0] buffalo: Starting application at :3000
> ✔ Scaled 1 instance in sfo1 [14s]
> Success! Deployment ready
```

That's using a [multi stage build](https://docs.docker.com/develop/develop-images/multistage-build/) to build Athens inside a `Dockerfile`, on Zeit's infrastructure. And I inception-ed it because I used a [hosted Athens module proxy](https://github.com/gomods/athens/issues/772) to build Athens itself inside the build, on Zeit's infrastructure.

<img src="/images/athens-zeit/diagram.png" />

Here's what it looks like after it's deployed:

```
➜  athens git:(zeit-now) now list
> 5 total deployments found under arschles [314ms]
> To list more deployments for an app run `now ls [app]`

  app       url                         inst #    type      state    age
  athens    athens-ukiespgiyz.now.sh         -    DOCKER    READY    4m
```

There's a UI for all this on the zeit.co site too. They even stream build logs to the site too #AMAZE

# Trying it Out

So basically I ran a few commands and got a ton of output on the CLI. Machines are doing my bidding; success!

Now that it's running, I wanted to see how it performed. I started small with [bld](https://github.com/arschles/bld) ([blog post](https://arschles.com/blog/az-acr-build-or-docker-build/) for background on that tool if you're interested):

```console
➜  bld git:(master) export GOPROXY=https://athens-ukiespgiyz.now.sh
➜  bld git:(master) sudo rm -r $GOPATH/pkg/mod
➜  bld git:(master) time go build
go: finding github.com/magefile/mage v1.4.0
go: finding github.com/tdewolff/parse v2.3.3+incompatible
go: finding github.com/spf13/cobra v0.0.3
go: finding github.com/mitchellh/go-homedir v1.0.0
go: finding github.com/mattn/go-isatty v0.0.3
go: finding github.com/spf13/viper v1.2.1
go: finding github.com/dlclark/regexp2 v1.1.6
go: finding github.com/tdewolff/minify v2.3.5+incompatible
go: finding github.com/spf13/pflag v1.0.3
go: finding github.com/magiconair/properties v1.8.0
go: finding gopkg.in/yaml.v2 v2.2.1
go: finding github.com/fsnotify/fsnotify v1.4.7
go: finding golang.org/x/text v0.3.0
go: finding golang.org/x/sys v0.0.0-20180906133057-8cf3aee42992
go: finding github.com/pelletier/go-toml v1.2.0
go: finding github.com/spf13/jwalterweatherman v1.0.0
go: finding github.com/mitchellh/mapstructure v1.0.0
go: finding github.com/spf13/pflag v1.0.2
go: finding gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405
go: finding github.com/spf13/afero v1.1.2
go: finding github.com/spf13/cast v1.2.0
go: finding github.com/hashicorp/hcl v1.0.0
go: finding github.com/davecgh/go-spew v1.1.1
go: downloading github.com/spf13/viper v1.2.1
go: downloading github.com/spf13/cobra v0.0.3
go: downloading github.com/mitchellh/go-homedir v1.0.0
go: downloading github.com/magefile/mage v1.4.0
go: downloading github.com/mitchellh/mapstructure v1.0.0
go: downloading github.com/hashicorp/hcl v1.0.0
go: downloading github.com/pelletier/go-toml v1.2.0
go: downloading gopkg.in/yaml.v2 v2.2.1
go: downloading github.com/spf13/jwalterweatherman v1.0.0
go: downloading github.com/magiconair/properties v1.8.0
go: downloading github.com/fsnotify/fsnotify v1.4.7
go: downloading github.com/spf13/pflag v1.0.3
go: downloading github.com/spf13/cast v1.2.0
go: downloading github.com/spf13/afero v1.1.2
go: downloading golang.org/x/sys v0.0.0-20180906133057-8cf3aee42992
go: downloading golang.org/x/text v0.3.0
go build  4.97s user 2.06s system 7% cpu 1:32.16 total
```

So, that kinda took a while. Building Athens itself took _forever_ as well (about 6:30), but there are a few reasons for that, and ways to get big ole speedups from experience hosting Athens:

1. Run this with more than one instance
1. Hook this up to external storage
    - You'd need to do this for #1 to work anyway
    - This is using local disk and I have no idea what or how fast that is on Zeit
1. Pre-seed the Athens cache with things
    - Everything I showed was from a cold local (`$GOPATH/pkg/mod`) and hosted (on the Zeit Athens server) cache

All in all, this is a pretty rad platform. I give it pretty rad. Not mega rad - that's reserved for [clippy](https://en.wikipedia.org/wiki/Office_Assistant) level things, but wayyy higher than semi-rad.

<img src="/images/athens-zeit/clippy.jpg" height="500px" />
<div style="text-align:center;font-size:70%;">Zeit: not quite clippy level, but pretty darn close!</div>

Keep on rockin' everybody!
