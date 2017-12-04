## Kibana

Easy installation of kibana plugins via environment variable `KIBANA_PLUGINS`.

### Usage

Set the `KIBANA_PLUGINS` environment variable with semi-colon `;` separated plugin addresses.

e.g. will install both [Sentinl](https://github.com/sirensolutions/sentinl) and [Logtrail](https://github.com/sivasamyk/logtrail)

```
KIBANA_PLUGINS="https://github.com/sirensolutions/sentinl/releases/download/tag-5.6.2/sentinl-v5.6.4.zip;https://github.com/sivasamyk/logtrail/releases/download/v0.1.23/logtrail-5.6.4-0.1.23.zip"
```
