#!/usr/bin/env groovy

node {

    stage('checkout') {
        checkout scm
    }

    stage('build go') {
        sh "ls -las"
        echo "PR test mb"
        sh "cp -R docker/* ."
        sh "go build main.go"
        sh "./main"
    }

}
