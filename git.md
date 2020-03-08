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
