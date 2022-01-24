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
        mysqlConnect = ""demo:admin123456@tcp(mysql:3306)/food_delivery?parseTime=true""

        stage('checkout code') {
            checkout scm
            sh "git checkout ${env.BRANCH_NAME} && git reset --hard origin/${env.BRANCH_NAME}"
        }

        stage('build') {
            sh "docker build -t ${imageName}:${env.BRANCH_NAME}-build-${buildNumber} -f ./Dockerfile ."
        }

        stage('push') {
            sh "docker run -d --name ${imageName}:${env.BRANCH_NAME}-build-${buildNumber} --network my-net -e mySqlConnect=${mysqlConnect} -p 3500:8080 ${project}"
        }
    } catch (e) {
        currentBuild.result = "FAILED"
        throw e
    }
}
