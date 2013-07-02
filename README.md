IronZebra
----------

Written in Go, based on the Revel framework. It's nice and fast, and a fun experiment!

To get this running on Heroku, I had to run one extra command not given in the Revel docs:

    heroku config:set BUILDPACK_URL=https://github.com/robfig/heroku-buildpack-go-revel.git

The URL didn't work though, it ran into this problem:

 >  github.com/robfig/revel/cache/inmemory.go:5:2: cannot find package "github.com/robfig/go-cache" in any of:
 >  ...
 >  !     Push rejected, failed to compile Revel app


However, I fixed it in my own repository, so updating the buildpack URL to my repo fixes the problem:

    heroku config:set BUILDPACK_URL=https://github.com/hermanschaaf/heroku-buildpack-go-revel.git