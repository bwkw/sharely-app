output "next_js_repository_url" {
  value = aws_ecr_repository.app["next_js"].repository_url
}

output "go_repository_url" {
  value = aws_ecr_repository.app["go"].repository_url
}