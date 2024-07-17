$ErrorActionPreference = 'Stop'

[string]$Version

if (-not $Version) {
    Write-Host "Fetching the latest version..."
    try {
        $response = Invoke-WebRequest -Uri "https://api.github.com/repos/tttol/mergectl/releases/latest" -UseBasicParsing
        $json = $response.Content | ConvertFrom-Json
        $Version = $json.tag_name
        if (-not $Version) {
            Write-Error "Failed to fetch the latest version"
            exit 1
        }
    } catch {
        Write-Error "Failed to fetch the latest version"
        exit 1
    }
}

# TODO arm64, i386
$url = "https://github.com/tttol/mergectl/releases/download/$Version/mergectl_Windows_x86_64.zip"
$output = "$env:temp\mergectl.zip"
$installDir = "$env:LocalAppData\mergectl"

Write-Host "Downloading mergectl version $Version from $url"
Invoke-WebRequest -Uri $url -OutFile $output

Write-Host "Extracting mergectl"
Expand-Archive -Path $output -DestinationPath $installDir

Write-Host "Adding mergectl to PATH"
$oldPath = [Environment]::GetEnvironmentVariable('Path', [System.EnvironmentVariableTarget]::User)
$newPath = "$oldPath;$installDir"
[Environment]::SetEnvironmentVariable('Path', $newPath, [System.EnvironmentVariableTarget]::User)

Write-Host "Installation complete. Please restart your terminal."
