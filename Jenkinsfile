pipeline {
    agent any
    environment {
      imageName = 'rayan147/nodejsapp'
      registryCredentialSet = 'dockerhub'
      dockerImage = '' 
  
}
    stages {
          stage('Cloning our Git') { 
            steps { 
                git 'https://github.com/rayan147/nodejs_k8s_pipeline.git' 
            }
        } 
          stage('Building our image') { 
            steps { 
                script { 
                    dockerImage = docker.build registry + ":$BUILD_NUMBER" 

                }
            } 
        }
        stage('Test') {
            steps {
                echo 'Testing...'
            }
        }
          stage('Deploy our image') { 
            steps { 
                script { 
                    docker.withRegistry( '', registryCredential ) { 
                        dockerImage.push() 
                    }
                } 
            }
        } 
    }
}