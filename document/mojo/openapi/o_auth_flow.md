| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `authorizationUrl` | `string` | `Url` | Y |  | The authorization URL to be used for this flow |
| `tokenUrl` | `string` | `Url` | Y |  | The token URL to be used for this flow. |
| `refreshUrl` | `string` | `Url` | N |  | The URL to be used for obtaining refresh tokens. |
| `scopes` | `Map<string, string>` |  | Y |  | The available scopes for the OAuth2 security scheme.A map between the scope name and a short description for it. |
