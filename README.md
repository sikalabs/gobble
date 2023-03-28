# Gobble: Like Ansible but in Go

## What is Gobble?

Gobble is a config management tool like Ansible but written in Go. It is designed to be simple and easy to use.

## Installation

Mac

```
brew install sikalabs/tap/gobble
```

Linux

```
curl -fsSL https://raw.githubusercontent.com/sikalabs/gobble/master/install.sh | sudo sh
```

Using [slu](https://github.com/sikalabs/slu)

```
slu install-bin gobble
```

## Example Usage

Create `gobblefile.yml`

```yaml
meta:
  schema_version: 3
global:
  no_strict_host_key_checking: true
  vars:
    prefix: '[gobble]'
    msg: Global Gobble!
hosts:
  all:
    - ssh_target: root@gobble-example.sikademo.com
      vars:
        ip: gobble-example.sikademo.com
        msg: Host Gobble!
plays:
  - name: Install Nginx
    hosts: [all]
    sudo: false
    tags: [setup]
    tasks:
      - name: Install Nginx
        apt_install:
          name: nginx
  - name: Index.html
    hosts: [all]
    sudo: false
    tags: [content]
    tasks:
      - name: Create index.html
        template:
          path: /var/www/html/index.html
          template: <h1>{{- if .Vars.prefix -}}{{.Vars.prefix}} {{ end -}}{{.Vars.msg}}</h1>{{ printf "\n" }}
  - name: See
    hosts: [all]
    sudo: false
    tags: [content]
    tasks:
      - name: See
        print:
          template: |
            See: http://{{.Vars.ip}}
```

Run

```
gobble run
```
