variable "region" {
  description = "AWS region where thee EC2 instance will be deployed"
  type        = string
}

variable "ami" {
  description = "AMI image to be deployed"
  type        = string
}

variable "ec2_instance_type" {
  description = "AWS EC2 instance type"
  default     = "t3.medium"
  type        = string
}

variable "security_groups" {
  description = "List of security groups"
  type        = list(string)
}

variable "key_pair_name" {
  description = "Private SSH key to provisioning. It is the private key from the AWS key pair"
  type        = string
}

variable "availability_zone" {
  description = "AWS availability zone"
  type        = string
}

variable "ec2_instance_tags" {
  description = "Tags to be applied to the EC2 instance"
  type        = map(string)
}

variable "root_volume_size" {
  description = "Size of root block device"
  default     = 8
  type        = number
}

variable "root_volume_delete_on_termination" {
  description = "If set to true deletes the root block device after deleting the EC2 instance"
  default     = true
  type        = bool
}
