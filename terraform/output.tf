output "server_arn" {
  description = "server ARN"
  value       = aws_instance.app_server.arn
}

output "server_name" {
  description = "Name (id) of the server"
  value       = aws_instance.app_server.id
}