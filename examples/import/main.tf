terraform {
  required_providers {
    hashicups = {
      version = "0.2"
      source  = "hashicorp.com/edu/hashicups"
    }
  }
}

provider "hashicups" {
  username = "education"
  password = "test123"
}

resource "hashicups_order" "sample" {}

resource "hashicups_order" "sample_4" {}

output "sample_order" {
  value = hashicups_order.sample
}
