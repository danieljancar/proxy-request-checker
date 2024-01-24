# Release

We use [semantic versioning](https://semver.org/) for our releases, together with other CI/CD tools, ran by GitHub Actions.

## Publishing a new release

To publish a release make sure to have all features stable on develop branch, then push to master branch. This will trigger a new release with semantic versioning found in `.github/workflows/cd.yml`.

