# Configuration

For our proxy request checker to work, we need to set up some personalized configurations, to fit your needs. We have
several configuration locations and options.

# Environment Variables

We have a few environment variables that you can set to configure the proxy request checker in general. The variables
are saved in the `.env` file at the root of the project.

| Variable                  | Description                                                 | Default |
|---------------------------|-------------------------------------------------------------|---------|
| `CREATE_REPORT`           | Whether or not to create a report file.                     | `1`     |
| `PROXY_HANDLING`          | Whether or not to handle proxies.                           | `0`     |
| `PROXY_EXPECTED_RESPONSE` | The expected response from the proxy, to start handling it. | `403`   |

# Configuration Files

There are two configuration files, one for the links to check, and one for the proxies to handle the intercepted
returned bodies.

## Links

Configure all the links to be checked in the `config/links/links-config.json` file. The format is as follows:

```json
[
  {
    "url": "https://xxx.xxx.xxx.xxx:xxxx",
    "expectedStatus": 200
  },
  {
    "url": "https://xxx.xxx.xxx.xxx:xxxx",
    "expectedStatus": 403
  }
]
```

> **Remember:** As you see above, if a link is expected to and returns a `403` status code, it will be handled by the
> proxy handler.

## Proxies

Configure the proxies to handle the intercepted returned bodies in the `config/proxy/proxy-config.json` file. We can
read out defined values from the returned HTML body of your proxy interceptor service.

```json
{
  "html_parsing": {
    "depth": 2,
    "tag": "h1"
  }
}
```

### What does this mean?

- `depth`: The depth of the HTML tree to search for the tag.
- `tag`: The tag to search for.

#### Example

```html
<html>
  <body>
    <h1>My Value</h1>
  </body>
</html>
```

> **Note:** The above example is a very simple HTML structure, but it is just to show you how the depth and tag work.


