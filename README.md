# log-pruner

A simple go server that monitors a directory and prunes files based on criteria.  It is intended to be deployed as a docker sidecar container in a Kubernetes cluster for containers that don't do log management themselves (namely Java servers).

## Docker Usage

The docker contains only the executable with the entrypoint being `log-pruner`.  You mount in directories you want to scan.

```
docker run -t -v /host/path/to/logs:/logs quay.io/jkamenik/log-pruner:1.0 watch -p /logs
```

## Command Usage

The program provides a lot of inline help: `log-pruner help`.

To scan a directory but not actually do anything destructive run `log-pruner prune --dry-run`.  If your want to monitor a directory for pruning then run `log-prune watch`.

### `log-prune help`

Provides useful information about how to use log-prune including help about the commands.

### `log-prune prune`

Run the pruner once.  This is useful for running a manual catchup or for seeing what it would do.

### `log-prune watch`

Watch for changes in log files and then run a prune.  This is normally what is run.

## Deployment

```yaml
TODO: Add kubernetes stateful set eqiv.
```

## Releases Themes

### Release 1.0

Initial release.  Simple linear scanner.

## Contributing

1. Create a branch based on the intented target version (see Versioning)
2. Add your name to the CONTRIBUTORS file
3. Add your code
4. Submit a PR

## Branching / Versioning

Unlike gitflow (which is a good branching scheme, just not the one we use) there is no single "Development" branch.  Instead each major release version gets its own long lived branch.  All the patch releases are done in order on that release branch.

The benefit of this is that multiple version of the tool can be worked at the same time.  The drawback is slightly more complicated the gitflow, which can appear complicated until you actually do it.

So the rule of thumb is that for any 1.0.x release just treat the "release-1.0" as if was "develop" branch in the gitflow diagrams.  Then everything should make sense.  For 1.1.x features go into the "release-1.1" branch, etc...

One other change is that because multiple branches are done in parallel tagging is done on the release branch before the merge, not on the master branch after the merge.

The choice of what features go into a release is in the release section above.