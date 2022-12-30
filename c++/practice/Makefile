CXXFLAGS = -std=c++17
objects := $(patsubst %.cc,%.o,$(wildcard *.cc))
main : $(objects)
	    $(CXX) $(CXXFLAGS) -o $@ $(objects)
