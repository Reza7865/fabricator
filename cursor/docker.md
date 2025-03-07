<think>To understand how the Docker file works for the Fabric project, I need to first locate it and examine its contents. Let me search for Dockerfile in the repository.</think>Let me search for the Dockerfile in the repository to understand how containerization works in this project.

Ran tool
Great! Let's look at the Dockerfile to understand how the container is set up:

Read file: Dockerfile
Let's also check if there are any Docker Compose files that might provide additional context:

Ran tool
Let's examine the Docker Compose file:

Read file: docker-compose.yml
Let's also look for the ENV file referenced in the Dockerfile:

Ran tool
Let's examine the ENV file:

Read file: ENV
Let's see if there's any documentation about Docker usage:

Ran tool
Let me check if there's a .dockerignore file to understand what files are excluded from the Docker build:

Read file: .dockerignore
Based on the files I've examined, I now have a good understanding of how the Docker setup works in the Fabric project. Let me explain:

## How Fabric's Docker Setup Works

### Docker Build Process

The `Dockerfile` uses a multi-stage build approach:

1. **Builder Stage**:
   - Uses `golang:1.23.3-alpine` as the base image
   - Sets up the working directory at `/app`
   - Copies Go module files (`go.mod` and `go.sum`) and downloads dependencies
   - Copies the entire source code and builds the application
   - Produces a statically linked binary named `fabric`

2. **Final Stage**:
   - Uses `alpine:latest` as a lightweight base image
   - Copies only the compiled binary from the builder stage
   - Copies the `patterns` directory containing prompt patterns
   - Sets up a clean configuration directory at `/root/.config/fabric`
   - Adds the ENV file with environment variables
   - Exposes port 8080 for the API server
   - Sets the entrypoint to run the fabric binary with the `--serve` flag

### Docker Compose Configuration

The `docker-compose.yml` file:
- Defines a service named `fabric-api`
- Builds the Docker image using the current directory's Dockerfile
- Maps port 8080 on the host to port 8080 in the container
- Mounts the ENV file as a read-only volume
- Sets the GIN_MODE environment variable to "release" for production mode

### Environment Variables

The `ENV` file contains essential configuration:
- Default AI vendor (OpenRouter)
- Default AI model (gemini-2.0-flash-exp)
- Model context length
- Pattern repository configuration
- API keys for various services:
  - OpenRouter
  - YouTube
  - Jina AI
  - Gemini
  - OpenAI

### Files Excluded from Docker Build

The `.dockerignore` file excludes:
- Git-related files (.git, .gitignore)
- Environment files (.env)
- Documentation (README.md)
- Docker-related files (docker-compose.yml, Dockerfile)

### Running the Docker Container

To run Fabric using Docker:

1. **Using Docker Compose (recommended)**:
   ```bash
   docker-compose up
   ```
   This will build the image if needed and start the container with the API server running on port 8080.

2. **Using Docker directly**:
   ```bash
   # Build the image
   docker build -t fabric .
   
   # Run the container
   docker run -p 8080:8080 -v ./ENV:/root/.config/fabric/.env:ro fabric
   ```

### Key Points to Remember

1. The Docker setup focuses on running Fabric in API server mode (`--serve`)
2. The container includes all pattern definitions from the repository
3. API keys and configuration are managed through environment variables
4. The configuration is stored in the container at `/root/.config/fabric`
5. You can customize the environment variables by modifying the ENV file

This Docker setup makes it easy to deploy Fabric as a standalone API server without having to install dependencies directly on your host system.
