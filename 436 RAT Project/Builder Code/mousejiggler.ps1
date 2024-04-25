Add-Type -AssemblyName System.Windows.Forms
$latmax = [System.Windows.Forms.Screen]::GetBounds(1) | Select-Object -ExpandProperty Width
$longmax = [System.Windows.Forms.Screen]::GetBounds(1) | Select-Object -ExpandProperty Height
while (1) { 
  $lat = Get-Random -Minimum 0 -Maximum ($latmax - 500)
  $long = Get-Random -Minimum 0 -Maximum ($longmax - 500)
  [Windows.Forms.Cursor]::Position = New-Object System.Drawing.Point(($lat), ($long));
  Start-Sleep -Milliseconds 50000
}