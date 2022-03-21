# Download 
- copy link `https://example.tar.gz` <a href="https://go.dev/doc/install">go-dev-doc-install</a>
# install
- `sudo tar -C /usr/local -xzf go1.18.linux-amd64.tar.gz`
- `echo $PATH` vemos si existe variable de entorno
- `.bashrc` nvim editamos y escribimos
```bash
export PATH=$PATH:/usr/local/go/bin
```
- actualizamos
- `source .bashrc`
- comprobamos la variable de entorno 
- `echo #PATH`
- go version
