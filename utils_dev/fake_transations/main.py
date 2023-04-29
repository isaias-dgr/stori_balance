from const import USER_PRODUCTS, CODES_COMPANIES
from random import randint, choice, random
from datetime import date
import csv
import boto3
import os

if __name__ == "__main__":
    client = boto3.client('s3')
    cont_general = 0
    for m in range(1, 2):
        start_date = date(2022, m, 1)
        end_date = None
        if m == 12:
            end_date = date(2023, 1, 1)
        else:
            end_date = date(2022, m+1, 1)
        
        num_days = (end_date - start_date).days
        print(end_date, start_date, m, num_days)

        for user, products in USER_PRODUCTS.items():
            file = f"./tmp/balance_{end_date}_{user}.csv"
            key = f"{end_date.year}/{m}/balance_{user}.cvs"

            with open(file, "w+") as f:
                temp_cvs = csv.writer(f)
                cont = 0
                for transaction in range(randint(1, 100)):
                    code, companie = choice(CODES_COMPANIES)
                    negative_or_positive = choice((-1, 1))
                    product = choice(products)
                    amount = round(random()*1000*negative_or_positive, 2)
                    day = randint(1, num_days)
                    date_short = f"{m}/{day}"

                    temp_cvs.writerow([transaction, user,
                                       product, date_short, code, amount, companie])
                    print([transaction, user,
                         product, date_short, code, amount, companie])
                    cont += 1
                print(f"sending: {key} transactions{cont}")

                cont_general += cont
            client.upload_file(
                file, 'bucket-story-balance-ingest-logs-dev', key)
            os.remove(file)
    print(cont_general)
