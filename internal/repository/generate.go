package repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate ../../bin/minimock -i UserRepository -o ./mocks/ -s "_minimock.go"
//go:generate ../../bin/minimock -i UserDataRepository -o ./mocks/ -s "_minimock.go"
