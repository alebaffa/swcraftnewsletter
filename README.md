# Software Craftsmanship Newsletter 
[http://www.softwarecraftsmanship.news/](http://www.softwarecraftsmanship.news/)

## Run locally

- Change PORT in _server.go_

## Deploy on Heroku

The first deploy on Heroku requires the following steps:
- run _godep save ./..._ 
- change the PORT (in case it's still on local test)

## Current state
[![Build Status](https://travis-ci.org/alebaffa/swcraftnewsletter.svg?branch=master)](https://travis-ci.org/alebaffa/swcraftnewsletter)
Deploy on Heroku is done automatically via Travis CI.