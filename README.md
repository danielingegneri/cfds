# cfds - Coldfusion data source management CLI
This is a simple tool to create `neo-datasource.xml` for Coldfusion 2016. It may work for CF 11 too.

## Usage:
### Generate neo-datasource.xml:
Both `input` and `output` have defaults and are not mandatory.

`cfds generate --seed=1234567890123456 [--input=datasources.yml] [--output=neo-datasource.xml]`

### Decrypt password:
`cfds decrypt --seed=1234567890ABCDEF <password>`

Where `<password>` is a Base64 encoded password, usually taken directly from 
the `neo-datasource.xml` file. Outputs to `STDOUT`.

### Encrypt password:
`cfds encrypt --seed=1234567890ABCDEF <password>`

Where `<password>` is the cleartext password. Outputs Base64 encoded password 
to `STDOUT`.