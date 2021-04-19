# Install

```
kubectl apply -f https://download.elastic.co/downloads/eck/1.3.1/all-in-one.yaml
```


# SETUP mappings

```
curl --location --request PUT 'https://localhost:9200/games/' \
--header 'Authorization: Basic ZWxhc3RpYzo3bkM5ak8zbkcwM3lSODk3ajBpbXRzOWY=' \
--header 'Content-Type: application/json' \
--data-raw '{
    "mappings": {
        "properties": {
            "name": {
                "type": "text"
            },
            "url": {
                "type": "keyword"
            },
            "maxPlayers": {
                "type": "integer"
            },
            "suggest": {
                "type": "completion"
            }
        }
    }
}'
```

# SETUP Games

````
curl --location --request POST 'https://localhost:9200/games/_doc' \
--header 'Authorization: Basic ZWxhc3RpYzo3bkM5ak8zbkcwM3lSODk3ajBpbXRzOWY=' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Pacman",
    "url": "/games/pacman",
    "maxPlayers": 1,
    "suggest": ["Pacman"]
}'

curl --location --request POST 'https://localhost:9200/games/_doc' \
--header 'Authorization: Basic ZWxhc3RpYzo3bkM5ak8zbkcwM3lSODk3ajBpbXRzOWY=' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Texas Holdem Poker",
    "url": "/games/poker",
    "maxPlayers": 10,
    "suggest": ["Texas", "Holdem", "Poker"]
}'

curl --location --request POST 'https://localhost:9200/games/_doc' \
--header 'Authorization: Basic ZWxhc3RpYzo3bkM5ak8zbkcwM3lSODk3ajBpbXRzOWY=' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Slot Machine",
    "url": "/games/slot",
    "maxPlayers": 1,
    "suggest": ["Slot", "Machine"]
}'
```