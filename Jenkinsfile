#!/usr/bin/env groovy

node {
    properties([disableConcurrentBuilds()])

    try {

        project = "fd-service"
        dockerRepo = "leequang198"
        dockerFile = "Dockerfile"
        imageName = "${dockerRepo}/${project}"
        buildNumber = "${env.BUILD_NUMBER}"

        stage('checkout code') {
            echo "stage checkout testttttt"
        }

        stage('build') {
            echo "stage build test"
        }
        stage('push') {
            echo "stage push"
        }
    } catch (e) {
        currentBuild.result = "FAILED"
        throw e
    }
}
