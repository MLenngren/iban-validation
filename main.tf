resource "aws_vpc" "main" {
  cidr_block           = "10.123.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true
  tags_all = {
    "Name" = "dev-male"
  }
  tags = {
    Name = "dev-male"
  }
}

resource "aws_subnet" "main_public_subnet" {
  vpc_id                  = aws_vpc.main.id
  cidr_block              = "10.123.1.0/24"
  map_public_ip_on_launch = true
  availability_zone       = "eu-north-1b"

  tags = {
    Name = "dev-male-public"
  }
}

resource "aws_internet_gateway" "main_internet_gateway" {
  vpc_id = aws_vpc.main.id

  tags = {
    "Name" = "dev-male-lgw"
  }
}

resource "aws_route_table" "dev_male_main_public_rt" {
  vpc_id = aws_vpc.main.id

  tags = {
    "Name" = "dev_male_public_rt"
  }
}

resource "aws_route" "default_route" {
  route_table_id         = aws_route_table.dev_male_main_public_rt.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.main_internet_gateway.id

}

resource "aws_route_table_association" "main_public_association" {
  subnet_id      = aws_subnet.main_public_subnet.id
  route_table_id = aws_route_table.dev_male_main_public_rt.id
}

resource "aws_security_group" "main_sg" {
  name        = "main_sg"
  description = "div security group"
  vpc_id      = aws_vpc.main.id

  ingress {
    description      = "TLS from VPC"
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  tags = {
    Name = "allow_tls"
  }
}

resource "aws_key_pair" "main_auth" {
  key_name   = "mainkey"
  #public_key = file("~/.ssh/mainawskey.pub")
  public_key = var.MAINAWSKEYPUB
}

resource "aws_instance" "dev_go" {
  instance_type          = "t3.micro"
  ami                    = data.aws_ami.server_ami.id
  key_name               = aws_key_pair.main_auth.id
  vpc_security_group_ids = [aws_security_group.main_sg.id]
  subnet_id              = aws_subnet.main_public_subnet.id
  user_data              = file("userdata.tpl")

  root_block_device {
    volume_size = 10
  }

  tags = {
    Name = "dev-go"
  }

  provisioner "local-exec" {
      command = var.isdevbox ? templatefile("${var.host_os}-ssh-config.tpl", {
          hostname = self.public_ip,
          user = "ubuntu",
          identityfile = "~/.ssh/mainawskey"
      }) : ""
      interpreter = var.host_os == "linux" ? ["bash", "-c"] : ["powershell", "-Command"]
  }
}

