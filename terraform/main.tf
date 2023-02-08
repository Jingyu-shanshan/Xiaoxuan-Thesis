resource "aws_instance" "my_vm" {
  ami           = "ami-0c651f40f9861388f"
  instance_type = "t2.micro"

  tags = {
    Name = "My EC2 instance",
  }
}