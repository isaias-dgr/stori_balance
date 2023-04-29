resource "aws_security_group" "rds_mysql_sg" {
  name_prefix = "sg_${var.app_notify_name}"

  ingress {
    from_port   = 3306
    to_port     = 3306
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 0
    to_port     = 65535
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 65535
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_db_instance" "db_notify_app" {
  allocated_storage = 10
  engine            = "mysql"
  engine_version    = "8.0.28"
  instance_class    = "db.t3.micro"

  db_name  = "db_${var.app_notify_name}"
  username = "db_user_${var.app_notify_name}"
  password = var.app_notify_password
  multi_az = false

  parameter_group_name = "default.mysql8.0"
  skip_final_snapshot  = true

  publicly_accessible = true

  vpc_security_group_ids = [aws_security_group.rds_mysql_sg.id]
}

output "database_endpoint_address" {
  value = aws_db_instance.db_notify_app.endpoint
}


output "database_vpc" {
  value = data.aws_db_subnet_group.database.vpc_id
}



data "aws_db_subnet_group" "database" {
  name = aws_db_instance.db_notify_app.db_subnet_group_name
}
