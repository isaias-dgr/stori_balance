resource "aws_cloudwatch_event_rule" "event_rule_app" {
  name        = "event_${var.app_ingest_name}"
  description = "Trigger every 2 minutes"

  schedule_expression = "rate(12 hours)"
}

resource "aws_cloudwatch_event_target" "event_target_app" {
  rule      = aws_cloudwatch_event_rule.event_rule_app.name
  target_id = "target_${var.app_ingest_name}"
  arn       = aws_lambda_function.lambda_app.arn
}


