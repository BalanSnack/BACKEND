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
                        configName: "my-ssh-connection",
                        transfers: [
                            sshTransfer(
                                sourceFiles: 'backend',
                                remoteDirectory: '/usr/local/backend',
                                execCommand: 'cd /usr/local/backend && ./backend'
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