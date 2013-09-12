# GANGLIA-CACHER

[![Build Status](https://secure.travis-ci.org/jbuchbinder/ganglia-cacher.png)](http://travis-ci.org/jbuchbinder/ganglia-cacher)

Caching mechanism for [ganglia-web](https://github.com/ganglia/ganglia-web) cache information, independent of the PHP mechanism.

## BUILDING

```
go get github.com/reiver/go-php
go get github.com/bradfitz/gomemcache/memcache
go build
```

## RUNNING

Ideally, this should be cron'd instead of the PHP-based version which exists
within the [ganglia-web](https://github.com/ganglia/ganglia-web) project.

