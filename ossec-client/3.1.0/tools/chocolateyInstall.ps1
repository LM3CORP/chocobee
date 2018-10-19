$ErrorActionPreference = 'Stop'

$packageArgs = @{
  packageName    = 'ossec-client'
  url            = 'https://updates.atomicorp.com/channels/atomic/windows/ossec-agent-win32-3.1.0-5696.exe'
  installerType  = 'exe'
  checksum       = '5D84C7F6B75B9E9AB5FFE53553298DF08186D0819611892AB8219349388D4C02'
}

Install-ChocolateyPackage @packageArgs


