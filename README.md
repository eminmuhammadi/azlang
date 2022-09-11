# AzLang

[![](https://img.shields.io/static/v1?label=Sponsor&message=%E2%9D%A4&logo=GitHub&color=%23fe8e86)](https://github.com/sponsors/eminmuhammadi)
[![CodeQL](https://github.com/eminmuhammadi/azlang/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/eminmuhammadi/azlang/actions/workflows/codeql-analysis.yml)

CLI based tool for visualizing and pronouncing Azerbaijani language (Azərbaycan dili).

It includes different type of writing systems:

- **OFFICIAL_LATIN** => official version since 1991
- **OLD_LATIN** => 1929-1938 version; no longer in use; replaced by 1991 version
- **CYRILLIC** => 1958 version, still official in Dagestan

## The list of available commands

### transform

```bash
azlang transform --input="Salam, Dünya!" --from="OFFICIAL_LATIN" --to="CYRILLIC"
Салам, Дүнја!
```

```bash
azlang transform --input="Salam, Dünya!" --from="OFFICIAL_LATIN" --to="OLD_LATIN"
Salam, Dynja!
```

..To be continued
