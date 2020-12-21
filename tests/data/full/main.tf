provider "test" {
  version = "~1.0.0"
}

variable "test" {
  default = "Default value"
}

output "test" {
  value = var.variable
}

data "test_data" "test" {}

resource "test_resource" "test" {
  name = data.test_data.test.name
}

module "storage" {
  source = "./modules/storage"
}
