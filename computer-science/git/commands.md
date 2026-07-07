# Commands

Merge commits into branch while keeping only the changes on that branch

``` bash
git merge -s ours origin/branch -m "commit message"
```

## Stashing
``` bash
git stash push -m "My Message"
git stash 
```

## Releasing

with releasing the goal is to let main be equal to the release branch. 

``` bash
git checkout main
git pull
git merge -s ours develop --no-commit
git read-tree --reset -u develop
git commit -m "Release: make main match develop"
git push
```