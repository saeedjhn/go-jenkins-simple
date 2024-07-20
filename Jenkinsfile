pipeline {

    agent any

    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage('Checkout') {
                steps {
                    git 'https://github.com/saeedjhn/go-jenkins-simple.git'
                }
        }
        stage("unit-test") {
            steps {
                echo 'UNIT TEST EXECUTION STARTED'
                sh 'make unit-tests'
            }
        }
//         stage("functional-test") {
//             steps {
//                 echo 'FUNCTIONAL TEST EXECUTION STARTED'
//                 sh 'make functional-tests'
//             }
//         }
        stage("build") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'go version'
                sh 'make build'
            }
        }

//         stage('Docker Push') {
//             steps {
//                 withCredentials([usernamePassword(credentialsId: 'dockerhub', passwordVariable: 'dockerhubPassword', usernameVariable: 'dockerhubUser')]) {
//                 sh "docker login -u ${env.dockerhubUser} -p ${env.dockerhubPassword}"
//                 sh 'docker push go-jenkins-simple/go-micro'
//                 }
//             }
        stage("deploy") {
            steps {
                echo 'DEFINE YOUR DEPLOYMENT SCRIPT!'
                sh 'make run'
            }
        }
    }
}