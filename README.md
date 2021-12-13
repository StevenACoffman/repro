### Repro 

This repository reproduces a possible bug I found with google/ko in a private repository, or at least a misconfiguration that does not yield an immediate, helpful error.


If you run:

```
export GOPRIVATE=github.com/StevenACoffman
cd sub
ko publish -B --bare --local --platform=linux/amd64 --push=false .
```
The terminal will just hang there until you hit Ctrl-c.

At which point, it will then say:
```
^CError: error creating builder: 'builds': entry #0 does not contain a valid local import path (./.) for directory (.): err: signal: interrupt: stderr:
2021/12/12 21:00:19 error during command execution:error creating builder: 'builds': entry #0 does not contain a valid local import path (./.) for directory (.): err: signal: interrupt: stderr:
```

However, if you delete the .ko.yaml file in that directory and run it again, things work as you'd expect.
```
rm .ko.yaml
ko publish -B --bare --local --platform=linux/amd64 --push=false .
```

