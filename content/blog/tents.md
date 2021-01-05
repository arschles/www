---
author: "Aaron Schlesinger"
date: 2018-10-25T17:36:25-07:00
title: 'Tents - Kubernetes for Actual Developers'

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
# meta_img = "/images/image.jpg"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

(This post was adapted from https://gist.github.com/arschles/7a91b621c9b259bb17c371bb4a2a8773)

Kubernetes is way too hard for "regular" app developers. Can you imagine if GitHub was just getting started and wanted to run on Kubernetes? They would have a good reason to do it:

- Run anywhere, with a standard API
- Be able to scale
- Be able to run lots of things at once for cheap (i.e. the git servers vs. the web frontend)
- Get service discovery for free


But in return they have to figure out all this stuff:

- Write Dockerfile for each app
- Monorepo or microservices?
- Figure out how to docker build
- Get an image repository on Docker hub
- `docker push`
- Figure out what you need & write YAML
- Figure out helm?
- WTF is an ingress controller?
- HTTPS?

And that's just to deploy some code. Testing, A/B deploys and everything else is "advanced" stuff. Kubernetes is an amazing tool for "ops people" - it has abstracted away entire datacenters.

But app developers are left out in the cold. Tents is a deployment tool for cloud native apps and app developers.

# What it's Supposed to Do

If you're building a cloud native app, use Tents to:

- Quickly set up a dev environment to run your app in
- Spin up a new instance of your app for each PR
- Run your app in production, and keep it running
- Roll out new versions of your app in production

# Tents Does a Lot

Tents has a lot of features and solves a big problem for app developers, but it doesn't invent 
any new concepts and it isn't novel software. Instead, it "builds on the shoulders of giants"

## Tents is an Extraction, not an Invention

Tents assembles modern technologies and principles and makes them accessible to you. During assembly, Tents
took a very opionated view on how each piece should be used, and thus, Tents has an opinionated _holistic_ view 
on how apps should be run. There's a good reason why Tents doesn't do certain things, but if you do need to do something
outside the scope of Tents, there's always an "escape hatch" that you can use to get your work done and still use Tents for the other things.

Here are some of the principles that Tents borrows from:

- Organizations, repositories and teams from Github
- Traffic splitting
- Cookie-based routing
- Shadow deployments
- A/B testing

Here are some of the technologies and principles that Tents uses:

- Kubernetes
- Traefik
- Istio and Envoy
- Helm
- gRPC
- Go
- JSON Web Tokens
- Draft

# User Workflow with Tents

Tents has a standard feature set that gives you a powerful but easy workflow:

1. Write your code
1. Build your code into a Docker image (Tents can help with this too!)
1. Run `tents deploy`

## For Development

Tents also helps you do local development:

First, `tents dev init` will set up your local development environment:

- A local [Minikube](https://github.com/kubernetes/minikube) cluster optimized for a fast development cycle
- A Dockerfile suitable for building their development app
- A [pare](https://github.com/arschles/pare) configuration file suitable for building the entire application for development
- A `DEVELOPING.md` file that describes the basic development cycle

## For Production

Production app rollouts are special because your app can't go down and your users should not notice any issues while the 
rollout is in progress. Tents supports features to help during rollouts

- Cookie-based routing
- Traffic splitting
- A/B testing

Additionally, Tents has a powerful authentication/authorization system to make sure nobody can accidentally deploy
to production

# Tents Environments

The simplest cloud native application usually runs in three environments:

- Development: the environment you push code to very frequently, to do quick local testing
- Staging: the environment you push code to, so you can test right before you promote to production
- Production: the big show!

Good news! Tents supports those environments and any other ones you define.

Here's what an environment is to Tents:

- Code: the application that's running
- Configuration: credentials, environment variables, etc...
- Resources: the CPUs that run the application, databases, networking, etc...uses)

Environments are defined in Tents by the organization/application model described below in the
[Application Layout, Authentication and Authorization](./application-layout,-authentication-and-authorization) section.

You'll interact with environments primarily via the `tents deploy` tool. For example:

`tents deploy myorg/app1`

# Organizations, Applications and Teams

Organizations, applications and teams define how Tents organizes all your applications.

- Organization: a "bucket" that Tents keeps the following in:
    - Applications
    - Common configuration
    - Authentication (authn) rules
    - Authorization (authz) rules
- Application: an application! Each application has configuration that is merged with its organization's configuration, but the app config overrides the org config
- Teams: groupings of users (see below) inside an organization. An application can assign one or more teams to one of three roles:
    - `read`: the team can only see deployment history and app status
    - `deploy`: the team can do everything from `read`, and they can also deploy new application versions
    - `admin`: the team can do everything from `deploy` and they can change configuration or delete the application. They can also add additional teams or users directly to the application

Organization names have to be "globally" unique (there can only be one organization 'myorg' per Tents installation), and team and application names have to be unique inside each organization.

_Note: organizations, applications and teams in Tents are a concept modeled after organizations and repositories in GitHub_

# Authentication and Authorization

Authentication (authn) and Authorization (authz) are important to Tents because they make sure that nobody 
accidentally deploys to production! Also, we want to make sure that environments are split up so that everyone has their own environment to work in.

Authn and authz work alongside organizations, applications and teams (from above) to provide a really powerful way to organize and secure your applications.

## Users

Each person who interacts with a Tents installation gets a username and a user account. Authentication and authorization are done via a standard username/password, and users are issued [JSON Web Token](https://jwt.io/)s after they login, so that they don't need to continuously login (both on the CLI and the Tents dashboard).

Every new user gets an organization automatically created for them, and they have permanent `admin` rights for all applications created in that organization (they can't delete the organization though). Since usernames are also organization names, they have to be globally unique just like organizations.

As mentioned in the Organizations, Applications and Teams" section above, each user can get put into one or more teams inside each organization. Users can also be assigned directly to an application.

# Use Case: Pull Requests

Most developers submit a pull request (PR) to solicit reviews (code, design, ...) from their team mates. CI systems generally
run on PRs to test whether the code passes the tests from `master` and the ones that were introduced in the new branch. This
setup helps give clues to the team whether the branch is ready to be merged. 

We need more clues, though.

The entire application should be built from a branch and deployed to the cluster. At that point, the application will
be running in an environment that is as close to production as possible. The team should run their tests 
(UI automation, compatibility, etc...) against that version before the PR is merged.

The application still needs to be slightly different from production, though:

1. It needs to be created and deleted separately from production
1. It is rolled out differently (most likely all at once)
1. It needs to have separate configuration from production (so it can't talk to the production DB, for example)
1. It has to have a different URL (if it's internet-facing components) from production
1. It needs to have different resiliency and redundancy (probably less)

The organizations and applications feature works very well to achieve this PR workflow, and more. When a PR is created, 
you would have your CI system deploy to an application called `pr-$PR_NUM` under an appropriate tents organization. 
For example, if you open up a PR number 1234, your CI might deploy to `myorg/pr-1234`. 

After the deploy is done, your team can do whatever testing is necessary on the application. And on each push to the PR, 
your CI system would redeploy the app.

Then, when the PR is closed, you would have your CI system delete the `myorg/pr-1234` application using `tents delete myorg/pr/1234`.

