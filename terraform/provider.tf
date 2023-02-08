terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.19.0"
    }
  }

  backend "s3" {
    bucket = "tf-backend-state-xiaoxuan"
    key    = "state"
    region = "eu-north-1"
  }
}