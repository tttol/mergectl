# mergectl
![GitHub](https://img.shields.io/github/license/tttol/mos3)  
Merge multiple git branches.

# Install
```bash
curl -fsSL https://raw.githubusercontent.com/tttol/mergectl/main/install.sh | sh
```

# Usage
```bash
cd [path to git repository]
mergectl [source branch] [target branch 1] [target branch 2] [target branch 3] ...
```

> [!NOTE]
> `mergectl main test1` means following commands below
> ```bash
> git checkout test1
> git pull
> git merge remotes/origin/main --no-ff
> git push
> ```