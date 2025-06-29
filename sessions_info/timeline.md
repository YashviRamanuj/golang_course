# Course Timeline

1. Go setup + syntax basics + Data structures
   - intro to golang
   - key data structures, int, float, string, bool, rune, slice, maps
   - Focus: Variables, control flow, functions
   - GOPATH, GOROOT, packages
   - Project: "Hello GoBlogr" CLI

2. Deep dive into string pkg && kick off project

3. Structs, Interfaces, Error Handling
   - Focus: Define models like `User`, `Post`
   - Project: Domain models

4. Concurrency
   - Focus: Goroutines, channels, WaitGroup
   - Project: Simple concurrent email sender (mock)

5. HTTP servers & REST APIs
   - Focus: net/http, routing, JSON
   - Project: CRUD for posts (in-memory)

6. Web framework (Gin)
   - Focus: Router, middleware, validation
   - Project: Basic blog API with validation

7. Database
   - Focus: PostgreSQL/neon + GORM/sqlx, migrations
   - Project: Persistent posts and users

8. Auth
   - Focus: JWTs, password hashing (bcrypt), middleware
   - Project: Secure login/signup

9. Testing + Logging
   - Focus: `testing`, Testify, Logrus/Zap
   - Project: Unit tests, structured logs

10.*Project Structuring
    - Focus: Packages, clean architecture
    - Project: Refactored codebase

11.* Docker & Deployment
    - Focus: Dockerfile, Compose, railway/vercel backend hosting
    - Project: Deployed API

12.* GenAI Tools (Intro)
    - Focus: OpenAI/[gemini] Go SDK, basic prompt-based generation
    - Project: Endpoint: `/ai/draft`

13.* GenAI Features (Advanced)
    - Focus: Prompt tuning, context handling
    - Project: Endpoint: `/ai/suggest-tags`

14.* Docs & Final Presentation
    - Focus: README, OpenAPI docs, deployment instructions
    - Project: Polished repo + demo-ready


## pending

1. error handling and panic and ok in golang