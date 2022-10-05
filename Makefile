init-terraform:
	echo "bucket=$$(yq '.terraform.bucket' setting.yaml)" > terraform/.tfbackend

update-github-actions:
	ytt -f .github/templates/deploy_bottom.yaml --data-values-file setting.yaml > tmp.yaml
	cat .github/templates/deploy_upper.yaml tmp.yaml > .github/workflows/deploy.yaml
	rm tmp.yaml
