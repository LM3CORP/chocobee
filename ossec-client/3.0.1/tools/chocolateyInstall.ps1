$ErrorActionPreference = 'Stop'

$packageArgs = @{
  packageName    = 'ossec-client'
  url            = 'https://updates.atomicorp.com/channels/atomic/windows/ossec-agent-win32-3.0.1-5667.exe'
  installerType  = 'exe'
  checksum       = '1AD4FD843780E304AB60405EF5DB7AF5456948F3E4CB17205FDF774828E1FDFB'
  checksumType   = 'sha256'
  silentArgs     = '/S'
  validExitCodes = @(0)

}

Install-ChocolateyPackage @packageArgs


