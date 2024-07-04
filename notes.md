```
docker build -t n9mi/coba-go-actions:v1.0.0 .
docker run --rm -p 3000:3000 -e FOO=bar n9mi/coba-go-actions:v1.0.0
```