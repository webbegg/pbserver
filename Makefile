:build
	@go build -o release/central.exe .

:run build
	@release/central.exe serve
