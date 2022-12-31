CXXFLAGS = -std=c++17
objects := $(patsubst %.cc,%.o,$(wildcard *.cc))
main.out : $(objects)
	    $(CXX) $(CXXFLAGS) -o $@ $(objects)
