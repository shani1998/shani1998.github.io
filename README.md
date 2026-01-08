# Portfolio Website

Personal portfolio site built with Go and static HTML/CSS.

## Running Locally

```bash
# Install dependencies
go mod download

# Run the server
go run main.go

# Open http://localhost:8080
```

## Deployment

The site runs on any platform that supports Go (Heroku, Render, Fly.io, etc.). Set the `PORT` environment variable if needed.

```bash
# Build
go build -o server main.go

# Run
./server
```

## Project Structure

```
├── main.go           # Go server (serves static files + email endpoint)
├── index.html        # Main portfolio page
├── static/           # CSS, JS, images, fonts
└── templates/        # Additional pages (cloud, ml, webdev)
```

## Contact Form

The contact form sends emails via Gmail SMTP. Set these environment variables:

```bash
SENDER_EMAIL=your-email@gmail.com
SENDER_PASSWORD=your-app-password
```

## License

MIT — feel free to fork and customize.