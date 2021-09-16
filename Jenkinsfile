pipeline {
    agent any
    environment {
      registry = 'rayan147/nodejsapp:latest'
      registryCredential = '79c0e82e-d596-4999-90bd-d759317cda6f'
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
                    def version = readFile('VERSION')
                    def versions = version.split('\\.')
                    def major = versions[0]
                    def minor = versions[0] + '.' + versions[1]
                    def patch = version.trim()

                    dockerImage = docker.build(registry)

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
                        dockerImage.push(major)
                        dockerImage.push(minor)
                        dockerImage.push(patch)
                    }
                } 
            }
        } 
        
        stage('Deploy to k8s'){
            steps {
                script {
                    sh 'sudo kubectl apply -f nodejsapp.yaml'
                }
            }
        }
    }
}