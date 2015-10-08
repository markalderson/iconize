# iconize

A terminal command that allows you to create launchers for your Ubuntu apps.

## Example

``` bash
path/to/iconize -n "Sublime Text" -p /home/marco/Desktop/sublime/sublime -i /home/marco/Desktop/sublime/Icon/256x256/sublime_text.png
```

**Note**: If you want to be able to use `iconize` from any directory without
indicating the path to the executable, just type:

``` bash
echo 'export PATH=$PATH:/path/to/iconize' >> ~/.bashrc
source ~/.bashrc
```

Then you'll be able to call `iconize` from any directory:

``` bash
iconize -n "Sublime Text" -p /home/marco/Desktop/sublime/sublime -i /home/marco/Desktop/sublime/Icon/256x256/sublime_text.png
```

## Options

`iconize` wants a `name` (flag `-n`) to be used as a display name for the
launcher, and a path (flag `-p`) to the executable that actually
starts the app.

If `name` contains multiple words it has to be enclosed by double quotes.

You can also provide a custom `icon` for the launcher (flag `-i`).

Type `iconize -h` for help.

## How it works

`iconize` just creates a reasonable [Desktop Entry](http://standards.freedesktop.org/desktop-entry-spec/latest/) inside
`~/.local/share/applications`. The entry is named based on the `-n` flag, but
note that the name you provide will be _lower-cased_ and spaces will be
reaplaced with hyphens (e.g. `Sublime Text` -> `sublime-text.desktop`).

## Alternatives

`iconize` is just a personal excuse to learn a little bit of
[Go](https://golang.org/).

[Alacarte](https://en.wikipedia.org/wiki/Alacarte) and
[MenuLibre](https://smdavis.us/projects/menulibre/) are more powerful
solutions and also have a graphical user interface.
