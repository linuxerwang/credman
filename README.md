Credman
=====
A simple, powerful, secure and cross-platform credential management.

Features:
- Multiple local profiles
- AES256 encryption of profile data
- SHA256 hashing for profile password
- Multiple fields per site/serivce
- Auto generate password
- Sync with custom server and user management
- Cross platform
- Easy usage
- Only one encrypted file per profile that you can carry around!
- Encrypt files
- Restoring profile's password in case you manage to forget it!

## Installation
You can download binary releases from [here](https://github.com/ShrewdSpirit/credman/releases)

Or you can build from source by installing [Go](https://golang.org/) and then running:

```
$ go -u github.com/ShrewdSpirit/credman
$ go install github.com/ShrewdSpirit/credman
```

## Usage
Credman works on Linux, RPi, OSX and all Windows versions (BSD might work, but I'll never test).
It requires 'xsel' or 'xclip' to be installed on Linux otherwise copy function will not work.

First you must create a profile to store your sites in:

`$ credman profile add "profile name"`

It will prompt you for password. Enter a secure one and hit enter.

If it's the first profile you create, it will be set as default profile. You can add as many profiles as you want.
Commands that deal with sites require you to specify a profile name with '-p' option. If it's not specified, default profile will be used.

Next step is to add a site/service:

`$ credman site add "site name" --field=email="my email" --field=username="my username" --password-generate`

OR

`$ credman s a 'site name' -f=email='my email' -g`

It will create a site inside default profile with email, username and an auto generated password.
You can use `-p="profile"` to set a specific profile for site or omit to use default profile.
Site's fields are optional data that you can add to store extra stuff for each site.

All credman configs and profiles will be created at user's home directory under .credman directory.

You can add/delete/rename/change password for each site and profile.
Check `credman help` for a detailed help or read [docs](https://github.com/ShrewdSpirit/credman/blob/master/Docs.md) for a brief on commands.

## TODO
- [x] Implement local management
- [x] Pattern matching for site/profile names
- [x] Colorful output!
- [ ] File encryption
- [ ] Password restore
- [ ] Server sync
- [ ] Move the functionality out of commands so other packages can use credman as a library
- [ ] Desktop GUI frontend
- [ ] Android app
