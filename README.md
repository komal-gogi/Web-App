The project explains the step to step integration with CircleCi

Part one: Writing a shell script to read the messages

- create a directory with .circleci and add a config file like
  config.yml
- the indentation has to be written correctly 
- sign-up to circleci account with either Github or Bitbucket (build free)
- add the following contents to config.yml file

version: 2
jobs:
  build:
    docker:
      - image: alpine:3.7
    steps:
      - run:
        name: The First Step
        command: |
        echo 'Hello'
        echo 'this is delivery product'

Part Two: Adding checkout commands to get subsequent steps

- now add seconf run step and do ls -al

Part three: adding reference to docker image for build job