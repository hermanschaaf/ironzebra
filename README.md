IronZebra
=========

A blog engine written in Go, based on the Revel framework. It's nice and fast, and a fun experiment! You can see it live on the [IronZebra site](http://ironzebra.com).

Features
----------

This will be a growing feature list. If you feel like it's lacking something important to you, please feel free to fork and make pull requests!

 - Markdown posts
 - Uses MongoDB for storage 
 - A simple admin interface
 - Deploys effortlessly to Heroku

Important things that are still lacking: post tags, template caching, image uploads and storage, multiple authors. 

Running on Heroku
----------

To get this running on Heroku, I had to run one extra command not given in the Revel docs:

    heroku config:set BUILDPACK_URL=https://github.com/hermanschaaf/heroku-buildpack-go-revel.git

The original Revel URL didn't work, it ran into this problem:

    github.com/robfig/revel/cache/inmemory.go:5:2: cannot find package "github.com/robfig/go-cache" in any of:
    ...
    !     Push rejected, failed to compile Revel app

I fixed it in my branch of the Revel heroku repository, so updating the buildpack URL to my repo fixes the problem.

Once you have set this, you can follow the normal heroku deployment steps.

If it's the first time:

    heroku git:remote -a falling-wind-1624

otherwise just commit your changes to git and run

    git push heroku master

