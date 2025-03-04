# 🚀 User Profile Management System - Microservices Architecture

## 📌 Overview
The **User Profile Management System** is a distributed **microservices-based architecture** for managing user profiles, updates, and authentication securely. Each microservice is **independent**, deployed separately, and communicates using **REST, GraphQL, WebSocket, WebHook, RPC, and SOAP**.

## 🏗 Microservices Overview

### 1️⃣ List User Profiles Microservice
- **Purpose:** Retrieve user profiles with filtering, sorting, and pagination.
- **Security:** JWT authentication, CORS, rate limiting, DDoS protection.
- **Communication:** REST API, GraphQL API.

### 2️⃣ Update User Email Microservice
- **Purpose:** Securely updates user email addresses with verification.
- **Security:** JWT authentication, email encryption, rate limiting.
- **Communication:** REST API, WebHook.

### 3️⃣ Update User Password Microservice
- **Purpose:** Handles password updates with encryption and brute-force protection.
- **Security:** JWT authentication, AES & bcrypt encryption, IP blocking.
- **Communication:** REST API, RPC.

### 4️⃣ Update Username Microservice
- **Purpose:** Updates and validates unique usernames.
- **Security:** JWT authentication, rate limiting, uniqueness enforcement.
- **Communication:** REST API, SOAP.

### 5️⃣ View User Profile Microservice
- **Purpose:** Retrieves user profile data.
- **Security:** JWT authentication, encryption, rate limiting.
- **Communication:** REST API, WebSocket, GraphQL.

## 🔐 Security Measures
- **Authentication:** JWT-based authorization.
- **Data Protection:** Encryption for sensitive information.
- **DDoS & EDoS Mitigation:** Rate limiting, IP blocking.
- **CORS Implementation:** Controlled cross-origin access.
- **API Documentation:** Integrated **Swagger UI**.

## 🛠 Deployment & CI/CD
Each microservice is **independently deployed** on **separate EC2 instances** using **GitHub Actions**:
- **Automated builds & testing** via CI/CD pipelines.
- **Isolated deployment** for scalability and fault tolerance.
