# stori Balance

## Description

Aplicacion para consumir archivos CSV Crear un reporte o Balance bancario y ser enviado por email.

- El sistema cuenta con dos lambdas.
  - Lambda ingest-logs se ejecuta con un triger temporal cada mes. Este lambda consigue la lista de archivos con transacciones y envia la direcciones a una cola de mensajes.
  - Lambda notify-balance se ejecuta con un triger de sqs.
    - Obtiene todas las transacciones las guarda en una instancia Mysql RDS .
    - Crear un document(balance) con las transacciones ordenadas por fecha y por producto y es alojada en un bucket. (deberia ser un pdf pero fue html por cuestiones de tiempo)
    - Envia notificacion con un resumen del balance y un elace presigned del archivo en S3.
    - Envia un sms notificando el balance esta listo!.
    - // TODO hacer que este lambda se ejecute cada que se crea un archivo con logs en s3.

![Ingest balances](https://github.com/isaias-dgr/stori_balance/assets/89608187/3bffb9d0-755c-4153-ba8d-1cb06e994ecc)
[Ingest balances.pdf](https://github.com/isaias-dgr/stori_balance/files/11558977/Ingest.balances.pdf)

## Assumptions

- Un sistema externo crear archivos csv por cada usuario del banco con las transacciones creadas en un mes
- El archivo puede contener transacciones de varios productos (credito, debito etc).

### Assumptions files

Examples files cvs

```bash
0,2ca03e45-8d0e-25b3-bce3-2b1d299c10cc,5048379029392553,5/5,OXX,-49.66,Oxxo
1,2ca03e45-8d0e-25b3-bce3-2b1d299c10cc,5048379029392553,5/14,OXX,-815.25,Oxxo
2,2ca03e45-8d0e-25b3-bce3-2b1d299c10cc,5048379029392553,5/22,NFLX,-653.87,Netflix
3,2ca03e45-8d0e-25b3-bce3-2b1d299c10cc,5108758856586949,5/17,DIS,-456.72,The Walt Disney Company
4,2ca03e45-8d0e-25b3-bce3-2b1d299c10cc,5108758856586949,5/11,CQC,213.83,Cielito Querido Café
5,2ca03e45-8d0e-25b3-bce3-2b1d299c10cc,5108758856586949,5/26,UBR,-292.82,Uber
6,2ca03e45-8d0e-25b3-bce3-2b1d299c10cc,5048379029392553,5/21,WMT,-126.21,Walmart
7,2ca03e45-8d0e-25b3-bce3-2b1d299c10cc,5108758856586949,5/15,AAPL,80.68,Apple
8,2ca03e45-8d0e-25b3-bce3-2b1d299c10cc,5048379029392553,5/19,UBR,-559.78,Uber
9,2ca03e45-8d0e-25b3-bce3-2b1d299c10cc,5048379029392553,5/15,GEN,-469.22,Generico
10,2ca03e45-8d0e-25b3-bce3-2b1d299c10cc,5048379029392553,5/30,NFLX,816.97,Netflix
11,2ca03e45-8d0e-25b3-bce3-2b1d299c10cc,5048379029392553,5/5,OXX,-684.42,Oxxo
12,2ca03e45-8d0e-25b3-bce3-2b1d299c10cc,5048379029392553,5/26,DIS,159.89,The Walt Disney Company
13,2ca03e45-8d0e-25b3-bce3-2b1d299c10cc,5048379029392553,5/8,AAPL,-17.39,Apple
14,2ca03e45-8d0e-25b3-bce3-2b1d299c10cc,5048379029392553,5/30,TMT,157.14,Tecnológico de Monterrey
15,2ca03e45-8d0e-25b3-bce3-2b1d299c10cc,5108758856586949,5/10,AMZN,585.5,Amazon
16,2ca03e45-8d0e-25b3-bce3-2b1d299c10cc,5048379029392553,5/6,SEL,749.55,Seven eleven
17,2ca03e45-8d0e-25b3-bce3-2b1d299c10cc,5048379029392553,5/13,RAP,-680.67,Rappid
18,2ca03e45-8d0e-25b3-bce3-2b1d299c10cc,5108758856586949,5/9,RR,467.61,Rolls-Royce Holdings
19,2ca03e45-8d0e-25b3-bce3-2b1d299c10cc,5048379029392553,5/21,WMT,404.33,Walmart
20,2ca03e45-8d0e-25b3-bce3-2b1d299c10cc,5108758856586949,5/24,GEN,752.75,Generico
```

## Outputs

![Screenshot from 2023-05-24 14-47-29](https://github.com/isaias-dgr/stori_balance/assets/89608187/d3bda605-ba0e-4cf6-8196-d9a5bb70653f)

file:///home/isaias/Pictures/Screenshots/Screenshot%20from%202023-05-24%2014-47-59.png![image](https://github.com/isaias-dgr/stori_balance/assets/89608187/dfc68ed3-cd58-4394-a7c8-f96e9f155823)

### TODO

- [ ] Crear archivos de configuracion para la creacion de contenedores docker
- [ ] Agregar transacciones en DB
- [ ] Unittest/Coverage
- [ ] Pipeline de despliegue
- [ ] API para crear dashboard
- [ ] Nuevo esquema para soportar S3 como trigger de Lambda
- [ ] Mejorar terraform files
- [ ] Crear migraciones automaticos

## Requierments.

- Terraform
- Golang 1.19 +
- Docker (proximamente)
- AWS Acount (Localstack)
- Makefile

## Build Lambda Ingest Logs

```bash
make build-lambda-logs
```

## Build Lambda Notify Balance

```bash
make build-lambda-pdf
```

## Create infrastructure

```bash
make infra-init
make infra-plan
make infra-apply
```

## Delete infrastructure

```bash
make infra-destroy
```

## Deploy App

work in progress
