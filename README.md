# 🔐 Generate & Validate OTP With Go (Twilio)

A simple Go-based service to **send and verify OTPs via SMS** using the [Twilio Verify API](https://www.twilio.com/docs/verify). Ideal for phone number authentication and two-factor verification flows.

---

## ✨ Features

- ✅ Send OTP to phone via SMS
- ✅ Verify OTP submitted by the user
- ✅ Twilio Verify integration
- ✅ Config-driven with environment variables
- ✅ Minimal and clean REST API in Go

---

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/ParasJain0307/Generate-Validate-OTP-With-Go.git
cd Generate-Validate-OTP-With-Go

2. Install Dependencies
 go mod tidy

3. Configure Environment Variables
Create a .env file in the root directory:

TWILIO_ACCOUNT_SID=your_twilio_account_sid
TWILIO_AUTHTOKEN=your_twilio_auth_token
TWILIO_SERVICES_ID=your_twilio_verify_service_id
🔒 Do NOT commit this file to version control. Add .env to your .gitignore.


⚙️ API Endpoints

✅ POST /otp
Send an OTP to the given phone number.
Request Body:
    {
     "phone": "+919999999999"
    }

    {
     "sid": "VE1234567890abcdef" 
    }

✅ POST /verifyOTP

Verify the OTP for the given phone number.
Request Body:
   {
    "phone": "+919999999999",
    "code": "123456"
   }

Success Response:
  {
    "message": "OTP verified successfully"
  }

Error Response:
  {
    "error": "OTP verification failed"
  }


🧪 Testing With cURL

# Send OTP
curl -X POST http://localhost:8080/otp \
  -H "Content-Type: application/json" \
  -d '{"phone": "+919999999999"}'

# Verify OTP
curl -X POST http://localhost:8080/verifyOTP \
  -H "Content-Type: application/json" \
  -d '{"phone": "+919999999999", "code": "123456"}'


📁 Project Structure

├── api/
│   ├── config.go        # Configuration and Twilio client setup
│   ├── router.go        # API routes and HTTP server
│   ├── services.go      # Twilio OTP send/verify logic
│   ├── helper.go        # Utility functions (formatting, etc.)
│   ├── handler.go       # HTTP handlers using the services
│
├── data/
│   └── models.go        # Request/response structs and domain models
│
├── cmd/
│   ├── main.go          # Application entry point (initialization, launch)
│   └── .env             # Local environment variables (not committed)
│
├── go.mod / go.sum      # Go module and dependencies
├── LICENSE              # MIT License file
└── README.md            # Project documentation


⚠️ Notes
  Use phone numbers in E.164 format (e.g., +91XXXXXXXXXX)

  OTPs are single-use

  Codes expire after a short period (as configured in Twilio)


🛡️ Security Best Practices
✅ Never commit .env or secret keys

✅ Use HTTPS in production

✅ Rotate your Twilio Auth Token regularly

👨‍💻 Author
    Paras Jain

