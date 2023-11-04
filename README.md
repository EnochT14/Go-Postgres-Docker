
## Prerequisites

Before running this application, you need to have the following prerequisites:

1. Go installed on your system.
2. A PostgreSQL database with the necessary credentials and a "sensor_data" table to store the data.

## Installation

1. Clone this repository to your local machine.
2. Modify the database connection information in the `const` section of the `main.go` file according to your database configuration.

   ```go
   const (
       host          = "192.168.100.10" // Host IP
       port          = 5432
       user          = "postgres"      // PostgreSQL username
       password      = "password"      // PostgreSQL password
       dbname        = "postgres"      // Database name
       listenAddress = ":8080"
   )
   ```

3. Install the required Go packages by running:

   ```bash
   go get -u github.com/gorilla/mux
   go get -u github.com/lib/pq
   go get -u github.com/rs/cors
   ```

4. Build and run the application:

   ```bash
   go build
   ./your-application-name
   ```

