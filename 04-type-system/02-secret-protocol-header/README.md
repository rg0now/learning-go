# Setting secret protocol header

Consider a custom protocol that has the following fixed header for a `PUBLISH` packet:

| **Bit** | **7**                            | **6** | **5** | **4**        | **3**  | **2** | **1**     | **0**  |
|---------|----------------------------------|-------|-------|--------------|--------|-------|-----------|--------|
| byte 1  | PROTOCOL Connect Packet Type (2) |       |       | FirstAttempt | QoS(2) |       | Broadcast | Secure |
|         | 0                                | 1     | 0     | X            | X      | X     | X         | X      |

The definition of `QoS` is the following:

| **QoS value** | **Bit 2** | **Bit 1** | **Description**             |
|---------------|-----------|-----------|-----------------------------|
| 0             | 0         | 0         | At most once delivery       |
| 1             | 0         | 1         | At least once delivery      |
| 2             | 1         | 0         | Exactly once delivery       |
| -             | 1         | 1         | Reserved - must not be used |

Write a small function that accepts three arguments: `FirstAttempt`, `Broadcast` and `Secure`, and constructs the header for the `PUBLISH` packet. Each argument has type `bool`, where the value of the actual bit is assigned to the parameter; e.g., if the `Broadcast` bit should be `1`, then the value of the corresponding argument will be also `1`. Note that the length of the Packet type and the QoS fields is two bits.

The signature of the function should be like below:

``` go
func createPublishFixHeader(bool, bool, bool) byte
```

**Important!** The QoS setting is fix! In your case the QoS requirement is: **At least once delivery**.

Place your code into the file `exercise.go` near the placeholder `// INSERT YOUR CODE HERE`.
