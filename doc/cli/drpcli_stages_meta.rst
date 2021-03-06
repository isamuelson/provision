drpcli stages meta
==================

Gets metadata for the stage

Synopsis
--------

Gets metadata for the stage

::

    drpcli stages meta [id] [flags]

Options
-------

::

      -h, --help   help for meta

Options inherited from parent commands
--------------------------------------

::

      -d, --debug               Whether the CLI should run in debug mode
      -E, --endpoint string     The Digital Rebar Provision API endpoint to talk to (default "https://127.0.0.1:8092")
      -f, --force               When needed, attempt to force the operation - used on some update/patch calls
      -F, --format string       The serialzation we expect for output.  Can be "json" or "yaml" (default "json")
      -P, --password string     password of the Digital Rebar Provision user (default "r0cketsk8ts")
      -r, --ref string          A reference object for update commands that can be a file name, yaml, or json blob
      -T, --token string        token of the Digital Rebar Provision access
      -t, --trace string        The log level API requests should be logged at on the server side
      -Z, --traceToken string   A token that individual traced requests should report in the server logs
      -U, --username string     Name of the Digital Rebar Provision user to talk to (default "rocketskates")

SEE ALSO
--------

-  `drpcli stages <drpcli_stages.html>`__ - Access CLI commands relating
   to stages
-  `drpcli stages meta add <drpcli_stages_meta_add.html>`__ - Atomically
   add [key]:[val] to the metadata on [stages]:[id]
-  `drpcli stages meta get <drpcli_stages_meta_get.html>`__ - Get a
   specific metadata item from stage
-  `drpcli stages meta remove <drpcli_stages_meta_remove.html>`__ -
   Remove the meta [key] from [stages]:[id]
-  `drpcli stages meta set <drpcli_stages_meta_set.html>`__ - Set
   metadata [key]:[val] on [stages]:[id]
