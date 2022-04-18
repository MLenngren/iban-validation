terraform {
  backend "s3" {
    bucket = "maletest-pfc-ibanvalidator"
    key    = "default-infrastructure"
    region = "eu-north-1"
  }
}

resource "aws_s3_bucket" "terraform_state" {
  bucket = "maletest-pfc-ibanvalidator"
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
