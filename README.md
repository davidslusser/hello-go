# hello-go

Example Go web app.

## Docker publishing

The [Build and Push Docker Image workflow](.github/workflows/build_and_publish_docker.yaml)
publishes the same Docker image to GitHub Container Registry (GHCR) and Docker Hub.
It runs when a GitHub release is created, a supported version tag is pushed,
or the workflow is started manually.

- GHCR: `ghcr.io/davidslusser/hello-go:<tag>`
- Docker Hub: `docker.io/<DOCKERHUB_USERNAME>/hello-go:<tag>`

Stable version tags such as `1.0.0` also update `latest` in both registries.
Development tags such as `1.0.0-dev` and `1.0.0-dev.1` also update `dev`.
Release candidate tags such as `1.0.0-rc` and `1.0.0-rc.1`
publish their versioned tag only. Manual dispatch publishes `latest`.

### GitHub Actions secrets

Configure the following repository secrets under **Settings → Secrets and variables →
Actions**:

- `DOCKERHUB_USERNAME`: Docker Hub username (also used as the Docker Hub image namespace)
- `DOCKERHUB_TOKEN`: Docker Hub access token with permission to push to the `hello-go` repository

The GHCR login uses the built-in `GITHUB_TOKEN` with `packages: write` permission.
Make sure the Docker Hub image repository exists and that the token has push access
before running the workflow.
