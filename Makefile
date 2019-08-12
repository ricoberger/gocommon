.PHONY: release-major release-minor release-patch

release-major:
	$(eval MAJORVERSION=$(shell git describe --tags --abbrev=0 | sed s/v// | awk -F. '{print $$1+1".0.0"}'))
	git checkout master
	git pull
	git tag -a $(MAJORVERSION) -m 'release $(MAJORVERSION)'
	git push origin --tags

release-minor:
	$(eval MINORVERSION=$(shell git describe --tags --abbrev=0 | sed s/v// | awk -F. '{print $$1"."$$2+1".0"}'))
	git checkout master
	git pull
	git tag -a $(MINORVERSION) -m 'release $(MINORVERSION)'
	git push origin --tags

release-patch:
	$(eval PATCHVERSION=$(shell git describe --tags --abbrev=0 | sed s/v// | awk -F. '{print $$1"."$$2"."$$3+1}'))
	git checkout master
	git pull
	git tag -a $(PATCHVERSION) -m 'release $(PATCHVERSION)'
	git push origin --tags
