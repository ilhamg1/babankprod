# Babank: AI-Powered Quiz Generator Development Roadmap

This roadmap contains exactly 1000 atomic, uniformly weighted tasks. Each task represents roughly 4-6 hours of development effort, enabling highly accurate Agile velocity tracking.

## Epic 1: Workspace, CI/CD, and GCP Foundation

### Bazel Monorepo Initialization
*   Step 1: Initialize the root workspace directory and git repository
*   Step 2: Create the WORKSPACE.bazel file defining the root monorepo boundaries
*   Step 3: Install and configure Bazelisk as the global Bazel wrapper
*   Step 4: Define the root package.json for global Node dependencies
*   Step 5: Configure .gitignore for Bazel cache (`bazel-*`) and Node modules
*   Step 6: Create the standardized `apps/` and `libs/` directory structure
*   Step 7: Set up the `.bazelrc` for shared build flags and remote caching configurations
*   Step 8: Write the initial root BUILD.bazel file
*   Step 9: Configure strict cross-platform line ending rules in `.gitattributes`
*   Step 10: Verify monorepo structural integrity by running a baseline `bazel build //...`

### Angular Workspace Configuration
*   Step 11: Provision the `@babank/web` Angular application via the Angular CLI
*   Step 12: Eject or wrap the Angular CLI configuration into Bazel rules (`@bazel/concatjs`)
*   Step 13: Configure `angular.json` to map output directories into the Bazel `dist/` cache
*   Step 14: Establish standard TypeScript configurations (`tsconfig.json`) for strict mode compliance
*   Step 15: Configure angular-eslint with Google TypeScript Style (gts) rules
*   Step 16: Set up the Jasmine/Karma testing framework within the Bazel build graph
*   Step 17: Configure standard environment variable injection (dev, staging, prod)
*   Step 18: Write the initial `app.component.ts` shell with baseline routing
*   Step 19: Create Bazel macros for repetitive Angular component generation tasks
*   Step 20: Verify the Angular app builds and serves locally using `bazel run //apps/web:serve`

### Go Backend Workspace Configuration
*   Step 21: Initialize the root `go.mod` file for the Go backend services
*   Step 22: Integrate `rules_go` and `gazelle` into the root WORKSPACE.bazel
*   Step 23: Run Gazelle to automatically generate BUILD.bazel files for Go packages
*   Step 24: Establish the standard Go project layout (`cmd/`, `internal/`, `pkg/`)
*   Step 25: Configure `golangci-lint` for strict Go static analysis and styling
*   Step 26: Write a baseline `main.go` HTTP server for the core API entrypoint
*   Step 27: Configure structured JSON logging using standard Go libraries (e.g., `slog`)
*   Step 28: Set up Graceful Shutdown logic handling SIGINT and SIGTERM signals
*   Step 29: Write basic unit tests for the Go server boot sequence
*   Step 30: Verify the Go binary compiles and runs using `bazel run //apps/api:api`

### GitHub Repository Setup
*   Step 31: Provision a new repository in GitHub
*   Step 32: Configure local git remotes to point to the GitHub repository
*   Step 33: Add collaborators and configure repository access permissions
*   Step 34: Set up strict branch protection rules enforcing PR reviews on `main`
*   Step 35: Require passing status checks (build, test, lint) before merging
*   Step 36: Create standard Issue Templates (Bug, Feature) in the repository
*   Step 37: Create standard Pull Request templates requiring test coverage summaries
*   Step 38: Document the branching strategy (Feature Branching) in `CONTRIBUTING.md`
*   Step 39: Configure automated branch deletion rules post-merge
*   Step 40: Verify push and pull functionality with GitHub credentials

### Google Cloud Build CI Pipeline
*   Step 41: Create the foundational `cloudbuild.yaml` configuration file
*   Step 42: Define the Cloud Build trigger for pull requests targeting `main`
*   Step 43: Configure the CI pipeline step to run `bazel test //...`
*   Step 44: Configure the CI pipeline step to run `angular-eslint` and `golangci-lint`
*   Step 45: Set up Google Cloud Storage bucket for Bazel remote build caching
*   Step 46: Authenticate Cloud Build service account to access the remote cache bucket
*   Step 47: Configure CI steps to publish test coverage reports as build artifacts
*   Step 48: Implement a Slack webhook notification for failed CI builds
*   Step 49: Implement a Slack webhook notification for successful deployments
*   Step 50: Verify the end-to-end CI pipeline execution by submitting a test PR

### Artifact Registry and Docker Config
*   Step 51: Provision a Docker repository in Google Artifact Registry (GAR)
*   Step 52: Configure IAM roles allowing Cloud Build to push to GAR
*   Step 53: Write a multi-stage `Dockerfile` for the Go backend API
*   Step 54: Write a multi-stage `Dockerfile` for the Angular frontend (Nginx/Node)
*   Step 55: Integrate `rules_docker` (or `rules_oci`) into the Bazel workspace
*   Step 56: Configure Bazel to build optimized, distroless container images
*   Step 57: Add a Cloud Build step to tag images with the short Git commit SHA
*   Step 58: Add a Cloud Build step to push built images to the GAR repository
*   Step 59: Configure GAR lifecycle policies to clean up untagged or old images
*   Step 60: Verify image building and pushing via the Cloud Build UI

### Google Cloud IAM & Service Accounts
*   Step 61: Create the principle `babank-backend-sa` service account for Cloud Run
*   Step 62: Create the principle `babank-agones-sa` service account for Game Servers
*   Step 63: Create the principle `babank-ci-sa` service account for Cloud Build
*   Step 64: Assign principle of least privilege (PoLP) roles to the backend SA (Spanner read/write)
*   Step 65: Assign PoLP roles to the CI SA (Artifact Registry Writer, Logs Writer)
*   Step 66: Configure Workload Identity Federation for local developer authentication
*   Step 67: Audit existing project IAM bindings to remove overly broad permissions
*   Step 68: Set up Google Cloud Key Management Service (KMS) for secret encryption
*   Step 69: Document the exact IAM role mappings in the infrastructure runbook
*   Step 70: Verify service account token generation and impersonation via `gcloud`

### Security Command Center Vulnerability Scanning
*   Step 71: Enable Google Cloud Security Command Center (SCC) Standard/Premium tier
*   Step 72: Configure Artifact Registry automated vulnerability scanning for pushed images
*   Step 73: Set up SCC notifications targeting a dedicated security Slack channel
*   Step 74: Configure Bazel vulnerability scanning for third-party Go modules
*   Step 75: Configure `npm audit` integration within the CI pipeline to block builds on high CVEs
*   Step 76: Establish IAM boundaries preventing unauthorized modification of SCC policies
*   Step 77: Write a script to generate Software Bill of Materials (SBOM) during builds
*   Step 78: Configure automated patching policies for minor Node.js runtime updates
*   Step 79: Review baseline SCC findings for the initial project scaffolding
*   Step 80: Document the incident response protocol for critical zero-day alerts

### Developer Environment Emulators
*   Step 81: Install the Google Cloud SDK and standard emulators components locally
*   Step 82: Configure the Google Cloud Spanner local emulator Docker container
*   Step 83: Configure the Firebase Local Emulator Suite (Auth, Realtime DB)
*   Step 84: Write a `start-emulators.sh` script to boot all local infrastructure
*   Step 85: Configure local backend `.env` variables to point strictly to emulators
*   Step 86: Configure Angular `environment.development.ts` to target emulator ports
*   Step 87: Write a Bazel macro to seed the Spanner emulator with test data on boot
*   Step 88: Write a Bazel macro to seed Firebase Auth with dummy user accounts
*   Step 89: Document emulator bootstrapping processes in `README.md`
*   Step 90: Verify local end-to-end connectivity without touching production GCP APIs

### Cloud Monitoring & Telemetry Foundation
*   Step 91: Enable the Google Cloud Monitoring and Trace APIs
*   Step 92: Integrate the Google Cloud OpenTelemetry Go SDK into the backend
*   Step 93: Configure distributed tracing middleware for all incoming HTTP requests
*   Step 94: Set up custom metrics for critical business events (e.g., Match Completed)
*   Step 95: Configure structured payload logging to Google Cloud Logging
*   Step 96: Create a standard Cloud Monitoring Dashboard for API latency and error rates
*   Step 97: Create a standard Cloud Monitoring Dashboard for infrastructure utilization (CPU/Memory)
*   Step 98: Set up alerting policies for API error rates exceeding 1% over 5 minutes
*   Step 99: Set up alerting policies for Spanner CPU utilization exceeding 70%
*   Step 100: Verify telemetry ingestion by generating synthetic traffic locally

## Epic 2: Core Database Schema & Spanner Setup

### Spanner Instance Provisioning
*   Step 101: Provision a regional Google Cloud Spanner instance (e.g., us-central1)
*   Step 102: Configure Spanner processing units (PUs) tailored for initial launch scale
*   Step 103: Establish IAM roles restricting Spanner administrative access to DevOps
*   Step 104: Grant the backend service account `spanner.databaseUser` permissions
*   Step 105: Create the initial `babank-prod` database within the Spanner instance
*   Step 106: Configure Spanner point-in-time recovery (PITR) data retention policies
*   Step 107: Set up Cloud Monitoring alerts specifically for Spanner storage utilization
*   Step 108: Document the Spanner instance connection topology and connection limits
*   Step 109: Verify connectivity to the production instance using the GCP Console
*   Step 110: Verify connectivity to the production instance using the Go Spanner Client locally

### Global User Profiles Schema
*   Step 111: Design the Spanner DDL schema for the `Users` table (ID, Email, DisplayName)
*   Step 112: Design the DDL schema for user statistics (MatchesPlayed, WinRate, GlobalElo)
*   Step 113: Define Protobuf schemas representing the User Profile entity for API responses
*   Step 114: Implement Go backend CRUD operations utilizing Spanner mutations
*   Step 115: Implement exact-match secondary indexes for `Email` lookups in Spanner
*   Step 116: Implement prefix-search capabilities for `DisplayName` lookups
*   Step 117: Write Spanner read-write transactions for updating user statistics safely
*   Step 118: Write Go unit tests mocking the Spanner client for user profile logic
*   Step 119: Create initial database migration scripts for the Users schema
*   Step 120: Execute the DDL migration against the local Spanner emulator

### Match History and Statistics Schema
*   Step 121: Design the Spanner DDL schema for the `Matches` table (MatchID, Mode, Timestamp)
*   Step 122: Design the Spanner DDL schema for the `MatchParticipants` table (MatchID, UserID, Score, Placement)
*   Step 123: Define interleaved tables in Spanner to co-locate `MatchParticipants` physically with `Matches`
*   Step 124: Define Protobuf schemas representing Match History records
*   Step 125: Implement Go backend pagination logic for fetching a user's match history
*   Step 126: Implement Spanner transactions to commit match results atomically across all participants
*   Step 127: Create secondary indexes optimizing leaderboards by grouping historical match data
*   Step 128: Write Go integration tests ensuring match data integrity during concurrent inserts
*   Step 129: Create the DDL migration scripts for the Match History schemas
*   Step 130: Execute the DDL migration against the local Spanner emulator

### Study Groups (Guilds) Schema
*   Step 131: Design the Spanner DDL schema for the `StudyGroups` table (GroupID, Name, Emblem, Description)
*   Step 132: Design the Spanner DDL schema for the `StudyGroupMembers` table (GroupID, UserID, Role)
*   Step 133: Define interleaved tables in Spanner to co-locate members with their group
*   Step 134: Define Protobuf schemas representing Study Group entities
*   Step 135: Implement Go backend logic for Guild creation and Spanner transaction limits
*   Step 136: Implement Spanner queries to resolve a user's active Study Group status
*   Step 137: Implement Go logic to handle Group Member roles (Admin, Member) and permissions
*   Step 138: Write Spanner transactions handling atomic group joins and capacity constraints
*   Step 139: Create the DDL migration scripts for the Study Group schemas
*   Step 140: Execute the DDL migration against the local Spanner emulator

### Social Graphs (Friends & Blocks) Schema
*   Step 141: Design the Spanner DDL schema for the `Friendships` table (UserID1, UserID2, Status)
*   Step 142: Design the Spanner DDL schema for the `BlockList` table (BlockerID, BlockedID)
*   Step 143: Define Protobuf schemas representing social graph relationships
*   Step 144: Implement Go backend logic to dispatch friend requests and pending states
*   Step 145: Implement Spanner transactions ensuring bidirectional friendship states remain consistent
*   Step 146: Implement Go logic blocking interactions between users on the BlockList
*   Step 147: Optimize Spanner queries to rapidly fetch online friends using index scans
*   Step 148: Write Go integration tests covering complex block/friend race conditions
*   Step 149: Create the DDL migration scripts for the Social Graph schemas
*   Step 150: Execute the DDL migration against the local Spanner emulator

### Virtual Economy (Coins & Gems) Schema
*   Step 151: Design the Spanner DDL schema for the `Wallets` table (UserID, Coins, Gems)
*   Step 152: Design the Spanner DDL schema for the `Transactions` ledger (TxID, UserID, Amount, Type, Timestamp)
*   Step 153: Define Protobuf schemas representing financial balances and transaction logs
*   Step 154: Implement strict Spanner read-write transactions ensuring wallet balances cannot drop below zero
*   Step 155: Implement Go logic to generate idempotent transaction IDs preventing duplicate charges
*   Step 156: Implement Spanner queries to aggregate transaction history for user receipts
*   Step 157: Write Go integration tests simulating high-concurrency wallet deductions (race condition testing)
*   Step 158: Design Spanner IAM policies strictly limiting wallet modification access to specific internal services
*   Step 159: Create the DDL migration scripts for the Virtual Economy schemas
*   Step 160: Execute the DDL migration against the local Spanner emulator

### Leaderboards & Leagues Schema
*   Step 161: Design the Spanner DDL schema for the `Leaderboards` table (LeagueID, UserID, Rank, Points)
*   Step 162: Design the Spanner DDL schema for the `Leagues` metadata table (LeagueID, Name, PromotionThreshold)
*   Step 163: Define Protobuf schemas representing leaderboard rankings and league status
*   Step 164: Implement Spanner queries utilizing `ORDER BY Points DESC LIMIT N` for rapid top-100 retrieval
*   Step 165: Implement Go backend logic caching top leaderboard results in memory to reduce Spanner reads
*   Step 166: Implement Spanner batch updates for weekly league resets and promotions
*   Step 167: Write Go unit tests validating the league promotion and demotion mathematical logic
*   Step 168: Optimize Spanner secondary indexes specifically for the `Points` column
*   Step 169: Create the DDL migration scripts for the Leaderboard schemas
*   Step 170: Execute the DDL migration against the local Spanner emulator

### Database Migration Automation
*   Step 171: Select and configure a Go-native database migration tool compatible with Cloud Spanner (e.g., `golang-migrate`)
*   Step 172: Establish the `migrations/` directory storing sequential SQL DDL scripts
*   Step 173: Write a Bazel rule to execute database migrations against local emulators automatically
*   Step 174: Configure Cloud Build to execute migration scripts against staging databases pre-deployment
*   Step 175: Configure Cloud Build to execute migration scripts against production databases during deployment windows
*   Step 176: Implement Go logic to verify schema version compatibility before server boot
*   Step 177: Document the process for writing and testing backward-compatible migrations
*   Step 178: Test a rollback migration script on the local emulator
*   Step 179: Configure IAM roles ensuring Cloud Build can apply DDL changes safely
*   Step 180: Verify the automated migration pipeline on the GCP environment

### Spanner Backup & Disaster Recovery
*   Step 181: Configure scheduled automated backups for the Spanner instance (e.g., daily at 2AM UTC)
*   Step 182: Define backup retention policies (e.g., 30 days) within Google Cloud
*   Step 183: Document the manual backup trigger process in the disaster recovery runbook
*   Step 184: Document the exact steps required to restore a Spanner database from a backup
*   Step 185: Perform a dry-run restoration of a backup into a temporary development database
*   Step 186: Configure Cloud Monitoring alerts for failed automated backup jobs
*   Step 187: Implement a Go script to export critical, anonymized analytical data to Cloud Storage periodically
*   Step 188: Review GCP quotas to ensure backup storage limits accommodate growth
*   Step 189: Establish IAM policies restricting backup deletion capabilities
*   Step 190: Conduct a full disaster recovery tabletop exercise with the engineering team

### Database Connection Pooling & Configuration
*   Step 191: Configure the Go Spanner client `NumChannels` settings for optimal gRPC multiplexing
*   Step 192: Configure the Go Spanner client `MaxSessions` pool size based on expected instance concurrency
*   Step 193: Implement logic to gracefully close Spanner sessions during server shutdown
*   Step 194: Configure Spanner client timeouts and exponential backoff retry logic
*   Step 195: Instrument the Spanner connection pool with OpenTelemetry metrics (in-use sessions, waits)
*   Step 196: Create a Grafana/Cloud Monitoring dashboard specifically tracking session exhaustion
*   Step 197: Conduct local load testing to determine the optimal session pool baseline
*   Step 198: Implement Go context deadlines (Context.WithTimeout) for all Spanner database calls
*   Step 199: Document the connection pooling architecture and tuning parameters
*   Step 200: Deploy connection pool configurations to staging and monitor stability under load

## Epic 3: Authentication, Users, and Profiles

### Firebase Authentication Setup
*   Step 201: Provision a Firebase project and link it to the primary Google Cloud project
*   Step 202: Enable Email/Password and Identity Platform authentication providers in the Firebase console
*   Step 203: Configure the authorized authorized domains for production and staging environments
*   Step 204: Install the Firebase Admin SDK within the Go backend workspace
*   Step 205: Install the Firebase Web SDK within the Angular frontend workspace
*   Step 206: Initialize the Firebase client configuration in the Angular `environment.ts`
*   Step 207: Initialize the Firebase Admin client using GCP Application Default Credentials in Go
*   Step 208: Configure standard CORS headers for Firebase Auth API interactions
*   Step 209: Create the foundation Angular Authentication Service wrapping the Firebase SDK
*   Step 210: Verify initial frontend-to-Firebase connectivity in the local environment

### Email/Password Registration Flow
*   Step 211: Define Protobuf schemas for custom registration payloads (e.g., DisplayName capture)
*   Step 212: Implement the Angular Material registration form with reactive form validations
*   Step 213: Wire the registration form to the Angular Firebase Auth Service
*   Step 214: Implement Go backend Cloud Functions or Webhooks to trigger on new user creation
*   Step 215: Write Go logic to intercept the Firebase creation event and generate a Spanner User record
*   Step 216: Implement NgRx actions and reducers to handle registration loading, success, and error states
*   Step 217: Bind the Angular Material form UI to the NgRx state selectors
*   Step 218: Implement error mapping (e.g., 'Email already in use') to user-friendly UI snackbars
*   Step 219: Write Jasmine tests verifying the reactive form validation logic
*   Step 220: Conduct end-to-end registration flow tests verifying Spanner synchronization

### OAuth (Google, GitHub) Integration
*   Step 221: Configure OAuth consent screens and client IDs in the Google Cloud Console
*   Step 222: Enable Google and GitHub sign-in providers within the Firebase Authentication console
*   Step 223: Implement the Angular Material UI buttons for social SSO logins
*   Step 224: Wire the SSO buttons to the Firebase `signInWithPopup` methods in the Auth Service
*   Step 225: Implement Go backend logic to handle Spanner synchronization for first-time SSO users
*   Step 226: Define NgRx actions and reducers handling the distinct OAuth redirect and popup flows
*   Step 227: Implement logic to merge accounts if an OAuth email matches an existing password account
*   Step 228: Write Jasmine tests mocking the OAuth provider responses
*   Step 229: Update the privacy policy and terms of service links on the SSO screens
*   Step 230: Conduct end-to-end testing of the Google OAuth flow in the staging environment

### JWT Token Validation & Middleware
*   Step 231: Implement Go backend middleware to intercept incoming HTTP requests and extract the Authorization header
*   Step 232: Utilize the Firebase Admin SDK to verify the signature and validity of incoming JWT tokens
*   Step 233: Implement Go logic to extract the UID from the validated JWT and attach it to the request Context
*   Step 234: Configure the Angular HTTP Interceptor to automatically attach the Firebase ID Token to outgoing requests
*   Step 235: Implement token refresh logic in the Angular interceptor using RxJS when the token expires
*   Step 236: Write Go unit tests bypassing the middleware for unauthenticated routes (e.g., /health)
*   Step 237: Write Go unit tests validating middleware rejection of expired or malformed JWTs
*   Step 238: Implement NgRx logic to log the user out if a token is permanently revoked or invalid
*   Step 239: Verify JWT validation latency via OpenTelemetry distributed tracing
*   Step 240: Document the authentication architecture and context injection patterns

### User Profile Editing (Bio, Avatar)
*   Step 241: Define Protobuf schemas for User Profile Update requests (Bio string, Avatar URL)
*   Step 242: Implement the Go backend API route and validation logic for profile updates
*   Step 243: Write Spanner mutations to safely update specific profile fields without overwriting statistics
*   Step 244: Define NgRx state slices and actions to optimistic UI updates for profile changes
*   Step 245: Implement the Angular Material profile editing form and image upload placeholder
*   Step 246: Wire the profile form to the Angular HTTP Service calling the Go backend
*   Step 247: Implement strict backend validation preventing profanity in bios or display names
*   Step 248: Implement backend rate limiting to prevent spamming profile updates
*   Step 249: Write Jasmine tests validating the UI form state during submission
*   Step 250: Conduct end-to-end testing ensuring profile updates immediately reflect on the user dashboard

### Password Reset & Account Recovery
*   Step 251: Enable the Firebase Password Reset template in the Firebase Console
*   Step 252: Implement the Angular Material 'Forgot Password' request form
*   Step 253: Wire the 'Forgot Password' form to the Firebase `sendPasswordResetEmail` method
*   Step 254: Implement the Angular routing and view for the custom email action handler (`/auth/action`)
*   Step 255: Implement the Angular Material form to capture the new password securely
*   Step 256: Wire the new password form to the Firebase `confirmPasswordReset` method
*   Step 257: Implement rate limiting logic on the Go backend to prevent email spamming
*   Step 258: Design the UI snackbar notifications for successful or expired reset links
*   Step 259: Write Jasmine unit tests validating password strength requirements on the reset form
*   Step 260: Perform end-to-end testing of the complete account recovery lifecycle

### Two-Factor Authentication (2FA)
*   Step 261: Upgrade the Firebase project to Identity Platform to enable Multi-Factor Authentication (MFA)
*   Step 262: Implement the Angular Material UI for linking a phone number via SMS verification
*   Step 263: Wire the SMS verification UI to the Firebase `MultiFactorResolver` flow
*   Step 264: Update the Angular login flow to detect `auth/multi-factor-auth-required` errors
*   Step 265: Implement the secondary login screen to capture the SMS verification code
*   Step 266: Write Go backend middleware to verify MFA claims in the Firebase JWT token
*   Step 267: Implement Spanner logic to track a user's MFA enrollment status for analytics
*   Step 268: Write Jasmine tests simulating the MFA challenge and resolution paths
*   Step 269: Implement recovery code generation and secure display upon initial MFA enrollment
*   Step 270: Conduct staging environment tests of the 2FA enrollment and subsequent login

### User Privacy & Visibility Settings
*   Step 271: Design the Spanner DDL schema for the `UserPreferences` table (UserID, ProfileVisibility, AllowDMs)
*   Step 272: Define Protobuf schemas for Privacy Settings update requests
*   Step 273: Implement the Angular Material settings dashboard with toggle switches for privacy options
*   Step 274: Wire the privacy toggles to the Go backend API to persist preferences in Spanner
*   Step 275: Implement Go backend logic preventing users from viewing profiles marked as 'Private'
*   Step 276: Implement Go backend logic rejecting Direct Messages if the recipient disabled them
*   Step 277: Update the Matchmaking Director to anonymize display names if 'Hide Name' is enabled
*   Step 278: Define NgRx state slices caching the user's local privacy preferences
*   Step 279: Write Go integration tests ensuring privacy rules are respected across all API endpoints
*   Step 280: Verify privacy settings persistence and enforcement in the staging environment

### GDPR & Account Deletion Flows
*   Step 281: Implement the Angular Material confirmation dialog for permanent account deletion
*   Step 282: Wire the deletion dialog to a strict Go backend API endpoint requiring password re-authentication
*   Step 283: Write Spanner transactions to mark the user as 'Deleted' (soft delete) and wipe PII data
*   Step 284: Implement Go logic to trigger Firebase Authentication user deletion
*   Step 285: Write a Google Cloud Task queue worker to purge orphaned database records asynchronously
*   Step 286: Write a script to scrub user IP logs and analytical events after 30 days
*   Step 287: Implement a 'Download My Data' export feature generating a JSON artifact of user history
*   Step 288: Secure the export feature behind a generated, short-lived presigned URL
*   Step 289: Write Go unit tests validating the scrubbing of PII data during soft deletion
*   Step 290: Document the exact GDPR compliance architecture for future audits

### Admin Role & Impersonation Logic
*   Step 291: Define the `Admin` boolean flag in the Spanner `Users` schema
*   Step 292: Implement Go backend middleware restricting specific routes to Admin users
*   Step 293: Develop an internal Angular 'Admin Dashboard' protected by route guards
*   Step 294: Implement backend logic allowing an Admin to 'impersonate' a user for debugging
*   Step 295: Implement strict audit logging in Go whenever an impersonation session begins
*   Step 296: Update the Angular HTTP Interceptor to handle the `X-Impersonate-User` header
*   Step 297: Implement the Angular UI banner alerting the Admin that impersonation is active
*   Step 298: Write Go integration tests ensuring standard users cannot access Admin APIs
*   Step 299: Configure Cloud Logging alerts to notify security if impersonation is abused
*   Step 300: Conduct a security review of the impersonation JWT generation logic

## Epic 4: AI Quiz Generation Service & Vertex AI

### Vertex AI Project Provisioning
*   Step 301: Enable the Vertex AI API within the primary Google Cloud project
*   Step 302: Provision a dedicated `babank-ai-sa` service account for Vertex AI interactions
*   Step 303: Grant the AI service account the `Vertex AI User` IAM role
*   Step 304: Integrate the official Google Cloud Vertex AI Go SDK into the backend workspace
*   Step 305: Configure regional settings ensuring Vertex AI requests stay within the target geography
*   Step 306: Create a Cloud Monitoring dashboard tracking Vertex AI quota usage and limits
*   Step 307: Write a basic Go script to authenticate and test a simple completion prompt
*   Step 308: Implement exponential backoff retry logic in Go for handling Vertex AI `429 Too Many Requests`
*   Step 309: Configure budget alerts in GCP Billing tied directly to Vertex AI consumption
*   Step 310: Document the baseline generative models (e.g., Gemini 1.5 Pro) to be used

### Generative Model Prompt Engineering
*   Step 311: Establish a `prompts/` directory in the Go backend to store version-controlled text prompts
*   Step 312: Write the core 'System Prompt' dictating the persona and rules of the Quiz Generator
*   Step 313: Define strict prompt constraints demanding JSON-only output from the model
*   Step 314: Define prompt logic requiring diverse question types (multiple choice, true/false)
*   Step 315: Write few-shot examples within the prompt context to demonstrate the expected JSON structure
*   Step 316: Implement prompt injection defenses within the system instructions
*   Step 317: Create a Go utility to dynamically inject user topics and difficulties into the prompt template
*   Step 318: Write Go unit tests verifying the prompt template rendering logic
*   Step 319: Conduct manual prompt testing in Vertex AI Studio to evaluate hallucination rates
*   Step 320: Finalize version 1.0 of the Quiz Generation prompt configuration

### Topic to Question Generation Pipeline
*   Step 321: Define Protobuf schemas for the internal `GenerateQuizRequest` (Topic, Count, Difficulty)
*   Step 322: Implement the Go backend API route to accept quiz generation requests
*   Step 323: Wire the Go backend to dispatch the rendered prompt to the Vertex AI Gemini model
*   Step 324: Implement streaming response handling in Go to reduce perceived generation latency
*   Step 325: Implement a timeout Context in Go cancelling generation if it exceeds 30 seconds
*   Step 326: Develop the Angular UI input field capturing the user's desired quiz topic
*   Step 327: Wire the Angular UI to the Go backend generation endpoint
*   Step 328: Implement NgRx state to handle the 'Generating...' loading status in the UI
*   Step 329: Write backend unit tests mocking the Vertex AI API response
*   Step 330: Conduct end-to-end testing of the full generation pipeline from Angular to Vertex AI

### Difficulty Scaling Algorithm
*   Step 331: Expand the Vertex AI prompt to accept an explicit `Difficulty` parameter (Easy, Medium, Hard)
*   Step 332: Define few-shot examples illustrating exactly what constitutes a 'Hard' vs 'Easy' question
*   Step 333: Implement Go logic to calculate a dynamic temperature setting based on difficulty (higher temp for harder questions)
*   Step 334: Implement the Angular Material UI toggle allowing users to select the difficulty level
*   Step 335: Update the NgRx state slice to persist the user's preferred difficulty setting
*   Step 336: Write Go logic to analyze the generated vocabulary complexity using basic heuristics
*   Step 337: Implement Spanner logic tracking the average win-rate per difficulty tier to ensure accuracy
*   Step 338: Conduct A/B testing of prompt variations to ensure 'Hard' questions are not just obscure
*   Step 339: Write Jasmine unit tests for the Angular difficulty selection component
*   Step 340: Deploy the difficulty scaling feature and monitor generation error rates

### Question Validation & Fact-Checking
*   Step 341: Implement a secondary Vertex AI 'Fact Checker' prompt running in parallel to generation
*   Step 342: Configure the Fact Checker prompt to flag historically or factually inaccurate answers
*   Step 343: Write Go logic to discard specific generated questions if flagged by the Fact Checker
*   Step 344: Implement logic to request replacement questions if the batch size drops below the required count
*   Step 345: Define Protobuf schemas for validation error reporting (e.g., `FailedFactCheck`)
*   Step 346: Implement Go logic to sanitize generated strings, removing markdown formatting artifacts
*   Step 347: Write Go unit tests validating the sanitization and filtering algorithms
*   Step 348: Monitor the cost impact of the secondary fact-checking AI call
*   Step 349: Implement a user reporting UI allowing players to flag bad questions post-match
*   Step 350: Wire the reporting UI to a Spanner table for manual developer review

### Output Parsing & Protobuf Serialization
*   Step 351: Define the strict Protobuf schema for a generated `Quiz` containing `Questions` and `Options`
*   Step 352: Implement Go logic to parse the raw JSON string returned by Vertex AI
*   Step 353: Implement strict schema validation in Go ensuring every question has exactly one correct answer
*   Step 354: Implement logic to detect and reject malformed JSON (e.g., missing quotes, truncated arrays)
*   Step 355: Write Go unit tests intentionally feeding malformed JSON to the parser to verify error handling
*   Step 356: Implement a Go fallback loop to ask Vertex AI to correct the JSON if parsing fails
*   Step 357: Serialize the validated Go struct into the final Protobuf binary payload
*   Step 358: Update the Angular HTTP service to decode the Protobuf response payload
*   Step 359: Write Jasmine unit tests validating the Angular Protobuf decoding logic
*   Step 360: Conduct load testing to verify Protobuf serialization overhead under heavy concurrent generation

### Caching Generated Quizzes in Spanner
*   Step 361: Design the Spanner DDL schema for the `QuizCache` table (TopicHash, Difficulty, Payload, Expiry)
*   Step 362: Implement Go logic to normalize and hash the user's requested topic string (e.g., 'ww2' -> 'world war 2')
*   Step 363: Implement Go backend logic to query the `QuizCache` before calling Vertex AI
*   Step 364: Implement Spanner transactions to insert newly generated quizzes into the cache
*   Step 365: Set up a Spanner Time-To-Live (TTL) policy to automatically delete cached quizzes after 30 days
*   Step 366: Implement logic to bypass the cache if the user explicitly requests a 'Fresh' generation
*   Step 367: Write Go integration tests ensuring cache hits return identical payloads to fresh generations
*   Step 368: Instrument cache hit/miss ratios with OpenTelemetry metrics
*   Step 369: Create a Cloud Monitoring dashboard visualizing the Vertex AI cost savings from the cache
*   Step 370: Deploy the cache mechanism and monitor Spanner storage growth

### AI Rate Limiting & Cost Controls
*   Step 371: Implement Redis-backed rate limiting using Google Cloud Memorystore
*   Step 372: Configure the Go backend to restrict standard users to 5 quiz generations per hour
*   Step 373: Configure the Go backend to restrict premium users to 50 quiz generations per hour
*   Step 374: Implement HTTP 429 (Too Many Requests) error handling in the Go API
*   Step 375: Implement the Angular Material UI alert explaining the rate limit to the user
*   Step 376: Define NgRx logic to calculate and display the 'Time until next generation' countdown
*   Step 377: Configure a global circuit breaker in Go to halt Vertex AI calls if costs exceed $X/hour
*   Step 378: Write Go unit tests validating the Redis token bucket algorithm
*   Step 379: Conduct staging load tests to verify rate limits hold under synthetic DDoS attacks
*   Step 380: Document the cost control architecture and circuit breaker emergency procedures

### Quiz Translation & Localization
*   Step 381: Expand the Vertex AI prompt to explicitly demand output in the requested language ISO code
*   Step 382: Implement the Angular Material UI dropdown for selecting the quiz generation language
*   Step 383: Wire the language selection state into the NgRx store
*   Step 384: Update the Go backend `GenerateQuizRequest` Protobuf schema to include `LanguageCode`
*   Step 385: Modify the Spanner `QuizCache` hash to include the LanguageCode, preventing cross-language cache collisions
*   Step 386: Implement fallback Go logic to default to English if an unsupported language is requested
*   Step 387: Write Go unit tests validating language parameter propagation through the pipeline
*   Step 388: Conduct manual testing generating quizzes in Spanish, French, and Japanese
*   Step 389: Monitor Vertex AI response latency variations across different language outputs
*   Step 390: Finalize the list of officially supported generation languages for Launch

### Player Feedback & AI Fine-tuning Loop
*   Step 391: Design the Spanner DDL schema for the `QuestionFeedback` table (QuestionID, Rating, Reason)
*   Step 392: Implement the Angular Material UI 'Thumbs Up / Thumbs Down' buttons on post-match review screens
*   Step 393: Wire the feedback buttons to a Go backend API route
*   Step 394: Write a Google Cloud Run scheduled job to aggregate highly-rated questions into a 'Gold Standard' dataset
*   Step 395: Format the Gold Standard dataset into JSONL suitable for Vertex AI model fine-tuning
*   Step 396: Document the manual procedure for launching a Vertex AI fine-tuning job
*   Step 397: Write a script to evaluate the fine-tuned model's output quality against the baseline model
*   Step 398: Configure A/B testing infrastructure in Go to route 10% of traffic to the fine-tuned model
*   Step 399: Monitor player satisfaction metrics to determine if the fine-tuned model performs better
*   Step 400: Establish a monthly cadence for retraining the fine-tuned model on new player data

## Epic 5: Real-time Chat & Presence via Firebase

### Firebase Realtime Database Provisioning
*   Step 401: Enable the Firebase Realtime Database (RTDB) in the Firebase Console
*   Step 402: Configure the RTDB to deploy in the same primary region as the Spanner instance
*   Step 403: Establish the fundamental JSON structure for the RTDB (e.g., `/presence`, `/chats`)
*   Step 404: Write the initial Firebase Security Rules locking down read/write access
*   Step 405: Install the Firebase RTDB Go Admin SDK in the backend workspace
*   Step 406: Install the Firebase RTDB Angular SDK in the frontend workspace
*   Step 407: Write a Go script to verify Admin SDK connection and latency to the RTDB
*   Step 408: Create an Angular wrapper service specifically managing RTDB connection state
*   Step 409: Configure automated daily backups for the RTDB in the Google Cloud Console
*   Step 410: Document the RTDB data architecture in the infrastructure runbook

### Global Lobby Chat Channel
*   Step 411: Define the RTDB JSON path for the global lobby (e.g., `/chats/global/messages`)
*   Step 412: Implement the Angular Material UI chat window component for the main lobby
*   Step 413: Implement an Angular service subscribing to the RTDB `/chats/global/messages` stream
*   Step 414: Write Firebase Security Rules ensuring any authenticated user can read/write to the global chat
*   Step 415: Implement Go backend logic to periodically trim the global chat history (keeping only the last 100 messages)
*   Step 416: Define NgRx state to track unread messages while the chat window is minimized
*   Step 417: Implement UI logic to auto-scroll to the newest message upon receipt
*   Step 418: Write Jasmine unit tests validating the auto-scroll and unread badge logic
*   Step 419: Conduct load testing on the RTDB to verify latency under 1000 simulated global chatters
*   Step 420: Deploy the global chat channel and monitor RTDB bandwidth consumption

### Private Direct Messaging
*   Step 421: Define the RTDB JSON path for direct messages (e.g., `/chats/private/{user1}_{user2}`)
*   Step 422: Write strict Firebase Security Rules ensuring only the two participants can read/write the private node
*   Step 423: Implement the Angular Material UI for selecting a friend and opening a private chat tab
*   Step 424: Wire the Angular RTDB service to listen to the specific private chat path
*   Step 425: Implement Go backend logic to generate a consistent chat ID regardless of which user initiates
*   Step 426: Define Protobuf schemas for the Go backend API used to fetch historical private messages
*   Step 427: Implement NgRx actions for switching active chat contexts (Global vs Private)
*   Step 428: Write Jasmine unit tests validating the chat context switching logic
*   Step 429: Verify Security Rule enforcement by attempting unauthorized reads via a mocked client
*   Step 430: Deploy private messaging and monitor RTDB concurrent connection metrics

### Study Group (Guild) Chat Rooms
*   Step 431: Define the RTDB JSON path for study group chats (e.g., `/chats/groups/{groupID}`)
*   Step 432: Implement Go backend logic to securely push a list of a user's authorized Group IDs to their RTDB profile
*   Step 433: Write Firebase Security Rules relying on the user's RTDB profile to authorize access to the Group chat
*   Step 434: Implement the Angular Material UI channel list distinguishing between different Study Groups
*   Step 435: Wire the Angular RTDB service to multiplex subscriptions across multiple Group channels
*   Step 436: Define NgRx selectors calculating aggregate unread message counts across all Groups
*   Step 437: Implement an 'announcement' styling feature for Group Admins within the chat UI
*   Step 438: Write Go integration tests ensuring users removed from a Group lose RTDB access immediately
*   Step 439: Conduct staging environment testing of the complex Security Rule authorization flow
*   Step 440: Document the Group chat security architecture and authorization propagation

### Match-Specific Voice/Text Chat
*   Step 441: Define the RTDB JSON path for ephemeral match chats (e.g., `/chats/matches/{matchID}`)
*   Step 442: Write Firebase Security Rules restricting access exclusively to players holding the active match ticket
*   Step 443: Implement the Angular Material in-game chat overlay component
*   Step 444: Wire the in-game UI to the ephemeral RTDB path during the active match
*   Step 445: Implement Go backend logic executed by Agones to delete the ephemeral chat node entirely upon match completion
*   Step 446: Integrate a third-party WebRTC provider (e.g., LiveKit or Twilio) for the voice channel infrastructure
*   Step 447: Implement the Angular UI toggle to mute/unmute the microphone during the match
*   Step 448: Write Go backend logic to mint secure WebRTC access tokens for authenticated match participants
*   Step 449: Conduct end-to-end testing of the ephemeral chat lifecycle (creation to deletion)
*   Step 450: Monitor RTDB node churn rates (rapid creation/deletion) in Cloud Monitoring

### Online/Offline Presence System
*   Step 451: Define the RTDB JSON path for global presence tracking (e.g., `/presence/{userID}`)
*   Step 452: Implement the Firebase `onDisconnect` hook in Angular to automatically mark the user as 'offline' upon closure
*   Step 453: Implement Angular logic to push the 'online' state upon successful Firebase authentication
*   Step 454: Write Go backend logic utilizing the Admin SDK to query the real-time presence of a user's friends list
*   Step 455: Implement the Angular Material UI indicators (green dot) on the friends list and group roster
*   Step 456: Define NgRx state to cache presence status and prevent aggressive UI flickering
*   Step 457: Implement a 'Do Not Disturb' manual override toggle in the user settings UI
*   Step 458: Write Jasmine unit tests validating the presence mapping logic in the NgRx reducers
*   Step 459: Deploy the presence system and monitor RTDB concurrent connection limits
*   Step 460: Document the presence architecture and `onDisconnect` reliability considerations

### Typing Indicators & Read Receipts
*   Step 461: Define the RTDB JSON path for typing indicators (e.g., `/chats/private/{chatId}/typing/{userId}`)
*   Step 462: Define the RTDB JSON path for read receipts (e.g., `/chats/private/{chatId}/lastRead/{userId}`)
*   Step 463: Implement Angular RxJS logic (using `debounceTime`) to throttle typing indicator writes to the RTDB
*   Step 464: Implement the Angular Material UI 'User is typing...' animation bubble
*   Step 465: Implement Angular logic using Intersection Observer to dispatch read receipt writes when messages enter the viewport
*   Step 466: Implement the Angular Material UI 'Read' checkmark icon next to sent messages
*   Step 467: Write strict Firebase Security Rules ensuring users can only update their own typing/read states
*   Step 468: Write Jasmine unit tests verifying the `debounceTime` optimization for typing events
*   Step 469: Conduct manual UX testing to ensure typing indicators disappear reliably if the user closes the tab
*   Step 470: Deploy the typing and read receipt systems to the staging environment

### Chat Moderation & Profanity Filters
*   Step 471: Integrate a robust Go-based profanity filtering library or standard dictionary
*   Step 472: Implement Go backend logic subscribing to RTDB `onWrite` events for all chat paths (via Cloud Functions or local worker)
*   Step 473: Write Go logic to analyze incoming messages and replace profanity with asterisks (***)
*   Step 474: Implement a strike system in Spanner tracking repeated violations by specific users
*   Step 475: Write Go logic to automatically apply a temporary 'shadow ban' (read-only) for users exceeding 3 strikes
*   Step 476: Implement the Angular UI alert notifying the user of their shadow ban status and remaining duration
*   Step 477: Implement a 'Report Message' Angular UI context menu allowing users to flag inappropriate content
*   Step 478: Wire the 'Report Message' action to a Spanner table for manual administrator review
*   Step 479: Write Go unit tests validating the profanity masking algorithm against edge cases
*   Step 480: Deploy the moderation worker and verify performance under high chat volume

### Push Notifications (FCM) integration
*   Step 481: Enable Firebase Cloud Messaging (FCM) within the Firebase Console
*   Step 482: Generate a Web Push Certificate (VAPID key) for browser notification routing
*   Step 483: Implement the Angular service to request browser notification permissions from the user
*   Step 484: Implement logic to securely transmit and persist the user's FCM device token in Spanner
*   Step 485: Write a `firebase-messaging-sw.js` service worker to handle background message reception
*   Step 486: Implement Go backend logic to construct and dispatch FCM payloads via the Admin SDK
*   Step 487: Configure the Go backend to dispatch FCMs only for Direct Messages when the user is marked 'offline'
*   Step 488: Implement Angular UI toast notifications for foreground FCM message handling
*   Step 489: Write Go integration tests mocking the FCM Admin SDK dispatch
*   Step 490: Conduct end-to-end testing of push notifications across Chrome, Firefox, and Safari

### Chat History Archival to Spanner
*   Step 491: Design the Spanner DDL schema for the `ArchivedChats` table (ChatID, MessageID, Payload, Timestamp)
*   Step 492: Implement a Go worker process to scan the RTDB periodically (e.g., hourly) for messages older than 7 days
*   Step 493: Write Go logic to extract old messages, serialize them, and insert them into Spanner
*   Step 494: Write Go logic to execute a destructive delete on the archived RTDB nodes to reclaim space
*   Step 495: Implement a Go backend API route allowing clients to request paginated historical chat data from Spanner
*   Step 496: Update the Angular chat UI with 'Load Previous Messages' infinite scrolling logic
*   Step 497: Wire the infinite scrolling logic to stitch Spanner historical data with RTDB live data seamlessly
*   Step 498: Write Go integration tests ensuring zero message loss during the archival migration window
*   Step 499: Configure Cloud Monitoring alerts tracking the size of the RTDB to ensure archival is functioning
*   Step 500: Document the hybrid RTDB/Spanner data storage architecture in the engineering wiki

## Epic 6: Matchmaking Logic via Open Match

### Open Match Core Installation on GKE
*   Step 501: Provision a dedicated Google Kubernetes Engine (GKE) cluster for Open Match and Agones
*   Step 502: Install the Open Match core components (Frontend, Backend, Query) via Helm charts
*   Step 503: Configure Kubernetes namespaces isolating Open Match services from standard backend APIs
*   Step 504: Set up the Redis state store required by Open Match for ticket synchronization
*   Step 505: Implement a Go backend service to act as the Matchmaking Frontend, accepting player join requests
*   Step 506: Configure strict IAM roles allowing the Matchmaking Frontend to write to the Open Match Redis instance
*   Step 507: Implement OpenTelemetry tracing within the GKE cluster specifically tracking ticket lifecycle times
*   Step 508: Write a basic Go test client to submit a dummy ticket and verify Redis insertion
*   Step 509: Configure Kubernetes Horizontal Pod Autoscalers (HPA) for the Open Match core services
*   Step 510: Document the Open Match architecture and Redis dependency in the infrastructure runbook

### Matchmaking Ticket Creation API
*   Step 511: Define Protobuf schemas for `CreateTicketRequest` (PlayerID, GameMode, Topic, EloRating)
*   Step 512: Implement the Go backend API route receiving match requests from Angular clients
*   Step 513: Write logic validating that the player is not already in an active match or queue
*   Step 514: Map the validated request into an Open Match `Ticket` structure, attaching search fields
*   Step 515: Implement Go logic to dispatch the `Ticket` to the Open Match Frontend service via gRPC
*   Step 516: Return the Open Match `TicketID` to the Angular client for polling
*   Step 517: Implement an Angular HTTP service to dispatch the join request and handle 429 rate limits
*   Step 518: Implement NgRx state actions for transitioning the user UI into the 'Queueing' state
*   Step 519: Write Jasmine unit tests verifying the UI transition upon successful ticket creation
*   Step 520: Conduct load testing on the Ticket Creation API to ensure it handles 1000 requests/sec

### Normal Mode Match Director
*   Step 521: Implement a custom Go service acting as the Open Match 'Director' for Normal Mode
*   Step 522: Configure the Director to poll the Open Match Backend service for matches every 1 second
*   Step 523: Define the Open Match 'MatchProfile' querying tickets requesting 'Normal Mode'
*   Step 524: Write Go logic to group tickets into 5-player matches based on similar Elo ratings
*   Step 525: Implement an expanding search radius algorithm in Go (e.g., widen Elo bands if queue time > 30s)
*   Step 526: Upon forming a match, instruct the Director to request a dedicated Game Server from the Agones Allocator
*   Step 527: Write Go logic assigning the newly allocated Agones IP/Port back to the Open Match tickets
*   Step 528: Write Go unit tests validating the Elo grouping math within the MatchFunction
*   Step 529: Deploy the Normal Mode Director as a standalone pod in the GKE cluster
*   Step 530: Monitor Director polling latency and match formation times via Cloud Monitoring

### Speed Mode Match Director
*   Step 531: Implement a dedicated Go Director service exclusively for the 'Speed Mode' queue
*   Step 532: Define the Open Match 'MatchProfile' filtering strictly for 'Speed Mode' tickets
*   Step 533: Implement grouping logic prioritizing ping/latency over strict Elo balancing for faster matchmaking
*   Step 534: Configure the Director to request a specific 'Speed Mode' Game Server fleet from Agones
*   Step 535: Write Go logic handling backfill scenarios if a player drops out before the 5-player lobby fills
*   Step 536: Implement strict ticket deletion logic if a Speed Mode match cannot be formed within 2 minutes
*   Step 537: Write Go unit tests mocking the Open Match backend to verify fast-path grouping
*   Step 538: Deploy the Speed Mode Director to GKE on a distinct release cycle from Normal Mode
*   Step 539: Instrument the Director with custom metrics tracking the average queue wait time
*   Step 540: Conduct staging tests simulating sudden influxes of Speed Mode players

### Death Mode Match Director
*   Step 541: Implement a dedicated Go Director service exclusively for the 'Death Mode' queue
*   Step 542: Define the Open Match 'MatchProfile' filtering strictly for 'Death Mode' tickets
*   Step 543: Implement specialized grouping logic creating massive 20-player lobbies (Battle Royale style)
*   Step 544: Configure the Director to request high-capacity 'Death Mode' Game Servers from Agones
*   Step 545: Write Go logic to continuously widen the search parameters until the 20-player threshold is met
*   Step 546: Implement a 'force start' override if the lobby reaches 15 players after 3 minutes
*   Step 547: Write Go unit tests verifying the 20-player aggregation logic and force-start conditions
*   Step 548: Deploy the Death Mode Director to GKE
*   Step 549: Create a dedicated Cloud Monitoring dashboard comparing queue times across Normal, Speed, and Death modes
*   Step 550: Document the specific scaling heuristics for the Death Mode director in the runbook

### ELO/MMR Evaluation Functions
*   Step 551: Implement the custom Go `MatchFunction` (MFF) executed by Open Match during evaluations
*   Step 552: Write Go logic to calculate the absolute Elo delta between all tickets in a proposed match
*   Step 553: Implement rejection logic inside the MFF if the Elo delta exceeds acceptable thresholds
*   Step 554: Implement 'smurf detection' logic within the MFF to rapidly adjust bands for high-win-rate players
*   Step 555: Write highly parallelized Go code using goroutines to evaluate thousands of tickets concurrently
*   Step 556: Define Protobuf schemas for the MFF gRPC server configuration
*   Step 557: Write exhaustive Go unit tests covering edge cases (e.g., negative Elos, massive disparities)
*   Step 558: Deploy the MFF service to GKE, ensuring it is accessible by all Match Directors
*   Step 559: Configure Horizontal Pod Autoscaling specifically for the MFF based on CPU utilization
*   Step 560: Conduct performance profiling (pprof) on the MFF to eliminate memory leaks during heavy ticket evaluation

### Study Group vs Study Group (GvG) Matchmaking
*   Step 561: Implement a specialized Go Director for GvG (Guild vs Guild) matches
*   Step 562: Define Open Match Profiles demanding exactly two distinct Study Group IDs
*   Step 563: Implement Go logic to aggregate the average Elo of all 5 members of a Study Group ticket
*   Step 564: Write MFF logic ensuring two Study Groups are evenly matched based on aggregate Elo
*   Step 565: Implement logic to request a 'GvG' specific Agones fleet with extended match timers
*   Step 566: Implement a fallback logic in Go allowing a Study Group to play against a team of 5 randoms if no other Guild is queued
*   Step 567: Write Go integration tests mocking a 5-man party ticket submission
*   Step 568: Develop the Angular UI flow for a Guild Leader to queue the entire party simultaneously
*   Step 569: Wire the Angular party UI to NgRx state managing party readiness
*   Step 570: Conduct end-to-end testing of a full 5v5 GvG match formation

### Matchmaking Lobby UI & State
*   Step 571: Define NgRx state slices managing the user's active matchmaking Ticket ID and status
*   Step 572: Implement an Angular service polling the Go API every 2 seconds for ticket assignment updates
*   Step 573: Develop the Angular Material UI 'Queue Overlay' displaying elapsed time and estimated wait
*   Step 574: Implement Angular logic to handle the 'Match Found' transition, displaying the Agones IP and Port
*   Step 575: Develop a 'Cancel Search' button in the Angular UI
*   Step 576: Wire the 'Cancel Search' button to a Go API route executing an Open Match ticket deletion
*   Step 577: Write Jasmine unit tests validating the polling lifecycle and cancelation interruptions
*   Step 578: Implement an exponential backoff in the Angular polling service if the Go API returns 429s
*   Step 579: Design standard audio cues (using standard HTML5 Audio) for 'Queue Joined' and 'Match Found'
*   Step 580: Perform UX testing to ensure the queue overlay remains visible across all dashboard navigation

### Cross-Region Latency Optimization
*   Step 581: Update the Go Ticket Creation API to capture the user's approximate geographic region (e.g., NA-East, EU-West)
*   Step 582: Append the region string as a Search Field in the Open Match Ticket
*   Step 583: Configure the Match Directors to heavily weight region similarity during MFF evaluation
*   Step 584: Implement Go logic allowing the Director to request an Agones game server specifically in the matched region
*   Step 585: Configure GKE and Agones to span multiple GCP regions (or federated clusters)
*   Step 586: Write Go MFF logic slowly expanding the region search (e.g., NA-East -> NA-West) if queues are dead
*   Step 587: Define NgRx state tracking the player's selected or auto-detected region
*   Step 588: Implement the Angular UI dropdown allowing players to manually override their matchmaking region
*   Step 589: Write Go integration tests ensuring EU players are not placed in NA servers unless explicitly requested
*   Step 590: Deploy regional fleets and monitor latency metrics via Agones telemetry

### Matchmaker Telemetry & Metrics
*   Step 591: Integrate OpenCensus/OpenTelemetry into the Go Match Directors and MFFs
*   Step 592: Expose Prometheus `/metrics` endpoints from all custom Open Match services
*   Step 593: Configure a Google Cloud Managed Prometheus scraper to ingest the custom metrics
*   Step 594: Build a Grafana/Cloud Monitoring dashboard tracking 'Tickets Created per Minute'
*   Step 595: Build a dashboard tracking 'Average Time in Queue by Region and Mode'
*   Step 596: Build a dashboard tracking 'Agones Allocation Success/Failure Rates'
*   Step 597: Set up alerting policies if the 'Allocation Failure Rate' exceeds 5%
*   Step 598: Set up alerting policies if the 'Average Time in Queue' exceeds 3 minutes globally
*   Step 599: Implement structured JSON logging for every match formed, detailing the Elo deltas
*   Step 600: Conduct a simulated outage of the MFF and verify alerts fire correctly

## Epic 7: Game Server Orchestration via Agones

### Agones Installation on Google Kubernetes Engine
*   Step 601: Define the Kubernetes namespace and RBAC permissions required for the Agones controller
*   Step 602: Install the Agones platform on the dedicated GKE cluster via the official Helm chart
*   Step 603: Configure the Agones custom resource definitions (CRDs) for GameServers and Fleets
*   Step 604: Verify the Agones controller, allocator, and ping service pods are running healthy
*   Step 605: Configure GCP VPC firewall rules opening UDP/TCP port ranges (e.g., 7000-8000) for game traffic
*   Step 606: Configure a dedicated Node Pool in GKE specifically for Game Servers to isolate them from web APIs
*   Step 607: Apply Kubernetes taints and tolerations ensuring only GameServers schedule on the dedicated nodes
*   Step 608: Set up the Agones Allocator service with gRPC TLS certificates for secure Open Match communication
*   Step 609: Write a basic Go test client to request a dummy allocation from the Agones Allocator
*   Step 610: Document the Agones installation architecture and firewall requirements in the runbook

### Dedicated Go Game Server Initialization
*   Step 611: Initialize a new Go binary specifically for the authoritative Game Server (`cmd/server`)
*   Step 612: Integrate the official Agones Go SDK into the Game Server binary
*   Step 613: Implement the Go boot sequence calling `agones.Ready()` to mark the server as available
*   Step 614: Implement Graceful Shutdown logic calling `agones.Shutdown()` when the match concludes
*   Step 615: Configure a basic health check loop calling `agones.Health()` every 2 seconds
*   Step 616: Write a multi-stage Dockerfile optimized for the standalone Go Game Server
*   Step 617: Configure Cloud Build to compile and push the Game Server image to Artifact Registry
*   Step 618: Deploy a manual `GameServer` CRD to GKE to verify the image boots and registers correctly
*   Step 619: Implement structured logging outputting standard JSON to be captured by Google Cloud Logging
*   Step 620: Verify standard logs appear in the GCP Console under the specific GameServer pod ID

### WebSockets Connection Handling
*   Step 621: Implement a robust Go WebSocket server utilizing `gorilla/websocket` within the Game Server
*   Step 622: Configure the Agones `Fleet` definition to expose the WebSocket TCP port dynamically
*   Step 623: Implement Go logic to extract and validate the Firebase JWT passed during the WebSocket upgrade
*   Step 624: Implement a Go `Player` struct managing individual read/write goroutines for each connection
*   Step 625: Implement a thread-safe broadcast channel in Go to send real-time state updates to all connected players
*   Step 626: Implement Go logic handling dirty disconnects (e.g., sudden WiFi drop) and connection cleanup
*   Step 627: Implement an Angular WebSocket service using RxJS `webSocket` subject to connect to the assigned Agones IP/Port
*   Step 628: Wire the Angular WebSocket service to automatically inject the Firebase JWT on connection
*   Step 629: Write Jasmine unit tests validating the Angular reconnection backoff logic
*   Step 630: Conduct load testing on a single Go Game Server to ensure it handles 20 concurrent WebSockets smoothly

### Real-time Quiz State Machine (Question reveals, timers)
*   Step 631: Define the core Go State Machine phases: Lobby, Countdown, QuestionReveal, ScoreTally, MatchEnd
*   Step 632: Implement Go ticker logic executing precisely every 1 second to manage phase timers
*   Step 633: Define Protobuf payloads for the `StateUpdate` messages broadcasted to clients
*   Step 634: Implement Go logic fetching the pre-generated Quiz JSON from Spanner during the Lobby phase
*   Step 635: Implement Go logic broadcasting the current question (without the answer) to all clients
*   Step 636: Implement Go logic broadcasting the correct answer and scoreboard once the timer expires
*   Step 637: Implement NgRx reducers on the Angular frontend to parse `StateUpdate` payloads and update the UI
*   Step 638: Develop the Angular UI synchronized countdown timer component
*   Step 639: Write Go unit tests validating the exact timing transitions of the State Machine
*   Step 640: Write Jasmine tests verifying the Angular UI correctly hides/reveals questions based on NgRx state

### Player Input Validation & Scoring
*   Step 641: Define Protobuf schemas for the `SubmitAnswer` payload sent by the Angular client
*   Step 642: Implement Go logic intercepting the `SubmitAnswer` WebSocket message
*   Step 643: Write strict Go validation ensuring the answer is submitted before the server-side timer expires
*   Step 644: Write strict Go validation ensuring the player hasn't already answered the current question
*   Step 645: Implement scoring math in Go (e.g., base points + speed bonus for answering quickly)
*   Step 646: Implement anti-cheat logic in Go detecting impossible response times (e.g., answering in 10ms)
*   Step 647: Implement Go logic aggregating scores into a real-time leaderboard broadcasted every phase
*   Step 648: Implement the Angular Material UI for selecting an answer option
*   Step 649: Wire the UI selection to dispatch the Protobuf `SubmitAnswer` payload over the WebSocket
*   Step 650: Write Go integration tests simulating 5 concurrent players answering simultaneously to check for race conditions

### K-Factor ELO Calculation Engine
*   Step 651: Implement a Go package dedicated to Elo math (e.g., `pkg/elo`) within the Game Server
*   Step 652: Write the standard 1v1 Elo algorithm implementation
*   Step 653: Extend the Go logic into a Multi-player Elo algorithm (e.g., treating a 5-player FFA as a series of 1v1s)
*   Step 654: Implement dynamic K-Factor scaling in Go (higher K for new accounts, lower K for veterans)
*   Step 655: Implement Go logic calculating the Elo deltas for all participants at the `MatchEnd` phase
*   Step 656: Define Protobuf schemas for the `MatchResults` payload broadcasting the Elo changes to the clients
*   Step 657: Implement the Angular Material UI 'Results Screen' displaying the Elo gain/loss animations
*   Step 658: Write exhaustive Go unit tests verifying the mathematical correctness of the Multi-player Elo formula
*   Step 659: Write Go unit tests verifying K-Factor scaling correctly prevents extreme volatility at high ranks
*   Step 660: Integrate the calculated Elo deltas into the final Spanner transaction payload

### Post-Match Rewards & Coin Distribution
*   Step 661: Implement Go logic calculating the base Coin and Gem payouts based on final placement
*   Step 662: Implement Go logic multiplying rewards based on 'First Win of the Day' or active boosters
*   Step 663: Compile the final Match Record, Elo Deltas, and Economy Rewards into a single Go struct
*   Step 664: Write a massive, atomic Google Cloud Spanner transaction to commit the entire match result at once
*   Step 665: Implement Go retry logic for the Spanner transaction in case of transient database contention
*   Step 666: Broadcast the final `MatchResults` Protobuf to all clients only after Spanner confirms the commit
*   Step 667: Call `agones.Shutdown()` precisely after all results are broadcasted and saved
*   Step 668: Develop the Angular UI 'Rewards Screen' displaying coin cascades and level-up bars
*   Step 669: Write Go integration tests verifying that economy balances perfectly reflect the calculated match rewards
*   Step 670: Monitor Spanner transaction latencies at the end of matches to ensure no bottlenecks

### Server Fleet Autoscaling Policies
*   Step 671: Define the baseline Agones `Fleet` YAML configuration for production environments
*   Step 672: Configure the `FleetAutoscaler` utilizing a `Buffer` policy (e.g., maintain exactly 5 'Ready' servers)
*   Step 673: Implement Cloud Monitoring alerts if the buffer drops to 0, indicating allocator exhaustion
*   Step 674: Write a Bazel target to apply the Fleet configurations to GKE seamlessly
*   Step 675: Implement separate `Fleet` definitions for Normal Mode, Speed Mode, and Death Mode based on CPU requirements
*   Step 676: Configure GKE Cluster Autoscaler to rapidly provision new nodes when pending pods exceed limits
*   Step 677: Conduct a staging load test spawning 500 matches simultaneously to verify the GKE Cluster Autoscaler
*   Step 678: Optimize the Game Server Docker image size to accelerate pod startup times during scaling events
*   Step 679: Document the specific node instance types (e.g., e2-standard-4) selected for optimal Game Server packing
*   Step 680: Verify Agones scaling metrics appear correctly in the custom Grafana/Monitoring dashboard

### Game Server Crash Recovery & Reconnects
*   Step 681: Implement Go logic generating a unique `MatchSessionToken` passed to clients upon match formation
*   Step 682: Implement Go logic allowing a client to pass the `MatchSessionToken` on WebSocket connection (instead of just JWT)
*   Step 683: Write Go logic to seamlessly re-attach a returning player to their existing `Player` struct in memory
*   Step 684: Implement Angular RxJS logic automatically attempting a WebSocket reconnection if the connection drops unexpectedly
*   Step 685: Configure the Angular UI to display a 'Reconnecting...' overlay during network interruptions
*   Step 686: Write Go logic ensuring an offline player automatically receives 0 points for missed questions
*   Step 687: Implement strict Spanner logic ensuring a crashed Agones server doesn't leave 'ghost' matches in the database
*   Step 688: Write a scheduled Go worker scanning Spanner for matches stuck in 'In Progress' for over 1 hour and voiding them
*   Step 689: Conduct manual testing by disabling Wi-Fi mid-match and verifying clean reconnection
*   Step 690: Document the crash recovery architectures and edge-case behaviors in the engineering wiki

### Agones Dashboard & Fleet Monitoring
*   Step 691: Expose Prometheus metrics directly from the Go Game Server (e.g., players connected, memory allocated)
*   Step 692: Configure Agones to automatically inject Prometheus annotations into GameServer pods
*   Step 693: Deploy the official Agones metrics dashboards to Google Cloud Monitoring
*   Step 694: Create a custom dashboard tracking 'Average Match Duration' and 'Client Reconnect Rates'
*   Step 695: Create a custom dashboard tracking 'Spanner Commit Latency' originating from Game Servers
*   Step 696: Configure alerting policies for Game Servers that stay 'Allocated' for longer than 30 minutes (potential zombie servers)
*   Step 697: Implement Go logic calling `agones.SetLabel()` to tag servers with their specific Game Mode for easier filtering
*   Step 698: Verify all logs from short-lived Game Server pods are successfully shipped to Cloud Logging before termination
*   Step 699: Conduct a tabletop exercise responding to a massive GKE node failure
*   Step 700: Finalize the complete Game Server operations runbook

## Epic 8: Angular Shared UI Components & Material

### Angular Material Theme Generation
*   Step 701: Initialize Angular Material within the workspace (`ng add @angular/material`)
*   Step 702: Generate the custom SCSS theme defining the Primary, Accent, and Warn palettes (Deep Indigos, Electric Amber, Crimson)
*   Step 703: Configure the global `styles.scss` to import standard CSS resets and the compiled Material theme
*   Step 704: Implement a CSS variable strategy allowing for runtime switching between Light and Dark modes
*   Step 705: Create an Angular Service (`ThemeService`) utilizing RxJS to track the active theme state
*   Step 706: Implement an Angular Material UI toggle switch allowing the user to select their theme preference
*   Step 707: Wire the `ThemeService` to persist the selected theme in `localStorage`
*   Step 708: Write Jasmine unit tests validating the `ThemeService` initialization logic
*   Step 709: Update the root `index.html` to dynamically inject the correct theme class on the `<body>` tag
*   Step 710: Verify all standard Angular Material components render perfectly in both Light and Dark modes

### Typography and Color Palettes
*   Step 711: Select and integrate a custom Google Font (e.g., 'Inter' or 'Outfit') via `index.html`
*   Step 712: Define the custom Angular Material Typography config, overriding the default Roboto scales
*   Step 713: Implement specific CSS utility classes (`.text-h1`, `.text-body`) mapped directly to the typography scale
*   Step 714: Define SCSS mixins for standard spacing units (e.g., margins and paddings) to enforce vertical rhythm
*   Step 715: Implement an SCSS color map extracting specific palette hex codes for usage outside of standard Material components
*   Step 716: Create an Angular UI Style Guide component demonstrating every typography and color variant
*   Step 717: Write CSS regression tests (e.g., using a tool like BackstopJS) for the style guide
*   Step 718: Optimize the Google Font loading strategy (e.g., `font-display: swap`) to prevent Cumulative Layout Shift
*   Step 719: Audit contrast ratios across all primary and accent color combinations to ensure WCAG AA compliance
*   Step 720: Document the CSS architecture and utility class naming conventions in the frontend `README.md`

### Global Layouts & Navigation Shell
*   Step 721: Develop the main `SidenavComponent` using Angular Material `<mat-sidenav>`
*   Step 722: Develop the top `ToolbarComponent` using `<mat-toolbar>`, housing the user profile menu and notifications
*   Step 723: Implement Angular Router configuration to wrap all authenticated views within the Sidenav shell
*   Step 724: Create responsive SCSS logic using Media Queries to automatically collapse the Sidenav on mobile devices
*   Step 725: Implement NgRx state selecting the current Route to dynamically highlight the active Sidenav item
*   Step 726: Develop an Angular Directive managing swipe gestures on mobile devices to open/close the Sidenav
*   Step 727: Write Jasmine unit tests verifying the Sidenav collapse behavior across different mocked viewport sizes
*   Step 728: Implement the global loading bar (`<mat-progress-bar>`) bound to an HTTP Interceptor for network requests
*   Step 729: Conduct manual UX testing on actual mobile devices (iOS/Android) for navigation fluidity
*   Step 730: Configure strict Route Guards in Angular to prevent unauthenticated users from seeing the Navigation Shell

### Reusable Quiz Card Components
*   Step 731: Develop the base `<app-quiz-card>` component accepting standard Inputs (Title, Topic, Difficulty)
*   Step 732: Implement Angular Content Projection (`<ng-content>`) allowing custom actions to be embedded in the card footer
*   Step 733: Implement a hover elevation animation using standard CSS transitions
*   Step 734: Develop the `<app-question-renderer>` component handling markdown-formatted quiz questions
*   Step 735: Develop the `<app-answer-button>` component featuring active/selected/disabled states
*   Step 736: Write Jasmine unit tests verifying the Quiz Card emits the correct event outputs when clicked
*   Step 737: Implement Skeleton Loader variants for the Quiz Cards to display during initial data fetching
*   Step 738: Verify the responsive grid layout ensuring Quiz Cards wrap correctly on tablet and mobile viewports
*   Step 739: Integrate custom icon sets (e.g., Material Symbols) into the Quiz Card headers
*   Step 740: Create an internal component library demo view for testing various Quiz Card permutations

### Player Avatar & Badge Components
*   Step 741: Develop the `<app-player-avatar>` component accepting standard Inputs (ImageURL, Size, Status)
*   Step 742: Implement fallback UI logic displaying the user's initials if their ImageURL is null or fails to load
*   Step 743: Implement the presence indicator (green dot) anchored to the bottom right of the avatar
*   Step 744: Develop the `<app-league-badge>` component rendering SVGs representing the user's rank (Bronze, Silver, Gold)
*   Step 745: Implement an Angular Pipe mapping numeric Elo ratings to their corresponding string ranks
*   Step 746: Write Jasmine unit tests validating the initials extraction algorithm (e.g., 'John Doe' -> 'JD')
*   Step 747: Optimize avatar image loading using standard HTML `loading="lazy"` attributes
*   Step 748: Develop a `<app-player-tooltip>` component displaying high-level stats when hovering over an avatar
*   Step 749: Integrate the Avatar and Badge components into the main Navigation Toolbar
*   Step 750: Verify the Avatar component scales cleanly via CSS variables without losing resolution

### Dynamic Data Tables (Leaderboards)
*   Step 751: Develop the base `<app-leaderboard-table>` using Angular Material `<mat-table>`
*   Step 752: Implement Angular Material Pagination (`<mat-paginator>`) to handle thousands of ranked rows
*   Step 753: Implement Angular Material Sorting (`<mat-sort>`) allowing users to sort by Win Rate or Matches Played
*   Step 754: Wire the Table component to an NgRx DataSource supporting server-side sorting and pagination
*   Step 755: Implement a sticky header and fixed column layout ensuring the player's rank remains visible while scrolling horizontally on mobile
*   Step 756: Implement CSS logic highlighting the current user's specific row if they appear on the active page
*   Step 757: Write Jasmine unit tests validating the server-side pagination state calculations
*   Step 758: Develop a Skeleton Loader specifically for the Table component to prevent layout jumping
*   Step 759: Implement a fuzzy search input above the table allowing users to filter for specific display names
*   Step 760: Conduct performance profiling to ensure the table renders 100 rows without dropping frames

### Modal & Dialog Service Framework
*   Step 761: Develop a centralized Angular `DialogService` wrapping the Angular Material `MatDialog` API
*   Step 762: Implement standard configuration presets for Modals (e.g., standard width, disable close on backdrop click)
*   Step 763: Develop a generic `<app-confirmation-dialog>` component accepting title, body, and action button text inputs
*   Step 764: Implement a standard error dialog explicitly parsing and displaying HTTP error payloads
*   Step 765: Wire the `DialogService` into NgRx Effects to automatically display error dialogs on failed API calls
*   Step 766: Write Jasmine unit tests mocking `MatDialog` to verify the `DialogService` opens the correct components
*   Step 767: Develop a full-screen mobile dialog preset for complex forms (e.g., Quiz Generation settings)
*   Step 768: Implement focus trapping logic within all custom dialogs to ensure strict ADA/WCAG accessibility
*   Step 769: Verify Dialog keyboard navigation (Escape to close, Tab to focus) across all standard modals
*   Step 770: Document the `DialogService` usage patterns for other frontend developers

### Toast Notifications & Snackbars
*   Step 771: Develop a centralized Angular `ToastService` wrapping the Angular Material `MatSnackBar` API
*   Step 772: Implement custom styling classes mapping to Success (Green), Warning (Yellow), and Error (Red) states
*   Step 773: Configure standard duration settings (e.g., 3000ms for Success, 5000ms for Errors)
*   Step 774: Wire the `ToastService` into the central NgRx Effects to display success toasts globally (e.g., 'Profile Updated')
*   Step 775: Implement a custom snackbar component allowing an embedded action button (e.g., 'Undo' or 'View Details')
*   Step 776: Write Jasmine unit tests validating the `ToastService` dispatch parameters
*   Step 777: Configure the `MatSnackBar` positioning to anchor at the bottom-center on mobile, and top-right on desktop
*   Step 778: Implement logic to prevent multiple identical toasts from stacking simultaneously on the screen
*   Step 779: Integrate Firebase Cloud Messaging (FCM) payloads directly into the `ToastService` to display incoming background notifications
*   Step 780: Conduct UX testing to ensure toasts do not obscure critical UI elements like the primary action button

### Loading Spinners & Skeleton Screens
*   Step 781: Develop a global `<app-loading-overlay>` component utilizing `<mat-spinner>` for critical blocking operations
*   Step 782: Implement an RxJS Subject tracking the global 'isBlocking' state across the application
*   Step 783: Develop generic `<app-skeleton-text>` and `<app-skeleton-image>` directives for granular content loading
*   Step 784: Implement a shimmering CSS animation to apply uniformly across all skeleton components
*   Step 785: Wire the skeleton components to the NgRx `isLoading` selectors within the specific feature views
*   Step 786: Write Jasmine unit tests verifying the loading overlay prevents user interactions while active
*   Step 787: Implement a minimum display duration (e.g., 500ms) for spinners to prevent aggressive UI flickering on fast networks
*   Step 788: Develop a specialized Skeleton Screen layout perfectly matching the final Active Match UI
*   Step 789: Review all API routes to ensure they properly dispatch the 'Stop Loading' action even on HTTP failures
*   Step 790: Deploy the Skeleton architecture and monitor perceived performance metrics

### Form Validation & Input Directives
*   Step 791: Implement a centralized Angular Service parsing complex backend validation errors into human-readable strings
*   Step 792: Develop standard `<mat-form-field>` wrappers ensuring consistent error message placement
*   Step 793: Implement a custom Angular Validator for strong password enforcement (8 chars, 1 number, 1 symbol)
*   Step 794: Implement a custom Angular Validator ensuring two 'Confirm Password' fields match exactly
*   Step 795: Develop an Angular Directive automatically trimming whitespace on blur for all text inputs
*   Step 796: Write Jasmine unit tests covering every custom Validator and Directive extensively
*   Step 797: Implement dynamic form submission disabling (button greyed out) if the `FormGroup` is invalid or pending
*   Step 798: Wire the centralized error parser to the NgRx failure actions
*   Step 799: Conduct UX testing ensuring screen readers properly announce all form validation errors
*   Step 800: Document the standard Angular Reactive Forms architecture pattern used across the project

## Epic 9: Angular Application Views & NgRx State

### Landing Page & Marketing Views
*   Step 801: Develop the `<app-landing-page>` component serving as the unauthenticated root route (`/`)
*   Step 802: Implement high-performance hero section animations using standard CSS keyframes
*   Step 803: Develop the 'Features' grid highlighting AI Generation, Real-time PvP, and Guilds
*   Step 804: Implement SEO-friendly HTML `<meta>` tags and OpenGraph data for social sharing
*   Step 805: Wire the 'Get Started' Call to Action (CTA) buttons to the Registration routing flow
*   Step 806: Write Jasmine unit tests ensuring the Landing Page renders perfectly without an active authentication token
*   Step 807: Optimize the loading of hero background images using modern WebP formats and `srcset`
*   Step 808: Implement a basic 'Contact Us' footer form integrating with a Go backend webhook
*   Step 809: Perform Google Lighthouse audits targeting a 95+ score for Performance and SEO
*   Step 810: Deploy the Landing Page and verify social preview cards on Twitter and Discord

### User Dashboard & Hub
*   Step 811: Develop the `<app-dashboard>` component serving as the primary authenticated route (`/hub`)
*   Step 812: Implement the 'Welcome Back' header fetching the user's display name and current level from NgRx
*   Step 813: Develop the 'Recent Matches' widget fetching and displaying the last 5 Spanner match records
*   Step 814: Develop the 'Active Study Groups' widget fetching the user's guild affiliations
*   Step 815: Implement the 'Quick Play' CTA button routing the user directly into the Normal Mode Matchmaker
*   Step 816: Write Jasmine unit tests verifying the Dashboard properly aggregates data from multiple distinct NgRx selectors
*   Step 817: Implement pull-to-refresh functionality on mobile devices to forcefully update the dashboard state
*   Step 818: Develop empty states (e.g., 'You haven't played any matches yet!') with contextual CTAs
*   Step 819: Conduct performance profiling verifying the Dashboard renders rapidly upon successful login
*   Step 820: Deploy the Dashboard view to the staging environment

### Quiz Preparation & Mode Selection UI
*   Step 821: Develop the `<app-mode-selection>` component presenting Normal, Speed, and Death Mode options
*   Step 822: Implement informative tooltips explaining the specific scoring and timer rules for each mode
*   Step 823: Develop the 'Topic Selection' input allowing the user to type custom strings or pick trending topics
*   Step 824: Wire the Topic Selection input to an auto-complete Go backend API fetching popular Spanner queries
*   Step 825: Implement the 'Find Match' button dispatching the Open Match `CreateTicketRequest` Protobuf via HTTP
*   Step 826: Implement the visual transition from 'Selecting Mode' to the 'Matchmaking Queue Overlay'
*   Step 827: Write Jasmine unit tests verifying the 'Find Match' button is disabled if no topic is selected
*   Step 828: Implement deep linking support allowing a user to navigate directly to `/play/speed`
*   Step 829: Conduct UX testing ensuring the mode selection cards are easily tappable on small mobile screens
*   Step 830: Review the complete flow from Dashboard to Queue to ensure zero dead-ends

### Active Match Gameplay UI
*   Step 831: Develop the `<app-active-match>` component rendering the primary real-time quiz interface
*   Step 832: Implement the massive countdown timer UI, pulsing red when under 5 seconds
*   Step 833: Develop the Question Display container, handling dynamic text sizing for exceptionally long generated questions
*   Step 834: Implement the 4 Answer Buttons, changing color (Green/Red) immediately upon receiving the server's correct answer broadcast
*   Step 835: Develop the 'Mini-Scoreboard' pinning the current top 3 players' scores to the top of the screen
*   Step 836: Implement particle effect animations (e.g., confetti) triggering locally upon selecting the correct answer
*   Step 837: Write Jasmine unit tests mocking the Agones WebSocket data stream to verify state transitions
*   Step 838: Implement a 'Quit Match' confirmation dialog warning the user about Elo penalties
*   Step 839: Optimize Angular Change Detection (`OnPush`) across the entire Active Match view to ensure 60fps rendering
*   Step 840: Conduct exhaustive testing across unstable network conditions simulating dropped WebSocket frames

### Post-Match Summary & Results Screen
*   Step 841: Develop the `<app-match-results>` component rendering after the Agones Server broadcasts the final state
*   Step 842: Implement the podium UI (1st, 2nd, 3rd place) highlighting the winner's avatar
*   Step 843: Develop the Elo change animation (e.g., counting up from 1500 to 1525)
*   Step 844: Implement the 'Coin and Gem Rewards' widget detailing the payout breakdown (Base + Win Bonus)
*   Step 845: Develop the 'Review Questions' tab allowing players to see the correct answers for questions they missed
*   Step 846: Wire the 'Review Questions' tab to the Spanner 'QuestionFeedback' API to rate the AI generation
*   Step 847: Implement a 'Play Again' CTA rapidly re-queueing the user with the same mode and topic
*   Step 848: Write Jasmine unit tests ensuring the Results Screen safely handles edge cases (e.g., 0 score ties)
*   Step 849: Conduct manual testing ensuring the Results Screen remains visible even after the Agones server terminates
*   Step 850: Deploy the complete Match lifecycle (Selection -> Gameplay -> Results) to staging

### Study Group Management Dashboard
*   Step 851: Develop the `<app-guild-detail>` component displaying the Study Group roster and statistics
*   Step 852: Implement the Angular Material UI for the Group Admin to invite new members via search
*   Step 853: Implement the Angular Material UI for the Group Admin to kick members or promote them to Officers
*   Step 854: Wire the member management UI to the Go backend Spanner mutation APIs
*   Step 855: Develop the 'Group Chat' tab embedding the Firebase RTDB chat component
*   Step 856: Develop the 'Group Leaderboard' tab displaying internal rankings of guild members
*   Step 857: Implement the 'Edit Group Details' modal (Name, Emblem, Privacy) restricted to Admins
*   Step 858: Write Jasmine unit tests verifying UI elements (Kick Button) are completely hidden from standard members
*   Step 859: Conduct end-to-end testing of the complete Study Group lifecycle (Creation -> Invite -> Chat -> Delete)
*   Step 860: Document the Study Group component hierarchy in the frontend architecture wiki

### Friends List & Social Hub
*   Step 861: Develop the `<app-social-hub>` component rendering as a persistent right-side drawer on desktop
*   Step 862: Implement the 'Friends List' tab fetching the Spanner friend records and joining with Firebase Presence data
*   Step 863: Develop the 'Add Friend' input resolving user display names to IDs via the Go backend
*   Step 864: Implement the 'Pending Requests' tab allowing users to accept or decline incoming friendships
*   Step 865: Wire the 'Accept Request' button to the Go backend Spanner API and NgRx state updates
*   Step 866: Develop the 'Direct Message' action button opening the private Firebase chat context
*   Step 867: Implement a 'Block User' action hiding all future interactions and dispatching to Spanner
*   Step 868: Write Jasmine unit tests validating the sorting logic (Online friends appear at the top of the list)
*   Step 869: Implement real-time list resorting as Firebase Presence events stream in
*   Step 870: Deploy the Social Hub and verify it does not negatively impact Active Match performance

### Player Profile & Statistics View
*   Step 871: Develop the `<app-player-profile>` component displaying detailed historical statistics
*   Step 872: Implement the 'Win Rate' and 'Total Matches' charts utilizing a lightweight charting library (e.g., Chart.js)
*   Step 873: Develop the 'Match History' paginated table fetching Spanner records via the Go API
*   Step 874: Implement the 'Favorite Topics' widget aggregating the user's most frequently played quiz subjects
*   Step 875: Develop the public read-only variant of the profile view for inspecting other players
*   Step 876: Wire the profile view to the Go backend API, handling 404 errors if the requested profile doesn't exist
*   Step 877: Implement the 'Edit Profile' routing for the authenticated user viewing their own profile
*   Step 878: Write Jasmine unit tests verifying the private 'Edit' button does not appear on public profile views
*   Step 879: Optimize the Go API aggregating statistics to ensure the Profile view loads in under 300ms
*   Step 880: Deploy the Player Profile view and verify SEO indexing for public profiles (if enabled)

### Settings & Preferences View
*   Step 881: Develop the `<app-settings>` component housing tabs for Account, Privacy, Audio, and UI
*   Step 882: Implement the 'Audio Settings' tab containing sliders for Master, SFX, and Music volumes
*   Step 883: Wire the audio sliders to a central Angular `AudioService` and persist values to `localStorage`
*   Step 884: Implement the 'Account Settings' tab containing the Change Email and Change Password Firebase flows
*   Step 885: Implement the 'Privacy Settings' tab allowing users to toggle 'Allow Friend Requests'
*   Step 886: Wire the Privacy toggles to the Go backend API persisting preferences in Spanner
*   Step 887: Write Jasmine unit tests ensuring the Settings form properly marks itself as 'Dirty' when values change
*   Step 888: Implement an 'Unsaved Changes' route guard preventing the user from navigating away before saving
*   Step 889: Conduct manual testing of the complete Settings configuration lifecycle
*   Step 890: Document the local storage key definitions in the engineering wiki

### Error Boundaries & 404 Pages
*   Step 891: Develop the `<app-error-boundary>` global component catching all unhandled runtime exceptions in Angular
*   Step 892: Implement logic to gracefully display a generic 'Something went wrong' UI instead of a blank white screen
*   Step 893: Wire the Error Boundary to dispatch the stack trace to a remote error logging service (e.g., Sentry or Google Cloud Error Reporting)
*   Step 894: Develop the `<app-not-found>` component serving as the catch-all `**` wildcard route
*   Step 895: Implement a 'Return to Hub' CTA on the 404 page
*   Step 896: Write Jasmine unit tests forcing a component failure to verify the Error Boundary catches it
*   Step 897: Develop a specialized 'Maintenance Mode' view intercepting traffic if the Go backend returns a specific 503 status
*   Step 898: Implement HTTP interceptor logic to detect 401 Unauthorized errors and force a logout/redirect
*   Step 899: Conduct comprehensive testing of all error states (Network Offline, 500 Internal Error, 403 Forbidden)
*   Step 900: Deploy the finalized error architecture to production

## Epic 10: Monetization, Analytics, and Launch Prep

### Stripe Payments Integration
*   Step 901: Provision a Stripe Account and configure the API keys in Google Cloud Secret Manager
*   Step 902: Install the Stripe Go SDK in the backend workspace
*   Step 903: Install the Stripe.js library in the Angular frontend workspace
*   Step 904: Define Protobuf schemas for the `CreateCheckoutSession` Go API request
*   Step 905: Implement Go backend logic to generate a secure Stripe Checkout Session URL
*   Step 906: Write the Go Webhook handler intercepting `checkout.session.completed` events from Stripe
*   Step 907: Implement strict Webhook signature verification in Go to prevent fraudulent transactions
*   Step 908: Write Spanner transactions to reliably credit the user's Gem balance upon successful webhook processing
*   Step 909: Write Go unit tests mocking the Stripe SDK to verify checkout session generation
*   Step 910: Conduct end-to-end testing in the Stripe Test Environment simulating successful and failed payments

### Virtual Shop: Avatar Cosmetics
*   Step 911: Design the Spanner DDL schema for the `CosmeticItems` catalog (ItemID, Name, Cost, Type)
*   Step 912: Design the Spanner DDL schema for the `UserInventory` table (UserID, ItemID)
*   Step 913: Implement the Angular Material `<app-shop>` component featuring a grid of purchasable Avatars
*   Step 914: Implement Go backend logic validating the user has sufficient Gems before processing a purchase
*   Step 915: Write strict Spanner transactions atomically deducting the Gem balance and inserting the `UserInventory` record
*   Step 916: Implement the Angular UI 'Equip' button updating the user's active Profile Avatar in Spanner
*   Step 917: Define NgRx state synchronizing the user's active inventory and available balances
*   Step 918: Write Jasmine unit tests verifying the 'Purchase' button disables if the Gem balance is too low
*   Step 919: Conduct load testing on the purchase API simulating concurrent transactions on the same account
*   Step 920: Deploy the Avatar Cosmetics shop to the staging environment

### Virtual Shop: UI Borders & Emotes
*   Step 921: Expand the `CosmeticItems` catalog in Spanner to include 'UI Borders' and 'Emotes'
*   Step 922: Implement the Angular UI tab rendering the purchasable Borders and previewing them dynamically
*   Step 923: Implement the Angular UI tab rendering the purchasable Emotes (used in Active Matches)
*   Step 924: Update the Agones Go Game Server to broadcast Emote events triggered by players during a match
*   Step 925: Update the Active Match Angular UI to render floating Emote animations over player avatars
*   Step 926: Implement Go backend logic verifying a user actually owns an Emote before allowing them to equip it
*   Step 927: Write Go integration tests ensuring unowned items cannot be equipped via direct API calls
*   Step 928: Conduct UX testing ensuring the Shop interface is intuitive and purchase confirmations are clear
*   Step 929: Verify Spanner inventory tables optimize correctly for rapid read queries
*   Step 930: Deploy the Borders and Emotes functionality to production

### Coin Pack Purchasing System
*   Step 931: Define the predefined 'Gem Packs' (e.g., 500 Gems for $4.99, 1200 Gems for $9.99) in Stripe Products
*   Step 932: Implement the Angular `<app-gem-store>` component displaying the pricing tiers
*   Step 933: Wire the Gem Store UI to the `CreateCheckoutSession` API routing to the Stripe hosted page
*   Step 934: Implement the Angular Success route (`/shop/success`) displaying a 'Thank You' animation
*   Step 935: Implement the Angular Cancel route (`/shop/cancel`) returning the user gracefully to the Shop
*   Step 936: Configure Cloud Monitoring alerts to trigger immediately if the Stripe Webhook failure rate spikes
*   Step 937: Implement Go backend logic to handle Stripe refunds, deducting the Gems from the user's Spanner balance
*   Step 938: Write Go integration tests verifying the webhook idempotency (processing the same event twice safely)
*   Step 939: Verify all financial transactions comply with baseline PCI requirements (no card data touches our servers)
*   Step 940: Execute a full tabletop review of the monetization architecture with the engineering team

### Google Analytics 4 Setup
*   Step 941: Provision a new Google Analytics 4 (GA4) property in the Google Marketing Platform
*   Step 942: Integrate the GA4 tracking snippet (gtag.js) into the Angular `index.html`
*   Step 943: Configure the Angular Router to dispatch GA4 `page_view` events upon successful route transitions
*   Step 944: Implement a centralized Angular `AnalyticsService` abstracting the `gtag` API calls
*   Step 945: Wire the `AnalyticsService` to dispatch a standard event when a Matchmaking Ticket is created
*   Step 946: Wire the `AnalyticsService` to dispatch a standard event when a Stripe Checkout Session is initiated
*   Step 947: Configure GA4 User Properties allowing segmentation by 'Premium' vs 'Standard' users
*   Step 948: Implement a strict Cookie Consent Banner in Angular to comply with GDPR/CCPA before firing GA4
*   Step 949: Write Jasmine unit tests validating the `AnalyticsService` does not execute if consent is denied
*   Step 950: Verify event ingestion in the GA4 Real-time Debug View dashboard

### Custom Conversion Events Tracking
*   Step 951: Define the exact custom conversion events (e.g., `quiz_completed`, `account_created`, `purchase_completed`)
*   Step 952: Implement Go backend logic to dispatch Server-Side tracking events to GA4 (Measurement Protocol) for critical conversions (e.g., Stripe Webhook success)
*   Step 953: Implement logic ensuring server-side events perfectly deduplicate against client-side events
*   Step 954: Configure Google Tag Manager (GTM) to manage the custom event triggers dynamically
*   Step 955: Wire the Angular `AnalyticsService` to push specific payloads into the GTM `dataLayer`
*   Step 956: Write Go unit tests mocking the GA4 Measurement Protocol API endpoint
*   Step 957: Build a custom GA4 Exploration report analyzing the funnel from 'Registration' to 'First Match Played'
*   Step 958: Configure Cloud Logging to alert on excessive failures when hitting the GA4 Measurement Protocol
*   Step 959: Conduct a full data integrity audit ensuring test accounts are excluded from production GA4 metrics
*   Step 960: Document the complete analytics schema (Events and Parameters) in the marketing wiki

### Search Engine Optimization (SEO) Metadata
*   Step 961: Install `@angular/platform-browser` components `Meta` and `Title` services
*   Step 962: Implement an Angular `SEOService` centralizing the injection of `<title>` and `<meta>` tags
*   Step 963: Configure the default OpenGraph tags (og:title, og:image, og:description) for the root Landing Page
*   Step 964: Implement dynamic SEO injection for public Player Profiles (e.g., 'View [PlayerName] stats on Babank!')
*   Step 965: Generate a dynamic `sitemap.xml` using a Go backend worker parsing public profiles and public Study Groups
*   Step 966: Generate standard `robots.txt` explicitly preventing indexing of internal `/hub` and `/match` routes
*   Step 967: Write Jasmine unit tests verifying the `SEOService` injects the correct DOM elements
*   Step 968: Perform Google Search Console verification and submit the dynamic `sitemap.xml`
*   Step 969: Audit site performance using Google Lighthouse specifically targeting the 'SEO' scoring pillar
*   Step 970: Deploy the SEO enhancements to production and monitor crawl rates

### Load Testing with Distributed Clients
*   Step 971: Select a distributed load testing framework (e.g., Locust or k6)
*   Step 972: Write a Python/JS load testing script simulating standard user registration and login
*   Step 973: Write a complex load testing script simulating the WebSocket upgrade and trivia answering flow
*   Step 974: Provision a temporary, high-capacity Google Kubernetes Engine cluster specifically for generating load
*   Step 975: Execute a baseline load test simulating 10,000 concurrent users navigating the Angular UI
*   Step 976: Execute an extreme load test simulating 5,000 concurrent players in active Agones matches
*   Step 977: Monitor Spanner CPU utilization, Vertex AI quotas, and Cloud Run scaling latency during the test
*   Step 978: Identify and remediate any critical bottlenecks discovered (e.g., unoptimized Spanner queries)
*   Step 979: Execute a 'Soak Test' running 1,000 concurrent users continuously for 24 hours to detect memory leaks
*   Step 980: Document the maximum theoretical concurrency limits of the finalized architecture

### Production Infrastructure Hardening
*   Step 981: Configure Google Cloud Armor (WAF) policies blocking known malicious IPs and typical SQLi/XSS attack vectors
*   Step 982: Implement Cloud Armor rate limiting at the edge to mitigate simple Layer 7 DDoS attacks
*   Step 983: Configure Google Cloud CDN to aggressively cache the Angular frontend static assets
*   Step 984: Audit all Go backend APIs ensuring they properly validate CORS headers and origins
*   Step 985: Review all Firebase Security Rules verifying there are zero wildcard (`true`) read/write paths
*   Step 986: Rotate all critical infrastructure secrets (Stripe Keys, Vertex AI Keys) in Google Secret Manager
*   Step 987: Perform a final audit of Google Cloud IAM ensuring no developer has raw 'Owner' access to the production project
*   Step 988: Configure automated backups for all critical systems (Spanner, RTDB, Cloud Storage)
*   Step 989: Conduct a Penetration Test (internal or third-party) focusing on the Go backend JWT validation
*   Step 990: Remediate any final High/Critical security vulnerabilities identified by Security Command Center

### Final Release Deployment & Tagging
*   Step 991: Execute the final code freeze on the `main` Git branch
*   Step 992: Draft the comprehensive `v1.0.0` Release Notes detailing the Launch capabilities
*   Step 993: Create the immutable Git Tag `v1.0.0` triggering the final production Cloud Build pipeline
*   Step 994: Verify the successful deployment of the Angular frontend to Cloud Run / Firebase Hosting
*   Step 995: Verify the successful deployment of the Go API to Cloud Run
*   Step 996: Verify the successful deployment of the final Agones Fleets to the GKE production cluster
*   Step 997: Execute the final suite of production smoke tests (Registration, Matchmaking, Purchase)
*   Step 998: Verify all Cloud Monitoring Dashboards and Alerting Policies are green and active
*   Step 999: Officially route the production custom domain (e.g., `play.babank.com`) to the Google Cloud Load Balancer
*   Step 1000: Celebrate the successful launch of Babank v1.0.0!

