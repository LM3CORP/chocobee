$ErrorActionPreference = 'Stop'

$packageArgs = @{
  packageName    = 'ossec-client'
  url            = 'https://updates.atomicorp.com/channels/atomic/windows/ossec-agent-win32-3.0.0-5502.exe'
  installerType  = 'exe'
  checksum       = '63F5A6807383BE5B51BA5706300633372645E1FAFA00FE0B77E3862C40E20510'
  checksumType   = 'sha256'
  silentArgs     = '/S'
  validExitCodes = @(0)

}

Install-ChocolateyPackage @packageArgs


