### Repro 

This repository reproduces a possible bug I found with google/ko in a private repository, or at least a misconfiguration that does not yield an immediate, helpful error.


If you run:

```
cd sub
./ko.sh
```
The terminal will just hang there until you hit Ctrl-c.

At which point, it (at least on my MacOSX machine) will then say:
```
2021/12/16 19:48:53 Using base gcr.io/distroless/static:nonroot for github.com/StevenACoffman/repro/sub
2021/12/16 19:48:55 Using build config my-build for github.com/StevenACoffman/repro/sub
Error: failed to publish images: error building "ko://github.com/StevenACoffman/repro/sub": template: argsTmpl:1:60: executing "argsTmpl" at <.Env.APP>: map has no entry for key "APP"
2021/12/16 19:48:55 error during command execution:failed to publish images: error building "ko://github.com/StevenACoffman/repro/sub": template: argsTmpl:1:60: executing "argsTmpl" at <.Env.APP>: map has no entry for key "APP"
```
