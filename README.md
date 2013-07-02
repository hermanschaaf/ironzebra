IronZebra
----------

Written in Go, based on the Revel framework. It's nice and fast, and a fun experiment!

To get this running on Heroku, I had to run one extra command not given in the Revel docs:

    heroku config:set BUILDPACK_URL=https://github.com/robfig/heroku-buildpack-go-revel.git