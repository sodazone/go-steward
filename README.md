TODO documentation :)

```
jq --slurp . < assets.jsonl > assets.json
```

go run main.go download --http-url http://127.0.0.1:3000 | jq --slurp . > assets.json