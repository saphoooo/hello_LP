# hello_LP
Hello applications for Linux Pratique

## hellopy

```
$ docker run -p 5000:5000 saphoooo/hellopy
$ curl localhost:5000/hello/joe
<!doctype html>
<title>Hello from Flask</title>

  <h1>Hello joe!</h1>
```

## hellogo

```
$ docker run -p 8000:8000 saphoooo/hellopy
$ curl localhost:8000/hello/joe
<!doctype html>
<title>Hello from Go</title>
  <h1>Hello, joe!</h1>
```