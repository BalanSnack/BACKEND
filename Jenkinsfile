node{
    checkout scm
    def backendImage = docker.build("balansnack-backend-image")
    backendImage.push()
}

// stage('deploy'){
//             steps{
//                 sshPublisher(
//                     continueOnError: false, 
//                     failOnError: true,
//                     publishers: [
//                         sshPublisherDesc(
//                         configName: "backend",
//                         transfers: [
//                             sshTransfer(
//                                 sourceFiles: 'backend',
//                                 remoteDirectory: '.',
//                                 execCommand: 'chmod +x backend && ./backend'
//                             )
//                         ],
//                         verbose: true
//                         )
//                     ]
//                 )
//             }
//         }