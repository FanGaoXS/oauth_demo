## clone repository

```bash
git clone https://github.com/FanGaoXS/oauth_demo.git
```

## move to directory

```bash
cd oauth_demo 
```

## copy environment file

```bash
cp env.example .env
```

## update environment config

```
GITHUB_CLIENT_ID = {your github client id}
...
```

## install dependency

```bash
go mod download
```

## run

```bash
go run main.go
```