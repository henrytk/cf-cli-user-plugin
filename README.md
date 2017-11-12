# CF CLI User Plugin

This is a work-in-progress plugin for user management in Cloud Foundry. So far, it is only a proof of concept and provides a single `whoami` style function for printing the name, GUID, and email of the logged in CF CLI user.

## Requirements

* You must have Go installed
* You must have the Cloud Foundry CLI installed.
* Installation to the CF CLI is done using a `make` helper script.

## Installation

Run the following to download this repository and install the plugin to your CF CLI.

```
git clone https://github.com/henrytk/cf-cli-user-plugin.git
cd cf-cli-user-plugin
make install
```

## Usage

```
cf user whoami
```

For help:

```
cf user help
cf user whoami help
```