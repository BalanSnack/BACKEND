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
        stage('deploy'){
            steps{
                sshPublisher(
                    continueOnError: false, 
                    failOnError: true,
                    publishers: [
                        sshPublisherDesc(
                        configName: "backend",
                        transfers: [
                            sshTransfer(
                                sourceFiles: 'backend',
                                remoteDirectory: '/usr/local/backend',
                                execCommand: './backend'
                            )
                        ],
                        verbose: true
                        )
                    ]
                )
            }
        }
    }
}