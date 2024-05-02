#!/bin/env python

from bs4 import BeautifulSoup
import json
import re
import os

ASCII = {
    "NUL": 0x0,
    "SOH": 0x1,
    "STX": 0x2,
    "ETX": 0x3,
    "EOT": 0x4,
    "ENQ": 0x5,
    "ACK": 0x6,
    "BEL": 0x7,
    "BS": 0x08,
    "TAB": 0x09,
    "HT": 0x09,
    "LF": 0x10,
    "VT": 0x11,
    "FF": 0x12,
    "CR": 0x13,
    "SO": 0x14,
    "SI": 0x15,
    "DLE": 0x16,
    "DC1": 0x17,
    "DC2": 0x18,
    "DC3": 0x19,
    "DC4": 0x20,
    "NAK": 0x21,
    "SYN": 0x22,
    "ETB": 0x23,
    "CAN": 0x24,
    "EM": 0x25,
    "SUB": 0x26,
    "ESC": 0x27,
    "FS": 0x28,
    "GS": 0x29,
    "RS": 0x30,
    "US": 0x31,
    "SP": 0x32,
    "DEL": 0x7F,
    "LRM": 0xE2808E,
    "RLM": 0xE280AF,
    "SHY": 0xC2AD,
    "NBSP": 0xC2A0,
}

# Relation of ESC/POS CODE PAGE NUMBER : MY SHITTY PRINTER NUMBER
# 1 Katakana prints something different from the esc/pos
PAGE_NUMBERS = {
    0: 0,
    2: 2,
    3: 3,
    4: 4,
    5: 5,
    16: 16,
    17: 17,
    18: 18,
    46: 23,
    51: 25,
    37: 28,
    49: 32,
    35: 56,
    34: 60,
    13: 61,
    36: 62,
    14: 64,
    11: 65,
    45: 72,
    47: 90,
    48: 91,
    50: 92,
    32: 93,
    52: 94,
    33: 95,
}
# Where is the 16th ?


def get_hex_string(original: str):
    data = bytes(original, encoding='utf-8')
    strhex = ''.join(format(byte, '02X') for byte in data)

    return strhex


def get_unicode_codepoint_string(original: str):
    if len(original) != 1:
        return ''

    char = ord(original)
    val = format(char, '02X')
    return val


def clean_char_representation(input_string):
    # Regular expression pattern to match strings with square brackets and
    # contents inside
    result = input_string
    pattern = r'\[[^\]]+\]'

    # Remove the square brackets and contents inside from each match
    result = re.sub(pattern, '', result)

    # Check if the last 4 characters are hexadecimal
    if re.match(r'[0-9A-Fa-f]{4}$', result[-4:]):
        # If they are, remove them
        result = result[:-4]

    return result


def get_codepoint(input: str):
    # Regular expression pattern to match Unicode string representations
    unicode_pattern = r'U\+[0-9A-Fa-f]{4}'

    # Find all matches in the input string
    matches = re.findall(unicode_pattern, input)

    # Print the matches
    if matches:
        return matches[0]


def get_hex_code(codepoint: str):
    # Parse the hexadecimal value from the string representation
    # Skip the "U+" prefix and convert to integer
    if codepoint is None:
        return None

    hex_value = int(codepoint[2:], 16)

    # Convert the integer to a character
    char = chr(hex_value)

    # Encode the character to UTF-8 bytes
    utf8_bytes = char.encode('utf-8')

    # Convert each byte to its hexadecimal representation
    hex_representation = ''.join(format(byte, '02X') for byte in utf8_bytes)

    return hex_representation


def get_escpos(left: str, right: str):
    return left[0] + right


def make_charmap(file):
    # htmlTable
    with open(file, "r") as t:
        html_content = t.read()

    # Parse the HTML content
    soup = BeautifulSoup(html_content, 'html.parser')

    # We only have a html table in the html of the file so we can go with it
    table = soup.table

    # Extract data from each row of the table
    rows = table.select("tr")
    index = [x for x in rows[0].text.strip() if x != '\n']
    # print(index)

    charset = []
    for row in rows[1:]:
        cells = row.select("td")
        cur_row = cells[0].text.strip()

        counter = 0
        for cell in cells[1:]:
            escpos = get_escpos(cur_row, index[counter])
            if len(escpos) > 2:
                print("escpos representation cannot be longer than 2 bytes.")
                exit(1)

            codepoint = get_codepoint(cell.attrs.get('title'))
            hexcode = get_hex_code(codepoint)

            character = cell.text.strip()
            if len(character) > 1:
                character = clean_char_representation(character)

            int_escpos = int(escpos, 16)
            if int_escpos > 0xFF or int_escpos < 0x80:
                continue

            if hexcode is None:
                character = None
            else:
                hexcode = "0x" + character

            cur_data = {
                "escpos_hex": "0x" + escpos,
                # "utf_hex": hexcode,
                "unicode_cp": codepoint,
                "character": character,
            }

            print(cur_data)

            charset.append(cur_data)
            counter += 1

    file_name_metadata = (
        file.replace(".html", "").replace("./", "").split('_')
    )
    charmap = {
        "code-page-number": int(file_name_metadata[0]),
        "code-page-name": file_name_metadata[1],
        "charset": charset,
    }
    return charmap


def list_tables(dir_path):
    tables = []
    for f in os.listdir(dir_path):
        if f.endswith(".html"):
            tables.append(f)

    return tables


def main():
    dir = "./"
    tables = list_tables(dir)
    all_charmaps = []
    for t in tables:
        charmap = make_charmap(dir + t)
        all_charmaps.append(charmap)
        json_name = dir + t.replace('.html', '.json')
        # print(json.dumps(charmap))
        #
        # with open(json_name, 'w') as f:
        #     json.dump(charmap, f)

    characters = {}
    for charmap in all_charmaps:
        for charset in charmap["charset"]:
            data = {
                "code-page-number": charmap["code-page-number"],
                "escpos": charset["escpos_hex"],
            }
            if charset["unicode_cp"] in characters:
                characters[charset["unicode_cp"]].append(data)
            else:
                characters[charset["unicode_cp"]] = [data]

    # print(json.dumps(characters))
    # with open('5890_charmap.json', 'w') as f:
    #     json.dump(characters, f)


if __name__ == "__main__":
    main()


# Its missing:
# Katakana
# TCVN-3
# CP1125
# Some hebrew chars, will work on them later I guess

# Its working:
# For CP: 775, 850, 860, 863, 865, 851, 853, 857, 737, 866, 852,
#         720,855, 861, 862, 864, 869, 1098
#    ISO: 8859-7, 8859-2, 8859-15
#    WCP: 1252, 775, 1250, 1251, 1253, 1254, 1255, 1256, 1257, 1258
#    KZ: 1048

# The KZ 1048 is a table with the rows 8x to Bx and replaces those
# rows on the table wcp1251
