all:
	@make main
	@make foo
	@make foo-yq
	@make bar
	@make bar-yq

main:
	cat main.yml | gobble run -c -

foo:
	gobble utils yaml-merge main.yml foo.yml | gobble run -c -

bar:
	gobble utils yaml-merge main.yml bar.yml | gobble run -c -

foo-yq:
	yq eval-all 'select(fileIndex == 0) * select(fileIndex == 1)' main.yml foo.yml | gobble run -c -

bar-yq:
	yq eval-all 'select(fileIndex == 0) * select(fileIndex == 1)' main.yml bar.yml | gobble run -c -
