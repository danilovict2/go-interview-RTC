# Video Calling Interview Platform

A platform designed for conducting technical interviews with integrated video calling capabilities.

## 🚀 Quick Start

Ensure you have either [Docker](https://www.docker.com/get-started) or the following installed and configured:

- [Go](https://golang.org/doc/install)
- [Node.js](https://nodejs.org/)
- [PostgreSQL](https://www.postgresql.org/download/)

### Clone the project

```bash
git clone https://github.com/danilovict2/go-interview-RTC.git
cd go-interview-RTC
```

### Set environment variables

```bash
cp .env.example .env
cp /assets/vue/.env.example .env
```

### Create a Stream Account

1. Sign up for a [Stream](https://getstream.io/) account.
2. Obtain your API keys from the Stream dashboard.
3. Open the `.env` files and paste the keys into the corresponding fields.

### Run with Docker

```bash
docker compose up --build
```

### Run locally

First, open the `.env` file and comment out the existing `DATABASE_DSN` line. Uncomment the alternative `DATABASE_DSN` line. Ensure you have created the `interview-rtc` database in PostgreSQL.

Then, run the following command:

```bash
make run
```

### Access the Platform

Open your web browser and navigate to [http://127.0.0.1:8000](http://127.0.0.1:8000) to access the video calling interview platform.


## ✨ Features

- 📹 Video Calls.
- 💻 Real-time code collaboration with syntax highlighting.
- 🌐 Support for multiple programming languages.
- 🔐 User authentication and authorization.
- 🎥 Session recording and playback.
- ⚙️ Easy setup with Docker and local development options.
- 🌊 Stream API integration.
- 🎨 Styling with Tailwind & Shadcn.

## 🤝 Contributing

### Build the project

```bash
npm run --prefix assets/vue build
go build -o bin/app
```

### Run the project

```bash
./bin/app
```

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.