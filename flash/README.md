
Ideas:

  - Understand what's under the hood,

  - Expertise for troubleshooting,

  - Get program size before flash.

```bash
$ tinygo flash -target=arduino app.go
```

--------------------------------------------------------------------------------

```bash
$ tinygo build -o app.hex -target=arduino app.go
```

```bash
$ ls -l app.hex
-rw-rw-r-- 1 boisgera boisgera 3248 oct.  21 11:12 app.hex
```


Cf. [Intel Hex Format](https://en.wikipedia.org/wiki/Intel_HEX)


```bash
$ cat app.hex
:100000000C9434000C94FC010C94FC010C94FC0145
:100010000C94FC010C94FC010C94EE010C94FC017A
:100020000C94FC010C94FC010C94FC010C94FC015C
:100030000C94FC010C94FC010C94FC010C94FC014C
:100040000C9472000C94FC010C94B4000C94FC0110
:100050000C94FC010C94FC010C94FC010C94FC012C
:100060000C94FC010C94FC011124A0E0B3E0ADBFA2
:10007000BEBF0F92A0E0B3E0C4E6D3E0E2E3F4E059
:10008000AC17BD0719F005900D92FACF0E946901D7
:100090000E9473010E94C7010E946701DC01A03029
:1000A000B10511F08C9108950E946101DC01A0302E
:1000B000B10511F08C9108950E9461011F93162FD4
:1000C0000E945600982F912381E0903009F4812DF1
:1000D0001F910895DC01A030B10511F06C930895D3
:1000E0000E9461010F921F920FB60F9211242F935D
:1000F0003F934F935F936F937F938F939F938091E1
:10010000ED039091EE032091EF033091F0034091C5
:10011000F1035091F2036091F3037091F40380585E
:10012000914C2F4F3F4F4F4F5F4F6F4F7F4F50932B
:10013000F2034093F1037093F4036093F3039093FD
:10014000EE038093ED033093F0032093EF039F9130
:100150008F917F916F915F914F913F912F910F9070
:100160000FBE1F900F9018950F921F920FB60F920F
:1001700011241F932F933F934F935F936F937F931C
:100180008F939F93AF93BF93EF93FF9386EC90E091
:100190000E945600182F80EC90E06CE10E945E00F7
:1001A00081708030A1F48091E4039091E503891B74
:1001B000803869F08091E40383958093E403809113
:1001C000E4038F77A82FBB27AC59BC4F1C93FF913A
:1001D000EF91BF91AF919F918F917F916F915F91BF
:1001E0004F913F912F911F910F900FBE1F900F9035
:1001F0001895EF92FF920F931F938B017C018DE076
:1002000093E066E170E00E940D01C701B8010E9411
:100210000D010E9441010E9448018F929F92AF926E
:10022000BF92CF92DF92EF92FF920F931F936B01D9
:100230007C0180E090E08C015C01AC0181E00C1558
:100240001D0595014207530708F4812D8170803008
:1002500069F4D701A00FB11F8C914A010E944B0194
:10026000A4010F5F1F4F4F4F5F4FE8CF1F910F91BA
:10027000FF90EF90DF90CF90BF90AF909F908F90C6
:1002800008958DE00E944B018AE00E944B01089581
:10029000F8948895FECFFF920F931F93F82E00ECF1
:1002A00010E0C80160E20E945E0081708030C9F3F6
:1002B00086EC90E06F2D0E946A001F910F91FF90D5
:1002C000089583E293E067E170E00E94F9000E94E4
:1002D0004801A4E6B3E085EF93E0A817B90711F051
:1002E0001D92FBCF08950F931F9385EC90E0612D35
:1002F0000E946A0084EC90E067E60E946A0081EC4C
:1003000090E068E90E946A0082EC90E066E00E945A
:100310006A0080E090E09093F4038093F30390935D
:10032000F2038093F1039093F0038093EF03909393
:10033000EE038093ED030EE610E0C8010E944E002C
:10034000887F80936E0016BC8FEF87BD83E084BDED
:1003500081E085BDC8010E944E00816080936E00DF
:100360008FB7F8942091F3033091F4032091F103B7
:100370003091F2032091EF033091F0032091ED03CF
:100380003091EE038FBF78941F910F9108950F93D2
:100390001F9381E08093EA0325EF33E000E019E04A
:1003A000C801821B930B809661E270E00E94FE01FF
:1003B000AB01041B150B1093E7030093E603C80180
:1003C000612D0E9412028CE493E068E170E00E94CB
:1003D0000D010E9441011F910F9108950F93002775
:1003E000A89504BF00916000086100936000002799
:1003F000009360000F9118958895FECFAA1BBB1B38
:1004000051E107C0AA1FBB1FA617B70710F0A61B14
:10041000B70B881F991F5A95A9F780959095BC0135
:10042000CD010895DC0101C06D9341505040E0F7CB
:1004300008956F7574206F66206D656D6F727970A9
:10044000616E69633A2072756E74696D65206572BC
:10045000726F723A206E696C20706F696E7465728B
:100460002064657265666572656E6365696E646554
:1004700078206F7574206F662072616E67654865BD
:100480006C6C6F2066726F6D2041726475696E6F5F
:060490002120F09F918B7A
:00000001FF
```



```bash
$ avrdude -C /etc/avrdude.conf -p atmega328p -c arduino -P /dev/ttyACM0 -D -U flash:w:app.hex:i
```

cf. `man avrdude`


`hex.py` utils to get the flashed data size.