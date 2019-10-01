FROM mysql-server
RUN sudo apt-get update
RUN sudo apt-get golang-go
WORKDIR SkillMatrix/
CMD go run main.go
