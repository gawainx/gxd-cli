# gxd-cli is gawainx' docker client

Create docker containers by writing a TOML config file!

Inspired by docker-compose.

## Support

- Edit .toml file rather than type a long long command to create your service comtainers.
- Support multi ports binding. e.g.: `-p 8080:80`
- Support multi volumes bingding. e.g.: `-v /use/bin:/code`
- Modify your own container name.
- Cross platform,work well in macOS and Linux.

## Not support and TODOs

- No Windows `.exe` file. And may never support.
- Not support interactive containers(e.g.:`bash` in your container)
- [TODO] `detach` option not take effect.
- [TODO] Create multi-services in just one `*.toml` file.
- [TODO] `pwd` support while mounting vols.

## How to use

- Install gxd-cli
- change dir to your work dir.
- type `gxd-cli init` to create example toml. Or you can create your own toml file.
- `gxd-cli -f {service.toml}` to make it work.

## How to build

To build from source you need to install golang and [Mage :: Mage](https://magefile.org/)

```shell
git clone git@github.com:gawainx/gxd-cli.git
cd gxd-cli
mage build # build for both linux and macOS
mage darwin # build for macOS only
```