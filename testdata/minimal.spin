[image]
# Consider making .packages implicit?
packages = "minimal.packages"
type = "liveos"

# LiveOS specific options
[liveos]
compression = "gzip"
filename = "Solus-1.2.1.iso"
bootloaders = ["syslinux"]

# Branding particulars
[branding]
title = "Solus 1.2.1"
start_string = "Start Solus"
