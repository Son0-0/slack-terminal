# slack-terminal

![slt](https://user-images.githubusercontent.com/81317358/189177174-c5971d8b-ac78-4ead-b433-4b6c331e6575.jpg)

## Description
- Slack 채널에 명령어를 입력하면 서버 터미널에서 해당 명령어 실행 후 결과 반환
  - API 서버 코드 Go로 작성 (net/http)
  - AWS EC2 Ubuntu 20.04 LTS 환경에서 동작 확인
- /server tail out.log 명령어 입력시 log 반환

## Installation (Slack App 생성 후 Slash Command 설정 필수)
1. ```$ git clone https://github.com/Son0-0/slack-terminal.git```
2. ```$ go build main.go```
3. ```$ ./main```

## Settings
1. PORT_NUM 설정 (기본 설정 9999)
2. [Slack App Slash Command 설정](https://velog.io/@supssson/Go-Slack을-터미널로-사용하기)
  - https://api.slack.com 에서 App 생성
  - Slash Command 설정

## How to run Go Server
1. main.go 파일 실행
  - 1.1 ```$ go run main.go```
  - 1.2 ```$ go build main.go && nohup ./main &```  
    - 서버 내 백그라운드 실행하고 싶은 경우

## Slack Message
<img width="407" alt="스크린샷 2022-09-09 오전 1 35 14" src="https://user-images.githubusercontent.com/81317358/189177214-eb48ff5b-337d-4cb8-a0b7-24e88e44fbef.png">
