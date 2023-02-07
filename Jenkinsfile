node{
    stage('clone repository'){
        checkout scm
    }
    stage('build image'){
        image = docker.build("balansnack/balansnack-backend")
    }
    stage('push image'){
        docker.withRegistry('https://registry.hub.docker.com', 'dockerhub'){
            image.push()
        }
    }
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