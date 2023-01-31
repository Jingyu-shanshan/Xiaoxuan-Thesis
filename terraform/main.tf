variable "AWS_REGION" {
  type = string
}

# Configure the AWS provider
provider "aws" {
  region = var.AWS_REGION
}

# Create an EC2 instance
resource "aws_instance" "primary-aws-instance" {
  ami           = "ami-785db401"
  instance_type = "t2.micro"
}