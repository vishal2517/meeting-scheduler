provider "aws" {
  region = "us-east-1"
}

resource "aws_instance" "meeting_api" {
  ami           = "ami-12345678"
  instance_type = "t2.micro"

  provisioner "remote-exec" {
    inline = [
      "docker run -d -p 8080:8080 meeting-scheduler"
    ]
  }
}
