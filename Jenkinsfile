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
            echo "stage checkout"
        }

        stage('build') {
            echo "stage build"
        }
        stage('push') {
            echo "stage push"
        }
    } catch (e) {
        currentBuild.result = "FAILED"
        throw e
    }
}
