
init-tfvers:
	cp default.tfvars terraform.tfvars
	cp default.tfbackend terraform.tfbackend

init-all:
	terraform -chdir=resources/main/image/ init -reconfigure -backend-config=../../../terraform.tfbackend

plan-all:
	terraform -chdir=resources/main/image/ plan -var-file ../../../terraform.tfvars

apply-all:
	terraform -chdir=resources/main/image/ apply -var-file ../../../terraform.tfvars
