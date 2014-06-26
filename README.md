# Writing a new post:

1. add a file to the _posts directory called `$year-$month-$day-$title.markdown`
2. copy the header from an existing post

```bash
vagrant up
vagrant ssh
cd /vagrant
bundle exec jekyll -w serve
```

When done, go to [localhost:4000](http://localhost:4000) to see the new post.

When you're done making changes, commit and push. Your new post should be
live on the list at [arschles.github.io](http://arschles.github.io).

Or specifically, at `arschles.github.io/$year/$month/$day/$title.html`.
