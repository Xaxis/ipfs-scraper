# ipfs-scraper

# Setting up Docker and Running the App

## Prerequisites
- Docker and Docker Compose installed on your system.

## Steps to Run (First Time)

1. **Clone the Repository** (optional if you have the project files)
   ```bash
   git clone https://github.com/Xaxis/ipfs-scraper
   cd ipfs-scraper
   ```
   
2. **Build the Docker Image & Run the Containers**
   ```bash
    docker-compose up --build
    ```
   
3. **Verify the Application is Running**
   - Ensure that both the application and database containers are running.
   - The API will be accessible at http://localhost:8080.

4. **Interact with the Database**
   - The database is accessible at `localhost:5432` with the following credentials:
     - Username: `postgres`
     - Password: `password`
     - Database: `ipfs_scraper`

5. **Shutting Down**
   ```bash
   docker-compose down
   ```
   
## Steps to Run (Subsequent Times)

1. **Clear volumes** (run this if you want to start fresh)
   ```bash
   docker-compose down -v
   ```
   
2. **Start the Containers**
   ```bash
   docker-compose up --build
   ```