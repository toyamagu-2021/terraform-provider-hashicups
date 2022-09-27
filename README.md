# Terraform Provider Hashicups

## Source

- [Terraform official tutorial](https://learn.hashicorp.com/tutorials/terraform/provider-use?in=terraform/providers)

## Test sample configuration

First, build and install the provider.

```shell
make install
```

Then, run the docker container

```shell
make run
```

Finally, run the following command to initialize the workspace and apply the sample configuration.

```shell
terraform init && terraform apply
```
