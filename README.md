# Nigerian Banks

![hero](https://res.cloudinary.com/ichtrojan/image/upload/v1594159123/ngbanks_kzboia.png)

## Introduction

This project aims to simplify bank data used by developers for Fintech APIs (Flutterwave and Paystack). Paystack uses `bank_code` to process transfers to accounts while Flutterwave uses `bank_slug` for the same operation. As a bonus, bank logos are also returned. How is this useful? -- can be used to spice up your bank dropdown/list UI, It doesn't have to be boring.

## Usage

Make a get request to `https://nigerianbanks.xyz`

## Prerequisites
* [Go](https://golang.org) installed on your machine
* [Docker](https://docker.com) installed on your machine (optional)

## Installation

* Clone this repository 🤷‍♂️ (obviously)

```bash
git clone https://github.com/ichtrojan/nigerian-banks.git
```

* Change directory

```bash
cd nigerian-banks
```

* Duplicate `.env.example` to `.env`

```bash
cp .env.example .env
```

* Run application

```bash
go run server.go
```

Alternatively, if you are a Docker fanboy, you can run:

```bash
docker-composer up
```

Your application will be served on port `9090` by default, you can change that by modifying the `.env` file.

## Contributors

* Deji Ajibola - [Twitter](https://twitter.com/damndeji) [GitHub](https://github.com/youthtrouble)

## Authors

* Kamsi Oleka - [Twitter](https://twitter.com/Eze_Mmuo) [Github](https://github.com/kamsy)
* Michael Trojan Okoh - [Twitter](https://twitter.com/ichtrojan) [Github](https://github.com/ichtrojan)
