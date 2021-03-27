# Targets:
# make, make all:                     build all solutions
# make py, make cpp, etc:             build all solutions for one language
# make py/a, make cpp/b, etc:         build solution for one part and one lang
# make run, make run/all:             run all solutions
# make run/py, make run/cpp, etc:     run all solutions for one lang
# make run/py/a, make run/cpp/b, etc  run solution for one part and one lang
# make clean:                         remove all temporary files

langs = py cpp go # scm
runlangs = $(addprefix run/,$(langs))
runs = $(addsuffix /a,$(runlangs)) $(addsuffix /b,$(runlangs))

CXX = g++-10
CXXFLAGS = -std=c++20 -Wall -Wextra -pedantic-errors -g

GO = go build

SCM = scheme
SCMFLAGS = --optimize-level 3

.PHONY: all $(langs)
all: $(langs)
$(langs): %: %/a %/b

py/%: %.py
	mkdir -p py
	ln $< $@

cpp/%: %.cpp
	mkdir -p cpp
	$(CXX) -o $@ $< $(CXXFLAGS)

go/%: %.go
	mkdir -p go
	$(GO) -o $@ $<

scm/%: %.scm
	mkdir -p scm
	echo '(compile-program "$<")' | $(SCM) -q $(SCMFLAGS)
	mv $*.so $@

.PHONY: run run/all $(runlangs) $(runs)
run: run/all
run/all: $(runlangs)
$(runlangs): run/%: run/%/a run/%/b
$(runs): run/%: %
	$< < input.txt

.PHONY: clean
clean:
	rm -rfv *~ $(langs)
