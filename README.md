
# Folder Monitor Service

This service monitors a specified directory for new files and integrates with Kafka to send messages about new files detected, specifically targeting `.csv` files.

## Prerequisites

- Go 1.20 or higher
- Apache Kafka 3.6.1 (can be installed using Homebrew)

## Configuration

Before running the service, ensure that `constants.go` is correctly configured. The most important constant to set is `watchDir`, which specifies the directory to monitor. Other constants have default values but can be modified as needed.

## Installation

### Kafka Setup

Ensure Kafka is installed and running on your local machine. You can install Kafka using Homebrew:

```sh
brew install kafka
```

Start Kafka (assuming Zookeeper is running on default settings):

```sh
zookeeper-server-start /usr/local/etc/kafka/zookeeper.properties & kafka-server-start /usr/local/etc/kafka/server.properties
```

### Building the Service

Navigate to the service directory and build the service:

```sh
go build
```

## Running the Service

Once built, you can run the service with:

```sh
./[service-name]
```

Replace `[service-name]` with the name of the generated executable.

## Usage

The service will start monitoring the directory specified in `watchDir`. When a new `.csv` file is detected, it will validate the file and send a message to Kafka if validation is successful. Non-csv files and invalid csv files will be logged but not sent to Kafka.

## Logging

The service provides logs for various events such as service start-up, file detection, file validation, and errors.

---

Ensure Kafka and Zookeeper are running before starting the service. Modify `constants.go` as per your setup requirements.

## TODO

Refactor code into packages. Inlcude unit tests.
