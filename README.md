* Add station: `depts add <name> <code>`, e.g. `depts add Forskningsparken NSR:StopPlace:59600`
* List stations: `depts ls`
* Show departures for a certain station: `depts show [name]`.
  * full name: `depts show Forskningsparken`
  * prefix name: `depts show Fo`. As long as the prefix uniquely identifies a station, it is valid.
  * By wifi network name: See section below.

## Wifi network name

We can associate a station to a wifi connection. If my home network is named `HomeNetwork` and it
is close to `Frydenlund`, and the work network is named `WorkNetwork` and it is close to `Forskningsparken`, I can create an association `HomeNetwork -> Frydenlund` and `WorkNetwork -> Forskningsparken`. Depending on which network I am connected to, `depts show` will pick the appropriate station.

To associate a network to a station we need to know the station ID first. For Frydenlund it is 1:

```
$ depts ls
1   Frydenlund            NSR:StopPlace:58405
2   Forskningsparken      NSR:StopPlace:59600
3   Skullerud             NSR:StopPlace:58227
```

Then I can associate the current network with (say) Frydenlund:
```
$ depts associate 1
OK, associated station 'Frydenlund' to network with ssid 'HomeNetwork' 
```

If the station name is not specified and we have such association, the command `depts show` will
list departures from that station (equivalent to `depts show Frydenlund`):

```
$ depts show
19  Majorstuen                 -1.4m
18  Rikshospitalet              1.6m
17  Rikshospitalet             11.6m
18  Rikshospitalet             21.6m
19  Majorstuen                 18.6m
```

## Live updates
We can just use the beloved `watch` command to get a nice live table:
```
watch -tn5 depts show Frydenlund
```
