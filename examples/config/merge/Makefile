all:
	@make main
	@make foo
	@make bar

main:
	cat main.yml | gobble run -c -

foo:
	yq eval-all 'select(fileIndex == 0) * select(fileIndex == 1)' main.yml foo.yml | gobble run -c -

bar:
	yq eval-all 'select(fileIndex == 0) * select(fileIndex == 1)' main.yml bar.yml | gobble run -c -
