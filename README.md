# mergectl
![GitHub](https://img.shields.io/github/license/tttol/mos3)  
Merge multiple git branches.

# Install
## Mac, Linux
Run the following command on terminal.
```bash
curl -fsSL https://raw.githubusercontent.com/tttol/mergectl/main/install.sh | sh
```

## Windows
Run the following command on PowerShell.
```ps1
powershell -Command "Invoke-WebRequest -Uri https://raw.githubusercontent.com/tttol/mergectl/main/install.ps1 -OutFile install.ps1; .\install.ps1"
```

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