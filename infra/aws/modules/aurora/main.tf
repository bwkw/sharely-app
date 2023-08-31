# ---------------------------
# Aurora RDS Cluster
# ---------------------------
resource "aws_rds_cluster" "aurora_cluster" {
  cluster_identifier      = "${var.environment}-${var.app_name}-aurora-cluster"
  engine                  = "aurora-mysql"
  database_name           = "${var.environment}${var.app_name}"
  master_username         = var.db_username
  master_password         = var.db_password
  backup_retention_period = 5
  preferred_backup_window = "07:00-09:00"
  skip_final_snapshot     = true
  vpc_security_group_ids  = var.security_group_ids
  db_subnet_group_name    = aws_db_subnet_group.aurora.name
  availability_zones      = [var.availability_zone_a, var.availability_zone_c]
}

resource "aws_rds_cluster_instance" "aurora_instance_1a" {
  identifier         = "${var.environment}-${var.app_name}-aurora-instance-1a"
  cluster_identifier = aws_rds_cluster.aurora_cluster.id
  instance_class     = var.instance_class
  availability_zone  = var.availability_zone_a
  engine = "aurora-mysql"
}

resource "aws_rds_cluster_instance" "aurora_instance_1c" {
  identifier         = "${var.environment}-${var.app_name}-aurora-instance-1c"
  cluster_identifier = aws_rds_cluster.aurora_cluster.id
  instance_class     = var.instance_class
  availability_zone  = var.availability_zone_c
  engine = "aurora-mysql"
}

# ---------------------------
# DB Subnet Group
# ---------------------------
resource "aws_db_subnet_group" "aurora" {
  name       = "${var.environment}-${var.app_name}-aurora-subnet-group"
  subnet_ids = [
    var.subnet_pri2_1a_id,
    var.subnet_pri2_1c_id
  ]

  tags = {
    Name = "${var.environment}-${var.app_name}-aurora-subnet-group"
  }
}
