all: gen

clean:
	rm -rf informers kubernetes listers

gen: clean
	cp -r expansions/* ./
	bash cmd/autogenerate.sh

