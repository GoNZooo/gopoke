gopoke
------

# What is gopoke?

Gopoke is an application that will use a HTTP GET to ping one or more sites and return the results.

# Why make it?

Because I needed something to do in Go and it was a natural thing.

I had previously thought to do it in Elixir/Erlang, but it seemed perfectly fine to do it in Go as it turned out to be even easier.

It's basically an extended version of the application that Russ Cox makes in [A tour of Go](https://www.youtube.com/watch?v=ytEkHepK08c).

# What more?

One could harness the power of Erlang or Elixir to automatically supervise the sites one specifies, meaning you could essentially use it as a watchdog application for whichever servers you run.

More information about the facilities to do this kind of thing can be found [here](https://github.com/alco/porcelain).

Obviously these are not Go resources, but the same kind of thing could most likely be done very easily in Go.
