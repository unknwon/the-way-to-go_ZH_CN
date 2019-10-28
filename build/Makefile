OUTPUT_FILENAME = the-way-to-go_ZH_CN

all: epub
.PHONY : all

epub:
	pandoc -o $(OUTPUT_FILENAME).epub \
		--toc-depth=2 \
		--resource-path=../eBook \
		../eBook/*.md
.PHONY : epub

clean:
	rm -f $(OUTPUT_FILENAME).epub
.PHONY : clean

help:
	@echo "... all"
	@echo "... epub"
	@echo "... clean"
	@echo "... help"
.PHONY : help
