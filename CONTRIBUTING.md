# Contributing Guide

Have you created a bundle that you would like to share as an example?
In your pull request:

1. Create a directory in this repository with a unique name that identifies the use case for the bundle.
1. Put your porter.yaml and supporting files in that directory.
1. Include a README.md that explains what the example demonstrates, what the bundle does, and how to use the bundle.
1. Include two tables: one for parameters and another for credentials. This helps people see info about the bundle without having to run porter explain.

## Build an example

To build an example just like we do in CI,

```
mage BuildExample NAME
```

Run the following command to build all examples:

```
mage Build
```
