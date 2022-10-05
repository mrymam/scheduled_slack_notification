init: rsrc=
init:
	terraform -chdir=resources/main/$(rsrc) init

init-all:
	make init rsc=services
	make init rsc=image
	make init rsc=secrets
	make init rsc=job
	make init rsc=schedule

plan: rsrc=
plan:
	terraform -chdir=resources/main/$(rsrc) plan

plan-all:
	make plan rsc=services
	make plan rsc=image
	make plan rsc=secrets
	make plan rsc=job
	make plan rsc=schedule

apply: rsrc=
apply:
	terraform -chdir=resources/main/$(rsrc) apply -auto-approve

apply-all:
	make apply rsc=services
	make apply rsc=image
	make apply rsc=secrets
	make apply rsc=job
	make apply rsc=schedule
