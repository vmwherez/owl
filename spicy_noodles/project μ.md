> Adopted this name from `krj` to encompass all of the music and DSP things.
> Music as a focal point to constrain the fields C++ and signal processing.
> 
> Specifically, µ introduces C (not C++) at the same time as the `*.wav` specification and other DSP concepts
> 



"I just need to dump everything out before it’s too late, avoiding mental shortcuts as much as possible, but for you to understand my derivations, you will still need several prerequisite knowledge that I no longer remember how I acquired them:

1. C programming and UNIX shell utilities [[WAV format]]
2. [[Integral calculus]]
3. [[Complex numbers]] (you need to know what ![](data:image/svg+xml;base64,PHN2ZyB2ZXJzaW9uPScxLjEnIHhtbG5zPSdodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZycgeG1sbnM6eGxpbms9J2h0dHA6Ly93d3cudzMub3JnLzE5OTkveGxpbmsnIHdpZHRoPScxNy44NTM5NzlwdCcgaGVpZ2h0PScxMy41NzAzNDZwdCcgdmlld0JveD0nLjY2NDE2NiAtMTAuMzYyNzI5IDEzLjk0ODQyMSAxMC42MDE4MzMnPgo8ZGVmcz4KPHBhdGggaWQ9J2cxLTE4MScgZD0nTTQuMDA5Nzc2LTIuOTI0OTY3QzQuMTA2MjA0LTMuMzkxMDMzIDQuMTc4NTI0LTMuODg5MjQyIDQuMTc4NTI0LTQuMzM5MjM3QzQuMTc4NTI0LTUuMjMxMTkxIDMuODg5MjQyLTUuOTYyNDMyIDIuODg0Nzg5LTUuOTYyNDMyQy45NjQyNzUtNS45NjI0MzIgLjI1NzE0LTMuMDkzNzE1IC4yNTcxNC0xLjQ4NjU5Qy4yNTcxNC0uNTcwNTI5IC41Nzg1NjUgLjEwNDQ2MyAxLjU2Njk0NyAuMTA0NDYzQzMuMDM3NDY2IC4xMDQ0NjMgMy42NjQyNDQtMS4zODIxMjcgNC4wMDk3NzYtMi45MjQ5NjdaTTMuMjc4NTM0LTMuMTY2MDM2SDEuMjY5NjI4QzEuNDE0MjctMy44NDkwNjQgMS44NTYyMjktNS42OTcyNTcgMi44MjA1MDQtNS42OTcyNTdDMy4zMjY3NDgtNS42OTcyNTcgMy40NzEzODktNS4xOTkwNDggMy40NzEzODktNC42MjA0ODNDMy40NzEzODktNC4wOTgxNjggMy4zNTA4NTUtMy41MDM1MzIgMy4yNzg1MzQtMy4xNjYwMzZaTTMuMjE0MjQ5LTIuODc2NzUzQzMuMDEzMzU5LTEuOTI4NTUgMi41NjMzNjQtLjEzNjYwNiAxLjYxNTE2LS4xMzY2MDZDMS4xMjQ5ODctLjEzNjYwNiAuOTgwMzQ2LS42MTA3MDcgLjk4MDM0Ni0xLjIwNTM0M0MuOTgwMzQ2LTEuNzU5ODAyIDEuMTA4OTE2LTIuNDE4NzIzIDEuMjA1MzQzLTIuODc2NzUzSDMuMjE0MjQ5WicvPgo8dXNlIGlkPSdnNi0xMDEnIHhsaW5rOmhyZWY9JyNnNS0xMDEnIHRyYW5zZm9ybT0nc2NhbGUoMS4zNjM2MzYpJy8+CjxwYXRoIGlkPSdnNS0xMDEnIGQ9J00zLjkyNzY3My0yLjQ0NjAyOEMzLjkyNzY3My0zLjQ1NDI0OCAzLjI3ODkwNS00LjExMTc4MiAyLjI5Njk4Ny00LjExMTc4MkMxLjk3MjYwMy00LjExMTc4MiAxLjU5NTYxNy0zLjk5NzgxIDEuMjM2MTY1LTMuNzg3Mzk5Qy41MTcyNi0zLjM3NTM0NCAuMjI3OTQ1LTIuODMxNzgyIC4yMjc5NDUtMS45MjAwMDFDLjIyNzk0NS0xLjM2NzY3MiAuMzUwNjg1LS44OTQyNDcgLjU5NjE2NS0uNTYxMDk2Qy45MzgwODMtLjA5NjQzOCAxLjUzNDI0NyAuMTc1MzQzIDIuMTkxNzgyIC4xNzUzNDNDMi41MTYxNjUgLjE3NTM0MyAyLjg0MDU0OSAuMTA1MjA2IDMuMjAwMDAxLS4wNTI2MDNDMy40MzY3MTQtLjE0OTA0MSAzLjYyMDgyMy0uMjYzMDE0IDMuNjU1ODkyLS4zMTU2MTdMMy44MzEyMzQtLjYxMzY5OUwzLjcxNzI2Mi0uNzAxMzdDMy4xNTYxNjUtLjM2ODIxOSAyLjk1NDUyMi0uMjg5MzE1IDIuNTc3NTM1LS4yODkzMTVDMi4wMDc2NzItLjI4OTMxNSAxLjUzNDI0Ny0uNTQzNTYyIDEuMjg4NzY4LS45NzMxNTFDMS4xMjIxOTItMS4yNjI0NjYgMS4wNjA4MjItMS41MDc5NDYgMS4wNDMyODgtMi4wMTY0MzlIMi4zMjMyODhDMi45MjgyMi0yLjAxNjQzOSAzLjMwNTIwNy0yLjA0Mjc0IDMuOTEwMTM4LTIuMTM5MTc5QzMuOTE4OTA1LTIuMjYxOTE5IDMuOTI3NjczLTIuMzQwODIzIDMuOTI3NjczLTIuNDQ2MDI4Wk0zLjE0NzM5OC0yLjM0MDgyM0wyLjE2NTQ4LTIuMzE0NTIxQzEuNzYyMTkyLTIuMzE0NTIxIDEuNTM0MjQ3LTIuMzIzMjg4IDEuMDYwODIyLTIuMzY3MTI0QzEuMDYwODIyLTIuNzc5MTc5IDEuMDk1ODkxLTIuOTcyMDU2IDEuMjA5ODYzLTMuMjAwMDAxQzEuMzkzOTczLTMuNTg1NzU1IDEuNzcwOTYtMy44MjI0NjcgMi4xOTE3ODItMy44MjI0NjdDMi40ODEwOTctMy44MjI0NjcgMi43MDkwNDItMy43MDg0OTQgMi44NjY4NS0zLjQ3MTc4MkMzLjA1OTcyNy0zLjE4MjQ2NyAzLjE0NzM5OC0yLjkyODIyIDMuMTQ3Mzk4LTIuMzQwODIzWicvPgo8cGF0aCBpZD0nZzUtMTA1JyBkPSdNMi4zNzU4OTEgLjAyNjMwMVYtLjIzNjcxMkwxLjk3MjYwMy0uMjYzMDE0QzEuNjc0NTIxLS4yODA1NDggMS42NDgyMi0uMzQxOTE4IDEuNjQ4MjItLjg5NDI0N1YtNC4wNzY3MTRMMS41NjkzMTYtNC4xMTE3ODJDMS4xMjIxOTItMy45Mjc2NzMgLjY2NjMwMi0zLjgwNDkzMyAuMjAxNjQ0LTMuNzUyMzNWLTMuNTA2ODUxSC41MjYwMjhDLjg3NjcxMy0zLjUwNjg1MSAuOTExNzgxLTMuNDQ1NDgxIC45MTE3ODEtMi44NjY4NVYtLjg5NDI0N0MuOTExNzgxLS4zNDE5MTggLjg4NTQ4LS4yODA1NDggLjU4NzM5Ny0uMjYzMDE0TC4xODQxMS0uMjM2NzEyVi4wMjYzMDFDLjU1MjMyOSAuMDA4NzY3IC45MTE3ODEgMCAxLjI4IDBTMi4wMDc2NzIgLjAwODc2NyAyLjM3NTg5MSAuMDI2MzAxWk0xLjcxODM1Ny01LjU4NDY1OUMxLjcxODM1Ny01LjgzODkwNiAxLjU0MzAxNC02LjAyMzAxNiAxLjI4ODc2OC02LjAyMzAxNkMxLjAxNjk4Ny02LjAyMzAxNiAuODQxNjQ0LTUuODQ3NjczIC44NDE2NDQtNS41ODQ2NTlTMS4wMTY5ODctNS4xNDYzMDMgMS4yOC01LjE0NjMwM1MxLjcxODM1Ny01LjMyMTY0NiAxLjcxODM1Ny01LjU4NDY1OVonLz4KPC9kZWZzPgo8Zz4KPHVzZSB4PScuNjY0MTY2JyB5PScwJyB4bGluazpocmVmPScjZzYtMTAxJy8+Cjx1c2UgeD0nNy4zOTU3ODknIHk9Jy00LjMzOTcxNCcgeGxpbms6aHJlZj0nI2c1LTEwNScvPgo8dXNlIHg9JzEwLjQzNDA2MycgeT0nLTQuMzM5NzE0JyB4bGluazpocmVmPScjZzEtMTgxJy8+CjwvZz4KPC9zdmc+) means [geometrically](https://www.youtube.com/watch?v=ZxYOEwM6Wbk))"

## [[Euler's Formula]]

#### [[Integral calculus]]
##### sine.c

```c
//
/* $ cc sine.c $(pkg-config --cflags --libs jack) -lm */
//
#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>
#include <signal.h>
#include <math.h>
#include <jack/jack.h>
#define PI_F 3.14159265f

static jack_client_t *client = NULL;
static jack_port_t *port_out = NULL;
static volatile int done = 0;

static void
die(const char *msg)
{
  fprintf(stderr, "[error] %s\n", msg);
  if (client)
    jack_client_close(client);
  exit(EXIT_FAILURE);
}

static void
info(const char *msg)
{
  fprintf(stderr, "[info] %s\n", msg);
}

static void
on_shutdown(void *arg)
{
  client = NULL;
  die("jack server is down, exiting...");
}

static void
on_signal(int signum)
{
  done = 1;
}

static int
on_process(jack_nframes_t nframes, void *arg)
{
  static float phs = 0;
  jack_default_audio_sample_t *out;
  jack_nframes_t i;

  out = jack_port_get_buffer(port_out, nframes);
  for (i = 0; i < nframes; ++i) {
    out[i] = 0.1f * sinf(2*PI_F*phs);
    phs += 0.01f;
    while (phs >= 1) phs--;
  }

  return 0;
}

int main(void)
{
  const char **ports;

  client = jack_client_open("sine", JackNoStartServer, NULL);
  if (!client)
    die("fail to open client");
  info("jack client opened");

  jack_on_shutdown(client, on_shutdown, NULL);
  info("shutdown callback registered");

  if (jack_set_process_callback(client, on_process, NULL))
    die("fail to set up process callback");
  info("process callback registered");

  port_out = jack_port_register(client, "out", JACK_DEFAULT_AUDIO_TYPE,
                                JackPortIsOutput, 0);
  if (!port_out)
    die("fail to register audio output port");
  info("output port registered");

  if (jack_activate(client))
    die("fail to activate client");
  info("jack client activated");

  ports = jack_get_ports(client, NULL, NULL,
                         JackPortIsInput|JackPortIsPhysical);
  if (ports) {
    int i;
    for (i = 0; ports[i]; ++i)
      jack_connect(client, jack_port_name(port_out), ports[i]);
    jack_free(ports);
  }

  signal(SIGINT, on_signal);
  signal(SIGTERM, on_signal);
#ifndef _WIN32
  signal(SIGQUIT, on_signal);
  signal(SIGHUP, on_signal);
#endif

  info("done! press ctrl-c to exit");

  while (!done)
    sleep(1);

  jack_deactivate(client);
  jack_client_close(client);
  return 0;
}


```


https://mu.krj.st

---

[[µ++]]



