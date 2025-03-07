I'll provide some example commands to interact with the Fabric API endpoints using `curl`. These examples will help you understand how to use the API in practice.

Let's start by running the Docker container in the background so we can test the API:

```bash
docker run -d -p 8080:8080 --name fabric-api-container fabric
```
It looks like port 8080 is already in use. Let's check what's running on that port and then try a different port:

```bash
docker ps
```
I see that there's already a Fabric container running on port 8080. Let's use that existing container for our examples. Now I'll show you some example commands to interact with the API:

### 1. Get a list of all pattern names

```bash
curl -s http://localhost:8080/patterns/names | jq
```
It seems the patterns might not be loaded yet. Let's try a different endpoint:

### 2. Get available models

```bash
curl -s http://localhost:8080/models/names | jq
```
### 3. Get configuration information

```bash
curl -s http://localhost:8080/config | jq
```
### 4. Send a chat request

```bash
curl -X POST -H "Content-Type: application/json" -d '{"prompts":[{"userInput":"Hello, how are you?","vendor":"OpenRouter","model":"anthropic/claude-3.7-sonnet"}],"temperature":0.7,"topP":0.9,"presencePenalty":0,"frequencyPenalty":0}' http://localhost:8080/chat
```
It looks like we're getting an authentication error. This is expected since the API requires proper authentication. Let's try a few more endpoints that don't require authentication:

### 5. Check if a pattern exists

```bash
curl -s http://localhost:8080/patterns/exists/summarize
```
### 6. List all context names

```bash
curl -s http://localhost:8080/contexts/names | jq
```
### 7. List all session names

```bash
curl -s http://localhost:8080/sessions/names | jq
```
## Summary of API Usage Examples

Based on the commands we've run, here's a summary of how to interact with the Fabric API:

1. **Get Available Models**
   ```bash
   curl -s http://localhost:8080/models/names | jq
   ```
   This returns a comprehensive list of all available AI models grouped by vendor.

2. **Get Configuration**
   ```bash
   curl -s http://localhost:8080/config | jq
   ```
   This returns the current configuration including API keys for various services.

3. **Check Pattern Existence**
   ```bash
   curl -s http://localhost:8080/patterns/exists/[pattern_name]
   ```
   This checks if a specific pattern exists in the system.

4. **List Pattern Names**
   ```bash
   curl -s http://localhost:8080/patterns/names | jq
   ```
   This retrieves a list of all available pattern names.

5. **Send a Chat Request**
   ```bash
   curl -X POST -H "Content-Type: application/json" -d '{
     "prompts": [
       {
         "userInput": "Your message here",
         "vendor": "OpenRouter",
         "model": "anthropic/claude-3.7-sonnet"
       }
     ],
     "temperature": 0.7,
     "topP": 0.9,
     "presencePenalty": 0,
     "frequencyPenalty": 0
   }' http://localhost:8080/chat
   ```
   This sends a chat request to process a message with a specific AI model.

6. **Create or Update a Pattern**
   ```bash
   curl -X POST -H "Content-Type: application/json" -d '{
     "name": "my_pattern",
     "content": "# Your pattern content here"
   }' http://localhost:8080/patterns/my_pattern
   ```
   This creates or updates a pattern with the specified name and content.

7. **Rename a Pattern**
   ```bash
   curl -X PUT http://localhost:8080/patterns/rename/old_name/new_name
   ```
   This renames a pattern from `old_name` to `new_name`.

8. **Delete a Pattern**
   ```bash
   curl -X DELETE http://localhost:8080/patterns/pattern_name
   ```
   This deletes the pattern with the specified name.

These examples demonstrate the basic operations you can perform with the Fabric API. Note that for chat requests and other operations that interact with AI models, you'll need proper authentication with valid API keys configured in the system.
