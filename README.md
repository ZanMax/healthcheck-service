# Healthcheck Service

Healthcheck helps you check if your services are still running.

### Installation

```bash
./build.sh
```

```bash
cp bin/healthcheck /usr/local/bin
```

### Supervisor

```bash
sudo apt update
sudo apt install -y supervisor
sudo service supervisor start
sudo supervisorctl status
```

```bash
sudo vim /etc/supervisor/conf.d/healthcheck.conf
```

```
[program:healthcheck]
directory=/usr/local
command=/usr/local/bin/healthcheck
autostart=true
autorestart=true
stderr_logfile=/var/log/healthcheck/err.log
stdout_logfile=/var/log/healthcheck/out.log
environment=HEALTHCHECK_ENV=prod
```

```bash
sudo supervisorctl reload
```

### Usage

http://localhost:8080/healthcheck
or
http://localhost:8080/healthcheck/ext