## 共通のクラスタ
resource "aws_ecs_cluster" "cluster" {
  name = var.cluster_name
}

# Go サービス定義
resource "aws_ecs_service" "go_service" {
  name            = "go-service"
  cluster         = aws_ecs_cluster.cluster.id
  task_definition = aws_ecs_task_definition.go_task.arn
  launch_type     = "FARGATE"
  desired_count   = var.desired_count

  network_configuration {
    subnets          = var.subnets
    security_groups  = [aws_security_group.ecs_tasks_sg.id]
    assign_public_ip = var.assign_public_ip
  }
}

# Go タスク定義
resource "aws_ecs_task_definition" "go_task" {
  family                   = "go-task-family"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = var.task_cpu
  memory                   = var.task_memory
  execution_role_arn       = var.ecs_execution_role_arn

  container_definitions = jsonencode([{
    name  = "go-container"
    image = aws_ecr_repository.app["go"].repository_url
  }])
}

# Go タスクのオートスケーリング
resource "aws_appautoscaling_target" "go_target" {
  service_namespace  = "ecs"
  resource_id        = "service/${aws_ecs_cluster.cluster.name}/${aws_ecs_service.go_service.name}"
  scalable_dimension = "ecs:service:DesiredCount"

  min_capacity = 1
  max_capacity = 2
}

resource "aws_appautoscaling_policy" "go_cpu_scale_up" {
  name               = "cpu-scale-up"
  service_namespace  = aws_appautoscaling_target.go_target.service_namespace
  scalable_dimension = aws_appautoscaling_target.go_target.scalable_dimension
  resource_id        = aws_appautoscaling_target.go_target.resource_id
  policy_type        = "TargetTrackingScaling"

  target_tracking_scaling_policy_configuration {
    target_value       = 80
    predefined_metric_specification {
      predefined_metric_type = "ECSServiceAverageCPUUtilization"
    }

    scale_out_cooldown = 60
    scale_in_cooldown  = 300
  }
}

# Next.js サービス定義
resource "aws_ecs_service" "next_js_service" {
  name            = "next-js-service"
  cluster         = aws_ecs_cluster.cluster.id
  task_definition = aws_ecs_task_definition.next_js_task.arn
  launch_type     = "FARGATE"
  desired_count   = var.desired_count

  network_configuration {
    subnets          = var.subnets
    security_groups  = [aws_security_group.ecs_tasks_sg.id]
    assign_public_ip = var.assign_public_ip
  }
}

# Next.js タスク定義
resource "aws_ecs_task_definition" "next_js_task" {
  family                   = "next-js-task-family"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = var.task_cpu
  memory                   = var.task_memory
  execution_role_arn       = var.ecs_execution_role_arn

  container_definitions = jsonencode([{
    name  = "next-js-container"
    image = aws_ecr_repository.app["next_js"].repository_url
  }])
}

# Next.js タスクのオートスケーリング
resource "aws_appautoscaling_target" "next_js_target" {
  service_namespace  = "ecs"
  resource_id        = "service/${aws_ecs_cluster.cluster.name}/${aws_ecs_service.next_js_service.name}"
  scalable_dimension = "ecs:service:DesiredCount"

  min_capacity = 1
  max_capacity = 2
}

resource "aws_appautoscaling_policy" "next_js_cpu_scale_up" {
  name               = "cpu-scale-up"
  service_namespace  = aws_appautoscaling_target.next_js_target.service_namespace
  scalable_dimension = aws_appautoscaling_target.next_js_target.scalable_dimension
  resource_id        = aws_appautoscaling_target.next_js_target.resource_id
  policy_type        = "TargetTrackingScaling"

  target_tracking_scaling_policy_configuration {
    target_value       = 80
    predefined_metric_specification {
      predefined_metric_type = "ECSServiceAverageCPUUtilization"
    }

    scale_out_cooldown = 60
    scale_in_cooldown  = 300
  }
}