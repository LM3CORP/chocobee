$ErrorActionPreference = 'Stop'

$packageArgs = @{
  packageName    = 'ossec-client'
  url            = 'https://updates.atomicorp.com/channels/atomic/windows/ossec-agent-win32-2.9.4-5177.exe'
  installerType  = 'exe'
  checksum       = '503E93C7A1E71A3B0FBA19B9BD36C00938231FD8625FB2BCE57EBED23AA67626'
  checksumType   = 'sha256'
  silentArgs     = '/S'
  validExitCodes = @(0)

}

Install-ChocolateyPackage @packageArgs


