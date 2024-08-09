# azurenum

## Overview
A command line tool for enumerating and monitoring Azure environments for security misconfigurations and exposed assets 

## Features

* Get keyvaults with expired secrets
* Get exposed storage buckets
* Output as a table or json

## Usage
```yaml
Usage:
  azurenum [command]

Available Commands:
  completion      Generate the autocompletion script for the specified shell
  help            Help about any command
  keyvault        Get all keyvaults
  resource-groups Get all resource groups
  storage-account Get all Storage Accounts
  subscriptions   Get all subscriptions
  tenants         Get all tenants

Flags:
  -h, --help                  help for azurenum
  -s, --subscription string   Subscription ID
      --teams string          Send result to teams
  -t, --tenant string         Tenant name or ID

Use "azurenum [command] --help" for more information about a command.
```

## Installation
```sh
go install -v github.com/mattytmn/azurenum@latest
```
## Running
```sh
azurenum -h
```