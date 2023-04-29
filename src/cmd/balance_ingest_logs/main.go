package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	adapter "github.com/isaias-dgr/story-balance/src/internal/adapters/input/lambda"
	queue "github.com/isaias-dgr/story-balance/src/internal/adapters/output/queue"
	storage "github.com/isaias-dgr/story-balance/src/internal/adapters/output/storage"
	usecase "github.com/isaias-dgr/story-balance/src/internal/core/use_case"
	"github.com/sirupsen/logrus"
)

func SetUpLog() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{})
	logger.SetOutput(os.Stdout)
	return logger
}

func SetUpStorage(l *logrus.Logger) *storage.Storage {
	l.Info("🪣 Set storage")
	bucket := os.Getenv("BUCKET")
	return storage.NewStorage(bucket, l)
}

func SetUpQueue(l *logrus.Logger) *queue.Queue {
	l.Info("🛣️ Set queue")
	bucket := os.Getenv("SQS_QUEUE_URL")
	return queue.NewQueue(bucket, l)
}

func main() {
	l := SetUpLog()
	l.Info("😎 SetUp Ingest log transacction and send to queue.")
	storage_service := SetUpStorage(l)
	queue_servie := SetUpQueue(l)
	use_case := usecase.NewIngestUseCase(l, storage_service, queue_servie)
	handler := adapter.NewfuncIngest(l, use_case)
	lambda.Start(handler.Handler)
}
