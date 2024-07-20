pipeline {

    agent any

//     tools {
// //         go 'go1.22'
//     }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }

    stages {
        stage("unit-test") {
            steps {
                echo 'UNIT TEST EXECUTION STARTED'
                sh 'make unit-tests'
            }
        }
        stage("functional-test") {
            steps {
                echo 'FUNCTIONAL TEST EXECUTION STARTED'
                sh 'make functional-tests'
            }
        }
        stage("build") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'go version'
                sh 'go get ./...'
                sh 'docker build . -t go-jenkins-simple/go-micro'
            }
        }
//         stage("server-running") {
//             steps {
//                 echo 'SERVER RUNNING STARTED'
//                 sh 'go version'
//                 sh 'go get ./...'
//                 sh 'docker build . -t go-jenkins-simple/go-micro'
//             }
//         }
        stage("deploy") {
            steps {
                echo 'DEFINE YOUR DEPLOYMENT SCRIPT!'
            }
        }

//         stage('Docker Push') {
//             agent any
//             steps {
//                 withCredentials([usernamePassword(credentialsId: 'dockerhub', passwordVariable: 'dockerhubPassword', usernameVariable: 'dockerhubUser')]) {
//                 sh "docker login -u ${env.dockerhubUser} -p ${env.dockerhubPassword}"
//                 sh 'docker push go-jenkins-simple/go-micro'
//                 }
//             }
        }
    }
}