# Course Timeline

1. **Week 1: Go setup + syntax basics + Data structures + error handling**
   - intro to golang
   - key data structures, int, float, string, bool, rune, slice, maps
   - Focus: Variables, control flow, functions
   - GOPATH, GOROOT, packages
   - Project: "Hello GoBlogr" CLI

2. **Week 2: Structs & Interfaces**
   - Focus: Define models like `User`, `Post`
   - Project: Domain models

3. **Week 3: Concurrency**
   - Focus: Goroutines, channels, WaitGroup
   - Project: Simple concurrent email sender (mock)

4. **Week 4: HTTP servers & REST APIs**
   - Focus: net/http, routing, JSON
   - Project: CRUD for posts (in-memory)

5. **Week 5: Web framework (Gin)**
   - Focus: Router, middleware, validation
   - Project: Basic blog API with validation

6. **Week 6: Database**
   - Focus: PostgreSQL/neon + GORM/sqlx, migrations
   - Project: Persistent posts and users

7. **Week 7: Auth**
   - Focus: JWTs, password hashing (bcrypt), middleware
   - Project: Secure login/signup

8. **Week 8: Testing + Logging**
   - Focus: `testing`, Testify, Logrus/Zap
   - Project: Unit tests, structured logs

9. **Week 9: Project Structuring**
    - Focus: Packages, clean architecture
    - Project: Refactored codebase

10. **Week 10: Docker & Deployment**
    - Focus: Dockerfile, Compose, railway/vercel backend hosting
    - Project: Deployed API

11. **Week 11: GenAI Tools (Intro)**
    - Focus: OpenAI/[gemini] Go SDK, basic prompt-based generation
    - Project: Endpoint: `/ai/draft`

12. **Week 12: GenAI Features (Advanced)**
    - Focus: Prompt tuning, context handling
    - Project: Endpoint: `/ai/suggest-tags`

13. **Week 13: Docs & Final Presentation**
    - Focus: README, OpenAPI docs, deployment instructions
    - Project: Polished repo + demo-ready
