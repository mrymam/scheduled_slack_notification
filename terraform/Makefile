init: rsrc=
init:
	terraform -chdir=resources/main/$(rsrc) init -backend-config="./.tfbackend"

init-all:
	make init rsrc=services
	make init rsrc=image
	make init rsrc=secrets
	make init rsrc=job
	make init rsrc=schedule

plan: rsrc=
plan:
	terraform -chdir=resources/main/$(rsrc) plan

plan-all:
	make plan rsrc=services
	make plan rsrc=image
	make plan rsrc=secrets
	make plan rsrc=job
	make plan rsrc=schedule

apply: rsrc=
apply:
	terraform -chdir=resources/main/$(rsrc) apply -auto-approve

apply-all:
	make apply rsrc=services
	make apply rsrc=image
	make apply rsrc=secrets
	make apply rsrc=job
	make apply rsrc=schedule
