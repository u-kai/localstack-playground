provider "aws" {
  region = "ap-northeast-1"

}

resource "aws_dynamodb_table" "playground_table" {
  name         = "playground_table"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "id"
  attribute {
    name = "id"
    type = "S"
  }
  tags = {
    Name = "playground_table"
  }
}
