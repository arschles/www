.PHONY: socialgen
socialgen:
	cd socialgen && \
		go build -o socialgen . && \
		./socialgen -title ${TITLE} -out ../cover.png
		mkdir -p static/images/${SLUG}
		mv cover.png static/images/${SLUG}/cover.png
		@echo "cover.png moved to static/images/${SLUG}/cover.png"
		@echo "Set this in your front matter:"
		@echo "meta_img: '/images/${SLUG}/cover.png'"