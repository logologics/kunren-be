# kunren-be
Kunren backend

# mongo
```
docker run --rm -p 27017:27017 --name some-mongo mongo:latest
...
docker exec -ti some-mongo bash
mongo

```

# remember
```
curl -vX POST -H 'contenttype: application/json' --data-binary @resource/words/word.home.json localhost:9876/remember
```

# search

```
curl -H "ContentType: application/json" localhost:9876/search/jisho？q＝home
```

# list vocab
```
curl -H "ContentType: application/json" 'localhost:9876/vocabs/list?k='
```

# find vocab
```
curl -H "ContentType: application/json" 'localhost:9876/vocabs/find?l=ja&k=猫'
```


# Atlas

# Kagome
```
docker run -d  -p 6060:6060 ikawaha/kagome:latest server
curl -XPUT localhost:6060/tokenize -d'{"sentence":"すもももももももものうち", "mode":"normal"}' | jq 
```

# docker
```
cp -r ../kunren-fe/build/* web
docker build . -t kunren
docker run -v /Users/kaiser/gitp/kunren-be/config:/config -p 9876:9876 kunren
```

# Deployment
- Build container
- Deploy on fargate
- Expose through static IP in public subnet
- change DNS record
- create cert
- write lambda function to run certbot



