
update-github-actions:
	.github/ytt/ytt -f .github/templates/deploy.yaml --data-values-file setting.yaml > .github/workflows/deploy.yaml