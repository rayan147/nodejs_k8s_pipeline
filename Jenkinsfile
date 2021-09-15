pipeline {
    agent any
    environment {
  imageName = 'rayan147/nodejsapp'
  registryCredentialSet = 'dockerhub'
  
}
    stages {
           stage('Build Docker Image') {
            steps {
                script {
                  sh 'docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t rayan147/nodejsapp:latest --push .'
                }
            }
        }
        stage('Test') {
            steps {
                echo 'Testing...'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}