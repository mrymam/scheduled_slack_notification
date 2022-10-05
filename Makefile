
update-github-actions:
	.github/ytt/ytt -f .github/templates/deploy_bottom.yaml --data-values-file setting.yaml > tmp.yaml
	cat .github/templates/deploy_upper.yaml tmp.yaml > .github/workflows/deploy.yaml
	rm tmp.yaml