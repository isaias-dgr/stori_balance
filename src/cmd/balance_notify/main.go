package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"

	adapter "github.com/isaias-dgr/stori-balance/src/internal/adapters/input/lambda"
	mailer "github.com/isaias-dgr/stori-balance/src/internal/adapters/output/mailer"
	notifier "github.com/isaias-dgr/stori-balance/src/internal/adapters/output/notifier"
	repo "github.com/isaias-dgr/stori-balance/src/internal/adapters/output/repository"
	storage "github.com/isaias-dgr/stori-balance/src/internal/adapters/output/storage"

	writter "github.com/isaias-dgr/stori-balance/src/internal/adapters/output/document"
	usecase "github.com/isaias-dgr/stori-balance/src/internal/core/use_case"
	"github.com/sirupsen/logrus"
)

func SetUpLog() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{})
	logger.SetOutput(os.Stdout)
	return logger
}

func SetUpStorageIng(l *logrus.Logger) *storage.Storage {
	l.Info("🪣 Set storage")
	bucket := os.Getenv("BUCKET_LOG")
	return storage.NewStorage(bucket, l)
}

func SetUpStorageNot(l *logrus.Logger) *storage.Storage {
	l.Info("🪣 Set storage")
	bucket := os.Getenv("BUCKET_NOTIFY")
	return storage.NewStorage(bucket, l)
}

func SetUpRepository(l *logrus.Logger) *repo.Repository {
	l.Info("🐬 Set Repository")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	url := os.Getenv("DB_URL")
	db_name := os.Getenv("DB_NAME")
	return repo.NewRepository(l, user, password, url, db_name)
}

func SetUpNotifier(l *logrus.Logger) *notifier.Notifier {
	l.Info("📲 Set Notifier")
	return notifier.NewNotifier(l)
}

func SetUpMailer(l *logrus.Logger) *mailer.Mailer {
	l.Info("✉️ Set Mailer")
	email_source := os.Getenv("EMAIL_SOURCE")
	template := os.Getenv("EMAIL_TEMPLATE")
	l.Infof("✉️ Set Mailer %s %s", email_source, template)
	return mailer.NewMailer(email_source, template, l)
}

func SetUpWriter(l *logrus.Logger) *writter.Writter {
	l.Info("✉️ Set Writter")
	return writter.NewWritter(l)
}

func main() {
	l := SetUpLog()
	l.Info("😎 SetUp Notify Balanace")
	storage_ing_service := SetUpStorageIng(l)
	storage_not_service := SetUpStorageNot(l)
	repository_service := SetUpRepository(l)
	notifier_service := SetUpNotifier(l)
	mailer_servicer := SetUpMailer(l)
	writter_servicer := SetUpWriter(l)

	use_case := usecase.NewNotifyBalance(l,
		repository_service,
		mailer_servicer,
		storage_ing_service,
		storage_not_service,
		notifier_service,
		writter_servicer)

	handler := adapter.NewfuncNotify(l, use_case)
	lambda.Start(handler.Handler)
}
