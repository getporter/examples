# WhalesayD

Take [getporter/examples/images/whalesay][whalesay] and throw it behind a http endpoint for great glory!

## Key Info

* Port: 8080
* Environment Variables
  * DEFAULT_MESSAGE: Set the default message printed by the delightful whale
* Endpoint: <http://localhost:8080?msg=MESSAGE>

## Try it out

```bash
docker run -d -e DEFAULT_MESSAGE="whale aren't you just precious?" -p 8080:8080 ghcr.io/getporter/examples/images/whalesayd
```

```bash
$ curl http://localhost:8080?msg=time+for+a+nap!
 _________________
< time for a nap! >
 -----------------
    \
     \
      \
                    ##        .
              ## ## ##       ==
           ## ## ## ##      ===
       /""""""""""""""""___/ ===
  ~~~ {~~ ~~~~ ~~~ ~~~~ ~~ ~ /  ===- ~~~
       \______ o          __/
        \    \        __/
          \____\______/
```

[whalesay]: https://github.com/orgs/getporter/packages/container/package/examples%2Fimages%2Fwhalesay
