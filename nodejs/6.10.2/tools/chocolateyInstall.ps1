$version       = '6.10.2'
$packageName   = 'nodejs'
$url           = "https://nodejs.org/dist/v$version/node-v$version-x64.msi"
$installertype = 'msi'
$silentArgs    = '/quiet'
$validExitCodes = @(0)

Install-ChocolateyPackage "$packageName" "$installertype" "$silentargs" "$url" -validExitCodes $validExitCodes
