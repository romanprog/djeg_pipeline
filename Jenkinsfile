#!/usr/bin/env groovy

node {

    stage('checkout') {
        checkout scm
    }

    stage('build go') {
        sh "ls -las"
        sh "cp -R docker/* ."
        sh "go build main.go"
        sh "./main"
        echo 'Its new jenkinsfile for test 981..'
    }

}
