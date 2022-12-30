CXXFLAGS = -std=c++11
objects := $(patsubst %.cc,%.o,$(wildcard *.cc))
main : $(objects)
	    $(CXX) $(CXXFLAGS) -o $@ $(objects)
