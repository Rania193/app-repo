pipeline {

    agent { label "jenkins_slave"}

    // parameters {
    //     choice(name: 'Choose either to build or deploy', choices: ['Build', 'Deploy'])
    // } 


    
    stages {
        stage('Build') {
            steps {
                echo 'building...'
                script{

                    //if (params.ENV == "Build") {
                        sh '''
                            echo 'debugging'
                        '''
                        withCredentials([usernamePassword(credentialsId: 'dockeriti', usernameVariable: 'USERNAME_ITI', passwordVariable: 'PASSWORD_ITI')]) {
                            sh '''
                                docker login -u ${USERNAME_ITI} -p ${PASSWORD_ITI}
                                echo 'logged in to docker'
                                docker build -t rania199/final_project:v${BUILD_NUMBER} .
                                echo 'built image'
                                docker push rania199/final_project:v${BUILD_NUMBER}
                                echo 'pushed image'
                                echo ${BUILD_NUMBER} > ../build.txt
                            '''
                     }
                    //}

                }

            }
        }



        stage('Deploy') {

            
            steps {
                echo 'Deploying...'


                script {
                    //if (params.ENV == "Deploy") {
                        withCredentials([file(credentialsId: 'jenkinsiti', variable: 'KITI')]) {
                            sh '''
                                export BUILD_NUMBER=$(cat ../build.txt)
                                mv deployment/deploy.yaml deployment/deploy.yaml.tmp
                                cat deployment/deploy.yaml.tmp | envsubst > deployment/deploy.yaml
                                rm -f deployment/deploy.yaml.tmp
                                kubectl apply -f deployment --kubeconfig ${KITI} -n app
                                echo 'done'
                             '''
                        }
                    }
                }


            }

            
       // }
    }


}