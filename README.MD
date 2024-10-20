# **FiberJobs**

A simple and efficient web scraper built in **Go** with the **Fiber framework**, collecting and storing **remote job data** from **UN Jobs** in **MongoDB**. This project demonstrates **HTTP scraping with proxy support**, **error handling with retries and backoff**, and **rotating proxies** to avoid IP blocking.

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

## **Test the Scraper**
```bash
curl http://localhost:8080/scrape-jobs
```
