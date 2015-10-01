### Servo example

**Note: Currently this is my personal playground for the [servo development][1].**
The example here can change at any time and without extra notice.

#### Structure

The example project is configured via the files in the [config](config) directory.
Have a close look at [config/types.yml](config/types.yml) to see how the application is wired up.
The individual routes are configured in [config/routes.yml](config/routes.yml).

If you would deploy this application these files would remain unchanged.
The [config/config.yml] in contrast contains actual configuration parameters that may change depending on the environment.

Additionally the example uses some of the experimental bundles I already created. Checkout [main.go](main.go) for details.

#### Run the example

To run the example do:

```golang
go run *.go
```

[1]: https://github.com/fgrosse/servo
