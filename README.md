servo
========

[![Build Status](https://secure.travis-ci.org/fgrosse/servo.png?branch=develop)](http://travis-ci.org/fgrosse/servo)
[![Coverage Status](https://coveralls.io/repos/fgrosse/servo/badge.svg?branch=develop)](https://coveralls.io/r/fgrosse/servo?branch=master)
[![GoDoc](https://godoc.org/github.com/fgrosse/servo?status.svg)](https://godoc.org/github.com/fgrosse/servo)
[![license](https://img.shields.io/badge/license-MIT-4183c4.svg)](https://github.com/fgrosse/servo/blob/master/LICENSE)

Servo is a web framework that heavily relies on dependency injection to wire up your application.

**Note: This project is still at the early stages of its development.**

### Example Implementation
 
You can have a look at an example that uses servo to build a simple web application at **https://github.com/fgrosse/servo-example**.

### First Bundles

Servo is build in bundles. This means that servo applications pull in dependencies by using other bundles.
For an example see how bundles are used [in the example][2]

So far the following bundloes are implemented:
* https://github.com/fgrosse/servo-logxi
* https://github.com/fgrosse/servo-gorilla

### Dependencies

* go 1.3 or higher
* [gopkg.in/yaml.v2][1] (LGPLv3)

[1]: https://github.com/go-yaml/yaml/tree/v2
[2]: https://github.com/fgrosse/servo-example/blob/master/main.go#L15
