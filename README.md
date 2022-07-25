# TagNow

TagNow is a command line utility tool to help create new semantic versioned tags. 

### Usage 

```
$: tagnow major -push
```

Running `tagnow major` would bump the major version of your repository's tags by incrementing from the newest tag, if no tags exist tagnow creates a new tag `v1.0.0`. Using the `-push` flag would automatically push the new tag to origin.

### TODO

[] Support for suffixed tags such as `v1.0.0-prod`
