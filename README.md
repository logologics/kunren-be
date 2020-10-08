# kunren-be
Kunren backend

# mongo
```
docker run --rm -p 27017:27017 --name some-mongo mongo:latest
```

# remember
```
curl -vX POST -H 'contenttype: application/json' --data-binary @resource/words/jisho.home.json localhost:9876/remember
```

# search
```
curl -H "ContentType: application/json" localhost:9876/search/jisho/home
```