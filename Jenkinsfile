pipeline {
    agent any
    tools {
        go 'go1.14'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
    }
    stages {        
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
              
            }
        }
        
        stage('Execute') {
            steps {
                echo 'Compiling and building'
                sh 'go run main.go'
            }
        }  
}
}