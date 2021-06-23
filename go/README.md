# Golang Project

This project should be run using golang version 1.15.6

## Install Dependencies

`$> make deps`

## Running Tests

`$> make test`

### Outputs

The following outputs should be preserved by the build pipeline:

* ./output/golang-test.out
* ./output/golang-coverage.out

## Building the Artefact

`$> make build`

### Outputs

The following outputs should be preserved by the build pipeline:

* ./output/golang-artefact

## Publishing the Artefact

This command mocks publishing an artefact and should only be run on all commits to the master branch

`$> make publish`
