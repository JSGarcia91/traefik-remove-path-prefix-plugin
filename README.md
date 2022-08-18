# Remove Path Prefix

`Remove Path Prefix` is a middleware plugin for [Traefik](https://github.com/traefik/traefik) which remove the prefix from an URL request.

Before and after processing a URL through `Remove Path Prefix` plugin:

- localhost/foo => localhost
- localhost/foo/bar => localhost/bar

`Remove Path Prefix` plugin allows to force slash in paths:
- localhost/foo => localhost/
- localhost/foo/bar => localhost/bar/

## Configuration

### Dynamic

To configure the `Remove Path Prefix` plugin you should create a [middleware](https://doc.traefik.io/traefik/middlewares/overview/) in your dynamic configuration as explained [here](https://doc.traefik.io/traefik/middlewares/overview/).
