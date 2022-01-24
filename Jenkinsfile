#!/usr/bin/env groovy

node {
    properties([disableConcurrentBuilds()])

    try {

        project = "food-delivery"
        dockerRepo = "leequang198"
        dockerFile = "Dockerfile"
        imageName = "${dockerRepo}/${project}"
        buildNumber = "${env.BUILD_NUMBER}"
        registryCredential = "DockerHubAccount"

        stage('checkout code') {
            checkout scm
            sh "git checkout ${env.BRANCH_NAME} && git reset --hard origin/${env.BRANCH_NAME}"
        }

        stage('build') {
            sh "docker build -t ${imageName}:${env.BRANCH_NAME}-build-${buildNumber} -f ./Dockerfile ."
        }

        stage('push') {
            sh "docker.withRegistry( '', registryCredential ) {
                docker push ${imageName}:${env.BRANCH_NAME}-build-${buildNumber}
            }"
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
