resource "aws_s3_bucket" "s3" {
  bucket = var.bucket_name
  region = var.AWS_REGION

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}

resource "aws_instance" "my_vm" {
  ami           = "ami-065deacbcaac64cf2"
  instance_type = "t2.micro"
  region        = var.AWS_REGION

  tags = {
    Name = "My EC2 instance",
  }
}