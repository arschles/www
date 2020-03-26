# arschles.github.com

This repository has the source for http://arschles.com. 

## Just getting started?

First, check out this repository and initialize submodules:

```console
$ git@github.com:arschles/www.git
$ git submodule update --init
```

>The submodule in this repository contains the [Hugo theme](https://themes.gohugo.io/) that this site uses. The `git submodule` command will pull down the source for that theme, so you can run the site

Next, download [hugo](https://gohugo.io) version 0.68.3 or higher. See the [Hugo download page](https://gohugo.io/getting-started/installing) for installation instructions appropriate for your platform.

Finally, run the below command from the root of this repository:

```console
$ hugo server
```