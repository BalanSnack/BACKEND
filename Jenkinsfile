pipeline{
    agent any
    stages{
        stage('git clone') {
            steps {
                checkout scm
                echo "git clone..."
            }
        }
        stage('go build'){
            steps{
                sh 'go build'
                echo "go build..."
            }
        }
    }
}