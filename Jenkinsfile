#!/usr/bin/env groovy
// Jenkinsfile for datacore_interview monorepo
// Based on https://stackoverflow.com/questions/48751013/jenkins-multibranch-pipeline-only-for-subfolder
// Author: Alexey Pechenkin <alexey74@gmail.com>
//

def GO_VERSION = '1.15.6'
def PYTHON_VERSION = '3.7.7'

pipeline {
    agent any

    stages {
        stage('Build Go code') {
            agent {
                docker {
                    image "golang:${GO_VERSION}"
                    reuseNode true
                    args '-u root --privileged'
                }
            }
            when {
                beforeAgent true
                anyOf {
                    changeset "go/**"
                    changeset "Jenkinsfile"
                    branch 'master'
                    triggeredBy cause: "UserIdCause"
                }
            }
            steps {
                sh 'cd go && make all'
            }
        }
        stage('Build Python code') {
            agent {
                docker {
                    image "python:${PYTHON_VERSION}"
                    reuseNode true
                    args '-u root --privileged'
                }
            }
            when {
                beforeAgent true
                anyOf {
                    changeset "python/**"
                    changeset "Jenkinsfile"
                    branch 'master'
                    triggeredBy cause: "UserIdCause"
                }
            }
            steps {
                sh '''
                    cd python &&
                    pip install -r requirements.txt &&
                    pipenv install --dev --deploy --ignore-pipfile &&
                    pipenv run make all
                    '''
            }
        }
        stage('Publish Go artifacts') {
            agent {
                docker {
                    image "golang:${GO_VERSION}"
                    reuseNode true
                }
            }
            when {
                beforeAgent true
                branch 'master'
                anyOf {
                    changeset "go/**"
                    triggeredBy cause: "UserIdCause"
                }
            }
            steps {
                milestone 1
                lock('go-artifact-publish') {
                    sh 'cd go && make publish'
                }
            }
        }
        stage('Publish Python artifacts') {
            agent {
                docker {
                    image "python:${PYTHON_VERSION}"
                    reuseNode true
                }
            }
            when {
                beforeAgent true
                branch 'master'
                anyOf {
                    changeset "python/**"
                    triggeredBy cause: "UserIdCause"
                }
            }
            steps {
                milestone 2
                lock('python-artifact-publish') {
                    sh '''
                        cd python &&
                        make publish
                    '''
                }
            }
        }
    }

    post {
        always {
            archiveArtifacts(
                artifacts:   'go/output/golang-test.out,'
                           + 'go/output/golang-coverage.out,'
                           + 'go/output/golang-artefact,'
                           + 'python/output/python-test.out,'
                           + 'python/output/python-coverage.tar.gz,'
                           + 'python/output/takehome-*.tar.gz'
                , fingerprint: true, allowEmptyArchive: true
            )
        }
    }
}
