#!/usr/bin/env groovy

node {

    stage('checkout') {
       checkout([$class: 'GitSCM',
          branches: [[name:  env.BRANCH_NAME]],
          doGenerateSubmoduleConfigurations: false,
          extensions: [],
          submoduleCfg: [],
          userRemoteConfigs: [[credentialsId: 'gitsshkey', url: 'https://github.com/romanprog/djeg_pipeline']]
       ])
    }

    def dockerImage
    def docketTag = env.BRANCH_NAME.replaceAll("/", "-")

    stage('build docker') {
        sh "ls -las"
        sh "cp -R docker/* ."
        dockerImage = docker.build("go-daemon:${docketTag}", '.')
    }

    stage('push docker'){
        docker.withRegistry('https://registry.djeg-test.shalb.com') {
            dockerImage.push(docketTag)
            if ("${env.push_latest}" == "true"){
                 dockerImage.push('latest')
            }
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
