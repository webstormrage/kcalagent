# Setup
1. Install golang 1.23.4
2. Setup dependencies
```
go mod tidy
```
3. create PostgreSQL empty Database
4. create .env file in project root
```
DATA_SOURCE_NAME='host=$dbhost port=$dbport user=$dbuser password=$dbpassword dbname=$dbname'
GEN_AI_API_KEY=$gemini_api_token
```
3. create logs/ folder in project root
4. create empty input.txt file in logs/ folder


# Usage

1. Print products you what to add in meals journal in logs/input.txt
2. Execute program
```
go run cmd/main.go
```
3. Follow program instructions