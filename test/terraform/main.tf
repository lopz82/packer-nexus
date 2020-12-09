terraform {
}

provider "aws" {
  region  = var.region
  version = ">= 3.2.0"
}

resource "aws_instance" "ec2" {
  ami               = var.ami
  instance_type     = var.ec2_instance_type
  security_groups   = var.security_groups
  key_name          = var.key_pair_name
  availability_zone = var.availability_zone
  tags              = var.ec2_instance_tags

  root_block_device {
    volume_size           = var.root_volume_size
    delete_on_termination = var.root_volume_delete_on_termination
  }
}
