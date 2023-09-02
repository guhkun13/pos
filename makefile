dev: 
	bin/air -c .air.toml

generate-table:
	tables-to-go -v -d pizza_app -u groot -p iamgroot -of internal/domain/$(domain) -
