📦 Distributed Content Moderation System

A scalable, modular, and intelligent content moderation system designed to analyze and flag inappropriate content including text, images, and videos using an event-driven microservices architecture.


🚀 Features
🧠 AI-Powered Moderation (text, image, video)

⚙️ Microservices Architecture for scalability and maintainability

📩 Event-driven communication with Kafka or RabbitMQ

📊 Moderator Dashboard to review flagged content

🔁 Feedback loop & retraining for continuous model improvement

🔐 API Gateway for centralized authentication and request routing

🗃️ Object Storage Integration for media files (e.g. AWS S3 o Cloudinary)


📐 High-Level Architecture
+------------------------+         +------------------------+
|    User Device/Client  | <-----> |       API Gateway       |
+------------------------+         +------------------------+
                                        |        ^
                                        v        |
                          +-------------------------------+
                          |      Content Upload Service   |
                          +-------------------------------+
                                        |        
                                 (Event) |
                                        v        
                        +--------------------------+   
                        |     Kafka/RabbitMQ       |
                        +--------------------------+
                             |          |          |
                          Text        Image/Video  Flagging
                          Moderation  Moderation   & Review
                             |          |           |
                             v          v           |
                       +-------------------------+  |
                       |   Text Analysis Service |  |
                       +-------------------------+  |
                       |  Image/Video Analysis   |  |
                       |      Service            |  |
                       +-------------------------+  |
                                                     v
                                +-----------------------------+
                                |   Content Flagging Service  |
                                +-----------------------------+
                                                     |
                                                     v
                                +-----------------------------+
                                |  Moderator Review Service   |
                                +-----------------------------+
                                                     |
                                                     v
                                +-----------------------------+
                                |     Moderation Dashboard    |
                                +-----------------------------+


🧱 Tech Stack
Backend: Go (Golang for all services)

Event Streaming: Kafka / RabbitMQ

Storage: AWS S3 / Google Cloud Storage / Cloudinary

Database: PostgreSQL / MongoDB

Authentication: JWT / OAuth

Containerization: Docker + Kubernetes

Dashboard: React (Vite)


🛠️ Microservices

Service	                      Description
--------                      -------------
API Gateway	                  Central entry point for routing and security
Content Upload Service	      Handles user uploads and triggers moderation events
Text Moderation Service	      Uses NLP to detect spam, hate speech, profanity
Image Moderation Service	    Uses CV models or APIs to detect explicit imagery
Video Moderation Service	    Processes videos for inappropriate content
Flagging Service	            Stores and tracks flagged content for moderator review
Moderation Dashboard	        Web UI for moderators to approve or reject flagged content
Retraining Service	          Improves models based on moderator feedback

🌐 API Gateway
The API Gateway is the central entry point for all client-facing requests. It abstracts internal microservices, manages authentication and authorization, enforces rate limits, and routes requests appropriately.

🔧 Responsibilities
🔐 Authentication & Authorization
Validates JWTs or OAuth tokens to ensure only authorized users and services can access the system.

🚦 Request Routing
Directs incoming HTTP requests to the appropriate microservices (e.g., content-upload, moderation services).

📊 Rate Limiting & Throttling
Protects services from abuse and ensures fair usage.

📈 Logging & Metrics
Captures structured logs and exposes metrics for observability.

🧪 Request Validation
Optionally validate payloads and parameters before routing.

🛠 Tech Stack
Language: Go

Framework: Gin or Fiber

Auth: JWT middleware

Reverse Proxy: Built-in routing or external (optional: Kong/Nginx)


# Clone the repo
git clone https://github.com/thoraf20/content-moderation-system.git

# Run using Docker Compose
cd content-moderation-system
docker-compose up --build

⚙️ Setup Instructions

Install dependencies:

Go

Docker & Docker Compose

Kafka or RabbitMQ

Configure environment variables:

Copy .env.example to .env in each service folder and fill in the required values.

Run all services:

docker-compose up --build

Access the system:

Dashboard UI: http://localhost:3000

API Gateway: http://localhost:8080


🧪 Sample API Routes (API Gateway)

Method

Endpoint

Description

POST

/api/upload

Upload content for moderation

GET

/api/content/:id/status

Get moderation status of content

GET

/api/flags

Retrieve all flagged content


🧪 Testing Strategy

Unit Tests: Per service using Go's built-in testing framework

Integration Tests: Ensure message passing between services via Kafka

End-to-End Tests: Simulate upload → moderation → flagging

Test Tools: Go test, Postman/Newman, Mockery