# Tailnet Portal

TODO: description

## Configuration

Here is a template for `config.json` which needs to be in the working directory.

The configuration displays the devices that are **included** in this file and in the specified order.

To put a custom icon for the device's web app, put it into `static/icons/<hostname>.png`.

```json
{
  "tailnet": "mytailnet.ts.net",
  "apiKey": "tskey-api-SECRET",
  "port": "5000",
  "devices": [
    {
      "hostname": "subsonicserver",
      "alias": "Music"
    },
    {
      "hostname": "passwordmanager",
      "alias": "Passwords"
    },
    {
      "hostname": "sftpwebclient",
      "alias": "Files"
    }
  ]
}
```
