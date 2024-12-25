provider "aws" {
  region = "us-east-1" # Change to your desired region
}

resource "aws_instance" "my_instance" {
  count         = 5
  ami           = "ami-0c02fb55956c7d316" # Replace with a valid AMI ID for your region
  instance_type = "t2.micro"

  key_name      = "sid.pem" # Replace with your key pair name

  tags = {
    Name = "Instance-${count.index + 1}"
  }
}

output "instance_public_ips" {
  value = aws_instance.my_instance[*].public_ip
}

output "instance_ids" {
  value = aws_instance.my_instance[*].id
}
