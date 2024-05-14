ifndef f
$(error file is not defined. Pass it as an file to make)
endif

filename := $(firstword $(f))

run: 
	g++ -std=c++11 -o main $(filename) && ./main