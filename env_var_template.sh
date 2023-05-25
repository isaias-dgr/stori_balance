echo "🛠️ Load Var Env"
export COMPANY=stori
export DEPTO=balance
export APP_INGEST="ingest_logs"
export APP_NOTIFY="notify"
export ENV=dev
export PASSWORD=0fc89521-652b-4746-a760-60f684dfa300


export APP_INGEST_SHORT="${DEPTO}_${APP_INGEST}"
export APP_NOTIFY_SHORT="${DEPTO}_${APP_NOTIFY}"

export APP_INGEST_BUCKET="${COMPANY}-${DEPTO}-ingest-logs-${ENV}"
export APP_NOTIFY_BUCKET="${COMPANY}-${DEPTO}-notify-${ENV}"

export APP_INGEST_NAME="${COMPANY}_${APP_INGEST_SHORT}_${ENV}"
export APP_NOTIFY_NAME="${COMPANY}_${APP_NOTIFY_SHORT}_${ENV}"

export TF_VAR_app_ingest_short="${APP_INGEST_SHORT}"
export TF_VAR_app_notify_short="${APP_NOTIFY_SHORT}"

export TF_VAR_app_ingest_name="${APP_INGEST_NAME}"
export TF_VAR_app_notify_name="${APP_NOTIFY_NAME}"

export TF_VAR_app_ingest_bucket="${APP_INGEST_BUCKET}"
export TF_VAR_app_notify_bucket="${APP_NOTIFY_BUCKET}"

export TF_VAR_app_notify_password="${PASSWORD}"

echo "🏁 End Load Var Env ${APP_INGEST_NAME}"
echo "🏁 End Load Var Env ${APP_NOTIFY_NAME}"