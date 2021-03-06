# tfmodule-azure-actiongroup

> This project was generated by [generator-aag-terraform-module](https://github.com/nweddle/generator-aag-terraform-module)

## Overview

This module provides the Alaska Airlines standards for creating an Azure Action Group following our esablished naming conventions. The examples show how to use the module. To stay up to date on our latest changes, visit our [Change Log](./docs/CHANGELOG.md)

## Usage

```hcl
module "tfmodule-azure-actiongroup" {
  source = "github.com/AlaskaAirlines/tfmodule-azure-actiongroup"
}
```

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| azurerm | n/a |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| appName | The base name of the application used in the naming convention. | `string` | n/a | yes |
| emailAddress | email address to send alerts to | `string` | `""` | no |
| emailName | Friendly Name for email address | `string` | `""` | no |
| enableEmail | Enable email alert capabilities | `bool` | `false` | no |
| enableSMS | Enable Texting Alerts | `bool` | `false` | no |
| enableWebHook | Enable Web Hook Alerts | `bool` | `false` | no |
| environment | Name of the environment ex (Dev, Test, QA, Prod) | `string` | n/a | yes |
| resource-group-name | Name of the resource group that exists in Azure | `string` | n/a | yes |
| shortName | Required shorthand name for SMS texts. | `string` | n/a | yes |
| smsCountryCode | Country Code for phone number | `number` | `1` | no |
| smsName | Friendly Name for phone number | `string` | `""` | no |
| smsPhoneNumber | Phone number for text alerts | `number` | `0` | no |
| webhookName | Friendly Name for web hook | `string` | `""` | no |
| webhookServiceUri | The full URI for the webhook | `string` | `""` | no |

## Outputs

| Name | Description |
|------|-------------|
| action\_group\_id | n/a |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->

## Development

### Prerequisites

- [terraform](https://learn.hashicorp.com/terraform/getting-started/install#installing-terraform)
- [terraform-docs](https://github.com/segmentio/terraform-docs)
- [pre-commit](https://pre-commit.com/#install)
- [golang](https://golang.org/doc/install#install)
- [golint](https://github.com/golang/lint#installation)

### Configurations

- Configure pre-commit hooks

```sh
pre-commit install
```

- Configure golang deps for tests

```sh
> go get github.com/gruntwork-io/terratest/modules/terraform
> go get github.com/stretchr/testify/assert
```

### Tests

- Tests are available in `test` directory
- In the test directory, run the below command

```sh
go test
```

## Authors

This project is authored by below people

- James Jackson

## Contributing

Please review our [Contribution](./docs/Contribution.md) guidelines to contribute to this project.

