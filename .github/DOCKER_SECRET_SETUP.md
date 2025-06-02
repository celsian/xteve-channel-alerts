# Setting up Docker Hub Secrets for GitHub Actions 🛠️

The *Build & Publish* workflow (`.github/workflows/docker-publish.yml`) pushes multi-architecture images to Docker Hub.  
For it to authenticate, two repository secrets must be present:

| Secret name | Purpose |
|-------------|---------|
| `DOCKERHUB_USERNAME` | Your Docker Hub account/organisation name (e.g. `celsian`) |
| `DOCKERHUB_TOKEN`    | A **read/write** access token generated in Docker Hub |

Follow the steps below once per repository.

---

## 1 · Create a Docker Hub Access Token

1. Sign-in to **[hub.docker.com](https://hub.docker.com/)**  
2. Click your avatar → **Account Settings**  
3. In the left menu choose **Security** → **New Access Token**  
4. Fill in:
   * **Token Description** – e.g. *GitHub Actions*
   * **Access Permissions** – choose **Read & Write**  
5. Press **Generate** and copy the token **once** – you will not see it again!

---

## 2 · Add GitHub Repository Secrets

1. Open your GitHub repo → **Settings** → **Secrets → Actions**  
2. Click **New repository secret** twice and add:

| Name | Value |
|------|-------|
| `DOCKERHUB_USERNAME` | Your Docker Hub username/organisation |
| `DOCKERHUB_TOKEN` | *The token copied in step 1-5* |

Make sure there are **no leading/trailing spaces**.

---

## 3 · Verify the Setup

1. Push a commit to **`master`** (or trigger **Actions → docker-publish → Run workflow** manually).  
2. In **Actions** you should see **“Build and Publish Docker Image”** start.
3. After it finishes:
   * The log should contain `docker login` succeeded messages.
   * Docker Hub → **Repositories** should show the new tags (e.g. `latest`, commit SHA).  
   * The **“Platforms”** tab on Docker Hub will list `linux/amd64` and `linux/arm64`.

If the workflow fails on “Authenticate to Docker Hub”:

* Confirm the token is still valid (regenerate if unsure).  
* Ensure the secrets are spelled exactly `DOCKERHUB_USERNAME` and `DOCKERHUB_TOKEN`.

---

### FAQ

**Q: Can I use a password instead of a token?**  
A: No. Docker Hub deprecated password logins for CI; tokens are required.

**Q: Can I scope the token to a single repository?**  
A: Docker Hub tokens are account-wide. For finer control create a separate
organisation and team with limited repository access.

**Q: How do I rotate the token?**  
1. Create a new token in Docker Hub.  
2. Replace `DOCKERHUB_TOKEN` in GitHub Secrets.  
3. Delete the old token in Docker Hub.

---

You’re all set – the GitHub Actions workflow will now build **and publish** images for **Intel/AMD (amd64)** and **Apple/ARM (arm64)** under the shared `latest` tag every time you push to *master* 🚀
