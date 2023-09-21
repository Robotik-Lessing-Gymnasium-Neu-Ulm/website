# Wie man die Website ändert

```bash
git pull --all

git checkout -b <featurebranch> production 

#änderungen vornehmen
git add <geänderte files>

git commit -m "commit msg"

git push -u origin <featurebranch>

#optional wenn fertig und featurebranch nicht mehr branch
git checkout production

git branch -d <featurebranch>

```

Danach in GitHub Pullrequest basierend auf faeture Branch erstellen und mergen...
