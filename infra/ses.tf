data "aws_s3_object" "template_object_s3" {
  bucket = aws_s3_bucket.templates_notify_app.id
  key    = "notify_app/balance/mail_balance.html"
}

resource "aws_ses_template" "MyTemplate" {
  name    = "mail_balance"
  subject = "Balance account!"
  html    = data.aws_s3_object.template_object_s3.body
  text    = "Balance account!"
}
