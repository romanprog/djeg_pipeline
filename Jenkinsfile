#!/usr/bin/env groovy

node {

    stage('checkout') {
        checkout scm
    }

    stage('build go') {
        sh "ls -las"
        sh "cp -R docker/* ."
        sh "go get github.com/sevlyar/go-daemon"
        sh "go build main.go"
    }


}
