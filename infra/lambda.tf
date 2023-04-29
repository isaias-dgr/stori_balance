# Lambda ingest files logs
data "archive_file" "lambda_app_zip" {
  type        = "zip"
  source_file = "../bin/${var.app_ingest_short}/main"
  output_path = "../bin/${var.app_ingest_short}/main.zip"
}

resource "aws_lambda_function" "lambda_app" {
  function_name    = var.app_ingest_name
  role             = aws_iam_role.role_app.arn
  handler          = "main"
  runtime          = "go1.x"
  filename         = "../bin/${var.app_ingest_short}/main.zip"
  source_code_hash = filebase64sha256("../bin/${var.app_ingest_short}/main.zip")

  environment {
    variables = {
      SQS_QUEUE_URL = aws_sqs_queue.queue_app.url,
      BUCKET        = "bucket-${var.app_ingest_bucket}"
    }
  }

  tracing_config {
    mode = "Active"
  }

  timeout     = 60
  memory_size = 128
}

resource "aws_lambda_permission" "permission_app" {
  statement_id  = "permission_${var.app_ingest_name}"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda_app.function_name
  principal     = "events.amazonaws.com"

  source_arn = aws_cloudwatch_event_rule.event_rule_app.arn
}


# Lambda send notify files logs
data "archive_file" "lambda_notify_zip" {
  type        = "zip"
  source_file = "../bin/${var.app_notify_short}/main"
  output_path = "../bin/${var.app_notify_short}/main.zip"
}

resource "aws_lambda_function" "lambda_notify" {
  function_name    = var.app_notify_name
  role             = aws_iam_role.role_notify_app.arn
  handler          = "main"
  runtime          = "go1.x"
  filename         = "../bin/${var.app_notify_short}/main.zip"
  source_code_hash = filebase64sha256("../bin/${var.app_notify_short}/main.zip")

  environment {
    variables = {
      SQS_QUEUE_URL  = aws_sqs_queue.queue_app.url,
      BUCKET_LOG     = "bucket-${var.app_ingest_bucket}"
      BUCKET_NOTIFY  = "bucket-${var.app_notify_bucket}"
      DB_URL         = aws_db_instance.db_notify_app.endpoint
      DB_PASSWORD    = var.app_notify_password
      DB_USER        = "db_user_${var.app_notify_name}"
      DB_NAME        = "db_${var.app_notify_name}"
      EMAIL_SOURCE   = "isaiasd.garciar@gmail.com"
      EMAIL_TEMPLATE = "mail_balance"
    }
  }

  tracing_config {
    mode = "Active"
  }

  timeout     = 60
  memory_size = 128
}


resource "aws_lambda_permission" "permission_notify_app" {
  statement_id  = "permission_${var.app_notify_name}"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda_notify.function_name
  principal     = "events.amazonaws.com"

  source_arn = aws_sqs_queue.queue_app.arn
}

resource "aws_lambda_event_source_mapping" "sqs_trigger" {
  event_source_arn = aws_sqs_queue.queue_app.arn
  function_name    = aws_lambda_function.lambda_notify.function_name
  batch_size       = 1
}
