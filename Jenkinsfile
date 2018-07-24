#!/usr/bin/env groovy

node {

    stage('checkout') {
        checkout scm
    }

    def dockerImage
    def docketTag = 'latest'

    stage('build go') {
        sh "ls -las"
        sh "cp -R docker/* ."
        sh "go get github.com/sevlyar/go-daemon"
        sh "go build main.go"
    }

    stage('push docker'){
        docker.withRegistry('https://registry.djeg-test.shalb.com') {
            dockerImage.push(docketTag)
        }
    }

   def docker_host = '/dev/null'
   def stack_name = 'cools'

//    stage ('init params'){
//        switch("${env.env}"){
//            case 'cools-dev':
//                docker_host='10.10.1.209'
//                stack_name='cools-dev'
//                break;
//                            }
//          }

//     stage('deploy'){
//         sh "docker -H tcp://${docker_host}:2376 service update --image cools-registry:5000/siteapp-engine:${docketTag} ${stack_name}_siteapp-engine"
//        }


}
