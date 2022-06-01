For the CLI app:

1. We build the go project for each micros
2. We do the shitty docker cmd
3. Docker launch each binary into alpine (It's a linux OS)
4. Each micros create a "/app" at the root of their own OS (alpine) "/" and copy it's binary into
5. The binary is executed from "/app/binary" by docker, so the app run into "/" (root of OS)
6. PS: since the binary is only dealing with go files, if you need others file to be accessible by the goApp you'll have to copy them into the root of the Docker OS (alpine), e.g in a "assets" dir, and since the app run in "/", the way to access it from the app is "./assets/..." even if in development these two objects are not linked by this path
