# Policy for lambda ingest
resource "aws_iam_role" "role_app" {
  name = "role_${var.app_ingest_name}"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_policy" "policy-app" {
  name = "policy_${var.app_ingest_name}"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "events:DescribeRule",
          "events:ListTargetsByRule"
        ]
        Effect = "Allow"
        Resource = [
          "arn:aws:events:::*"
        ]
      },
      {
        Action = [
          "s3:ListBucket",
          "s3:GetBucketLocation"
        ]
        Effect = "Allow"
        Resource = [
          "arn:aws:s3:::bucket-${var.app_ingest_bucket}"
        ]
      },
      {
        Effect : "Allow",
        Action : [
          "s3:GetObject"
        ],
        Resource : [
          "arn:aws:s3:::bucket-${var.app_ingest_bucket}/*"
        ]
      },
      {
        Action = [
          "sqs:SendMessage"
        ]
        Effect = "Allow"
        Resource = [
          aws_sqs_queue.queue_app.arn
        ]
      },
      {
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ]
        Effect = "Allow"
        Resource = [
          "arn:aws:logs:*:*:*"
        ]
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "policy-app_attachment" {
  policy_arn = aws_iam_policy.policy-app.arn
  role       = aws_iam_role.role_app.name
}

# Policy for lambda ingest
resource "aws_iam_role" "role_notify_app" {
  name = "role_${var.app_notify_name}"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_policy" "policy_notify_app" {
  name = "policy_${var.app_notify_name}"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "s3:ListBucket",
          "s3:GetBucketLocation",
          "s3:GetObject"
        ]
        Effect = "Allow"
        Resource = [
          aws_s3_bucket.bucket_app.arn,
          "${aws_s3_bucket.bucket_app.arn}/*"
        ]
      },
      {
        Action = [
          "s3:PutObject",
          "s3:GetObject",
          "s3:GetBucketLocation",
        ]
        Effect = "Allow"
        Resource = [
          aws_s3_bucket.bucket_notify_app.arn,
          "${aws_s3_bucket.bucket_notify_app.arn}/*"
        ]
      },
      {
        Effect : "Allow",
        Action : [
          "rds-db:connect",
          "rds-data:ExecuteSql",
          "rds-data:ExecuteStatement",
          "rds-data:BatchExecuteStatement",
        ],
        Resource : [
          aws_db_instance.db_notify_app.arn
        ]
      },
      {
        Effect : "Allow",
        Action : [
          "ec2:CreateNetworkInterface",
          "ec2:DescribeNetworkInterfaces",
          "ec2:DeleteNetworkInterface"
        ],
        Resource : [
          "*"
        ]
      },
      {
        Action = [
          "sqs:ReceiveMessage",
          "sqs:GetQueueAttributes",
          "sqs:GetQueueUrl",
          "sqs:DeleteMessage"
        ]
        Effect = "Allow"
        Resource = [
          aws_sqs_queue.queue_app.arn
        ]
      },
      {
        Action = [
          "ses:SendTemplatedEmail",
          "ses:SendEmail"
        ]
        Effect = "Allow"
        Resource = [
          "*"
        ]
      },
      {
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ]
        Effect = "Allow"
        Resource = [
          "arn:aws:logs:*:*:*"
        ]
      },
      {
        Action = [
          "sns:Publish",
        ]
        Effect = "Allow"
        Resource = [
          "*"
        ]
      }
    ]
  })
}



resource "aws_iam_role_policy_attachment" "policy_app_notify_attachment" {
  policy_arn = aws_iam_policy.policy_notify_app.arn
  role       = aws_iam_role.role_notify_app.name
}
