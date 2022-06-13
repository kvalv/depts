# Introduction

A CLI tool for transport departures. This is really just a wrapper around [entur's graphql API](https://developer.entur.org/pages-journeyplanner-journeyplanner) (which is really great). The primary motivation for me is to just have a project where I can have a project and set up various infrastructure (documentation, automated tests, build artifacts, etc.).

## Commands


```
Usage: depts <command>

Flags:
  -h, --help     Show context-sensitive help.
      --debug    Enable debug mode

Commands:
  show [<station>]
    List info for station

  add <name> <code>
    add a new station to database

  rm <id>
    Remove a station by its id

  associate <id>

  ls
    List stored stations
```


???+ example "list all favourite stations"

    ```bash
    $ depts ls
    1   Frydenlund            NSR:StopPlace:58405
    2   Forskningsparken      NSR:StopPlace:59600
    3   Skullerud             NSR:StopPlace:58227
    ```

## Matching station in `depts show <name>`

The `<name>` for the station need not be exact. To figure out which station we're using, three checks are done in order:

1. If `<name>` is empty, use the current network name and use any station bound to that network. If no such station exists, return an error.
2. If `<name>` is not empty, check if there exists a station that exactly matches the name. For example, if the name is `skullerud` and there exists a station with that name, use the `skullerud` station.
3. Otherwise, try to match on prefix. For example, if `sku` is provided, and there exists only one station with that prefix (`skullerud`), that station will be matched.
