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

## How to install

- Download package from [Releases gawainx/gxd\-cli](https://github.com/gawainx/gxd-cli/releases)
- `mv gxd-cli-{system}-{version} gxd-cli` and `sudo cp gxd-cli /usr/local/bin` or `sudo mv gxd-cli-{system}-{version} /usr/local/bin/gxd-cli`
- Enjoy it.

## How to use

- Install gxd-cli
- change dir to your work dir.
- type `gxd-cli init` to create example toml. Or you can create your own toml file.
- `gxd-cli up {service.toml}` to make it work.

## How to build

To build from source you need to install golang and [Mage :: Mage](https://magefile.org/)

```shell
git clone git@github.com:gawainx/gxd-cli.git
cd gxd-cli
mage build # build for both linux and macOS
mage darwin # build for macOS only
```

## More

If you meet any issues welcome to contact me.

## Thanks

- [urfave/cli: A simple, fast, and fun package for building command line apps in Go](https://github.com/urfave/cli)
- [magefile/mage: a Make/rake\-like build tool using Go](https://github.com/magefile/mage)
- [moby/moby: Moby Project \- a collaborative project for the container ecosystem to assemble container\-based systems](https://github.com/moby/moby)