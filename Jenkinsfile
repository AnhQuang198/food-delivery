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
            checkout scm
            sh "git checkout ${env.BRANCH_NAME} && git reset --hard origin/${env.BRANCH_NAME}"
        }

        stage('build') {
            echo "stage build test"
        }
        stage('push') {
            echo "stage push"
        }
        switch (env.BRANCH_NAME) {
                    case 'develop':
                        stage('deploy-dev') {
                            echo "push dev-kubernetes"
                        }
                        break
                    case 'master':
                        stage('deploy-prod') {
                            echo "push prod-kubernetes"
                        }
                        break
                }
    } catch (e) {
        currentBuild.result = "FAILED"
        throw e
    }
}
