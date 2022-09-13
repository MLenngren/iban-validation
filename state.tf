terraform {
  backend "s3" {
    bucket = "pfcdev-male-ibanvalidator"
    key    = "default-infrastructure"
    region = "eu-north-1"
    profile = "pfcdev"
  }
}

resource "aws_s3_bucket" "terraform_state" {
  bucket = "pfcdev-male-ibanvalidator"
  tags = {
    Environment = "dev"
  }
}
resource "aws_s3_bucket_acl" "example" {
  bucket = aws_s3_bucket.terraform_state.id
  acl    = "private"
}
resource "aws_s3_bucket_versioning" "versioning_example" {
  bucket = aws_s3_bucket.terraform_state.id
  versioning_configuration {
    status = "Enabled"
  }
}
