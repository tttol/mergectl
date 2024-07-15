# mergectl
![GitHub](https://img.shields.io/github/license/tttol/mos3)  
Merge multiple git branches.

# Install
## Mac, Linux
```bash
curl -fsSL https://raw.githubusercontent.com/tttol/mergectl/main/install.sh | sh
```

## Windows
mergectl CANNOT run on the cmd and PowerShell.  
Please install WSL2 then install with command for Linux.  
  
We are currently working on making it possible to run on cmd and PowerShell. 

# Usage
```bash
cd [path to git repository]
mergectl exec [source branch] [target branch 1] [target branch 2] [target branch 3] ...
```

> [!NOTE]
> `mergectl exec main test1` means following commands below
> ```bash
> git checkout test1
> git pull
> git merge remotes/origin/main --no-ff
> git push
> ```