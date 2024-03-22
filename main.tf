# This is a comment, you can use # to add comments in your OpenTofu files

# Define a resource block for your development environment
terraform {
  required_version = ">= 1.0.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.38"
    }
  }
}

# Documentation: https://www.terraform.io/docs/language/providers/requirements.html
provider "aws" {
  region  = "us-east-1"
  profile = "my-go-terraform"
  default_tags {
    tags = {
      service = "terraformGo"
    }
  }
}