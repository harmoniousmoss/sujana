## **Sujana – Scalable Web Scraper for Remote Job Data**

Sujana is a lightweight web scraper built to collect remote job listings from the UN Jobs portal. It features proxy rotation, retry logic with backoff, and resilient error handling to ensure reliable data collection.
---

## **Case Study**
This scraper collects **remote job data** from [UN Jobs](https://unjobs.org/search/remote).

---

## **Features**
- **Go with Fiber Framework**: Fast and minimalist web framework.
- **MongoDB Integration**: Stores scraped data for persistence.
- **Rotating Proxies**: Bypasses IP restrictions with proxy rotation.
- **Retry Mechanism with Backoff**: Handles transient errors gracefully.
- **Exponential Backoff with Jitter**: Avoids overloading the server with retries.
- **Tested Proxy Connections**: Verifies proxies before using them.

---

## **Project Structure**
```bash
fiberjobs/
├── config/            # MongoDB connection logic
├── handlers/          # Fiber route handlers
├── scraper/           # Scraper logic with proxy rotation
├── routes/            # API routes registration
├── .env               # Environment variables
├── go.mod             # Go module dependencies
├── main.go            # Application entry point
└── README.md          # Project documentation
```

---

## **ENV Example**
```bash
MONGODB_URI=
PROXY_URL=
```

---

## **Test the Scraper**
```bash
curl http://localhost:8080/scrape-jobs
```
