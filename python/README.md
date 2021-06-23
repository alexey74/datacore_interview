# Golang Project

This project should be run using python version 3.7.7

## Install Dependencies

`$> make deps`

## Running Tests

`$> make test`

### Outputs

The following outputs should be preserved by the build pipeline:

* ./output/python-test.out
* ./output/python-coverage.tar.gz

## Building the Artefact

`$> make build`

### Outputs

The following outputs should be preserved by the build pipeline:

* ./output/takehome-*.tar.gz

## Publishing the Artefact

This command mocks publishing an artefact and should only be run on all commits to the master branch

`$> make publish`
