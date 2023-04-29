resource "aws_sqs_queue" "queue_app" {
  name                       = "queue_${var.app_ingest_name}"
  visibility_timeout_seconds = 60
}
