# cmdtwitter

command line twitter client for go for Titter API v1.1

## library install

```
go get github.com/ChimeraCoder/anaconda
```

## Build

```
go build
```

## How to use

```
# [Tweet]
cmdtwitter \
    -key=(API Key) -sec=(API Secret) \
    -acc_tok=(Access token) -acc_sec=(Access token secret) \
    -cmd=tweet text=(Text)
```

When you make shell script:

```
#!/bin/sh

key=xxx
sec=xxx
acc_tok=xxx
acc_sec=xxx

# tweet
go run cmdtwitter.go \
  -key=$key -sec=$sec \
  -acc_tok=$acc_tok -acc_sec=$acc_sec \
  -cmd=tweet -text=Test

# get usertimeline
go run cmdtwitter.go \
  -key=$key -sec=$sec \
  -acc_tok=$acc_tok -acc_sec=$acc_sec \
  -cmd=usertimeline -count=3
```


