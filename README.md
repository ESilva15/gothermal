# About
This is still an early stage of the project. Its usable but it still has no
proper build system up.
I use it daily as it is, but I do hope I can make it something more accessible
in this learning adventure.


# TODO
- Submit the profile of the HZ-8360 thermal printer to the escpos-printer-db
project (I still need to properly configure the printer with its full features)
- Add a build system
- ...


# Usage
## Configuration
Configure your `config.yaml` in the `src` directory:
```yaml
use: "socket"
socket:
  host: "<printer-ip>"
  port: "<printer-port>"
usb:
  path: "<printer-usb-device>"
encoding:
  capabilities-file: ./capabilities.json
```


## About the `capabilities.json` file
For this, I decided to use the [escpos-printer-db](https://github.com/receipt-print-hq/escpos-printer-db)
project already available on GitHub. It's a project where one can configure its
thermal printer and then use a script to collate its code pages data.

I have copied the profile for my printer into this repo `HZ-8360.yml`, if you
have a different thermal printer you should check the original repo for
something compatible and if there isn't, write your own based on other profiles.

Copy the `capabilities.json` file you generated into the `src` directory of this
project.


## Running it
Go to the `src` directory and try it with:
`$ echo "Something, something" | go run .`


# Relevant links
https://github.com/receipt-print-hq/escpos-printer-db
https://download4.epson.biz/sec_pubs/pos/reference_en/escpos/dle_eot.html
https://download4.epson.biz/sec_pubs/pos/reference_en/escpos/commands.html

https://en.wikipedia.org/wiki/Code_page_860
https://en.wikipedia.org/wiki/Windows-1258
