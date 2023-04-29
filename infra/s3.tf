resource "aws_s3_bucket" "bucket_app" {
  bucket        = "bucket-${var.app_ingest_bucket}"
  force_destroy = true
}

resource "aws_s3_bucket" "bucket_notify_app" {
  bucket        = "bucket-${var.app_notify_bucket}"
  force_destroy = true
}

resource "aws_s3_bucket" "templates_notify_app" {
  bucket        = "templates-${var.app_notify_bucket}"
  force_destroy = true
}

resource "aws_s3_object" "templates_notify_app_object" {
  bucket = aws_s3_bucket.templates_notify_app.id
  key    = "notify_app/balance/mail_balance.html"
  source = "../templates/mail_balance.html"
}
