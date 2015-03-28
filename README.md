jqouch is an experimental Couchdb view server capable of understanding
(and using) jq filters.

More info:

* [CouchDB view server
protocol](http://docs.couchdb.org/en/1.6.1/query-server/protocol.html)
* [jq](https://stedolan.github.io/jq/) and [its
manual](https://stedolan.github.io/jq/manual/)

jqouch is really simple for the moment, all it understands is "reset",
"add_fun" and "map_doc".

jqouch is actually a midleman between CouchDB and jq. It takes input
from CouchDB, spawns jq processes for each "add_fun" it sees, and
forwards documents from "map_doc" to the process. When jq responds the
response is sent back (with proper formatting) to CouchDB.

Some benchmarks with the views in this directory run on the documents
loaded by load.sh, on a i5-3210M (2.5 GHz) with a Seagate Momentus Thin
(ST320LT007) (that's a laptop hard drive running at 7200RPM):

```
22925 documents

by_title_js: 35.436s
by_title_jq: 18.943s

by_subject_js: 32.779s
by_subject_jq: 16.957s

by_date_js: 32.592s
by_date_jq: 19.347s

by_amended_by_js: 36.530s
by_amended_by_jq: 19.588s
```

Views are run a few times to eliminate outstanding values and focus on
more average numbers. The views are calculated as such:

```shell
$ time curl http://localhost:5984/test_jq/_design/<designname>/_view/<viewname> > /dev/null
```

The number is the total time taken. This benchmark is entirely
unscientific and serves more to give an idea than to give a rough idea
of what to expect.

# Installation

* Get the Go toolchain
* Get the jq binary
* Compile jqouch with
```shell
$ go build
```
* Add the view server in your /etc/couchdb/local.ini:
```ini
[query_servers]
jq = /path/to/jqouch
```
* You're done !
