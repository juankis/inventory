resource "aws_key_pair" "keypair" {
  key_name   = "terraform-keypair"
  public_key = "${var.public_key}"
}