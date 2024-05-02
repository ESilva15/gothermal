#!/bin/env python

import json
import subprocess

PATH = "/home/esilva/Desktop/projetos/escpos-printer-db/dist/capabilities.json"
COMMAND = """iconv -f {charset} -t UTF-8 <<EOF | grep -v "^\\s*$"
$(printf '%b' "$(seq 128 255 | awk '{{printf("%c", $1)}}')")
EOF
"""

with open(PATH) as f:
    capabilites = json.load(f)


# .profiles has the printer profiles
# our printer is .profiles.HZ-8360
# our code pages are in .profiles.HZ-8360.codePages
# .encodings has the available encodings and the iconv name

encodings = capabilites["encodings"]
codepages = capabilites["profiles"]["HZ-8360"]["codePages"]
for key in codepages:
    name = encodings[codepages[key]]["name"] 
    if name.startswith("Unimplemented"):
        continue

    if "iconv" not in encodings[codepages[key]]:
        continue

    # if "data" not in encodings[codepages[key]]:
    #     continue

    cp = encodings[codepages[key]]
    cmd = COMMAND.format(charset = cp["iconv"])
    subprocess.run(cmd, shell=True)
    # print(cp["iconv"])



# print(json.dumps(capabilites["profiles"]["HZ-8360"]["codePages"]))
