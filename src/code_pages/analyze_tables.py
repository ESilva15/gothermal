#!/bin/env python


import os


def get_tables(dir_path):
    txt_files = []
    for f in os.listdir(dir_path):
        if f.endswith("1255.txt"):
            txt_files.append(f)

    return txt_files


def debug(d):
    hex = ":".join("{:02x}".format(ord(c)) for c in d[1])
    # if hex == 0x5b0:
    print(hex, d)


def main():
    dir = "./"
    charset = {}
    for f in get_tables(dir):
        with open(dir + f, "r") as data:
            for line in data.readlines():
                d = line.strip('\n').split(' ')

                debug(d)

                if len(d) < 2:
                    continue

                if d[1] in charset:
                    charset[d[1]].append(f)
                else:
                    charset[d[1]] = [f]

    # for k in charset:
    #     print(k, ":".join("{:02x}".format(ord(c)) for c in k), charset[k])


if __name__ == "__main__":
    main()
