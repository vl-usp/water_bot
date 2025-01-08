package storage

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate ../../bin/minimock -i User -o ./mocks/ -s "_minimock.go"
//go:generate ../../bin/minimock -i UserCache -o ./mocks/ -s "_minimock.go"
//go:generate ../../bin/minimock -i Reference -o ./mocks/ -s "_minimock.go"
