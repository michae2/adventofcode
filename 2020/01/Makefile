# Targets:
# make, make all:                 build all solutions
# make py, make cpp:              build all solutions for one language
# make py/a, make cpp/b:          build solution for one part and one language
# make run, make run/all:         run all solutions
# make run/py, make run/cpp:      run all solutions for one language
# make run/py/a, make run/cpp/b:  run solution for one part and one language
# make clean:                     remove all temporary files

langs = py cpp
runlangs = $(addprefix run/,$(langs))
runs = $(addsuffix /a,$(runlangs)) $(addsuffix /b,$(runlangs))

CXX = g++-10
CXXFLAGS = -std=c++20 -Wall -Wextra -pedantic-errors -g

.PHONY: all $(langs)
all: $(langs)
$(langs): %: %/a %/b

py/%: %.py
	mkdir -p py
	ln $< $@

cpp/%: %.cpp
	mkdir -p cpp
	$(CXX) $< -o $@ $(CXXFLAGS)

.PHONY: run run/all $(runlangs) $(runs)
run: run/all
run/all: $(runlangs)
$(runlangs): run/%: run/%/a run/%/b
$(runs): run/%: %
	$< < input.txt

.PHONY: clean
clean:
	rm -rfv *~ $(langs)
