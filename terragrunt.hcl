
# Include the root `terragrunt.hcl` configuration. The root configuration contains settings that are common across all  components and environments, such as how to configure remote state.
include "root" {
  path = find_in_parent_folders()
}

# Include the envcommon configuration for the component. The envcommon configuration contains settings that are common for the component across all environments.
include "envcommon" {
  path   = "${dirname(find_in_parent_folders())}/_envcommon/storage.hcl"
  expose = true
}


inputs = {
  # Certificate
  ca_cert_identifier = "rds-ca-rsa2048-g1"
  ami_id             = "ami-2222222222222"

  # VM Configurations
  vm_config = {
    instance_type = "t3a.medium"
    volume_size   = 30
    count         = 4
  }
}
