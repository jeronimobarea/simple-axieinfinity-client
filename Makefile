
run_tests:
	docker build -t go_test_env -f Dockerfile .
	docker run \
    	--rm \
    	go_test_env
