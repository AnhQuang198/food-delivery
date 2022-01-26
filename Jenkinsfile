#!/usr/bin/env groovy

node {
    properties([disableConcurrentBuilds()])

    try {

        project = "food-delivery"
        dockerRepo = "leequang198"
        dockerFile = "Dockerfile"
        imageName = "${dockerRepo}/${project}"
        buildNumber = "${env.BUILD_NUMBER}"

        stage('checkout code') {
            checkout scm
            sh "git checkout ${env.BRANCH_NAME} && git reset --hard origin/${env.BRANCH_NAME}"
        }

        stage('build') {
            sh "docker build -t ${imageName}:${env.BRANCH_NAME}-build-${buildNumber} -f ./Dockerfile ."
        }

        stage('push') {
            docker.withRegistry('https://registry.hub.docker.com', 'DockerHubAccount') {
                        app.push("${env.BUILD_NUMBER}")
                        app.push("latest")
                    }
        }

        switch (env.BRANCH_NAME) {
                    case 'develop':
                        stage('deploy-dev') {
                            echo "deploy-dev k8s"
                        }
                        break
                    case 'master':
                        stage('deploy-prod') {
                            echo "deploy-prod k8s"
                        }
                        break
                }
    } catch (e) {
        currentBuild.result = "FAILED"
        throw e
    }
}
