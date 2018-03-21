# Docker WhoAmI

Simple HTTP docker service that prints it's container ID/Hostname and IP

This solution is based on [whoami](https://github.com/jwilder/whoami) project, developed by &copy;[Jason Wilder](https://github.com/jwilder).

---

* [Docker Image](#docker-image)
* [Requirements](#requirements)
* [Limitations](#limitations)
* [How-to](#how-to)
* [To-Do](#to-do)

---
### Docker Image

[![](https://images.microbadger.com/badges/image/mgvazquez/whoami.svg)](https://microbadger.com/images/mgvazquez/whoami "Get your own image badge on microbadger.com")[![](https://images.microbadger.com/badges/version/mgvazquez/whoami.svg)](https://microbadger.com/images/mgvazquez/whoami "Get your own version badge on microbadger.com")[![](https://images.microbadger.com/badges/commit/mgvazquez/whoami.svg)](https://microbadger.com/images/mgvazquez/whoami "Get your own commit badge on microbadger.com")[![](https://images.microbadger.com/badges/license/mgvazquez/whoami.svg)](https://microbadger.com/images/mgvazquez/whoami "Get your own license badge on microbadger.com")
---

### Requirements
* `docker-engine` >= 1.0

---

### Limitations
* -

---

### How-to

#### Manually (with `docker run`)

You can manually run these commands:

```bash
$ docker run -d -p 8000:8000 --name whoami -t mgvazquez/whoami
736ab83847bb12dddd8b09969433f3a02d64d5b0be48f7a5c59a594e3a6a3541

$ curl $(hostname --all-ip-addresses | awk '{print $1}'):8000
I'm cff764675a6c
My IP is 10.1.1.1
```

#### Tuning

* -

---

### To-Do
* -

---

<p align="center"><img src="http://www.catb.org/hacker-emblem/glider.png" alt="hacker emblem"></p>