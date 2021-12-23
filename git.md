* After removes files in .gitignore, use the following command to remove tracking files.
```bash
git rm -r --cached .
git add .
```
* git rm <--> git add.
* git reset breaks history while git revert add new commits.
* git rebase doesn't keep history changes but only the result, while git merge keep the whole history.
* To remove unwanted commit:
```bash
git reset --hard [LAST_WANTED_HASH]
git push -f
```
* To remove unwanted commit while keeping some:
```bash
git revert [WANTED]
git revert [UNWANTED]
git cherry-pick
```
* Change branch name
```bash
git branch -m [new_branch_name]
```
* To browse sorted tag
```bash
git tag --sort=v:refname
```
or
```bash
git tag --sort=creatordate
```
* When working on a fork
  * to keep up with the original repo
  ```bash
  git pull [original_repo_url] [branch]
  ```
  * to also keep up with tags in original repo
  ```bash
  git fetch --tags [original_repo_url]
  ```
* Change repository cloned with https int SSH: (change remote urls)
`git remote set-url origin git@github.com:[Path]/[Repo].git`
