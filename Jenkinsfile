pipeline {
    agent any
    tools {
        go 'Go 1.13'
    }
    environment {
        GO111MODULE = 'on'
    }

    stages {
        stage('Compile') {
            steps {
                sh 'go build'
            }
        }
        stage('Test') {
            steps {
                sh 'go test ./... -coverprofile=coverage.txt'
               }
        }
    }
}