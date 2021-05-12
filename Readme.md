[![Build Status](https://www.travis-ci.com/mes1234/syncbrok.svg?branch=master)](https://www.travis-ci.com/mes1234/syncbrok)
<br>
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=mes1234_syncbrok&metric=alert_status)](https://sonarcloud.io/dashboard?id=mes1234_syncbrok)

# Syncbrok

Syncbrok is a simple fun project to learn some Golang concepts

# Usage
Compile syncbrok:
```
cmd/syncbrok/main.go
```
Compile HTTP subscriber:
```
cmd/subscriber/subs.go
```

Run both programs make sure to put inside the same folder :
```
config.yml
```

## Send new message:
```
http://172.30.83.211:10000/msg
```
Add headers:
```
queue: <queueName>
ParentId: <parentId> - optional
```
Body:
```
<content> 
```

## Send new queue:
```
http://172.30.83.211:10000/queue
```
Add headers:
```
queue: <queueName>
```
## Send new subscriber:
```
http://172.30.83.211:10000/subscrib
```
Add headers:
```
queue: <queueName>
endpoint: <address> eg. http://localhost:20000
```

