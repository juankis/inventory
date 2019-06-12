resource "aws_instance" "web" {
  ami           = "${var.ami}" #"ami-0a313d6098716f372" #"ami-0de53d8956e8dcf80"  ubuntu : "ami-0a313d6098716f372"
  instance_type = "${var.type_instance}"
  vpc_security_group_ids =  [ "${aws_security_group.rules_firewall.id}" ] #This will be static getting permissions
  key_name = "${aws_key_pair.keypair.key_name}" #to connect ssh use ssh -i ../watchdog ec2-user@ec2-3-91-94-206.compute-1.amazonaws.com
}

output "public_ip" { value = "${aws_instance.web.public_ip}" }
