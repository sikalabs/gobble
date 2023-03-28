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
  schema_version: 2
hosts:
  all:
    - ssh_target: root@gobble-example.sikademo.com
      vars:
        msg: Gobble Gobble!
plays:
  - name: Nginx Example
    hosts: [all]
    tasks:
      - name: Install Nginx
        apt_install:
          name: nginx
      - name: Create index.html
        template:
          path: /var/www/html/index.html
          template: <h1>{{.Vars.msg}}</h1>{{ printf "\n" }}
```

Run

```
gobble run
```
