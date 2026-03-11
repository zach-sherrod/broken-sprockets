# Broken Sprockets

**Broken Sprockets** is a Go library that provides **type-safe engineering units and conversions**.

The goal of the project is to eliminate common unit-conversion mistakes by enforcing **strong typing of physical quantities** such as distance, time, angle, speed, pressure, and more.

Instead of passing raw `float64` values throughout a codebase and hoping everyone remembers the units, this library encodes the unit system directly into the type system.

This makes code:

- safer
- clearer
- harder to misuse
- easier to maintain

---

# Motivation

Engineering software frequently suffers from subtle unit mistakes.

For example:

```go
distance := 100
time := 5
speed := distance / time
```

What units are these?

- meters?
- feet?
- seconds?
- milliseconds?

These ambiguities cause real-world bugs.

Broken Sprockets solves this by requiring explicit units:

```go
d := distance.Meters(100)
t := duration.Seconds(5)

v := speed.FromDistanceOverDuration(d, t)
```

Now the units are explicit and enforced by the compiler.

---

# Design Principles

The library follows a few simple design rules.

## Canonical Units

Each physical quantity is stored internally using a **canonical SI unit**.

| Quantity | Internal Unit |
|--------|--------|
| Distance | meters |
| Duration | seconds |
| Angle | radians |
| Speed | meters/sec |
| Angular Velocity | radians/sec |
| Acceleration | meters/sec² |
| Frequency | hertz |
| Pressure | pascals |

Using canonical units ensures conversions remain consistent and predictable.

---

## Separation of Types and Conversions

The library separates **quantity types** from **unit conversion packages**.

```
quantity/
    distance.go
    duration.go
    angle.go
    speed.go
    angular_velocity.go
    acceleration.go
    frequency.go
    pressure.go

distance/
duration/
angle/
speed/
angularvelocity/
acceleration/
frequency/
pressure/
```

- `quantity` defines physical types
- unit packages provide constructors and conversions

This separation keeps the library organized and scalable.

---

## Explicit Constructors

Units are created using explicit constructors.

Example:

```go
distance.Meters(100)
distance.Feet(50)

duration.Seconds(5)
duration.Milliseconds(100)

angle.Degrees(90)
angle.ArcSeconds(15)
```

---

## Explicit Conversions

Conversions are always explicit.

```go
meters := distance.ToMeters(d)
feet := distance.ToFeet(d)

seconds := duration.ToSeconds(t)
ms := duration.ToMilliseconds(t)
```

This prevents accidental mixing of units.

---

# Installation

If published as a module:

```
go get github.com/zacharysherrod/broken_sprockets
```

Import packages as needed:

```go
import (
    "github.com/zacharysherrod/broken_sprockets/distance"
    "github.com/zacharysherrod/broken_sprockets/duration"
    "github.com/zacharysherrod/broken_sprockets/speed"
)
```

---

# Example Usage

## Distance

```go
d := distance.Feet(100)

meters := distance.ToMeters(d)
feet := distance.ToFeet(d)
```

---

## Duration

```go
t := duration.Seconds(2)

ms := duration.ToMilliseconds(t)
```

---

## Speed

```go
d := distance.Meters(100)
t := duration.Seconds(10)

v := speed.FromDistanceOverDuration(d, t)

fmt.Println(speed.ToMetersPerSecond(v))
```

---

## Angular Velocity

```go
w := angularvelocity.DegreesPerSecond(15)

rad := angularvelocity.ToRadiansPerSecond(w)
```

---

## Acceleration

```go
a := acceleration.G(1)

fmt.Println(acceleration.ToMetersPerSecondSquared(a))
```

---

## Frequency

```go
f := frequency.Hertz(50)

period := frequency.ToPeriod(f)

fmt.Println(duration.ToMilliseconds(period))
```

---

## Pressure

```go
p := pressure.PSI(14.7)

fmt.Println(pressure.ToAtmospheres(p))
```

---

# Relationships Between Quantities

The library supports common physical relationships.

## Distance / Time → Speed

```go
speed.FromDistanceOverDuration(distance, duration)
```

## Angle / Time → Angular Velocity

```go
angularvelocity.FromAngleOverDuration(angle, duration)
```

## Speed / Time → Acceleration

```go
acceleration.FromSpeedOverDuration(speed, duration)
```

## Frequency ↔ Period

```
frequency = 1 / duration
```

```go
frequency.FromPeriod(duration)
frequency.ToPeriod(frequency)
```

---

# Example: Motion Calculation

```go
d := distance.Meters(100)
t := duration.Seconds(4)

v := speed.FromDistanceOverDuration(d, t)

fmt.Println(speed.ToMetersPerSecond(v))
```

---

# Example: Angular Motion

```go
a := angle.Degrees(90)
t := duration.Seconds(3)

w := angularvelocity.FromAngleOverDuration(a, t)

fmt.Println(angularvelocity.ToDegreesPerSecond(w))
```

---

# Running Tests

Run all tests in the module:

```
go test ./...
```

---

# Philosophy

Broken Sprockets favors:

- **clarity over cleverness**
- **explicit units**
- **simple APIs**
- **predictable conversions**

The goal is to make unit usage obvious in code and prevent subtle mistakes.

---

# Planned Future Units

The following packages are planned or under consideration.

| Quantity | Status |
|--------|--------|
| Mass | planned |
| Force | planned |
| Energy | planned |
| Power | planned |
| Temperature | planned |
| Area | planned |
| Volume | planned |

---

# Contributing

Contributions are welcome.

Potential improvements include:

- additional unit systems
- additional derived quantities
- improved documentation
- performance improvements
- expanded test coverage

---